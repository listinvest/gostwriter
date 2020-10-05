package point

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	p := Point{1, 2, 3}
	p2 := Point{3, 2, 1}
	p.Add(p2)
	if p.X != 4 || p.Y != 4 || p.Z != 4 {
		t.Errorf("Wanted {4,4,4} got %+v", p)
	}
}

func TestSub(t *testing.T) {
	p := Point{1, 2, 3}
	p2 := Point{3, 2, 1}
	p.Sub(p2)
	if p.X != -2 || p.Y != 0 || p.Z != 2 {
		t.Errorf("Wanted {4,4,4} got %+v", p)
	}
}

func TestMul(t *testing.T) {
	p := Point{1, 2, 3}
	p.Mul(5)
	if p.X != 5 || p.Y != 10 || p.Z != 15 {
		t.Errorf("Wanted {5,10,15} got %+v", p)
	}
}

func TestRotX(t *testing.T) {
	p := Point{0, 36, 228.3}
	p2 := Point{0, 36, 100.1}
	p.Sub(p2)
	p.RotX(-90 * DegToRad)
	p.Add(p2)
	if math.Abs(p.X-0) > 0.001 || math.Abs(p.Y-164.2) > 0.001 || math.Abs(p.Z-100.1) > 0.001 {
		t.Errorf("Wanted {0,164.2,100.1} got %+v", p)
	}
}

func TestRotY(t *testing.T) {
	p := Point{0, 36, 228.3}
	p2 := Point{0, 36, 100.1}
	p.Sub(p2)
	t.Log(p)
	p.RotY(90 * DegToRad)
	t.Log(p)
	p.Add(p2)
	if math.Abs(p.X-128.2) > 0.001 || math.Abs(p.Y-36) > 0.001 || math.Abs(p.Z-100.1) > 0.001 {
		t.Errorf("Wanted {128,36,100.1} got %+v", p)
	}
}

func TestRotZ(t *testing.T) {
	p := Point{0, 36, 100.1}
	p.RotZ(90 * DegToRad)
	if math.Abs(p.X-128.2) > 0.001 || math.Abs(p.Y-36) > 0.001 || math.Abs(p.Z)-100.1 > 0.001 {
		t.Errorf("Wanted {-36,0,100.1} got %+v", p)
	}
}

func TestIsZero(t *testing.T) {
	p := Point{0, 36, 100.1}
	p2 := Point{}
	if !p2.IsZero() {
		t.Errorf("Wanted {0,0,0} got %+v", p)
	}
	if p.IsZero() {
		t.Errorf("Wanted {0, 36, 100.1} got %+v", p)
	}
}

func TestNormalize(t *testing.T) {
	p := Point{1, 2, 3}
	p.Normalize()
	if math.Abs(p.X-0.2672) > 0.001 || math.Abs(p.Y)-0.5345 > 0.001 || math.Abs(p.Z)-0.8017 > 0.001 {
		t.Errorf("wanted {X:0.2672612419124244 Y:0.5345224838248488 Z:0.8017837257372732} got %+v", p)
	}
}

func TestMagnitude(t *testing.T) {
	p := Point{1, 2, 3}
	out := p.Magnitude()
	if out != 3.7416573867739413 {
		t.Errorf("wanted 3.7416573867739413 got %+v", out)
	}
}

func TestDotProduct(t *testing.T) {
	p := Point{1, 2, 3}
	p2 := Point{3, 2, 1}
	d := p.DotProduct(p2)
	if d != 10 {
		t.Errorf("Wanted 10 got %+v", d)
	}
}

func TestCrossProduct(t *testing.T) {
	p := Point{1, 2, 3}
	p2 := Point{3, 2, 1}
	out := Point{}
	out.CrossProduct(p, p2)
	if math.Abs(out.X)-4 > 0.001 || math.Abs(out.Y)-8 > 0.001 || math.Abs(out.Z)-4 > 0.001 {
		t.Errorf("wanted {-4, 8 -4} got %+v", out)
	}
}

func TestPointToEuler(t *testing.T) {
	p := Point{4, 4, 4}
	p.ToEuler()
	t.Errorf("wanted {-4, 8 -4} got %+v", p)
}
