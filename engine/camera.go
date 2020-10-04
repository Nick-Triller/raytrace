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

func ConstructCamera(lookfrom Point, lookat Point, viewUp Vec, vFoVDegrees, aspectRatio float64) *Camera {
	c := &Camera{}
	c.AspectRatio = aspectRatio
	c.FocalLength = 1
	// fov
	c.VFoVDegrees = vFoVDegrees
	theta := degreesToRadians(vFoVDegrees)
	h := math.Tan(theta / 2)
	c.ViewportHeight = 2 * h
	c.ViewportWidth = c.AspectRatio * c.ViewportHeight
	c.Vertical = Vec{0, c.ViewportHeight, 0}

	w := lookfrom.Subtract(lookat).Normalize()
	u := Cross(viewUp, w).Normalize()
	v := Cross(w, u)
	c.Origin = lookfrom
	c.Horizontal = u.MultiplyScalar(c.ViewportWidth)
	c.Vertical = v.MultiplyScalar(c.ViewportHeight)
	c.LowerLeftCorner = c.Origin.
		Subtract(c.Horizontal.DivideScalar(2)).
		Subtract(c.Vertical.DivideScalar(2)).
		Subtract(w)
	return c
}

func (c *Camera) getRay(s, t float64) *Ray {
	return &Ray{
		Origin:    c.Origin,
		Direction: c.LowerLeftCorner.
			Add(c.Horizontal.MultiplyScalar(s)).
			Add(c.Vertical.MultiplyScalar(t)).
			Subtract(c.Origin),
	}
}
