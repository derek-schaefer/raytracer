package raytracer

import (
	"bytes"
	"fmt"
	"math"
)

var (
	ColorBlack = NewColor(0, 0, 0)
	ColorWhite = NewColor(1, 1, 1)

	intensity = NewInterval(0.0, 0.999)
)

type Color struct {
	V Vec3
}

func NewColor(r, g, b float64) Color {
	return Color{NewVec3(r, g, b)}
}

func NewColorV(v Vec3) Color {
	return Color{v}
}

// Convert the color from linear space to gamma space.
func (c Color) LinearToGamma() Color {
	var v Vec3

	for i := 0; i < len(c.V); i++ {
		v[i] = math.Sqrt(c.V[i])
	}

	return NewColorV(v)
}

// Write the color in PPM format to a buffer.
func (c Color) WritePPM(buf *bytes.Buffer) {
	r := uint8(256 * intensity.Clamp(c.V.X()))
	g := uint8(256 * intensity.Clamp(c.V.Y()))
	b := uint8(256 * intensity.Clamp(c.V.Z()))

	buf.WriteString(fmt.Sprintf("%d %d %d\n", r, g, b))
}
