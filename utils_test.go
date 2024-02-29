package raytracer_test

import (
	"math"
	"testing"

	r "github.com/derek-schaefer/raytracer"
	"github.com/stretchr/testify/assert"
)

func TestRandomFloat64(t *testing.T) {
	min := 2.0
	max := 3.0

	n := r.RandFloat64(min, max)

	assert.True(t, min <= n && n <= max)
}

func TestNearlyEqual(t *testing.T) {
	assert.False(t, r.NearlyEqual(0, 1))
	assert.True(t, r.NearlyEqual(math.Pow(2, 53), math.Pow(2, 53)-1))
	assert.True(t, r.NearlyEqual(math.Pow(2, 53), math.Pow(2, 53)+1))
}
