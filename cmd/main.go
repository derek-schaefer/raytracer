package main

import (
	"bytes"
	"fmt"
	"math/rand"

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
	random := rand.New(rand.NewSource(1))

	world := r.NewHittables()

	materialGround := r.NewLambertian(r.LambertianOptions{Albedo: r.NewColor(r.NewVec3(0.8, 0.8, 0.0)), Random: random})
	materialCenter := r.NewLambertian(r.LambertianOptions{Albedo: r.NewColor(r.NewVec3(0.1, 0.2, 0.5)), Random: random})
	materialLeft := r.NewDielectric(r.DielectricOptions{IndexOfRefraction: 1.5, Random: random})
	materialRight := r.NewMetal(r.MetalOptions{Albedo: r.NewColor(r.NewVec3(0.8, 0.6, 0.2)), Fuzz: 0.0, Random: random})

	world.Add(r.NewSphere(r.NewPoint3(0.0, -100.5, -1.0), 100, materialGround))
	world.Add(r.NewSphere(r.NewPoint3(0.0, 0.0, -1.0), 0.5, materialCenter))
	world.Add(r.NewSphere(r.NewPoint3(-1.0, 0.0, -1.0), 0.5, materialLeft))
	world.Add(r.NewSphere(r.NewPoint3(-1.0, 0.0, -1.0), -0.4, materialLeft))
	world.Add(r.NewSphere(r.NewPoint3(1.0, 0.0, -1.0), 0.5, materialRight))

	camera := r.NewCamera(
		r.CameraOptions{
			AspectRatio:    aspectRatio,
			Center:         r.NewPoint3(0, 0, 0),
			FocalLength:    focalLength,
			ImageWidth:     imageWidth,
			MaxDepth:       maxDepth,
			Random:         random,
			Samples:        samplesPerPixel,
			ViewportHeight: viewportHeight,
		},
	)

	var buffer bytes.Buffer
	camera.Render(world).WritePPM(&buffer)
	fmt.Print(&buffer)
}
