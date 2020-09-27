package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
	"errors"

	"github.com/d2r2/go-i2c"
	"github.com/urfave/cli/v2"
	"github.com/x86ed/gostwriter/pca9685"
)

const (
	// ServoMin minimum servo value
	ServoMin = 150
	// ServoMax maximum servo value
	ServoMax = 600
	// USMin minimum pulse size
	USMin = 600
	// USMax maximum pulse size
	USMax = 2400
	// ServoFreq servo operationg frequency
	ServoFreq = 50
)

func setServoPulse(n uint8, pulse float64, c pca9685.Context){
	var pulselength float64 = 1000000
	pulselength = pulselength/ServoFreq
	fmt.Printf("%d μs per period\n", pulselength)
	pulselength = pulselength / 4096
	fmt.Printf("%d μs per bit\n", pulselength)
	pulse = pulse * 1000000
	pulse = pulse/pulselength
	fmt.Println(pulse)
	c.SetPWM(n, 0, uint16(pulse))
}

// ServoTest routine
func ServoTest(){
	itwoc, err := i2c.NewI2C(pca9685.I2CAddress,1)
	if err != nil {
		log.Fatal(err)
	}
	pwm := pca9685.Context{Debug: true}
	pwm.PWMServoDriver(1,itwoc)
	pwm.Begin()
	pwm.SetOscillatorFrequency(27000000)
	pwm.SetPWMFrequency(ServoFreq)
	time.Sleep(time.Millisecond *10)
	for servonum := uint8(10);servonum < 16; servonum++ {
		fmt.Println("servo: ",servonum)

		for pulselen := ServoMin; pulselen < ServoMax; pulselen++{
			pwm.SetPWM(servonum, 0, uint16(pulselen))
		}

		time.Sleep(time.Millisecond * 500)

		for pulselen := ServoMax; pulselen > ServoMin; pulselen--{
			pwm.SetPWM(servonum, 0, uint16(pulselen))
		}

		time.Sleep(time.Millisecond * 500)

		// Drive each servo one at a time using writeMicroseconds(), it's not precise due to calculation rounding!
		// The writeMicroseconds() function is used to mimic the Arduino Servo library writeMicroseconds() behavior. 
		for microsec := USMin; microsec < USMax; microsec++ {
			pwm.WriteMicroseconds(servonum, uint16(microsec));
		}

		time.Sleep(time.Millisecond * 500)
		for microsec := USMax; microsec > USMin; microsec-- {
			pwm.WriteMicroseconds(servonum, uint16(microsec));
		}

		time.Sleep(time.Millisecond * 500)
		if (servonum < 10) { 
			servonum = 10
		}
	}
	pwm.Bus.Close()
}

// ServoSet set an individual servo's value
func ServoSet(servo uint8, val uint16){
	itwoc, err := i2c.NewI2C(pca9685.I2CAddress,1)
	if err != nil {
		log.Fatal(err)
	}
	pwm := pca9685.Context{Debug: true}
	pwm.PWMServoDriver(1,itwoc)
	pwm.Begin()
	pwm.SetOscillatorFrequency(27000000)
	pwm.SetPWMFrequency(ServoFreq)
	time.Sleep(time.Millisecond *10)
	pwm.SetPWM(servo, 0, val)
	pwm.Bus.Close()
}

// Reset the I2C chip
func Reset(){
	itwoc, err := i2c.NewI2C(pca9685.I2CAddress,1)
	if err != nil {
		log.Fatal(err)
	}
	pwm := pca9685.Context{Debug: true}
	pwm.PWMServoDriver(1,itwoc)
	pwm.Begin()
	pwm.SetOscillatorFrequency(27000000)
	pwm.SetPWMFrequency(ServoFreq)
	time.Sleep(time.Millisecond *10)
	pwm.Reset()
	pwm.Bus.Close()
}

// Sleep the I2C chip
func Sleep(){
	itwoc, err := i2c.NewI2C(pca9685.I2CAddress,1)
	if err != nil {
		log.Fatal(err)
	}
	pwm := pca9685.Context{Debug: true}
	pwm.PWMServoDriver(1,itwoc)
	pwm.Begin()
	pwm.SetOscillatorFrequency(27000000)
	pwm.SetPWMFrequency(ServoFreq)
	time.Sleep(time.Millisecond *10)
	pwm.Sleep()
	pwm.Bus.Close()
}

// WakeUp the I2C chip
func WakeUp(){
	itwoc, err := i2c.NewI2C(pca9685.I2CAddress,1)
	if err != nil {
		log.Fatal(err)
	}
	pwm := pca9685.Context{Debug: true}
	pwm.PWMServoDriver(1,itwoc)
	pwm.Begin()
	pwm.SetOscillatorFrequency(27000000)
	pwm.SetPWMFrequency(ServoFreq)
	time.Sleep(time.Millisecond *10)
	pwm.WakeUp()
	pwm.Bus.Close()
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
		  {
			Name:    "ServoTest",
			Aliases: []string{"st"},
			Usage:   "Run the servo test program",
			Action:  func(c *cli.Context) error {
				ServoTest()
			  return nil
			},
		  },
		  {
			Name:    "ServoSet",
			Aliases: []string{"ss"},
			Usage:   "set a servo to a set value",
			Action:  func(c *cli.Context) error {
				if c.Args().Len() < 2 {
					log.Fatalln("Not enough arguments.")
					return errors.New("Not enough arguments")
				}
				srv, err :=  uint8(strconv.Atoi(c.Args().Get(0)))
				if err != nil {
					log.Fatalln(err)
					return err
				}
				val, err := uint16(strconv.Atoi(c.Args().Get(1)))
				if err != nil {
					log.Fatalln(err)
					return err
				}
				ServoSet(srv,val)
			  return nil
			},
		  },
		  {
			Name:    "Reset",
			Aliases: []string{"r"},
			Usage:   "reset the device",
			Action:  func(c *cli.Context) error {
				Reset()
			  return nil
			},
		  },
		  {
			Name:    "Sleep",
			Aliases: []string{"s"},
			Usage:   "sleep the device",
			Action:  func(c *cli.Context) error {
				Sleep()
			  return nil
			},
		  },
		  {
			Name:    "Wake",
			Aliases: []string{"w"},
			Usage:   "wake the device",
			Action:  func(c *cli.Context) error {
				WakeUp()
			  return nil
			},
		  },
		},
	  }
	
	  sort.Sort(cli.CommandsByName(app.Commands))
	
	  err := app.Run(os.Args)
	  if err != nil {
		log.Fatal(err)
	  }
}