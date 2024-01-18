package raytracer

import (
	"bytes"
	"fmt"
)

type Image struct {
	Width  int
	Height int
}

func WritePPM(image Image) *bytes.Buffer {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", image.Width, image.Height))

	for j := 0; j < image.Height; j++ {
		for i := 0; i < image.Width; i++ {
			r := float64(i) / (float64(image.Width) - 1)
			g := float64(j) / (float64(image.Height) - 1)
			b := 0.0

			ir := int32(255.999 * r)
			ig := int32(255.999 * g)
			ib := int32(255.999 * b)

			buffer.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
		}
	}

	return &buffer
}
