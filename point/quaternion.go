package point

import (
	"math"
)

// Quaternion in struct form
type Quaternion struct {
	W, X, Y, Z float64
}

// FromEuler converts Euler angles into a Quaternion
func (q *Quaternion) FromEuler(p Point) {
	ϕ := p.X // roll
	θ := p.Y // pitch
	ψ := p.Z // yaw

	cy := math.Cos(ψ * 0.5)
	sy := math.Sin(ψ * 0.5)
	cp := math.Cos(θ * 0.5)
	sp := math.Sin(θ * 0.5)
	cr := math.Cos(ϕ * 0.5)
	sr := math.Sin(ϕ * 0.5)
	q.W = cr*cp*cy + sr*sp*sy
	q.X = sr*cp*cy - cr*sp*sy
	q.Y = cr*sp*cy + sr*cp*sy
	q.Z = cr*cp*sy - sr*sp*cy
}

// ToEulerAngles converts a Quaternion to Euler Angles
func (q *Quaternion) ToEulerAngles() Point {
	var out Point

	// roll (x-axis rotation)
	sinrcosp := 2 * (q.W*q.X + q.Y*q.Z)
	cosrcosp := 1 - 2*(q.X*q.X+q.Y*q.Y)
	out.X = math.Atan2(sinrcosp, cosrcosp)

	// pitch (y-axis rotation)
	sinp := 2 * (q.W*q.Y - q.Z*q.X)
	if math.Abs(sinp) >= 1 {
		out.Y = math.Copysign(math.Pi/2, sinp) // use 90 degrees if out of range
	} else {
		out.Y = math.Sin(sinp)
	}

	// yaw (z-axis rotation)
	sinycosp := 2 * (q.W*q.Z + q.X*q.Y)
	cosycosp := 1 - 2*(q.Y*q.Y+q.Z*q.Z)
	out.Z = math.Atan2(sinycosp, cosycosp)

	return out
}

// Normalize normalizes a quaternion
func (q *Quaternion) Normalize() {
	m := q.Magnitude()
	q.X /= m
	q.Y /= m
	q.Z /= m
	q.W /= m
}

// Magnitude gives q's magnitude
func (q *Quaternion) Magnitude() float64 {
	return math.Sqrt(q.X*q.X + q.Y*q.Y + q.Z*q.Z + q.W*q.W)
}

// Invert a given quaternion
func (q *Quaternion) Invert() {
	q.W = -q.W
	q.X = -q.X
	q.Y = -q.Y
	q.Z = -q.Z
}

// Inverse of a quaternion
func (q *Quaternion) Inverse() {
	m := q.Magnitude()
	q.Conjugate()
	q.X /= m
	q.Y /= m
	q.Z /= m
	q.W /= m
}

// Conjugate a given quaternion
func (q *Quaternion) Conjugate() {
	q.X = -q.X
	q.Y = -q.Y
	q.Z = -q.Z
}

// Multiply Multiplies q x q2 in that order
func (q *Quaternion) Multiply(q2 Quaternion) {
	q1 := *q
	q.X = q1.X*q2.W + q1.Y*q2.Z - q1.Z*q2.Y + q1.W*q2.X
	q.Y = -q1.X*q2.Z + q1.Y*q2.W + q1.Z*q2.X + q1.W*q2.Y
	q.Z = q1.X*q2.Y - q1.Y*q2.X + q1.Z*q2.W + q1.W*q2.Z
	q.W = -q1.X*q2.X - q1.Y*q2.Y - q1.Z*q2.Z + q1.W*q2.W
}

// DotProduct of q . q2
func (q *Quaternion) DotProduct(q2 Quaternion) float64 {
	return (q.W * q2.W) + (q.X * q2.X) + (q.Y * q2.Y) + (q.Z * q2.Z)
}

// Similar returns true if (q) & (q2) are within (moe) of eachother
func (q *Quaternion) Similar(q2 Quaternion, moe float64) bool {
	if q.DotProduct(q2) < moe {
		return true
	}
	return false
}

// QuatRotate rotates a Quaternion on (axis) with (angle)
func QuatRotate(angle float64, axis Point) Quaternion {

	c, s := float64(math.Cos(float64(angle/2))), float64(math.Sin(float64(angle/2)))
	axis.Mul(s)
	return Quaternion{c, axis.X, axis.Y, axis.Z}
}

// RotationBetweenVectors averages (start) with (dest)
func RotationBetweenVectors(start Point, dest Point) Quaternion {
	start.Normalize()
	dest.Normalize()

	cosTheta := start.DotProduct(dest)
	var rotationAxis Point

	if cosTheta < -1+0.001 {
		// special case when vectors in opposite directions:
		// there is no "ideal" rotation axis
		// So guess one; any will do as long as it's perpendicular to start
		rotationAxis.CrossProduct(Point{}, start)
		if rotationAxis.Magnitude() < 0.01 { // bad luck, they were parallel, try again!
			rotationAxis.CrossProduct(Point{1, 0, 0}, start)
		}

		rotationAxis.Normalize()
		return QuatRotate(180.0*DegToRad, rotationAxis)
	}

	rotationAxis.CrossProduct(start, dest)

	s := math.Sqrt((1 + cosTheta) * 2)
	invs := 1 / s

	return Quaternion{
		W: s * 0.5,
		X: rotationAxis.X * invs,
		Y: rotationAxis.Y * invs,
		Z: rotationAxis.Z * invs,
	}

}
