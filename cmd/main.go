package main

import (
	"bytes"
	"fmt"
	"math/rand"

	r "github.com/derek-schaefer/raytracer"
)

const (
	aspectRatio     = 16.0 / 9.0
	defocusAngle    = 10.0
	fieldOfView     = 20
	focalLength     = 1.0
	focusDistance   = 3.4
	imageWidth      = 400
	maxDepth        = 50
	samplesPerPixel = 100
	viewportHeight  = 2.0
)

var (
	lookFrom = r.NewPoint3(-2, 2, 1)
	lookAt   = r.NewPoint3(0, 0, -1)
	vup      = r.NewPoint3(0, 1, 0)
)

func main() {
	random := rand.New(rand.NewSource(1))

	world := r.NewHittables()

	materialGround := r.NewLambertian(r.LambertianOptions{Albedo: r.NewColor(r.NewVec3(0.8, 0.8, 0.0)), Random: random})
	materialCenter := r.NewLambertian(r.LambertianOptions{Albedo: r.NewColor(r.NewVec3(0.1, 0.2, 0.5)), Random: random})
	materialLeft := r.NewDielectric(r.DielectricOptions{IndexOfRefraction: 1.5, Random: random})
	materialRight := r.NewMetal(r.MetalOptions{Albedo: r.NewColor(r.NewVec3(0.8, 0.6, 0.2)), Fuzz: 0.0, Random: random})

	world.Add(r.NewSphere(r.SphereOptions{Center: r.NewPoint3(0.0, -100.5, -1.0), Radius: 100, Material: materialGround}))
	world.Add(r.NewSphere(r.SphereOptions{Center: r.NewPoint3(0.0, 0.0, -1.0), Radius: 0.5, Material: materialCenter}))
	world.Add(r.NewSphere(r.SphereOptions{Center: r.NewPoint3(-1.0, 0.0, -1.0), Radius: 0.5, Material: materialLeft}))
	world.Add(r.NewSphere(r.SphereOptions{Center: r.NewPoint3(-1.0, 0.0, -1.0), Radius: -0.4, Material: materialLeft}))
	world.Add(r.NewSphere(r.SphereOptions{Center: r.NewPoint3(1.0, 0.0, -1.0), Radius: 0.5, Material: materialRight}))

	camera := r.NewCamera(
		r.CameraOptions{
			AspectRatio:   aspectRatio,
			DefocusAngle:  defocusAngle,
			FieldOfView:   fieldOfView,
			FocusDistance: focusDistance,
			ImageWidth:    imageWidth,
			LookAt:        lookAt,
			LookFrom:      lookFrom,
			MaxDepth:      maxDepth,
			Random:        random,
			Samples:       samplesPerPixel,
			ViewUp:        vup,
		},
	)

	var buffer bytes.Buffer
	camera.Render(world).WritePPM(&buffer)
	fmt.Print(&buffer)
}
