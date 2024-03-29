package raytracer_test

import (
	"testing"

	r "github.com/derek-schaefer/raytracer"
	"github.com/stretchr/testify/assert"
)

func TestNewCamera(t *testing.T) {
	t.Skip()
}

func TestCameraRender(t *testing.T) {
	world := r.NewHittables()

	material := r.NewMetal(r.MetalOptions{Albedo: r.NewColor(0.8, 0.6, 0.2), Fuzz: 0.0})

	world.Add(r.NewSphere(r.SphereOptions{Center: r.NewPoint3(0, 0, -1), Radius: 0.5, Material: material}))

	camera := r.NewCamera(
		r.CameraOptions{
			AspectRatio: 16.0 / 9.0,
			FieldOfView: 20,
			ImageWidth:  400,
			LookAt:      r.NewPoint3(0, 0, -1),
			LookFrom:    r.NewPoint3(-2, 2, 1),
			MaxDepth:    2,
			Samples:     2,
			ViewUp:      r.NewPoint3(0, 1, 0),
		},
	)

	image := camera.Render(world)

	assert.Equal(t, image.Width, 400)
	assert.Equal(t, image.Height, 225)
	assert.Equal(t, len(image.Pixels), 400*225)
}
