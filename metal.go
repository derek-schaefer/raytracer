package raytracer

type Metal struct {
	Albedo Color
	Fuzz   float64
}

func NewMetal(albedo Color, fuzz float64) Metal {
	var f float64

	if f < 1 {
		f = fuzz
	} else {
		f = 1
	}

	return Metal{Albedo: albedo, Fuzz: f}
}

func (m Metal) Scatter(in Ray, hit Hit) (Ray, Color, bool) {
	reflected := in.Direction.Unit().Reflect(hit.N)

	scattered := Ray{Origin: hit.P, Direction: reflected.Add(RandomUnitVec3().Multiply(m.Fuzz))}

	ok := scattered.Direction.Dot(hit.N) > 0

	return scattered, m.Albedo, ok
}
