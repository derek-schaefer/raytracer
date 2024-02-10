package raytracer_test

import (
	"bytes"
	"testing"

	r "github.com/derek-schaefer/raytracer"
)

func TestColorWritePPM(t *testing.T) {
	c1 := r.NewColor(r.Vec3{2e-2, 4e-2, 8e-2})

	var b bytes.Buffer

	c1.WritePPM(&b)

	if b.String() != "5 10 20\n" {
		t.Fail()
	}

	c2 := r.NewColor(r.NewVec3(-1, 0, 2))

	b.Reset()

	c2.WritePPM(&b)

	if b.String() != "0 0 255\n" {
		t.Fail()
	}
}
