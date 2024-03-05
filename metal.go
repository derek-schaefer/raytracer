package raytracer

import "math/rand"

type MetalOptions struct {
	Albedo Color
	Fuzz   float64
}

type Metal struct {
	MetalOptions
}

func NewMetal(options MetalOptions) Metal {
	if options.Fuzz > 1 {
		options.Fuzz = 1
	}

	return Metal{MetalOptions: options}
}

func (m Metal) Scatter(random *rand.Rand, in Ray, hit Hit) (Ray, Color, bool) {
	reflected := in.Direction.Unit().Reflect(hit.N)

	scattered := Ray{Origin: hit.P, Direction: reflected.Add(RandomUnitVec3(random).Multiply(m.Fuzz))}

	ok := scattered.Direction.Dot(hit.N) > 0

	return scattered, m.Albedo, ok
}
