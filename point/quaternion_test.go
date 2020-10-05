package point

import (
	"math"
	"testing"
)

func TestFromEuler(t *testing.T) {
	p := Point{45 * DegToRad, 30 * DegToRad, 60 * DegToRad}
	q := Quaternion{}
	q.FromEuler(p)
	if math.Abs(q.W)-0.822 > 0.001 || math.Abs(q.X)-0.200 > 0.001 || math.Abs(q.Y)-0.391 > 0.001 || math.Abs(q.Z)-0.360 > 0.001 {
		t.Errorf("Wanted {W:0.8223631719059993 X:0.200562121146575 Y:0.39190383732911993 Z:0.36042340565035597} got %+v", q)
	}
}

func TestToEuler(t *testing.T) {
	q := Quaternion{W: 0.8223631719059993, X: 0.200562121146575, Y: 0.39190383732911993, Z: 0.36042340565035597}
	p := q.ToEulerAngles()
	if math.Abs(p.X*RadToDeg)-45 > 0.001 || math.Abs(p.Y*RadToDeg)-30 > 0.001 || math.Abs(p.Z*RadToDeg)-60 > 0.001 {
		t.Errorf("Wanted {W:0.8223631719059993 X:0.200562121146575 Y:0.39190383732911993 Z:0.36042340565035597} got %+v", p)
	}
}

func TestQuadNormalize(t *testing.T) {
	q := Quaternion{W: 0.8223631719059993, X: 0.200562121146575, Y: 0.39190383732911993, Z: 0.36042340565035597}
	q.Normalize()
	if math.Abs(q.W)-0.822 > 0.001 || math.Abs(q.X)-0.200 > 0.001 || math.Abs(q.Y)-0.391 > 0.001 || math.Abs(q.Z)-0.360 > 0.001 {
		t.Errorf("Wanted {W:0.8223631719059994 X:0.20056212114657504 Y:0.39190383732912 Z:0.360423405650356} got %+v", q)
	}
}

func TestQuadmagnitude(t *testing.T) {
	q := Quaternion{W: 0.8223631719059993, X: 0.200562121146575, Y: 0.39190383732911993, Z: 0.36042340565035597}
	out := q.Magnitude()
	if out-1 > 0.001 {
		t.Errorf("Wanted 1 got %+v", out)
	}
}

func TestQuadInvert(t *testing.T) {
	q := Quaternion{W: 0.8223631719059993, X: 0.200562121146575, Y: 0.39190383732911993, Z: 0.36042340565035597}
	q2 := Quaternion{W: 0.8223631719059993, X: 0.200562121146575, Y: 0.39190383732911993, Z: 0.36042340565035597}
	q.Invert()
	if q.W+q2.W > 0.001 || q.X+q2.X > 0.001 || q.Y+q2.Y > 0.001 || q.Z+q2.Z > 0.001 {
		t.Errorf("Wanted {W: -0.8223631719059993, X: -0.200562121146575, Y: -0.39190383732911993, Z: -0.36042340565035597} got %+v", q)
	}
}

func TestQuadConj(t *testing.T) {
	q := Quaternion{W: 0.8223631719059993, X: 0.200562121146575, Y: 0.39190383732911993, Z: 0.36042340565035597}
	q2 := Quaternion{W: 0.8223631719059993, X: 0.200562121146575, Y: 0.39190383732911993, Z: 0.36042340565035597}
	q.Conjugate()
	if q.W-q2.W > 0.001 || q.X+q2.X > 0.001 || q.Y+q2.Y > 0.001 || q.Z+q2.Z > 0.001 {
		t.Errorf("Wanted {W: -0.8223631719059993, X: -0.200562121146575, Y: -0.39190383732911993, Z: -0.36042340565035597} got %+v", q)
	}
}

func TestQuadInverse(t *testing.T) {
	q := Quaternion{W: 0.8223631719059993, X: 0.200562121146575, Y: 0.39190383732911993, Z: 0.36042340565035597}
	q2 := Quaternion{W: 0.8223631719059993, X: 0.200562121146575, Y: 0.39190383732911993, Z: 0.36042340565035597}
	q.Inverse()
	if q.W-q2.W > 0.001 || q.X+q2.X > 0.001 || q.Y+q2.Y > 0.001 || q.Z+q2.Z > 0.001 {
		t.Errorf("Wanted {W: -0.8223631719059993, X: -0.200562121146575, Y: -0.39190383732911993, Z: -0.36042340565035597} got %+v", q)
	}
}

func TestQuadMultiply(t *testing.T) {
	q := Quaternion{W: 0.8223631719059993, X: 0.200562121146575, Y: 0.39190383732911993, Z: 0.36042340565035597}
	p := Point{30 * DegToRad, 30 * DegToRad, 30 * DegToRad}
	q2 := Quaternion{}
	q2.FromEuler(p)
	q.Multiply(q2)
	p = q.ToEulerAngles()
	if math.Abs(p.X*RadToDeg)-89.1 > 0.01 || math.Abs(p.Y*RadToDeg)-23.15 > 0.01 || math.Abs(p.Z*RadToDeg)-107 > 0.01 {
		t.Errorf("Wanted {X:1.555183740711392 Y:0.40412421382012875 Z:1.8589623465872065} got %+v", p)
	}
}

func TestQuadDotProduct(t *testing.T) {
	q := Quaternion{0, 1, 2, 3}
	q2 := Quaternion{3, 2, 1, 0}
	d := q.DotProduct(q2)
	if d != 4 {
		t.Errorf("Wanted 4 got %+v", d)
	}
}

func TestQuadSimilar(t *testing.T) {
	q := Quaternion{-20, 1, 2, 3}
	q2 := Quaternion{3, 2, 1, 0}
	q3 := Quaternion{3, 2, 1, 0}
	d := q.Similar(q2, .001)
	d2 := q2.Similar(q3, .001)
	if !d {
		t.Errorf("Wanted 4 got %+v", d)
	}
	if d2 {
		t.Errorf("Wanted 4 got %+v", d)
	}
}
