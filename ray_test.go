package raytracer_test

import (
	"testing"

	"github.com/derek-schaefer/raytracer"
)

func TestRayAt(t *testing.T) {
	r := raytracer.Ray{
		raytracer.Point3{1, 1, 1},
		raytracer.Vec3{2, 2, 2},
	}

	p := r.At(2)

	if p.X() != 5 {
		t.Fail()
	}

	if p.Y() != 5 {
		t.Fail()
	}

	if p.Z() != 5 {
		t.Fail()
	}
}
