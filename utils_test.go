package raytracer_test

import (
	"math"
	"testing"

	r "github.com/derek-schaefer/raytracer"
)

func TestRandomFloat64(t *testing.T) {
	min := 2.0
	max := 3.0

	n := r.RandFloat64(min, max)

	if !(min <= n && n <= max) {
		t.Fail()
	}
}

func TestNearlyEqual(t *testing.T) {
	if r.NearlyEqual(0, 1) {
		t.Fail()
	}

	if !r.NearlyEqual(math.Pow(2, 53), math.Pow(2, 53)-1) {
		t.Fail()
	}

	if !r.NearlyEqual(math.Pow(2, 53), math.Pow(2, 53)+1) {
		t.Fail()
	}
}
