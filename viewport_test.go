package raytracer_test

import (
	"testing"

	r "github.com/derek-schaefer/raytracer"
)

func TestViewportU(t *testing.T) {
	v := r.Viewport{Width: 16, Height: 9}

	if v.U() != (r.Vec3{16, 0, 0}) {
		t.Fail()
	}
}

func TestViewportV(t *testing.T) {
	v := r.Viewport{Width: 16, Height: 9}

	if v.V() != (r.Vec3{0, -9, 0}) {
		t.Fail()
	}
}

func TestViewportDeltaU(t *testing.T) {
	v := r.Viewport{Width: 16, Height: 9}

	if v.DeltaU(2) != (r.Vec3{8, 0, 0}) {
		t.Fail()
	}
}

func TestViewportDeltaV(t *testing.T) {
	v := r.Viewport{Width: 16, Height: 9}

	if v.DeltaV(3) != (r.Vec3{0, -3, 0}) {
		t.Fail()
	}
}

func TestViewportOrigin(t *testing.T) {
	v := r.Viewport{Width: 16, Height: 9}

	c := r.Camera{
		Center:      r.Point3{0, 0, 0},
		FocalLength: 1.0,
	}

	if v.Origin(c) != (r.Point3{-8, 4.5, -1}) {
		t.Fail()
	}
}
