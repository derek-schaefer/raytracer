package raytracer_test

import (
	"testing"

	r "github.com/derek-schaefer/raytracer"
	"github.com/stretchr/testify/assert"
)

func TestLambertianScatter(t *testing.T) {
	c := r.NewColor(r.NewVec3(0.5, 0.5, 0.5))
	l := r.NewLambertian(c)
	i := r.NewRay(r.NewPoint3(0, 0, 0), r.NewVec3(1, 1, 1))
	h := r.Hit{}

	s, a, ok := l.Scatter(i, h)

	assert.True(t, ok)
	assert.Equal(t, a, c)
	assert.Equal(t, s.Origin, i.Origin)
	assert.False(t, s.Direction.NearZero())
	assert.True(t, r.NearlyEqual(s.Direction.Length(), 1))
}
