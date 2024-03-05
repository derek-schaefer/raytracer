package raytracer

import "math/rand"

type Material interface {
	Scatter(r *rand.Rand, in Ray, hit Hit) (Ray, Color, bool)
}
