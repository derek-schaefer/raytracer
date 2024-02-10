package raytracer

type Viewport struct {
	Height float64
	Width  float64
}

// Vector across the horizontal viewport edge.
func (v Viewport) U() Vec3 {
	return Vec3{v.Width, 0, 0}
}

// Vector down the vertical viewport edge.
func (v Viewport) V() Vec3 {
	return Vec3{0, -v.Height, 0}
}

// Horizontal delta vector from pixel to pixel.
func (v Viewport) DeltaU(w float64) Vec3 {
	return v.U().Divide(w)
}

// Vertical delta vector from pixel to pixel.
func (v Viewport) DeltaV(h float64) Vec3 {
	return v.V().Divide(h)
}

// Location of the upper left pixel.
func (v Viewport) Origin(c *Camera) Point3 {
	return c.Center.Subtract(Vec3{0, 0, c.FocalLength}).
		Subtract(v.U().Divide(2)).
		Subtract(v.V().Divide(2))
}
