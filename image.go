package raytracer

import (
	"bytes"
	"fmt"
	"log"
)

type Image struct {
	Width  int
	Height int
}

func (image Image) WritePPM() *bytes.Buffer {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", image.Width, image.Height))

	for j := 0; j < image.Height; j++ {
		log.Printf("Scanlines remaining: %d\n", image.Height-j)

		for i := 0; i < image.Width; i++ {
			c := Color{
				float64(i) / (float64(image.Width) - 1),
				float64(j) / (float64(image.Height) - 1),
				0,
			}

			c.WritePPM(&buffer)
		}
	}

	log.Println("Done.")

	return &buffer
}
