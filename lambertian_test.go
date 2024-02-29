package raytracer_test

import (
	"testing"

	r "github.com/derek-schaefer/raytracer"
)

func TestLambertianScatter(t *testing.T) {
	c := r.NewColor(r.NewVec3(0.5, 0.5, 0.5))
	l := r.NewLambertian(c)
	i := r.NewRay(r.NewPoint3(0, 0, 0), r.NewVec3(1, 1, 1))
	h := r.Hit{}

	s, a, ok := l.Scatter(i, h)

	if !ok {
		t.Fail()
	}

	if a != c {
		t.Fail()
	}

	if s.Origin != i.Origin {
		t.Fail()
	}

	if s.Direction.NearZero() {
		t.Fail()
	}

	if !r.NearlyEqual(s.Direction.Length(), 1) {
		t.Fail()
	}
}
