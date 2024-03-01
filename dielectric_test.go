package raytracer_test

import (
	"testing"

	r "github.com/derek-schaefer/raytracer"
	"github.com/stretchr/testify/assert"
)

func TestDielectricScatter(t *testing.T) {
	d := r.NewDielectric(1.5)
	i := r.NewRay(r.NewPoint3(0, 0, 0), r.NewVec3(1, 1, 1))

	var h r.Hit
	var s r.Ray
	var a r.Color
	var ok bool
	var n float64

	s, a, ok = d.Scatter(i, h)

	n = 0.5773502691896258

	assert.True(t, ok)
	assert.Equal(t, a, r.ColorWhite)
	assert.Equal(t, s, r.NewRay(i.Origin, r.NewVec3(n, n, n)))

	h.F = true

	s, a, ok = d.Scatter(i, h)

	n = 0.3849001794597505

	assert.True(t, ok)
	assert.Equal(t, a, r.ColorWhite)
	assert.Equal(t, s, r.NewRay(i.Origin, r.NewVec3(n, n, n)))
}
