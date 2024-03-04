package raytracer

import (
	"fmt"
	"math"
	"math/rand"
)

type Vec3 [3]float64

func NewVec3(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

func RandomVec3(r *rand.Rand) Vec3 {
	return NewVec3(r.Float64(), r.Float64(), r.Float64())
}

func RandomRangeVec3(r *rand.Rand, min, max float64) Vec3 {
	return NewVec3(RandFloat64(r, min, max), RandFloat64(r, min, max), RandFloat64(r, min, max))
}

func RandomUnitSphereVec3(r *rand.Rand) Vec3 {
	for {
		p := RandomRangeVec3(r, -1, 1)

		if p.LengthSquared() < 1 {
			return p
		}
	}
}

func RandomUnitVec3(r *rand.Rand) Vec3 {
	return RandomUnitSphereVec3(r).Unit()
}

func RandomHemisphereVec3(r *rand.Rand, normal Vec3) Vec3 {
	v := RandomUnitVec3(r)

	if v.Dot(normal) > 0 {
		return v
	} else {
		return v.Multiply(-1)
	}
}

func (v Vec3) X() float64 {
	return v[0]
}

func (v Vec3) Y() float64 {
	return v[1]
}

func (v Vec3) Z() float64 {
	return v[2]
}

func (v *Vec3) SetX(x float64) *Vec3 {
	v[0] = x

	return v
}

func (v *Vec3) SetY(y float64) *Vec3 {
	v[1] = y

	return v
}

func (v *Vec3) SetZ(z float64) *Vec3 {
	v[2] = z

	return v
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

func (v Vec3) MultiplyV(o Vec3) Vec3 {
	var t Vec3

	for i := 0; i < len(v); i++ {
		t[i] = v[i] * o[i]
	}

	return t
}

func (v Vec3) Divide(t float64) Vec3 {
	return v.Multiply(1 / t)
}

func (v Vec3) LengthSquared() float64 {
	var t float64

	for i := 0; i < len(v); i++ {
		t += v[i] * v[i]
	}

	return t
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) Unit() Vec3 {
	return v.Divide(v.Length())
}

func (v Vec3) Dot(o Vec3) float64 {
	var t float64

	for i := 0; i < len(v); i++ {
		t += v[i] * o[i]
	}

	return t
}

func (u Vec3) Cross(v Vec3) Vec3 {
	return Vec3{
		u[1]*v[2] - u[2]*v[1],
		u[2]*v[0] - u[0]*v[2],
		u[0]*v[1] - u[1]*v[0],
	}
}

func (v Vec3) NearZero() bool {
	for i := 0; i < len(v); i++ {
		if !NearlyEqual(v[i], 0) {
			return false
		}
	}

	return true
}

func (v Vec3) Reflect(n Vec3) Vec3 {
	return v.Subtract(n.Multiply(2 * v.Dot(n)))
}

func (v Vec3) Refract(n Vec3, etaiOverEtat float64) Vec3 {
	cosTheta := math.Min(v.Multiply(-1).Dot(n), 1)
	rOutPerp := v.Add(n.Multiply(cosTheta)).Multiply(etaiOverEtat)
	rOutPara := n.Multiply(-math.Sqrt(math.Abs(1 - rOutPerp.LengthSquared())))
	return rOutPerp.Add(rOutPara)
}

func (v Vec3) String() string {
	return fmt.Sprintf("Vec3(%f, %f, %f)", v.X(), v.Y(), v.Z())
}
