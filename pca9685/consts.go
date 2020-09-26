// Package pca9685 is an addapter for the PCA9685 16-channel PWM & Servo Driver
//  Pick one up today in the adafruit shop!
//  ------> https://www.adafruit.com/product/815
//
// These driver use I2C to communicate, 2 pins are required to interface.
package pca9685

// REGISTER ADDRESSES
const (
	// Mode1 Mode Register 1
	Mode1 = iota
	// Mode2 Mode Register 2
	Mode2
	// I2C-bus subaddress 1
	SubAddr1
	// I2C-bus subaddress 2
	SubAddr2
	// I2C-bus subaddress 3
	SubAddr3
	// LED All Call I2C-bus address
	AllCallAddr
	// LED0 on tick, low byte
	LED0OnLow
	// LED0 on tick, high byte
	LED0OnHigh
	// LED0 off tick, low byte
	LED0OffLow
	// LED0 off tick, high byte
	LED0OffHigh
	// LED1 on tick, low byte
	LED1OnLow
	// LED1 on tick, high byte
	LED1OnHigh
	// LED1 off tick, low byte
	LED1OffLow
	// LED1 off tick, high byte
	LED1OffHigh
	// LED2 on tick, low byte
	LED2OnLow
	// LED2 on tick, high byte
	LED2OnHigh
	// LED2 off tick, low byte
	LED2OffLow
	// LED2 off tick, high byte
	LED2OffHigh
	// LED3 on tick, low byte
	LED3OnLow
	// LED3 on tick, high byte
	LED3OnHigh
	// LED3 off tick, low byte
	LED3OffLow
	// LED3 off tick, high byte
	LED3OffHigh
	// LED4 on tick, low byte
	LED4OnLow
	// LED4 on tick, high byte
	LED4OnHigh
	// LED4 off tick, low byte
	LED4OffLow
	// LED4 off tick, high byte
	LED4OffHigh
	// LED5 on tick, low byte
	LED5OnLow
	// LED5 on tick, high byte
	LED5OnHigh
	// LED5 off tick, low byte
	LED5OffLow
	// LED5 off tick, high byte
	LED5OffHigh
	// LED6 on tick, low byte
	LED6OnLow
	// LED6 on tick, high byte
	LED6OnHigh
	// LED6 off tick, low byte
	LED6OffLow
	// LED6 off tick, high byte
	LED6OffHigh
	// LED7 on tick, low byte
	LED7OnLow
	// LED7 on tick, high byte
	LED7OnHigh
	// LED7 off tick, low byte
	LED7OffLow
	// LED7 off tick, high byte
	LED7OffHigh
	// LED8 on tick, low byte
	LED8OnLow
	// LED8 on tick, high byte
	LED8OnHigh
	// LED8 off tick, low byte
	LED8OffLow
	// LED8 off tick, high byte
	LED8OffHigh
	// LED9 on tick, low byte
	LED9OnLow
	// LED9 on tick, high byte
	LED9OnHigh
	// LED9 off tick, low byte
	LED9OffLow
	// LED9 off tick, high byte
	LED9OffHigh
	// LED10 on tick, low byte
	LED10OnLow
	// LED10 on tick, high byte
	LED10OnHigh
	// LED10 off tick, low byte
	LED10OffLow
	// LED10 off tick, high byte
	LED10OffHigh
	// LED11 on tick, low byte
	LED11OnLow
	// LED11 on tick, high byte
	LED11OnHigh
	// LED11 off tick, low byte
	LED11OffLow
	// LED11 off tick, high byte
	LED11OffHigh
	// LED12 on tick, low byte
	LED12OnLow
	// LED12 on tick, high byte
	LED12OnHigh
	// LED12 off tick, low byte
	LED12OffLow
	// LED12 off tick, high byte
	LED12OffHigh
	// LED13 on tick, low byte
	LED13OnLow
	// LED13 on tick, high byte
	LED13OnHigh
	// LED13 off tick, low byte
	LED13OffLow
	// LED13 off tick, high byte
	LED13OffHigh
	// LED14 on tick, low byte
	LED14OnLow
	// LED14 on tick, high byte
	LED14OnHigh
	// LED14 off tick, low byte
	LED14OffLow
	// LED14 off tick, high byte
	LED14OffHigh
	// LED15 on tick, low byte
	LED15OnLow
	// LED15 on tick, high byte
	LED15OnHigh
	// LED15 off tick, low byte
	LED15OffLow
	// LED15 off tick, high byte
	LED15OffHigh
)

// AllLEDOnLow load all the LEDn_ON registers, low
const AllLEDOnLow = 0xFA

// AllLEDOnHigh load all the LEDn_ON registers, high
const AllLEDOnHigh = 0xFB

// AllLEDOffLow load all the LEDn_OFF registers, low
const AllLEDOffLow = 0xFC

// AllLEDOffHigh load all the LEDn_OFF registers,high
const AllLEDOffHigh = 0xFD

// PreScale Prescaler for PWM output frequency
const PreScale = 0xFE

// TestMode defines the test mode to be entered
const TestMode = 0xFF

// MODE1 bits

// Mode1AllCall respond to LED All Call I2C-bus address
const Mode1AllCall = 0x01

// Mode1SUB3 respond to I2C-bus subaddress 3
const Mode1SUB3 = 0x02

// Mode1SUB2 respond to I2C-bus subaddress 2
const Mode1SUB2 = 0x04

// Mode1SUB1 respond to I2C-bus subaddress 1
const Mode1SUB1 = 0x08

// Mode1Sleep Low power mode. Oscillator off
const Mode1Sleep = 0x10

// Mode1AutoIncrement Auto-Increment enabled
const Mode1AutoIncrement = 0x20

// Mode1EXTCLK Use EXTCLK pin clock
const Mode1EXTCLK = 0x40

// Mode1Restart Restart enabled
const Mode1Restart = 0x80

// MODE2 bits

// Mode2OutNE0 Active LOW output enable input
const Mode2OutNE0 = 0x01

// Mode2OutNE1 Active LOW output enable input - high impedience
const Mode2OutNE1 = 0x02

// Mode2OutDRV totem pole structure vs open-drain
const Mode2OutDRV = 0x04

// Mode2OCH Outputs change on ACK vs STOP
const Mode2OCH = 0x08

// Mode2INVRT Output logic state inverted
const Mode2INVRT = 0x10

// I2CAddress Default PCA9685 I2C Slave Address
const I2CAddress = 0x40

// FrequencyOscillator Int. osc. frequency in datasheet
const FrequencyOscillator = 25000000

// PrescaleMin minimum prescale value
const PrescaleMin = 3.0

// PrescaleMax maximum prescale value
const PrescaleMax = 255.0
