package raytracer

import (
	"fmt"
	"math"
)

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

func (v Vec3) LengthSquared() float64 {
	return v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z()
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) Unit() Vec3 {
	return v.Divide(v.Length())
}

func (v Vec3) String() string {
	return fmt.Sprintf("Vec3(%f, %f, %f)", v.X(), v.Y(), v.Z())
}
