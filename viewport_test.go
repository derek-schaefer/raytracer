package raytracer_test

import (
	"math/rand"
	"testing"

	r "github.com/derek-schaefer/raytracer"
	"github.com/stretchr/testify/assert"
)

func TestViewportU(t *testing.T) {
	v := r.Viewport{Width: 16, Height: 9}

	assert.Equal(t, v.U(), r.Vec3{16, 0, 0})
}

func TestViewportV(t *testing.T) {
	v := r.Viewport{Width: 16, Height: 9}

	assert.Equal(t, v.V(), r.Vec3{0, -9, 0})
}

func TestViewportDeltaU(t *testing.T) {
	v := r.Viewport{Width: 16, Height: 9}

	assert.Equal(t, v.DeltaU(2), r.Vec3{8, 0, 0})
}

func TestViewportDeltaV(t *testing.T) {
	v := r.Viewport{Width: 16, Height: 9}

	assert.Equal(t, v.DeltaV(3), r.Vec3{0, -3, 0})
}

func TestViewportOrigin(t *testing.T) {
	random := rand.New(rand.NewSource(1))

	v := r.Viewport{Width: 16, Height: 9}

	c := r.NewCamera(
		r.CameraOptions{
			Center:      r.Point3{0, 0, 0},
			FocalLength: 1.0,
			Random:      random,
		},
	)

	assert.Equal(t, v.Origin(c), r.Point3{-8, 4.5, -1})
}
