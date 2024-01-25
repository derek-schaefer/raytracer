package raytracer

import (
	"bytes"
	"fmt"
)

type Color struct {
	v Vec3
}

func NewColor(v Vec3) Color {
	return Color{v}
}

// Write the color in PPM format to a buffer.
func (c Color) WritePPM(buf *bytes.Buffer) {
	r := uint8(255.999 * c.v.X())
	g := uint8(255.999 * c.v.Y())
	b := uint8(255.999 * c.v.Z())

	buf.WriteString(fmt.Sprintf("%d %d %d\n", r, g, b))
}
