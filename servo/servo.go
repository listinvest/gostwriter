package servo

import "github.com/x86ed/gostwriter/point"

const (
	// X axis of rotation
	X = iota
	// Y axis of rotation
	Y
	// Z axis of rotation
	Z
)

// Servo representation of servo joint
type Servo struct {
	Axis    int
	Address uint8
	Angle   float64
	Upper   float64
	Lower   float64
	Name    string
	Offset  point.Point
	Next    *Servo
}

// SAxis base servo of arm
var SAxis = Servo{
	Axis:    Z,
	Address: 15,
	Name:    "S-1",
	Offset:  point.Point{X: 0, Y: 36, Z: 100.1},
	Next:    &LAxis,
}

// LAxis Lower Arm Servo
var LAxis = Servo{
	Axis:    Y,
	Address: 14,
	Name:    "L-2",
	Offset:  point.Point{X: 0, Y: 0, Z: 128.2},
	Next:    &UAxis,
}

// UAxis Upper Arm Servo
var UAxis = Servo{
	Axis:    Y,
	Address: 13,
	Name:    "U-3",
	Offset:  point.Point{X: 0, Y: 22, Z: 22},
	Next:    &RAxis,
}

// RAxis Rotational Servo for upper arm
var RAxis = Servo{
	Axis:    X,
	Address: 12,
	Name:    "U-4",
	Offset:  point.Point{X: 0, Y: 105, Z: 0},
	Next:    &BAxis,
}

// BAxis wrist pitch servo
var BAxis = Servo{
	Axis:    Y,
	Address: 11,
	Name:    "B-5",
	Offset:  point.Point{X: 0, Y: 22, Z: 0},
	Next:    &TAxis,
}

// TAxis wrist rotation servo
var TAxis = Servo{
	Axis:    X,
	Address: 10,
	Name:    "T-6",
	Offset:  point.Point{X: 0, Y: 15.1, Z: 36.5},
}
