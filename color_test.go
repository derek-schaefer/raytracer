package raytracer_test

import (
	"bytes"
	"math"
	"testing"

	r "github.com/derek-schaefer/raytracer"
	"github.com/stretchr/testify/assert"
)

func TestColorLinearToGamma(t *testing.T) {
	c1 := r.NewColor(9, 9, 9)

	c2 := c1.LinearToGamma()

	for i := 0; i < len(c2.V); i++ {
		assert.Equal(t, c1.V[i], math.Pow(c2.V[i], 2))
	}
}

func TestColorWritePPM(t *testing.T) {
	c1 := r.NewColor(2e-2, 4e-2, 8e-2)

	var b bytes.Buffer

	c1.WritePPM(&b)

	assert.Equal(t, b.String(), "5 10 20\n")

	c2 := r.NewColor(-1, 0, 2)

	b.Reset()

	c2.WritePPM(&b)

	assert.Equal(t, b.String(), "0 0 255\n")
}
