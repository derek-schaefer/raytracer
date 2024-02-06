package raytracer_test

import (
	"math"
	"testing"

	r "github.com/derek-schaefer/raytracer"
)

func TestNewHittables(t *testing.T) {
	s := r.Sphere{Center: r.Point3{0, 0, 1}, Radius: 1}

	hs := r.NewHittables(s)

	if len(hs.Objects) != 1 {
		t.Fail()
	}

	if hs.Objects[0] != s {
		t.Fail()
	}
}

func TestHittablesAdd(t *testing.T) {
	s := r.Sphere{Center: r.Point3{0, 0, 1}, Radius: 1}

	hs := r.NewHittables()

	if len(hs.Objects) != 0 {
		t.Fail()
	}

	hs.Add(s)

	if hs.Objects[0] != s {
		t.Fail()
	}
}

func TestHittablesClear(t *testing.T) {
	s := r.Sphere{Center: r.Point3{0, 0, 1}, Radius: 1}

	hs := r.NewHittables(s)

	if len(hs.Objects) != 1 {
		t.Fail()
	}

	hs.Clear()

	if len(hs.Objects) != 0 {
		t.Fail()
	}
}

func TestHittablesHit(t *testing.T) {
	s := r.Sphere{Center: r.Point3{0, 0, 1}, Radius: 1}

	hs := r.NewHittables(s)

	ray := r.Ray{Origin: r.Point3{0, 0, 0}, Direction: r.Vec3{0, 0, 1}}
	rayt := r.NewInterval(0, math.Inf(1))

	h, ok := hs.Hit(ray, rayt)

	if !ok {
		t.Fail()
	}

	p := r.Point3{0, 0, 2}
	n := r.Vec3{0, 0, -1}

	if h != (r.Hit{P: p, N: n, T: 2, F: false}) {
		t.Fail()
	}
}
