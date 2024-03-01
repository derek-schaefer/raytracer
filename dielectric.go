package raytracer

import (
	"math"
)

type Dielectric struct {
	IndexOfRefraction float64
}

func NewDielectric(indexOfRefraction float64) Dielectric {
	return Dielectric{IndexOfRefraction: indexOfRefraction}
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

	if cannotRefract {
		direction = unitDirection.Reflect(hit.N)
	} else {
		direction = unitDirection.Refract(hit.N, refractionRatio)
	}

	scattered := NewRay(hit.P, direction)

	return scattered, attenuation, true
}
