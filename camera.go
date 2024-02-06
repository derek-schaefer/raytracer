package raytracer

import "math"

type Camera struct {
	AspectRatio    float64
	Center         Point3
	FocalLength    float64
	ImageWidth     int
	ViewportHeight float64
}

func (c *Camera) Render(world *Hittables) *Image {
	imageHeight := int(float64(c.ImageWidth) / c.AspectRatio)

	viewport := Viewport{
		Height: c.ViewportHeight,
		Width:  c.ViewportHeight * float64(c.ImageWidth) / float64(imageHeight),
	}

	vorigin := c.Center.Subtract(Vec3{0, 0, c.FocalLength}).
		Subtract(viewport.U().Divide(2)).
		Subtract(viewport.V().Divide(2))

	image := NewImage(c.ImageWidth, imageHeight)

	iorigin := vorigin.Add(
		(viewport.DeltaU(float64(image.Width)).
			Add(viewport.DeltaV(float64(image.Height)))).
			Multiply(0.5),
	)

	deltaU := viewport.DeltaU(float64(image.Width))
	deltaV := viewport.DeltaV(float64(image.Height))

	for j := 0; j < image.Height; j++ {
		for i := 0; i < image.Width; i++ {
			pixelCenter := iorigin.Add(deltaU.Multiply(float64(i))).Add(deltaV.Multiply(float64(j)))
			rayDirection := pixelCenter.Subtract(c.Center)
			ray := Ray{Origin: c.Center, Direction: rayDirection}

			image.Set(c.rayColor(ray, world), i, j)
		}
	}

	return image
}

func (c *Camera) rayColor(ray Ray, world *Hittables) Color {
	if h, ok := world.Hit(ray, NewInterval(0, math.Inf(1))); ok {
		return NewColor(h.N.Add(Vec3{1, 1, 1}).Multiply(0.5))
	}

	a := 0.5 * (ray.Direction.Unit().Y() + 1.0)

	return NewColor(
		Vec3{1.0, 1.0, 1.0}.Multiply(1.0 - a).
			Add(Vec3{0.5, 0.7, 1.0}.Multiply(a)),
	)
}
