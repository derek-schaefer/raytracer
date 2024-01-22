package raytracer_test

import (
	"testing"

	"github.com/derek-schaefer/raytracer"
)

func TestX(t *testing.T) {
	v := raytracer.Vec3{1, 0, 0}

	if v.X() != 1 {
		t.Fail()
	}
}

func TestY(t *testing.T) {
	v := raytracer.Vec3{0, 1, 0}

	if v.Y() != 1 {
		t.Fail()
	}
}

func TestZ(t *testing.T) {
	v := raytracer.Vec3{0, 0, 1}

	if v.Z() != 1 {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	v1 := raytracer.Vec3{1, 2, 3}
	v2 := raytracer.Vec3{1, 1, 1}

	v3 := v1.Add(v2)

	if v3.X() != 2 {
		t.Fail()
	}

	if v3.Y() != 3 {
		t.Fail()
	}

	if v3.Z() != 4 {
		t.Fail()
	}
}

func TestSubtract(t *testing.T) {
	v1 := raytracer.Vec3{1, 2, 3}
	v2 := raytracer.Vec3{1, 1, 1}

	v3 := v1.Subtract(v2)

	if v3.X() != 0 {
		t.Fail()
	}

	if v3.Y() != 1 {
		t.Fail()
	}

	if v3.Z() != 2 {
		t.Fail()
	}
}

func TestMultiply(t *testing.T) {
	v1 := raytracer.Vec3{1, 2, 3}

	v2 := v1.Multiply(2)

	if v2.X() != 2 {
		t.Fail()
	}

	if v2.Y() != 4 {
		t.Fail()
	}

	if v2.Z() != 6 {
		t.Fail()
	}
}

func TestDivide(t *testing.T) {
	v1 := raytracer.Vec3{2, 4, 8}

	v2 := v1.Divide(2)

	if v2.X() != 1 {
		t.Fail()
	}

	if v2.Y() != 2 {
		t.Fail()
	}

	if v2.Z() != 4 {
		t.Fail()
	}
}
