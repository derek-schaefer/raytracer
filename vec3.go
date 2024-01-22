package raytracer

type Vec3 [3]float64

func (v *Vec3) X() float64 {
	return v[0]
}

func (v *Vec3) Y() float64 {
	return v[1]
}

func (v *Vec3) Z() float64 {
	return v[2]
}

func (v *Vec3) Add(o *Vec3) *Vec3 {
	v[0] += o[0]
	v[1] += o[1]
	v[2] += o[2]

	return v
}

func (v *Vec3) Multiply(t float64) *Vec3 {
	v[0] *= t
	v[1] *= t
	v[2] *= t

	return v
}
