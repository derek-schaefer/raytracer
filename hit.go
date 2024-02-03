package raytracer

type Hit struct {
	// Intersection point
	P Point3
	// Normal vector
	N Vec3
	// Ray parameter
	T float64
	// Front facing
	F bool
}

// Sets the hit record normal vector.
// NOTE: the parameter `outwardNormal` is assumed to have unit length.
func (h *Hit) SetFaceNormal(r Ray, outwardNormal Vec3) {
	h.F = r.Direction.Dot(outwardNormal) < 0

	if h.F {
		h.N = outwardNormal
	} else {
		h.N = outwardNormal.Multiply(-1)
	}
}
