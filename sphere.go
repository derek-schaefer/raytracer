package raytracer

import (
	"math"
)

type SphereOptions struct {
	Center   Point3
	Radius   float64
	Material Material
}

type Sphere struct {
	SphereOptions
}

func NewSphere(options SphereOptions) Sphere {
	return Sphere{SphereOptions: options}
}

func (s Sphere) Hit(r Ray, rt Interval) (Hit, bool) {
	oc := r.Origin.Subtract(s.Center)
	a := r.Direction.LengthSquared()
	halfB := oc.Dot(r.Direction)
	c := oc.LengthSquared() - s.Radius*s.Radius

	discriminant := halfB*halfB - a*c

	if discriminant < 0 {
		return Hit{}, false
	}

	sqrtd := math.Sqrt(discriminant)

	root := (-halfB - sqrtd) / a
	if !rt.Surrounds(root) {
		root = (-halfB + sqrtd) / a
		if !rt.Surrounds(root) {
			return Hit{}, false
		}
	}

	var hit Hit

	hit.T = root
	hit.P = r.At(hit.T)
	hit.Material = s.Material

	hit.SetFaceNormal(r, hit.P.Subtract(s.Center).Divide(s.Radius))

	return hit, true
}
