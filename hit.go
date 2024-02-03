package raytracer

type Hit struct {
	P Point3
	N Vec3
	T float64
}

type Hittable interface {
	Hit(r Ray, tmin, tmax float64) (Hit, bool)
}
