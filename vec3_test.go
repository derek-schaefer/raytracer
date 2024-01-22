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

	v1.Add(&v2)

	if v1.X() != 2 {
		t.Fail()
	}

	if v1.Y() != 3 {
		t.Fail()
	}

	if v1.Z() != 4 {
		t.Fail()
	}
}

func TestMultiply(t *testing.T) {
	v := raytracer.Vec3{1, 2, 3}

	v.Multiply(2)

	if v.X() != 2 {
		t.Fail()
	}

	if v.Y() != 4 {
		t.Fail()
	}

	if v.Z() != 6 {
		t.Fail()
	}
}
