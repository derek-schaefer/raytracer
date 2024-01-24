package raytracer_test

import (
	"math"
	"testing"

	r "github.com/derek-schaefer/raytracer"
)

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

func TestVec3String(t *testing.T) {
	v := r.Vec3{1, 2, 3}

	if v.String() != "Vec3(1.000000, 2.000000, 3.000000)" {
		t.Fail()
	}
}
