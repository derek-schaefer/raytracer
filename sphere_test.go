package raytracer_test

import (
	"math"
	"testing"

	r "github.com/derek-schaefer/raytracer"
)

func TestSphereHit(t *testing.T) {
	s1 := r.Sphere{Center: r.Point3{0, 0, -1}, Radius: 0.5}

	r1 := r.Ray{Origin: r.Vec3{}, Direction: r.Vec3{0, 0, -1}}

	if _, ok := s1.Hit(r1, 0, math.Inf(1)); !ok {
		t.Fail()
	}

	s2 := r.Sphere{Center: r.Point3{0, 0, -1}, Radius: 0.5}

	r2 := r.Ray{Origin: r.Vec3{}, Direction: r.Vec3{0, 0, 1}}

	if _, ok := s2.Hit(r2, 0, math.Inf(1)); ok {
		t.Fail()
	}
}
