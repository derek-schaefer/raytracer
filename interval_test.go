package raytracer_test

import (
	"math"
	"testing"

	r "github.com/derek-schaefer/raytracer"
	"github.com/stretchr/testify/assert"
)

func TestNewInterval(t *testing.T) {
	i := r.NewInterval(-1, 1)

	assert.Equal(t, i.Min, float64(-1))
	assert.Equal(t, i.Max, float64(1))
}

func TestIntervalContains(t *testing.T) {
	i := r.NewInterval(-1, 1)

	assert.True(t, i.Contains(-1))
	assert.True(t, i.Contains(0))
	assert.True(t, i.Contains(1))
	assert.False(t, i.Contains(-2))
	assert.False(t, i.Contains(2))
}

func TestIntervalSurrounds(t *testing.T) {
	i := r.NewInterval(-1, 1)

	assert.True(t, i.Surrounds(0))
	assert.False(t, i.Surrounds(-1))
	assert.False(t, i.Surrounds(1))
}

func TestIntervalClamp(t *testing.T) {
	i := r.NewInterval(-1, 1)

	assert.Equal(t, i.Clamp(-1.5), float64(-1))
	assert.Equal(t, i.Clamp(-1), float64(-1))
	assert.Equal(t, i.Clamp(0), float64(0))
	assert.Equal(t, i.Clamp(1), float64(1))
	assert.Equal(t, i.Clamp(1.5), float64(1))
}

func TestEmptyInterval(t *testing.T) {
	i := r.EmptyInterval()

	assert.True(t, math.IsInf(i.Min, 1))
	assert.True(t, math.IsInf(i.Max, -1))
}

func TestUniverseInterval(t *testing.T) {
	i := r.UniverseInterval()

	assert.True(t, math.IsInf(i.Min, -1))
	assert.True(t, math.IsInf(i.Max, 1))
}
