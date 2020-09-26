package fabrik

import (
	"github.com/x86ed/gostrider/point"
	"github.com/x86ed/gostrider/servo"
)

// FABRIK control struct
type FABRIK struct {
	Origin      point.Point
	EndEffector point.Point
	P2          []servo.Servo
	P           []servo.Servo
	Seed        []servo.Servo
	Forward     bool
	MOE         float64
}

func (f *FABRIK) Update() []servo.Servo {
	return f.Seed
}

func (f *FABRIK) MaxEnv() point.Point {
	out := point.Point{}
	// for _, v := range f.Seed {
	// 	v.
	// }
	return out
}
