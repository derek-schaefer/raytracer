package raytracer_test

import (
	"math/rand"
	"testing"

	r "github.com/derek-schaefer/raytracer"
	"github.com/stretchr/testify/assert"
)

func TestNewCamera(t *testing.T) {
	t.Skip()
}

func TestCameraRender(t *testing.T) {
	random := rand.New(rand.NewSource(1))

	world := r.NewHittables()

	world.Add(r.Sphere{Center: r.Point3{0, 0, -1}, Radius: 0.5})

	camera := r.NewCamera(
		r.CameraOptions{
			AspectRatio:    16.0 / 9.0,
			Center:         r.Point3{0, 0, 0},
			FocalLength:    1.0,
			ImageWidth:     400,
			Random:         random,
			ViewportHeight: 2.0,
		},
	)

	image := camera.Render(world)

	assert.Equal(t, image.Width, 400)
	assert.Equal(t, image.Height, 225)
	assert.Equal(t, len(image.Pixels), 400*225)
}
