package raytracer_test

import (
	"math"
	"testing"

	r "github.com/derek-schaefer/raytracer"
	"github.com/stretchr/testify/assert"
)

func TestSphereHit(t *testing.T) {
	s1 := r.Sphere{Center: r.Point3{0, 0, -1}, Radius: 0.5}

	r1 := r.Ray{Origin: r.Vec3{}, Direction: r.Vec3{0, 0, -1}}

	var ok bool

	_, ok = s1.Hit(r1, r.Interval{0, math.Inf(1)})

	assert.True(t, ok)

	s2 := r.Sphere{Center: r.Point3{0, 0, -1}, Radius: 0.5}

	r2 := r.Ray{Origin: r.Vec3{}, Direction: r.Vec3{0, 0, 1}}

	_, ok = s2.Hit(r2, r.Interval{0, math.Inf(1)})

	assert.False(t, ok)
}
