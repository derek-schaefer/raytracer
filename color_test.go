package raytracer_test

import (
	"bytes"
	"testing"

	"github.com/derek-schaefer/raytracer"
)

func TestColorWritePPM(t *testing.T) {
	c := raytracer.Color{2e-2, 4e-2, 8e-2}

	var b bytes.Buffer

	c.WritePPM(&b)

	if b.String() != "5 10 20\n" {
		t.Fail()
	}
}
