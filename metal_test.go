package raytracer_test

import (
	"testing"

	r "github.com/derek-schaefer/raytracer"
	"github.com/stretchr/testify/assert"
)

func TestMetalScatter(t *testing.T) {
	c := r.NewColor(r.NewVec3(0.5, 0.5, 0.5))
	m := r.NewMetal(r.MetalOptions{c, 0.5, random})
	i := r.NewRay(r.NewPoint3(0, 0, 0), r.NewVec3(1, 1, 1))
	h := r.Hit{}

	s, a, ok := m.Scatter(i, h)

	assert.Equal(t, a, c)
	assert.Equal(t, s.Origin, h.P)
	assert.True(t, ok || !r.NearlyEqual(s.Direction.Length(), 1))
}
