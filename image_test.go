package raytracer_test

import (
	"testing"

	r "github.com/derek-schaefer/raytracer"
)

func TestImageSet(t *testing.T) {
	img := r.NewImage(8, 8)

	img.Set(r.NewColor(r.Vec3{0, 0, 255}), 2, 3)

	if img.Pixels[19] != r.NewColor(r.Vec3{0, 0, 255}) {
		t.Fail()
	}
}

func TestImageGet(t *testing.T) {
	img := r.NewImage(8, 8)

	img.Pixels[19] = r.NewColor(r.Vec3{0, 0, 255})

	if img.Get(2, 3) != r.NewColor(r.Vec3{0, 0, 255}) {
		t.Fail()
	}
}

func TestImageWritePPM(t *testing.T) {
	// TODO
}
