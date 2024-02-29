package raytracer

type Metal struct {
	Albedo Color
}

func NewMetal(albedo Color) Metal {
	return Metal{Albedo: albedo}
}

func (m Metal) Scatter(in Ray, hit Hit) (Ray, Color, bool) {
	reflected := in.Direction.Unit().Reflect(hit.N)

	scattered := Ray{Origin: hit.P, Direction: reflected}

	return scattered, m.Albedo, true
}
