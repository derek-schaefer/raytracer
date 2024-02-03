package main

import (
	"bytes"
	"fmt"
	"math"

	r "github.com/derek-schaefer/raytracer"
)

const (
	imageWidth     = 400
	focalLength    = 1.0
	viewportHeight = 2.0
	aspectRatio    = 16.0 / 9.0
)

func hitSphere(center r.Point3, radius float64, r r.Ray) float64 {
	oc := r.Origin.Subtract(center)
	a := r.Direction.LengthSquared()
	halfB := oc.Dot(r.Direction)
	c := oc.LengthSquared() - radius*radius
	discriminant := halfB*halfB - a*c

	if discriminant < 0 {
		return -1.0
	} else {
		return (-halfB - math.Sqrt(discriminant)) / a
	}
}

func rayColor(ray r.Ray) r.Color {
	t := hitSphere(r.Point3{0, 0, -1}, 0.5, ray)

	if t > 0 {
		n := ray.At(t).Subtract(r.Vec3{0, 0, -1}).Unit()
		return r.NewColor(r.Vec3{n.X() + 1, n.Y() + 1, n.Z() + 1}.Multiply(0.5))
	}

	unitDirection := ray.Direction.Unit()
	a := 0.5 * (unitDirection.Y() + 1.0)
	c := r.Vec3{1.0, 1.0, 1.0}.Multiply(1.0 - a).Add(r.Vec3{0.5, 0.7, 1.0}.Multiply(a))

	return r.NewColor(c)
}

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

			image.Set(rayColor(ray), i, j)
		}
	}

	var buffer bytes.Buffer
	image.WritePPM(&buffer)
	fmt.Print(&buffer)
}
