package main

import (
	"fmt"
	"time"
	"log"

	"github.com/d2r2/go-i2c"
	"github.com/x86ed/gostwriter/pca9685"
)

const (
	ServoMin = 150
	ServoMax = 600
	USMin = 600
	USMax = 2400
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

func main() {
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
	for servonum := 10;servonum < 16; servonum++ {
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