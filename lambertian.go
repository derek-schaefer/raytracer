package raytracer

import "math/rand"

type LambertianOptions struct {
	Albedo Color
	Random *rand.Rand
}

type Lambertian struct {
	LambertianOptions
}

func NewLambertian(options LambertianOptions) Lambertian {
	if options.Random == nil {
		panic("LambertianOptions.Random should not be nil")
	}

	return Lambertian{LambertianOptions: options}
}

func (l Lambertian) Scatter(in Ray, hit Hit) (Ray, Color, bool) {
	direction := hit.N.Add(RandomUnitVec3(l.Random))

	// Catch degenerate scatter direction
	if direction.NearZero() {
		direction = hit.N
	}

	scattered := Ray{Origin: hit.P, Direction: direction}

	return scattered, l.Albedo, true
}
