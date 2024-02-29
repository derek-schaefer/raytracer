package raytracer_test

import (
	"testing"

	r "github.com/derek-schaefer/raytracer"
)

func TestMetalScatter(t *testing.T) {
	c := r.NewColor(r.NewVec3(0.5, 0.5, 0.5))
	m := r.NewMetal(c, 0.5)
	i := r.NewRay(r.NewPoint3(0, 0, 0), r.NewVec3(1, 1, 1))
	h := r.Hit{}

	s, a, ok := m.Scatter(i, h)

	if a != c {
		t.Errorf("Expected `a` to equal %s, got: %s", c, a)
	}

	if s.Origin != h.P {
		t.Errorf("Expected `s.Origin` to equal %s, got: %s", h.P, s.Origin)
	}

	if !(ok || !r.NearlyEqual(s.Direction.Length(), 1)) {
		t.Errorf("Expected the scattered ray to continue or be absorbed consistently")
	}
}
