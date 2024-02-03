package raytracer

import "math"

type Interval struct {
	Min float64
	Max float64
}

func NewInterval(min, max float64) Interval {
	return Interval{Min: min, Max: max}
}

func (i Interval) Contains(x float64) bool {
	return i.Min <= x && x <= i.Max
}

func (i Interval) Surrounds(x float64) bool {
	return i.Min < x && x < i.Max
}

func EmptyInterval() Interval {
	return Interval{math.Inf(1), math.Inf(-1)}
}

func UniverseInterval() Interval {
	return Interval{math.Inf(-1), math.Inf(1)}
}
