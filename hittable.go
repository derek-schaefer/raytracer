package raytracer

type Hittable interface {
	Hit(r Ray, rt Interval) (Hit, bool)
}
