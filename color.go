package raytracer

import (
	"bytes"
	"fmt"
	"math"
)

var (
	intensity = NewInterval(0.0, 0.999)

	ColorBlack = NewColor(NewVec3(0, 0, 0))
	ColorWhite = NewColor(NewVec3(1, 1, 1))
)

type Color struct {
	V Vec3
}

func NewColor(v Vec3) Color {
	return Color{v}
}

// Convert the color from linear space to gamma space.
func (c Color) LinearToGamma() Color {
	var v Vec3

	for i := 0; i < len(c.V); i++ {
		v[i] = math.Sqrt(c.V[i])
	}

	return NewColor(v)
}

// Write the color in PPM format to a buffer.
func (c Color) WritePPM(buf *bytes.Buffer) {
	r := uint8(256 * intensity.Clamp(c.V.X()))
	g := uint8(256 * intensity.Clamp(c.V.Y()))
	b := uint8(256 * intensity.Clamp(c.V.Z()))

	buf.WriteString(fmt.Sprintf("%d %d %d\n", r, g, b))
}
