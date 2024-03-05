package raytracer

import "math/rand"

type MetalOptions struct {
	Albedo Color
	Fuzz   float64
	Random *rand.Rand
}

type Metal struct {
	MetalOptions
}

func NewMetal(options MetalOptions) Metal {
	if options.Random == nil {
		panic("options.Random must not be nil")
	}

	if options.Fuzz > 1 {
		options.Fuzz = 1
	}

	return Metal{MetalOptions: options}
}

func (m Metal) Scatter(in Ray, hit Hit) (Ray, Color, bool) {
	reflected := in.Direction.Unit().Reflect(hit.N)

	scattered := Ray{Origin: hit.P, Direction: reflected.Add(RandomUnitVec3(m.Random).Multiply(m.Fuzz))}

	ok := scattered.Direction.Dot(hit.N) > 0

	return scattered, m.Albedo, ok
}
