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

	materialGround := r.NewLambertian(r.NewColor(r.NewVec3(0.8, 0.8, 0.0)))
	materialCenter := r.NewLambertian(r.NewColor(r.NewVec3(0.1, 0.2, 0.5)))
	materialLeft := r.NewDielectric(1.5)
	materialRight := r.NewMetal(r.NewColor(r.NewVec3(0.8, 0.6, 0.2)), 0.0)

	world.Add(r.NewSphere(r.NewPoint3(0.0, -100.5, -1.0), 100, materialGround))
	world.Add(r.NewSphere(r.NewPoint3(0.0, 0.0, -1.0), 0.5, materialCenter))
	world.Add(r.NewSphere(r.NewPoint3(-1.0, 0.0, -1.0), 0.5, materialLeft))
	world.Add(r.NewSphere(r.NewPoint3(1.0, 0.0, -1.0), 0.5, materialRight))

	camera := r.NewCamera(
		r.CameraOptions{
			AspectRatio:    aspectRatio,
			Center:         r.NewPoint3(0, 0, 0),
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
