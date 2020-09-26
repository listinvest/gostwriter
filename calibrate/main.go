package main

import (
	"time"

	"github.com/d2r2/go-i2c"
	"github.com/x86ed/gostrider/pca9685"
)

const (
	ServoMin = 150
	ServoMax = 600
	USMin = 600
	USMAX = 2400
	ServoFreq = 50
)

func main() {
	var servonum uint8 = 10
	for servonum < 16 {
		if (servonum < 10) { 
			servonum = 10
			continue
		}
	itwoc, err := i2c.NewI2C(pca9685.I2CAddress,1)
	if err != nil {
		log.Fatal(err)
	}
	pwm := pca9685.Context{}
	pwm.PWMServoDriver(1,itwoc)
	pwm.Begin()
	pwm.SetOscillatorFrequency(27000000)
	pwm.SetPWMFrequency(ServoFreq)
	time.Sleep(time.Millisecond *10)

	time.Sleep(time.Millisecond * 500)
	servonum++
}
}