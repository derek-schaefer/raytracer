package raytracer_test

import (
	"math"
	"testing"

	r "github.com/derek-schaefer/raytracer"
)

func TestNewInterval(t *testing.T) {
	i := r.NewInterval(-1, 1)

	if i.Min != -1 && i.Max != 1 {
		t.Fail()
	}
}

func TestIntervalContains(t *testing.T) {
	i := r.NewInterval(-1, 1)

	if !i.Contains(-1) {
		t.Fail()
	}

	if !i.Contains(0) {
		t.Fail()
	}

	if !i.Contains(1) {
		t.Fail()
	}

	if i.Contains(-2) {
		t.Fail()
	}

	if i.Contains(2) {
		t.Fail()
	}
}

func TestIntervalSurrounds(t *testing.T) {
	i := r.NewInterval(-1, 1)

	if !i.Surrounds(0) {
		t.Fail()
	}

	if i.Surrounds(-1) {
		t.Fail()
	}

	if i.Surrounds(1) {
		t.Fail()
	}
}

func TestEmptyInterval(t *testing.T) {
	i := r.EmptyInterval()

	if !math.IsInf(i.Min, 1) {
		t.Fail()
	}

	if !math.IsInf(i.Max, -1) {
		t.Fail()
	}
}

func TestUniverseInterval(t *testing.T) {
	i := r.UniverseInterval()

	if !math.IsInf(i.Min, -1) {
		t.Fail()
	}

	if !math.IsInf(i.Max, 1) {
		t.Fail()
	}
}
