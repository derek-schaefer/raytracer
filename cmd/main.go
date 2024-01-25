package main

import (
	"fmt"

	r "github.com/derek-schaefer/raytracer"
)

const (
	imageWidth     = 400
	focalLength    = 1.0
	viewportHeight = 2.0
	aspectRatio    = 16.0 / 9.0
)

func main() {
	imageHeight := int(imageWidth / aspectRatio)

	camera := r.Camera{
		FocalLength: focalLength,
		Center:      r.Point3{0, 0, 0},
	}

	viewport := r.Viewport{
		Height: viewportHeight,
		Width:  viewportHeight * imageWidth / float64(imageHeight),
	}

	vorigin := camera.Center.Subtract(r.Vec3{0, 0, camera.FocalLength}).
		Subtract(viewport.U().Divide(2)).
		Subtract(viewport.V().Divide(2))

	image := r.NewImage(imageWidth, imageHeight)

	iorigin := vorigin.Add(
		(viewport.DeltaU(float64(image.Width)).
			Add(viewport.DeltaV(float64(imageHeight)))).
			Multiply(0.5),
	)

	deltaU := viewport.DeltaU(float64(image.Width))
	deltaV := viewport.DeltaV(float64(image.Height))

	for j := 0; j < image.Height; j++ {
		for i := 0; i < image.Width; i++ {
			pixelCenter := iorigin.Add(deltaU.Multiply(float64(i))).Add(deltaV.Multiply(float64(j)))
			rayDirection := pixelCenter.Subtract(camera.Center)
			ray := r.Ray{Origin: camera.Center, Direction: rayDirection}

			unitDirection := ray.Direction.Unit()
			a := 0.5 * (unitDirection.Y() + 1.0)
			c := r.Vec3{1.0, 1.0, 1.0}.Multiply(1.0 - a).Add(r.Vec3{0.5, 0.7, 1.0}.Multiply(a))

			image.Set(r.NewColor(c), i, j)
		}
	}

	fmt.Print(image.WritePPM())
}
