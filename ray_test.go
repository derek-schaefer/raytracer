package raytracer_test

import (
	"testing"

	r "github.com/derek-schaefer/raytracer"
	"github.com/stretchr/testify/assert"
)

func TestRayAt(t *testing.T) {
	r := r.Ray{
		r.Point3{1, 1, 1},
		r.Vec3{2, 2, 2},
	}

	p := r.At(2)

	assert.Equal(t, p.X(), float64(5))
	assert.Equal(t, p.Y(), float64(5))
	assert.Equal(t, p.Z(), float64(5))
}

func TestRayString(t *testing.T) {
	r := r.Ray{
		r.Point3{1, 1, 1},
		r.Vec3{2, 2, 2},
	}

	assert.Equal(t, r.String(), "Ray(Vec3(1.000000, 1.000000, 1.000000), Vec3(2.000000, 2.000000, 2.000000))")
}
