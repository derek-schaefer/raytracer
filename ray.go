package raytracer

import "fmt"

type Ray struct {
	Origin    Point3
	Direction Vec3
}

func NewRay(origin Point3, direction Vec3) Ray {
	return Ray{Origin: origin, Direction: direction}
}

// Return the point on the ray at the specified value.
func (r Ray) At(t float64) Point3 {
	return r.Origin.Add(r.Direction.Multiply(t))
}

func (r Ray) String() string {
	return fmt.Sprintf("Ray(%s, %s)", r.Origin, r.Direction)
}
