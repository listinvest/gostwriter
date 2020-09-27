package pca9685

import (
	"fmt"
	"log"
	"time"

	"github.com/x86ed/go-i2c"
)

// Context of a i2c connection
type Context struct {
	Addr     int
	Bus      *i2c.I2C
	Prescale uint8
	Debug    bool
	OsFreq   float64
}

// PWMServoDriver Instantiates a new PCA9685 PWM driver chip with the I2C address on a TwoWire interface
func (c *Context) PWMServoDriver(addr int, cl *i2c.I2C) {
	c.Addr = addr
	c.Bus = cl
}

// Begin Setups the I2C interface and hardware
func (c *Context) Begin(prescale ...uint8) {
	c.Reset()
	var p uint8
	if len(prescale) > 0 {
		p = prescale[0]
	}
	if p > 0 {
		c.SetExtClk(p)
	} else {
		c.SetPWMFrequency(1000)
	}
	c.SetOscillatorFrequency(FrequencyOscillator)
}

// Reset sends a reset command to the PCA9685 chip over I2C
func (c *Context) Reset() {
	c.write8(Mode1, Mode1Restart)
	time.Sleep(time.Millisecond * 10)
}

// Sleep puts board into sleep mode
func (c *Context) Sleep() {
	awake := c.read8(Mode1)
	sleep := awake | uint8(Mode1Sleep)
	c.write8(Mode1, sleep)
	time.Sleep(time.Millisecond * 5)
}

// WakeUp wakes board from sleep
func (c *Context) WakeUp() {
	sleep := c.read8(Mode1)
	wakeup := sleep & ^uint8(Mode1Sleep)
	c.write8(Mode1, wakeup)
}

// SetExtClk sets EXTCLK pin to use the external clock
func (c *Context) SetExtClk(prescale uint8) {
	oldmode := c.read8(Mode1)
	newmode := (oldmode & ^uint8(Mode1Restart)) | uint8(Mode1Sleep)
	// go to sleep, turn off internal oscillator
	c.write8(Mode1, newmode)

	// This sets both the SLEEP and EXTCLK bits of the MODE1 register to switch to
	// use the external clock.
	c.write8(Mode1, (newmode | uint8(Mode1EXTCLK)))
	c.write8(Mode1, prescale)
	time.Sleep(time.Millisecond * 5)
	// clear the SLEEP bit to start
	c.write8(Mode1, (newmode & ^uint8(Mode1Sleep))|uint8(Mode1Restart)|uint8(Mode1AutoIncrement))
	if c.Debug {
		log.Printf("Mode now 0x%X\n", c.read8(Mode1))
	}
}

// SetPWMFrequency Sets the PWM frequency for the entire chip, up to ~1.6 KHz
func (c *Context) SetPWMFrequency(freq float64) {
	if c.Debug {
		log.Printf("Attempting to set freq %f\n", freq)
	}
	// Range output modulation frequency is dependant on oscillator
	if freq < 1.0 {
		freq = 1.0
	}
	if freq > 3500 {
		freq = 3500 // Datasheet limit is 3052=50MHz/(4*4096)
	}
	prescaleval := ((c.OsFreq / (freq * 4096.0)) + 0.5) - 1
	if prescaleval < PrescaleMin {
		prescaleval = PrescaleMin
	}
	if prescaleval > PrescaleMax {
		prescaleval = PrescaleMax
	}
	prescale := uint8(prescaleval)
	if c.Debug {
		log.Println("Final pre-scale: ", prescale)
	}
	oldmode := c.read8(Mode1)
	newmode := (oldmode & ^uint8(Mode1Restart)) | uint8(Mode1Sleep) // sleep
	c.write8(Mode1, newmode)                                        // go the fuck to sleep
	c.write8(PreScale, prescale)                                    // set the prescaler
	c.write8(Mode1, oldmode)
	time.Sleep(time.Millisecond * 5)
	c.write8(Mode1, oldmode|uint8(Mode1Restart)|uint8(Mode1AutoIncrement))
	if c.Debug {
		log.Printf("Mode now 0x%X\n", c.read8(Mode1))
	}
}

// SetOutputMode  sets the output mode of the PCA9685 to either open drain or push pull / (totempole) true. Warning: LEDs with integrated zener diodes should only be driven in open drain mode.
func (c *Context) SetOutputMode(totempole bool) {
	oldmode := c.read8(Mode2)
	var newmode uint8
	if totempole {
		newmode = oldmode | uint8(Mode2OutDRV)
	} else {
		newmode = oldmode & ^uint8(Mode2OutDRV)
	}
	c.write8(Mode2, newmode)
	if c.Debug {
		name := "open drain"
		if totempole {
			name = "totempole"
		}
		log.Printf("setting output mode: %s by setting Mode2 to %X", name, newmode)
	}
}

// GetPWM gets the PWM output of one of the PCA9685 pins (num) Pin 0-15
func (c *Context) GetPWM(num uint8) uint8 {
	twoc, err := i2c.NewI2C(I2CAddress, c.Addr)
	defer c.Bus.Close()
	if err != nil {
		log.Fatal(err)
	}
	c.Bus = twoc
	bb, i, err := c.Bus.ReadRegBytes(LED0OnLow+(4*num), 1)
	if err != nil || i != 1 {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", bb)
	return bb[0]
}

//SetPWM sets the PWM output of one of the PCA9685 pins (num) Pin 0-15 (on/off) at what point of the 4096-part cycle to turn the pin on/off
func (c *Context) SetPWM(num uint8, on uint16, off uint16) {
	if c.Debug {
		log.Printf("Setting PWM %d: %d -> %d\n", num, on, off)
	}
	twoc, err := i2c.NewI2C(I2CAddress, c.Addr)
	defer c.Bus.Close()
	if err != nil {
		log.Fatal(err)
	}
	c.Bus = twoc
	a, err := c.Bus.WriteBytes([]byte{
		LED0OnLow + (4 * num),
		byte(on >> 8),
		byte(on & 0x00ff),
		byte(off >> 8),
		byte(off & 0x00ff),
	})
	if err != nil || a != 5 {
		log.Fatal(err)
	}
}

// SetPin helper to set pin PWM output. Sets pin without having to deal with on/off tick placement and properly handles a zero value as completely off and 4095 as completely on.  Optional invert parameter supports inverting the pulse for sinking to ground.
func (c *Context) SetPin(num uint8, val uint16, invert ...bool) {
	val = min(val, 4095)
	if len(invert) > 0 && invert[0] {
		if val == 0 {
			// Special value for signal fully on.
			c.SetPWM(num, 4096, 0)
		} else if val == 4095 {
			c.SetPWM(num, 0, 4096)
		} else {
			c.SetPWM(num, 0, 4095-val)
		}
	} else {
		if val == 4095 {
			// Special value for signal fully on.
			c.SetPWM(num, 4096, 0)
		} else if val == 0 {
			// Special value for signal fully off.
			c.SetPWM(num, 0, 4096)
		} else {
			c.SetPWM(num, 0, val)
		}
	}
}

// ReadPrescale reads set Prescale from PCA9685
func (c *Context) ReadPrescale() uint8 {
	return c.read8(PreScale)
}

// WriteMicroseconds sets the PWM output of one of the PCA9685 pins based on the input microseconds, output is not precise
func (c *Context) WriteMicroseconds(num uint8, microseconds uint16) {
	if c.Debug {
		log.Printf("Setting PWM Via Microseconds on output%d: %d->\n", num, microseconds)
	}
	var pulse = float64(microseconds)
	var pulselength float64 = 1000000 // 1,000,000 us per second

	prescale := c.ReadPrescale()

	if c.Debug {
		log.Println(prescale, " PCA9685 chip prescale")
	}

	prescale++
	pulselength = float64(prescale) * pulselength
	pulselength = pulselength / c.OsFreq

	if c.Debug {
		log.Println(pulselength, " us per bit")
	}

	pulse = pulse / pulselength

	if c.Debug {
		log.Println(pulse, " pulse for PWM")
	}

	c.SetPWM(num, 0, uint16(pulse))
}

// SetOscillatorFrequency setter for the internally tracked oscillator used for freq calculations
func (c *Context) SetOscillatorFrequency(freq uint32) {
	c.OsFreq = float64(freq)
}

// GetOscillatorFrequency getter for the internally tracked oscillator used for freq calculations
func (c *Context) GetOscillatorFrequency() uint32 {
	return uint32(c.OsFreq)
}

func (c *Context) read8(addr uint8) uint8 {
	twoc, err := i2c.NewI2C(I2CAddress, c.Addr)
	defer c.Bus.Close()
	if err != nil {
		log.Fatal(err)
	}
	c.Bus = twoc
	bb, i, err := c.Bus.ReadRegBytes(addr, 1)
	if err != nil || i != 1 {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", bb)
	return bb[0]
}

func (c *Context) write8(addr uint8, d uint8) {
	twoc, err := i2c.NewI2C(I2CAddress, c.Addr)
	defer c.Bus.Close()
	if err != nil {
		log.Fatal(err)
	}
	c.Bus = twoc
	a, err := c.Bus.WriteBytes([]byte{addr, d})
	if err != nil || a != 2 {
		log.Fatal(err)
	}
}
