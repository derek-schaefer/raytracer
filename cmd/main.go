package main

import (
	"bytes"
	"fmt"
	"math/rand"

	r "github.com/derek-schaefer/raytracer"
)

const (
	aspectRatio     = 16.0 / 9.0
	defocusAngle    = 0.6
	fieldOfView     = 20
	focalLength     = 1.0
	focusDistance   = 10.0
	imageWidth      = 800
	maxDepth        = 50
	samplesPerPixel = 500
	viewportHeight  = 2.0
)

var (
	lookFrom = r.NewPoint3(13, 2, 3)
	lookAt   = r.NewPoint3(0, 0, 0)
	viewup   = r.NewPoint3(0, 1, 0)
)

func main() {
	random := rand.New(rand.NewSource(1))

	world := r.NewHittables()

	groundMaterial := r.NewLambertian(r.LambertianOptions{Albedo: r.NewColor(0.5, 0.5, 0.5)})
	world.Add(r.NewSphere(r.SphereOptions{Center: r.NewPoint3(0, -1000, 0), Radius: 1000, Material: groundMaterial}))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := random.Float64()
			center := r.NewPoint3(float64(a)+0.9*random.Float64(), 0.2, float64(b)+0.9*random.Float64())

			if center.Subtract(r.NewVec3(4, 0.2, 0)).Length() > 0.9 {
				var sphereMaterial r.Material

				if chooseMat < 0.8 {
					// diffuse
					albedo := r.NewColorV(r.RandomVec3(random))
					sphereMaterial = r.NewLambertian(r.LambertianOptions{Albedo: albedo})
				} else if chooseMat < 0.95 {
					// metal
					albedo := r.NewColorV(r.RandomRangeVec3(random, 0.5, 1))
					fuzz := r.RandFloat64(random, 0, 0.5)
					sphereMaterial = r.NewMetal(r.MetalOptions{Albedo: albedo, Fuzz: fuzz})
				} else {
					// glass
					sphereMaterial = r.NewDielectric(r.DielectricOptions{IndexOfRefraction: 1.5})
				}

				world.Add(r.NewSphere(r.SphereOptions{Center: center, Radius: 0.2, Material: sphereMaterial}))
			}
		}
	}

	material1 := r.NewDielectric(r.DielectricOptions{IndexOfRefraction: 1.5})
	world.Add(r.NewSphere(r.SphereOptions{Center: r.NewPoint3(0, 1, 0), Radius: 1.0, Material: material1}))

	material2 := r.NewLambertian(r.LambertianOptions{Albedo: r.NewColor(0.4, 0.2, 0.1)})
	world.Add(r.NewSphere(r.SphereOptions{Center: r.NewPoint3(-4, 1, 0), Radius: 1.0, Material: material2}))

	material3 := r.NewMetal(r.MetalOptions{Albedo: r.NewColor(0.7, 0.6, 0.5), Fuzz: 0.0})
	world.Add(r.NewSphere(r.SphereOptions{Center: r.NewPoint3(4, 1, 0), Radius: 1.0, Material: material3}))

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
			Samples:       samplesPerPixel,
			ViewUp:        viewup,
		},
	)

	var buffer bytes.Buffer
	camera.Render(world).WritePPM(&buffer)
	fmt.Print(&buffer)
}
