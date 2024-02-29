package raytracer

import (
	"math"
	"math/rand"
)

const (
	epsilon = 1e-16
)

func RandFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

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
