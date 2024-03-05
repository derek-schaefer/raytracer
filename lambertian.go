package raytracer

import "math/rand"

type LambertianOptions struct {
	Albedo Color
}

type Lambertian struct {
	LambertianOptions
}

func NewLambertian(options LambertianOptions) Lambertian {
	return Lambertian{LambertianOptions: options}
}

func (l Lambertian) Scatter(random *rand.Rand, in Ray, hit Hit) (Ray, Color, bool) {
	direction := hit.N.Add(RandomUnitVec3(random))

	// Catch degenerate scatter direction
	if direction.NearZero() {
		direction = hit.N
	}

	scattered := Ray{Origin: hit.P, Direction: direction}

	return scattered, l.Albedo, true
}
