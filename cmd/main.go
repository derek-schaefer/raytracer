package main

import (
	"fmt"

	"github.com/derek-schaefer/raytracer"
)

func main() {
	fmt.Print(raytracer.WritePPM(raytracer.Image{Width: 256, Height: 256}))
}
