package raytracer_test

import (
	"math"
	"math/rand"
	"testing"

	r "github.com/derek-schaefer/raytracer"
	"github.com/stretchr/testify/assert"
)

var (
	random = rand.New(rand.NewSource(1))
)

func TestNewVec3(t *testing.T) {
	v := r.NewVec3(1, 2, 3)

	assert.Equal(t, v.X(), float64(1))
	assert.Equal(t, v.Y(), float64(2))
	assert.Equal(t, v.Z(), float64(3))
}

func TestRandomVec3(t *testing.T) {
	min := 0.0
	max := 1.0

	v := r.RandomVec3(random)

	x := v.X()
	y := v.Y()
	z := v.Z()

	assert.True(t, min <= x && x <= max)
	assert.True(t, min <= y && y <= max)
	assert.True(t, min <= z && z <= max)
}

func TestRandomUnitSphereVec3(t *testing.T) {
	v := r.RandomUnitSphereVec3(random)

	assert.Less(t, v.LengthSquared(), float64(1))
}

func TestRandomUnitVec3(t *testing.T) {
	v := r.RandomUnitVec3(random)

	f := v.Length()

	assert.True(t, r.NearlyEqual(f, 1))

	for i := 0; i < len(v); i++ {
		assert.True(t, -1 <= v[i] && v[i] <= 1)
	}
}

func TestRandomHemisphereVec3(t *testing.T) {
	s := r.RandomUnitSphereVec3(random)

	v := r.RandomHemisphereVec3(random, s)

	assert.Greater(t, v.Dot(s), float64(0))
}

func TestRandomRangeVec3(t *testing.T) {
	min := 2.0
	max := 3.0

	v := r.RandomRangeVec3(random, min, max)

	x := v.X()
	y := v.Y()
	z := v.Z()

	assert.True(t, min <= x && x <= max)
	assert.True(t, min <= y && y <= max)
	assert.True(t, min <= z && z <= max)
}

func TestVec3X(t *testing.T) {
	v := r.Vec3{1, 0, 0}

	assert.Equal(t, v.X(), float64(1))
}

func TestVec3Y(t *testing.T) {
	v := r.Vec3{0, 1, 0}

	assert.Equal(t, v.Y(), float64(1))
}

func TestVec3Z(t *testing.T) {
	v := r.Vec3{0, 0, 1}

	assert.Equal(t, v.Z(), float64(1))
}

func TestVec3SetY(t *testing.T) {
	v := r.Vec3{0, 0, 0}

	v.SetY(1)

	assert.Equal(t, v.Y(), float64(1))
}

func TestVec3SetZ(t *testing.T) {
	v := r.Vec3{0, 0, 0}

	v.SetZ(1)

	assert.Equal(t, v.Z(), float64(1))
}

func TestVec3SetX(t *testing.T) {
	v := r.Vec3{0, 0, 0}

	v.SetX(1)

	assert.Equal(t, v.X(), float64(1))
}

func TestVec3Add(t *testing.T) {
	v1 := r.Vec3{1, 2, 3}
	v2 := r.Vec3{1, 1, 1}

	v3 := v1.Add(v2)

	assert.Equal(t, v3, r.Vec3{2, 3, 4})
}

func TestVec3Subtract(t *testing.T) {
	v1 := r.Vec3{1, 2, 3}
	v2 := r.Vec3{1, 1, 1}

	v3 := v1.Subtract(v2)

	assert.Equal(t, v3, r.Vec3{0, 1, 2})
}

func TestVec3Multiply(t *testing.T) {
	v1 := r.Vec3{1, 2, 3}

	v2 := v1.Multiply(2)

	assert.Equal(t, v2, r.Vec3{2, 4, 6})
}

func TestVec3MultiplyV(t *testing.T) {
	v1 := r.NewVec3(2, 3, 4)

	v2 := r.NewVec3(3, 4, 5)

	assert.Equal(t, v1.MultiplyV(v2), r.NewVec3(6, 12, 20))
}

func TestVec3Divide(t *testing.T) {
	v1 := r.Vec3{2, 4, 8}

	v2 := v1.Divide(2)

	assert.Equal(t, v2, r.Vec3{1, 2, 4})
}

func TestVec3Unit(t *testing.T) {
	n1 := 3.0
	v1 := r.Vec3{n1, n1, n1}

	n2 := n1 / math.Sqrt(math.Pow(n1, 3))
	v2 := v1.Unit()

	assert.Equal(t, v2, r.Vec3{n2, n2, n2})
}

func TestVec3LengthSquared(t *testing.T) {
	v := r.Vec3{2, 3, 4}

	assert.Equal(t, v.LengthSquared(), float64(29))
}

func TestVec3Length(t *testing.T) {
	v := r.Vec3{2, 3, 4}

	assert.Equal(t, v.Length(), math.Sqrt(29))
}

func TestVec3Dot(t *testing.T) {
	v1 := r.Vec3{1, 2, 3}
	v2 := r.Vec3{2, 3, 4}

	v3 := v1.Dot(v2)

	assert.Equal(t, v3, float64(20))
}

func TestVec3Reflect(t *testing.T) {
	v1 := r.NewVec3(0.1, 0.1, 0.1)
	v2 := r.NewVec3(0.2, 0.2, 0.2)

	n := 0.076

	assert.Equal(t, v1.Reflect(v2), r.NewVec3(n, n, n))
}

func TestVec3NearZero(t *testing.T) {
	v1 := r.NewVec3(0, 0, 0)

	assert.True(t, v1.NearZero())

	for i := 0; i < 3; i++ {
		var v2 r.Vec3

		v2[i] = math.SmallestNonzeroFloat64

		assert.False(t, v2.NearZero())
	}
}

func TestVec3Refract(t *testing.T) {
	v1 := r.NewVec3(0.1, 0.1, 0.1)
	v2 := r.NewVec3(0.2, 0.2, 0.2)

	n := -0.06270264507705078

	assert.Equal(t, v1.Refract(v2, 1.5), r.NewVec3(n, n, n))
}

func TestVec3String(t *testing.T) {
	v := r.Vec3{1, 2, 3}

	assert.Equal(t, v.String(), "Vec3(1.000000, 2.000000, 3.000000)")
}
