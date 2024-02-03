package raytracer

import (
	"math"
)

type Sphere struct {
	Center Point3
	Radius float64
}

func (s Sphere) Hit(r Ray, tmin, tmax float64) (Hit, bool) {
	var hit Hit

	oc := r.Origin.Subtract(s.Center)
	a := r.Direction.LengthSquared()
	halfB := oc.Dot(r.Direction)
	c := oc.LengthSquared() - s.Radius*s.Radius

	discriminant := halfB*halfB - a*c

	if discriminant < 0 {
		return hit, false
	}

	sqrtd := math.Sqrt(discriminant)

	root := (-halfB - sqrtd) / a
	if root <= tmin || tmax <= root {
		root = (-halfB + sqrtd) / a
		if root <= tmin || tmax <= root {
			return hit, false
		}
	}

	hit.T = root
	hit.P = r.At(hit.T)

	hit.SetFaceNormal(r, hit.P.Subtract(s.Center).Divide(s.Radius))

	return hit, true
}
