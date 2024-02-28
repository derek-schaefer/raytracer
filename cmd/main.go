package main

import (
	"bytes"
	"fmt"

	r "github.com/derek-schaefer/raytracer"
)

const (
	aspectRatio     = 16.0 / 9.0
	focalLength     = 1.0
	imageWidth      = 400
	maxDepth        = 50
	samplesPerPixel = 100
	viewportHeight  = 2.0
)

func main() {
	world := r.NewHittables()

	world.Add(r.Sphere{Center: r.Point3{0, 0, -1}, Radius: 0.5})
	world.Add(r.Sphere{Center: r.Point3{0, -100.5, -1}, Radius: 100})

	camera := r.NewCamera(
		r.CameraOptions{
			AspectRatio:    aspectRatio,
			Center:         r.Point3{0, 0, 0},
			FocalLength:    focalLength,
			ImageWidth:     imageWidth,
			MaxDepth:       maxDepth,
			Samples:        samplesPerPixel,
			ViewportHeight: viewportHeight,
		},
	)

	var buffer bytes.Buffer
	camera.Render(world).WritePPM(&buffer)
	fmt.Print(&buffer)
}
