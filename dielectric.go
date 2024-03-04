package raytracer

import (
	"math"
	"math/rand"
)

type DielectricOptions struct {
	IndexOfRefraction float64
	Random            *rand.Rand
}

type Dielectric struct {
	DielectricOptions
}

func NewDielectric(options DielectricOptions) Dielectric {
	return Dielectric{DielectricOptions: options}
}

func (d Dielectric) Scatter(in Ray, hit Hit) (Ray, Color, bool) {
	attenuation := ColorWhite

	var refractionRatio float64

	if hit.F {
		refractionRatio = 1 / d.IndexOfRefraction
	} else {
		refractionRatio = d.IndexOfRefraction
	}

	unitDirection := in.Direction.Unit()
	cosTheta := math.Min(unitDirection.Multiply(-1).Dot(hit.N), 1)
	sinTheta := math.Sqrt(1 - cosTheta*cosTheta)
	cannotRefract := refractionRatio*sinTheta > 1

	var direction Vec3

	if cannotRefract || Reflectance(cosTheta, refractionRatio) > d.Random.Float64() {
		direction = unitDirection.Reflect(hit.N)
	} else {
		direction = unitDirection.Refract(hit.N, refractionRatio)
	}

	scattered := NewRay(hit.P, direction)

	return scattered, attenuation, true
}
