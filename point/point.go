package point

import "math"

// Point in 3D space
type Point struct {
	X, Y, Z float64
}

// Add adds two points
func (p *Point) Add(p2 Point) {
	p.X += p2.X
	p.Y += p2.Y
	p.Z += p2.Z
}

// Sub subtracts two points
func (p *Point) Sub(p2 Point) {
	p.X -= p2.X
	p.Y -= p2.Y
	p.Z -= p2.Z
}

// Mul multiplies a point by a scalar
func (p *Point) Mul(s float64) {
	p.X *= s
	p.Y *= s
	p.Z *= s
}

// RotX rotate point around X axis (roll)
func (p *Point) RotX(ϕ float64) {
	tp := *p
	p.Y = (tp.Y * math.Cos(ϕ)) - (tp.Z * math.Sin(ϕ))
	p.Z = (tp.Y * math.Sin(ϕ)) + (tp.Z * math.Cos(ϕ))
}

// RotY rotate point around Y axis (pitch)
func (p *Point) RotY(θ float64) {
	p.Z = (p.Z * math.Cos(θ)) - (p.X * math.Sin(θ))
	p.X = (p.Z * math.Sin(θ)) + (p.X * math.Cos(θ))
}

// RotZ rotate point around Z axis (yaw)
func (p *Point) RotZ(ψ float64) {
	p.Y = (p.X * math.Cos(ψ)) - (p.Y * math.Sin(ψ))
	p.X = (p.X * math.Sin(ψ)) + (p.Y * math.Cos(ψ))
}

// IsZero lets you know if a point is zeroed
func (p *Point) IsZero() bool {
	if p.X == p.Y && p.Z == p.Y && p.Y == 0 {
		return true
	}
	return false
}

// Normalize normalizes a Point
func (p *Point) Normalize() {
	m := p.Magnitude()
	p.X /= m
	p.Y /= m
	p.Z /= m
}

// Magnitude gives p's magnitude
func (p *Point) Magnitude() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

// DotProduct of p . p2
func (p *Point) DotProduct(p2 Point) float64 {
	return (p.X * p2.X) + (p.Y * p2.Y) + (p.Z * p2.Z)
}

// CrossProduct the cross product of two points
func (p *Point) CrossProduct(p1, p2 Point) {
	p.X = (p1.Y * p2.Z) - (p1.Z * p2.Y)
	p.Y = (p1.Z * p2.X) - (p1.X * p2.Z)
	p.Z = (p1.X * p2.Y) - (p1.Y * p2.X)
}

// ToEuler converts a cartesian plot to euler angles
func (p *Point) ToEuler() {
	new := Point{}
	new.Y = math.Atan2(p.Z, math.Sqrt((p.X*p.X)+(p.Y*p.Y)))
	new.Z = math.Atan2(p.X, p.Y)
	p = &new
}
