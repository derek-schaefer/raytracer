package raytracer

import (
	"bytes"
	"fmt"
)

type Color = Vec3

func (c *Color) WritePPM(buf *bytes.Buffer) {
	r := uint8(255.999 * c.X())
	g := uint8(255.999 * c.Y())
	b := uint8(255.999 * c.Z())

	buf.WriteString(fmt.Sprintf("%d %d %d\n", r, g, b))
}
