package raytracer_test

import (
	"math"
	"testing"

	r "github.com/derek-schaefer/raytracer"
	"github.com/stretchr/testify/assert"
)

func TestNewHittables(t *testing.T) {
	s := r.Sphere{Center: r.Point3{0, 0, 1}, Radius: 1}

	hs := r.NewHittables(s)

	assert.Equal(t, len(hs.Objects), 1)
	assert.Equal(t, hs.Objects[0], s)
}

func TestHittablesAdd(t *testing.T) {
	s := r.Sphere{Center: r.Point3{0, 0, 1}, Radius: 1}

	hs := r.NewHittables()

	assert.Equal(t, len(hs.Objects), 0)

	hs.Add(s)

	assert.Equal(t, len(hs.Objects), 1)
	assert.Equal(t, hs.Objects[0], s)
}

func TestHittablesClear(t *testing.T) {
	s := r.Sphere{Center: r.Point3{0, 0, 1}, Radius: 1}

	hs := r.NewHittables(s)

	assert.Equal(t, len(hs.Objects), 1)

	hs.Clear()

	assert.Equal(t, len(hs.Objects), 0)
}

func TestHittablesHit(t *testing.T) {
	s := r.Sphere{Center: r.Point3{0, 0, 1}, Radius: 1}

	hs := r.NewHittables(s)

	ray := r.Ray{Origin: r.Point3{0, 0, 0}, Direction: r.Vec3{0, 0, 1}}
	rayt := r.NewInterval(0, math.Inf(1))

	h, ok := hs.Hit(ray, rayt)

	assert.True(t, ok)

	p := r.Point3{0, 0, 2}
	n := r.Vec3{0, 0, -1}

	assert.Equal(t, h, r.Hit{P: p, N: n, T: 2, F: false})
}
