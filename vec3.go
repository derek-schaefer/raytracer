package raytracer

type Vec3 [3]float64

func (v Vec3) X() float64 {
	return v[0]
}

func (v Vec3) Y() float64 {
	return v[1]
}

func (v Vec3) Z() float64 {
	return v[2]
}

func (v Vec3) Add(o Vec3) Vec3 {
	for i := 0; i < len(v); i++ {
		v[i] += o[i]
	}

	return v
}

func (v Vec3) Subtract(o Vec3) Vec3 {
	for i := 0; i < len(v); i++ {
		v[i] -= o[i]
	}

	return v
}

func (v Vec3) Multiply(t float64) Vec3 {
	for i := 0; i < len(v); i++ {
		v[i] *= t
	}

	return v
}

func (v Vec3) Divide(t float64) Vec3 {
	return v.Multiply(1 / t)
}
