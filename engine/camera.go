package engine

import (
	"math"
)

type Camera struct {
	AspectRatio     float64
	ViewportHeight  float64
	ViewportWidth   float64
	FocalLength     float64
	Origin          Point
	Horizontal      Vec
	Vertical        Vec
	LowerLeftCorner Vec
	// Vertical field of view in degrees
	VFoVDegrees float64
}

func ConstructCamera(cameraPos Point, vFoVDegrees, aspectRatio float64) *Camera {
	c := &Camera{}
	c.AspectRatio = aspectRatio
	c.Origin = cameraPos
	c.FocalLength = 1
	// fov
	c.VFoVDegrees = vFoVDegrees
	theta := degreesToRadians(vFoVDegrees)
	h := math.Tan(theta / 2)
	c.ViewportHeight = 2 * h
	c.ViewportWidth = c.AspectRatio * c.ViewportHeight
	c.Horizontal = Vec{c.ViewportWidth, 0, 0}
	c.Vertical = Vec{0, c.ViewportHeight, 0}
	c.LowerLeftCorner = c.Origin.
		Subtract(c.Horizontal.DivideScalar(2)).
		Subtract(c.Vertical.DivideScalar(2)).
		SubtractScalars(0, 0, c.FocalLength)
	return c
}

func (c *Camera) getRay(u, v float64) *Ray {
	return &Ray{
		Origin:    c.Origin,
		Direction: c.LowerLeftCorner.
			Add(c.Horizontal.MultiplyScalar(u)).
			Add(c.Vertical.MultiplyScalar(v)).
			Subtract(c.Origin),
	}
}
