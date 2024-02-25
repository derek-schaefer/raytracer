package raytracer_test

import (
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
