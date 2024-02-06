package raytracer_test

import (
	"testing"

	r "github.com/derek-schaefer/raytracer"
)

func TestCameraRender(t *testing.T) {
	world := r.NewHittables()

	world.Add(r.Sphere{Center: r.Point3{0, 0, -1}, Radius: 0.5})

	camera := r.Camera{
		AspectRatio:    16.0 / 9.0,
		Center:         r.Point3{0, 0, 0},
		FocalLength:    1.0,
		ImageWidth:     400,
		ViewportHeight: 2.0,
	}

	image := camera.Render(world)

	if image.Width != 400 {
		t.Fail()
	}

	if image.Height != 225 {
		t.Fail()
	}

	if len(image.Pixels) != 400*225 {
		t.Fail()
	}
}
