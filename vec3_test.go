package raytracer_test

import (
	"math"
	"testing"

	r "github.com/derek-schaefer/raytracer"
)

func TestNewVec3(t *testing.T) {
	v := r.NewVec3(1, 2, 3)

	if v.X() != 1 {
		t.Fail()
	}

	if v.Y() != 2 {
		t.Fail()
	}

	if v.Z() != 3 {
		t.Fail()
	}
}

func TestRandomVec3(t *testing.T) {
	min := 0.0
	max := 1.0

	v := r.RandomVec3()

	x := v.X()
	y := v.Y()
	z := v.Z()

	if !(min <= x && x <= max) {
		t.Fail()
	}

	if !(min <= y && y <= max) {
		t.Fail()
	}

	if !(min <= z && z <= max) {
		t.Fail()
	}
}

func TestRandomUnitSphereVec3(t *testing.T) {
	v := r.RandomUnitSphereVec3()

	if v.LengthSquared() >= 1 {
		t.Fail()
	}
}

func TestRandomUnitVec3(t *testing.T) {
	v := r.RandomUnitVec3()

	f := v.Length()

	if !r.NearlyEqual(f, 1) {
		t.Fail()
	}

	for i := 0; i < len(v); i++ {
		if !(-1 <= v[i] && v[i] <= 1) {
			t.Fail()
		}
	}
}

func TestRandomHemisphereVec3(t *testing.T) {
	s := r.RandomUnitSphereVec3()

	v := r.RandomHemisphereVec3(s)

	if v.Dot(s) <= 0 {
		t.Fail()
	}
}

func TestRandomRangeVec3(t *testing.T) {
	min := 2.0
	max := 3.0

	v := r.RandomRangeVec3(min, max)

	x := v.X()
	y := v.Y()
	z := v.Z()

	if !(min <= x && x <= max) {
		t.Fail()
	}

	if !(min <= y && y <= max) {
		t.Fail()
	}

	if !(min <= z && z <= max) {
		t.Fail()
	}
}

func TestVec3X(t *testing.T) {
	v := r.Vec3{1, 0, 0}

	if v.X() != 1 {
		t.Fail()
	}
}

func TestVec3Y(t *testing.T) {
	v := r.Vec3{0, 1, 0}

	if v.Y() != 1 {
		t.Fail()
	}
}

func TestVec3Z(t *testing.T) {
	v := r.Vec3{0, 0, 1}

	if v.Z() != 1 {
		t.Fail()
	}
}

func TestVec3SetY(t *testing.T) {
	v := r.Vec3{0, 0, 0}

	v.SetY(1)

	if v.Y() != 1 {
		t.Fail()
	}
}

func TestVec3SetZ(t *testing.T) {
	v := r.Vec3{0, 0, 0}

	v.SetZ(1)

	if v.Z() != 1 {
		t.Fail()
	}
}

func TestVec3SetX(t *testing.T) {
	v := r.Vec3{0, 0, 0}

	v.SetX(1)

	if v.X() != 1 {
		t.Fail()
	}
}

func TestVec3Add(t *testing.T) {
	v1 := r.Vec3{1, 2, 3}
	v2 := r.Vec3{1, 1, 1}

	v3 := v1.Add(v2)

	if v3 != (r.Vec3{2, 3, 4}) {
		t.Fail()
	}
}

func TestVec3Subtract(t *testing.T) {
	v1 := r.Vec3{1, 2, 3}
	v2 := r.Vec3{1, 1, 1}

	v3 := v1.Subtract(v2)

	if v3 != (r.Vec3{0, 1, 2}) {
		t.Fail()
	}
}

func TestVec3Multiply(t *testing.T) {
	v1 := r.Vec3{1, 2, 3}

	v2 := v1.Multiply(2)

	if v2 != (r.Vec3{2, 4, 6}) {
		t.Fail()
	}
}

func TestVec3MultiplyV(t *testing.T) {
	v1 := r.NewVec3(2, 3, 4)

	v2 := r.NewVec3(3, 4, 5)

	if v1.MultiplyV(v2) != r.NewVec3(6, 12, 20) {
		t.Fail()
	}
}

func TestVec3Divide(t *testing.T) {
	v1 := r.Vec3{2, 4, 8}

	v2 := v1.Divide(2)

	if v2 != (r.Vec3{1, 2, 4}) {
		t.Fail()
	}
}

func TestVec3Unit(t *testing.T) {
	n1 := 3.0
	v1 := r.Vec3{n1, n1, n1}

	n2 := n1 / math.Sqrt(math.Pow(n1, 3))
	v2 := v1.Unit()

	if v2 != (r.Vec3{n2, n2, n2}) {
		t.Fail()
	}
}

func TestVec3LengthSquared(t *testing.T) {
	v := r.Vec3{2, 3, 4}

	if v.LengthSquared() != 29 {
		t.Fail()
	}
}

func TestVec3Length(t *testing.T) {
	v := r.Vec3{2, 3, 4}

	if v.Length() != math.Sqrt(29) {
		t.Fail()
	}
}

func TestVec3Dot(t *testing.T) {
	v1 := r.Vec3{1, 2, 3}
	v2 := r.Vec3{2, 3, 4}

	v3 := v1.Dot(v2)

	if v3 != 20 {
		t.Fail()
	}
}

func TestVec3Reflect(t *testing.T) {}

func TestVec3NearZero(t *testing.T) {}

func TestVec3String(t *testing.T) {
	v := r.Vec3{1, 2, 3}

	if v.String() != "Vec3(1.000000, 2.000000, 3.000000)" {
		t.Fail()
	}
}
