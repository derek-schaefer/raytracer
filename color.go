package raytracer

import (
	"bytes"
	"fmt"
)

var (
	intensity = NewInterval(0.0, 0.999)
)

type Color struct {
	V Vec3
}

func NewColor(v Vec3) Color {
	return Color{v}
}

// Write the color in PPM format to a buffer.
func (c Color) WritePPM(buf *bytes.Buffer) {
	r := uint8(256 * intensity.Clamp(c.V.X()))
	g := uint8(256 * intensity.Clamp(c.V.Y()))
	b := uint8(256 * intensity.Clamp(c.V.Z()))

	buf.WriteString(fmt.Sprintf("%d %d %d\n", r, g, b))
}
