package main

import (
	"fmt"

	"github.com/derek-schaefer/raytracer"
)

const imageWidth = 256

func main() {
	img := raytracer.NewImage(imageWidth, imageWidth)

	for j := 0; j < img.Height; j++ {
		for i := 0; i < img.Width; i++ {
			c := raytracer.NewColor(
				raytracer.Vec3{
					float64(i) / (float64(img.Width) - 1),
					float64(j) / (float64(img.Height) - 1),
					0,
				},
			)

			img.Set(c, i, j)
		}
	}

	fmt.Print(img.WritePPM())
}
