package raytracer

import (
	"math"
	"math/rand"
)

const (
	epsilon = 1e-15
)

// Generate a random floating point number in the range [min, max).
func RandFloat64(r *rand.Rand, min, max float64) float64 {
	return min + r.Float64()*(max-min)
}

// See: https://stackoverflow.com/a/76386543
func NearlyEqual(a, b float64) bool {
	if a == b {
		return true
	}

	diff := math.Abs(a - b)
	if a == 0.0 || b == 0.0 || diff < math.SmallestNonzeroFloat64 {
		return diff < epsilon*math.SmallestNonzeroFloat64
	}

	return diff/(math.Abs(a)+math.Abs(b)) < epsilon
}

// Use Schlick's approximation for reflectance.
func Reflectance(cosine, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
