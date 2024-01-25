package raytracer_test

import (
	"testing"

	r "github.com/derek-schaefer/raytracer"
)

func TestRayAt(t *testing.T) {
	r := r.Ray{
		r.Point3{1, 1, 1},
		r.Vec3{2, 2, 2},
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

func TestRayString(t *testing.T) {
	r := r.Ray{
		r.Point3{1, 1, 1},
		r.Vec3{2, 2, 2},
	}

	if r.String() != "Ray(Vec3(1.000000, 1.000000, 1.000000), Vec3(2.000000, 2.000000, 2.000000))" {
		t.Fail()
	}
}
