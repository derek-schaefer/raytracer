package raytracer

import (
	"bytes"
	"fmt"
)

type Image struct {
	Width  int
	Height int
	Pixels []Color
}

func NewImage(w, h int) *Image {
	return &Image{
		Width:  w,
		Height: h,
		Pixels: make([]Color, w*h),
	}
}

func (img *Image) Get(x, y int) Color {
	return img.Pixels[x*img.Height+y]
}

func (img *Image) Set(c Color, x, y int) {
	img.Pixels[x*img.Height+y] = c
}

func (img *Image) WritePPM() *bytes.Buffer {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", img.Width, img.Height))

	for j := 0; j < img.Height; j++ {
		for i := 0; i < img.Width; i++ {
			img.Get(i, j).WritePPM(&buffer)
		}
	}

	return &buffer
}
