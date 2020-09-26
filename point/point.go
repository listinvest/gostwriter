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

// RotX rotate point around X axis
func (p *Point) RotX(θ float64) {
	p.Y = (p.Y * math.Cos(θ)) - (p.Z * math.Sin(θ))
	p.Z = (p.Y * math.Sin(θ)) + (p.Z * math.Cos(θ))
}

// RotY rotate point around Y axis
func (p *Point) RotY(θ float64) {
	p.X = (p.X * math.Cos(θ)) + (p.Z * math.Sin(θ))
	p.Z = (p.Z * math.Cos(θ)) - (p.X * math.Sin(θ))
}

// RotZ rotate point around Z axis
func (p *Point) RotZ(θ float64) {
	p.X = (p.X * math.Cos(θ)) - (p.Y * math.Sin(θ))
	p.Y = (p.X * math.Sin(θ)) + (p.Y * math.Cos(θ))
}
