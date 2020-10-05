package fabrik

import (
	"math"

	"github.com/x86ed/gostwriter/point"
	"github.com/x86ed/gostwriter/servo"
)

// FABRIK control struct
type FABRIK struct {
	Origin      point.Point
	EndEffector point.Point
	P2          servo.Servo
	P           servo.Servo
	Seed        servo.Servo
	Forward     bool
	MOE         float64 //Margin of Error
}

// MoveTo moves the end effector to a target point
func (f *FABRIK) MoveTo(target point.Point) servo.Servo {
	f.GetEffectorPos()
	f.SetPrimes()
	target = f.ConstrainToEnvelope(target)

	var cancel bool
	for !cancel {
		if f.InMOE(target, f.EndEffector) && f.InMOE(target, f.Origin) {
			cancel = true
		}
		f.Cycle(target)
	}
	if f.Forward {
		f.Seed = f.P2
	} else {
		f.Seed = f.P
	}
	f.P = servo.Servo{}
	f.P2 = servo.Servo{}
	return f.Seed
}

// Cycle iterates through the nodes to change the orentation of the joints
func (f *FABRIK) Cycle(t point.Point) {
	f.Forward = !f.Forward
	if f.Forward {
		//	offset := f.Origin
	}
}

// Closest calculates the closest position the offset cn be moved to given a (h)eading (o)rigin and (t)arget
func (f *FABRIK) Closest(s servo.Servo, h, o, t point.Point) (new, heading point.Point) {
	pos := o
	switch s.Axis {
	case servo.X:

		corrected := s.Angle + h.X
		new := s.Offset
		new.RotX(corrected)
		pos.Add(new)
		h.X = corrected
	case servo.Y:
		corrected := s.Angle + h.Y
		new := s.Offset
		new.RotY(corrected)
		pos.Add(new)
		h.Y = corrected
	case servo.Z:
		corrected := s.Angle + h.Z
		new := s.Offset
		new.RotZ(corrected)
		pos.Add(new)
		h.Z += corrected
	default:
		pos.Add(s.Offset)
	}
	new = pos
	heading = h
	return
}

// ConstrainToEnvelope Constrains the target to the active operation envelope
func (f *FABRIK) ConstrainToEnvelope(t point.Point) point.Point {
	return t
}

// InMOE determines if point t is within the Margin of Error of point b
func (f *FABRIK) InMOE(t, b point.Point) bool {
	if math.Pow(t.X-b.X, 2)+math.Pow(t.Y-b.Y, 2)+math.Pow(t.Z-b.Z, 2) > math.Pow(f.MOE, 2) {
		return false
	}
	return true
}

// SetPrimes sets the prime shadow arms values
func (f *FABRIK) SetPrimes() int {
	emp := servo.Servo{}
	var count int
	if f.Seed != emp {
		f.P2 = f.Seed
		i := f.Seed
		f.P = i
		f.P.Next = nil
		count++
		for i.Next != nil {
			i = *i.Next
			n := f.P
			f.P = i
			f.P.Next = &n
			count++
		}
	}
	return count
}

// GetEffectorPos gets the end effector position
func (f *FABRIK) GetEffectorPos() point.Point {
	if f.EndEffector.IsZero() {
		j := f.Seed
		pos := f.Origin
		var heading point.Point
		for j.Next != nil {
			switch j.Axis {
			case servo.X:
				corrected := j.Angle + heading.X
				new := j.Offset
				new.RotX(corrected)
				pos.Add(new)
				heading.X = corrected
			case servo.Y:
				corrected := j.Angle + heading.Y
				new := j.Offset
				new.RotY(corrected)
				pos.Add(new)
				heading.Y = corrected
			case servo.Z:
				corrected := j.Angle + heading.Z
				new := j.Offset
				new.RotZ(corrected)
				pos.Add(new)
				heading.Z += corrected
			default:
				pos.Add(j.Offset)
			}
			j = *j.Next
		}
		f.EndEffector = pos
	}
	return f.EndEffector
}
