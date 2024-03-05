package raytracer_test

import (
	"testing"

	r "github.com/derek-schaefer/raytracer"
	"github.com/stretchr/testify/assert"
)

func TestDielectricScatter(t *testing.T) {
	o := r.DielectricOptions{
		IndexOfRefraction: 1.5,
	}
	d := r.NewDielectric(o)
	i := r.NewRay(r.NewPoint3(0, 0, 0), r.NewVec3(1, 1, 1))

	var s r.Ray
	var a r.Color
	var ok bool
	var n float64

	h := r.Hit{
		P:        r.NewVec3(-1, -1, -1),
		N:        r.NewVec3(-1, -1, -1),
		Material: d,
	}

	s, a, ok = d.Scatter(random, i, h)

	n = -0.180354307934656

	assert.True(t, ok)
	assert.Equal(t, a, r.ColorWhite)
	assert.Equal(t, s, r.NewRay(h.P, r.NewVec3(n, n, n)))

	h.F = true

	s, a, ok = d.Scatter(random, i, h)

	n = 0.5910582031297084

	assert.True(t, ok)
	assert.Equal(t, a, r.ColorWhite)
	assert.Equal(t, s, r.NewRay(h.P, r.NewVec3(n, n, n)))
}
