package raytracer

type Point3 = Vec3

func NewPoint3(x, y, z float64) Point3 {
	return Point3{x, y, z}
}
