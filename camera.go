package raytracer

import (
	"math"
	"math/rand"
)

type CameraOptions struct {
	AspectRatio    float64
	Center         Point3
	FocalLength    float64
	ImageWidth     int
	MaxDepth       int
	Samples        int
	ViewportHeight float64
}

type Camera struct {
	CameraOptions

	deltaU      Vec3
	deltaV      Vec3
	imageHeight int
	iorigin     Vec3
	viewport    Viewport
	vorigin     Vec3
}

func NewCamera(options CameraOptions) *Camera {
	return &Camera{CameraOptions: options}
}

// Use the world of hittable objects to produce a new image.
func (c *Camera) Render(world *Hittables) *Image {
	c.initialize()

	image := NewImage(c.ImageWidth, c.imageHeight)

	for j := 0; j < image.Height; j++ {
		for i := 0; i < image.Width; i++ {
			var pixel Vec3

			for s := 0; s < c.Samples; s++ {
				ray := c.getRay(i, j)
				pixel = pixel.Add(c.rayColor(ray, c.MaxDepth, world).V)
			}

			scale := 1.0 / float64(c.Samples)
			pixel.SetX(pixel.X() * scale)
			pixel.SetY(pixel.Y() * scale)
			pixel.SetZ(pixel.Z() * scale)

			image.Set(i, j, NewColor(pixel).LinearToGamma())
		}
	}

	return image
}

// Assign a variety of properties used by multiple camera methods.
func (c *Camera) initialize() {
	c.imageHeight = int(float64(c.ImageWidth) / c.AspectRatio)

	c.viewport = Viewport{
		Height: c.ViewportHeight,
		Width:  c.ViewportHeight * float64(c.ImageWidth) / float64(c.imageHeight),
	}

	c.vorigin = c.Center.Subtract(Vec3{0, 0, c.FocalLength}).
		Subtract(c.viewport.U().Divide(2)).
		Subtract(c.viewport.V().Divide(2))

	c.iorigin = c.vorigin.Add(
		(c.viewport.DeltaU(float64(c.ImageWidth)).
			Add(c.viewport.DeltaV(float64(c.imageHeight)))).
			Multiply(0.5),
	)

	c.deltaU = c.viewport.DeltaU(float64(c.ImageWidth))
	c.deltaV = c.viewport.DeltaV(float64(c.imageHeight))
}

// Get a randomly sampled camera ray for the pixel at location i, j.
func (c *Camera) getRay(i, j int) Ray {
	di := c.deltaU.Multiply(float64(i))
	dj := c.deltaV.Multiply(float64(j))

	pixelCenter := c.iorigin.Add(di).Add(dj)
	pixelSample := pixelCenter.Add(c.pixelSampleSquare())

	rayOrigin := c.Center
	rayDirection := pixelSample.Subtract(rayOrigin)

	return Ray{Origin: rayOrigin, Direction: rayDirection}
}

// Returns a random point in the square surrounding a pixel at the origin.
func (c *Camera) pixelSampleSquare() Vec3 {
	px := -0.5 + rand.Float64()
	py := -0.5 + rand.Float64()

	dx := c.deltaU.Multiply(px)
	dy := c.deltaV.Multiply(py)

	return dx.Add(dy)
}

// Determine the ray color based on the object it hits, it any.
func (c *Camera) rayColor(ray Ray, depth int, world *Hittables) Color {
	if depth <= 0 {
		return NewColor(NewVec3(0, 0, 0))
	}

	// Near zero min value to avoid shadow acne due to floating point errors
	if h, ok := world.Hit(ray, NewInterval(0.001, math.Inf(1))); ok {
		direction := h.N.Add(RandomUnitVec3())
		return NewColor(c.rayColor(Ray{h.P, direction}, depth-1, world).V.Multiply(0.5))
	}

	a := 0.5 * (ray.Direction.Unit().Y() + 1.0)
	v := Vec3{1.0, 1.0, 1.0}.Multiply(1.0 - a).
		Add(Vec3{0.5, 0.7, 1.0}.Multiply(a))

	return NewColor(v)
}
