package raytracer_test

import (
	"math"
	"math/rand"
	"testing"

	r "github.com/derek-schaefer/raytracer"
	"github.com/stretchr/testify/assert"
)

var (
	random = rand.New(rand.NewSource(1))
)

func TestRandomFloat64(t *testing.T) {
	min := 2.0
	max := 3.0

	n := r.RandFloat64(random, min, max)

	assert.True(t, min <= n && n <= max)
}

func TestNearlyEqual(t *testing.T) {
	assert.False(t, r.NearlyEqual(0, 1))
	assert.True(t, r.NearlyEqual(math.Pow(2, 53), math.Pow(2, 53)-1))
	assert.True(t, r.NearlyEqual(math.Pow(2, 53), math.Pow(2, 53)+1))
}

func TestReflectance(t *testing.T) {
	c := math.Cos(1)
	n := r.Reflectance(c, 0.5)

	assert.Equal(t, n, 0.12935885194190405)
}
