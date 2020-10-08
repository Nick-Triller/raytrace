package engine

import "math"

// Related resource:
// https://www.scratchapixel.com/lessons/3d-basic-rendering/ray-tracing-rendering-a-triangle/geometry-of-a-triangle

type Triangle struct {
	V1   Point
	V2   Point
	V3   Point
	Material Material
	Normal Vec
}

func NewTriangle(v1, v2, v3 Point, m Material) *Triangle {
	// Calculate normal vector
	a := v2.Subtract(v1)
	b := v3.Subtract(v1)
	n := Cross(a, b).Normalize()
	return &Triangle{
		V1:       v1,
		V2:       v2,
		V3:       v3,
		Material: m,
		Normal:   n,
	}
}

// implement Hittable
func (t *Triangle) hit(ray *Ray, tMin float64, tMax float64) (*hitRecord, bool) {
	denominator := Dot(t.Normal, ray.Direction)
	// Protect from zero-division
	planeCenter := t.V1
	if math.Abs(denominator) > epsilon {
		time := Dot(planeCenter.Subtract(ray.Origin), t.Normal) / denominator
		if (time < tMax && time > tMin) && math.Abs(time) > epsilon {
			// Ray hits triangle plane, check if hit point is contained in triangle
			hitPoint := ray.Interpolate(time)
			if t.contains(hitPoint) {
				record := &hitRecord{
					p:         hitPoint,
					t:         time,
					material:  t.Material,
				}
				outwardNormal := t.Normal
				record.setFaceNormal(ray, outwardNormal)
				return record, true
			}
		}
	}
	return nil, false
}

func (t *Triangle) contains(p Point) bool {
	// edge1
	edge1 := t.V2.Subtract(t.V1)
	vp1 := p.Subtract(t.V1)
	c := Cross(edge1, vp1)
	if Dot(t.Normal, c) < 0 {
		return false
	}
	// edge2
	edge2 := t.V3.Subtract(t.V2)
	vp2 := p.Subtract(t.V2)
	c = Cross(edge2, vp2)
	if Dot(t.Normal, c) < 0 {
		return false
	}
	// edge3
	edge3 := t.V1.Subtract(t.V3)
	vp3 := p.Subtract(t.V3)
	c = Cross(edge3, vp3)
	if Dot(t.Normal, c) < 0 {
		return false
	}
	return true
}

// implement Hittable
func (t *Triangle) Translate(vec Vec) {
	t.V1 = t.V1.Add(vec)
	t.V2 = t.V2.Add(vec)
	t.V3 = t.V3.Add(vec)
}
