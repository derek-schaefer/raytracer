package raytracer

type Hittable interface {
	Hit(r Ray, tmin, tmax float64) (Hit, bool)
}
