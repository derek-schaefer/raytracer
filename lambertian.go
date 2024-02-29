package raytracer

type Lambertian struct {
	Albedo Color
}

func NewLambertian(albedo Color) Lambertian {
	return Lambertian{Albedo: albedo}
}

func (l Lambertian) Scatter(in Ray, hit Hit) (Ray, Color, bool) {
	direction := hit.N.Add(RandomUnitVec3())

	// Catch degenerate scatter direction
	if direction.NearZero() {
		direction = hit.N
	}

	scattered := Ray{Origin: hit.P, Direction: direction}

	return scattered, l.Albedo, true
}
