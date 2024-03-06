package raytracer

import (
	"log"
	"math"
	"math/rand"
)

type CameraOptions struct {
	// Ratio of image width over height
	AspectRatio float64
	// Vertical view angle (field of view)
	FieldOfView float64
	// Rendered image width in pixel count
	ImageWidth int
	// Point camera is looking at
	LookAt Point3
	// Point camera is looking from
	LookFrom Point3
	// Maximum number of ray bounces into scene
	MaxDepth int
	// Count of random samples for each pixel
	Samples int
	// Camera-relative "up" direction
	ViewUp Vec3
	// Variation angle of rays through each pixel
	DefocusAngle float64
	// Distance from camera lookfrom point to plane of perfect focus
	FocusDistance float64
	// Number of concurrent render workers
	Workers int
}

type Camera struct {
	CameraOptions

	imageHeight int

	center  Point3
	deltaU  Vec3
	deltaV  Vec3
	iorigin Vec3

	u Vec3
	v Vec3
	w Vec3

	defocusDiskU Vec3
	defocusDiskV Vec3
}

type job struct {
	x int
	y int
	r *rand.Rand
}

type result struct {
	x int
	y int
	c Color
}

func NewCamera(options CameraOptions) *Camera {
	if options.Workers == 0 {
		options.Workers = 1
	}

	return &Camera{CameraOptions: options}
}

// Use the world of hittable objects to produce a new image.
func (c *Camera) Render(world *Hittables) *Image {
	c.initialize()

	image := NewImage(c.ImageWidth, c.imageHeight)

	size := len(image.Pixels)
	jobs := make(chan job, size)
	results := make(chan result, size)

	log.Printf("worker count: %d\n", c.Workers)

	for w := 0; w < c.Workers; w++ {
		go c.renderWorker(world, jobs, results)
	}

	for y := 0; y < image.Height; y++ {
		for x := 0; x < image.Width; x++ {
			jobs <- job{x, y, rand.New(rand.NewSource(rand.Int63()))}
		}
	}

	close(jobs)

	log.Printf("job count: %d\n", size)

	for n := 1; n <= size; n++ {
		r := <-results

		image.Set(r.x, r.y, r.c)

		if n%(size/100) == 0 {
			log.Printf("completed: %d%%\n", int(float32(n)/float32(size)*100))
		}
	}

	log.Println("done")

	return image
}

// For each job, trace the indicated pixel with multi-sampling and send the result.
func (c *Camera) renderWorker(world *Hittables, jobs <-chan job, results chan<- result) {
	for j := range jobs {
		var pixel Vec3

		for s := 0; s < c.Samples; s++ {
			pixel = pixel.Add(c.rayColor(j.r, c.getRay(j.r, j.x, j.y), c.MaxDepth, world).V)
		}

		scale := 1.0 / float64(c.Samples)
		pixel.SetX(pixel.X() * scale)
		pixel.SetY(pixel.Y() * scale)
		pixel.SetZ(pixel.Z() * scale)

		results <- result{j.x, j.y, NewColorV(pixel).LinearToGamma()}
	}
}

// Assign a variety of properties used by multiple camera methods.
func (c *Camera) initialize() {
	c.imageHeight = int(float64(c.ImageWidth) / c.AspectRatio)

	if c.imageHeight < 1 {
		c.imageHeight = 1
	}

	c.center = c.LookFrom

	theta := c.FieldOfView * (math.Pi / 180.0)
	h := math.Tan(theta / 2)
	viewportHeight := 2 * h * c.FocusDistance
	viewportWidth := viewportHeight * (float64(c.ImageWidth) / float64(c.imageHeight))

	c.w = c.LookFrom.Subtract(c.LookAt).Unit()
	c.u = c.ViewUp.Cross(c.w).Unit()
	c.v = c.w.Cross(c.u)

	viewportU := c.u.Multiply(viewportWidth)
	viewportV := c.v.Multiply(-1).Multiply(viewportHeight)

	c.deltaU = viewportU.Divide(float64(c.ImageWidth))
	c.deltaV = viewportV.Divide(float64(c.imageHeight))

	viewportUpperLeft := c.center.
		Subtract(c.w.Multiply(c.FocusDistance)).
		Subtract(viewportU.Divide(2)).
		Subtract(viewportV.Divide(2))

	c.iorigin = viewportUpperLeft.Add(c.deltaU.Add(c.deltaV).Multiply(0.5))

	defocusRadius := c.FocusDistance * math.Tan(c.DefocusAngle/2*(math.Pi/180))

	c.defocusDiskU = c.u.Multiply(defocusRadius)
	c.defocusDiskV = c.v.Multiply(defocusRadius)
}

// Get a randomly-sampled camera ray for the pixel at location i,j, originating from the camera defocus disk.
func (c *Camera) getRay(random *rand.Rand, i, j int) Ray {
	di := c.deltaU.Multiply(float64(i))
	dj := c.deltaV.Multiply(float64(j))

	pixelCenter := c.iorigin.Add(di).Add(dj)
	pixelSample := pixelCenter.Add(c.pixelSampleSquare(random))

	var rayOrigin Vec3

	if c.DefocusAngle <= 0 {
		rayOrigin = c.center
	} else {
		rayOrigin = c.defocusDiskSample(random)
	}

	rayDirection := pixelSample.Subtract(rayOrigin)

	return Ray{Origin: rayOrigin, Direction: rayDirection}
}

// Returns a random point in the camera defocus disk.
func (c *Camera) defocusDiskSample(r *rand.Rand) Point3 {
	p := RandomUnitDiskVec3(r)

	return c.center.
		Add(c.defocusDiskU.Multiply(p.X())).
		Add(c.defocusDiskV.Multiply(p.Y()))
}

// Returns a random point in the square surrounding a pixel at the origin.
func (c *Camera) pixelSampleSquare(random *rand.Rand) Vec3 {
	px := -0.5 + random.Float64()
	py := -0.5 + random.Float64()

	dx := c.deltaU.Multiply(px)
	dy := c.deltaV.Multiply(py)

	return dx.Add(dy)
}

// Determine the ray color based on the object it hits, it any.
func (c *Camera) rayColor(random *rand.Rand, ray Ray, depth int, world *Hittables) Color {
	if depth <= 0 {
		return ColorBlack
	}

	// Near zero min value to avoid shadow acne due to floating point errors
	if h, ok := world.Hit(ray, NewInterval(0.001, math.Inf(1))); ok {
		if scattered, attenuation, ok := h.Material.Scatter(random, ray, h); ok {
			return NewColorV(attenuation.V.MultiplyV(c.rayColor(random, scattered, depth-1, world).V))
		}
		return ColorBlack
	}

	a := 0.5 * (ray.Direction.Unit().Y() + 1.0)
	v := Vec3{1.0, 1.0, 1.0}.Multiply(1.0 - a).
		Add(Vec3{0.5, 0.7, 1.0}.Multiply(a))

	return NewColorV(v)
}
