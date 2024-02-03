package raytracer

import (
	"math"
)

type Sphere struct {
	Center Point3
	Radius float64
}

func (s Sphere) Hit(r Ray, tmin, tmax float64) (Hit, bool) {
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
	if root <= tmin || tmax <= root {
		root = (-halfB + sqrtd) / a
		if root <= tmin || tmax <= root {
			return Hit{}, false
		}
	}

	var hit Hit

	hit.T = root
	hit.P = r.At(hit.T)
	hit.N = hit.P.Subtract(s.Center).Divide(s.Radius)

	return hit, true
}
