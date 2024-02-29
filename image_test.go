package raytracer_test

import (
	"bytes"
	"testing"

	r "github.com/derek-schaefer/raytracer"
	"github.com/stretchr/testify/assert"
)

func TestImageSet(t *testing.T) {
	img := r.NewImage(8, 8)

	img.Set(3, 2, r.NewColor(r.Vec3{0, 0, 255}))

	assert.Equal(t, img.Pixels[19], r.NewColor(r.Vec3{0, 0, 255}))
}

func TestImageGet(t *testing.T) {
	img := r.NewImage(8, 8)

	img.Pixels[19] = r.NewColor(r.Vec3{0, 0, 255})

	assert.Equal(t, img.Get(3, 2), r.NewColor(r.Vec3{0, 0, 255}))
}

func TestImageWritePPM(t *testing.T) {
	img := r.NewImage(2, 2)

	for j := 0; j < img.Height; j++ {
		for i := 0; i < img.Width; i++ {
			img.Set(i, j, r.NewColor(r.Vec3{float64(i) / 255, float64(j) / 255, 0}))
		}
	}

	var buf bytes.Buffer

	img.WritePPM(&buf)

	assert.Equal(t, buf.String(), "P3\n2 2\n255\n0 0 0\n1 0 0\n0 1 0\n1 1 0\n")
}
