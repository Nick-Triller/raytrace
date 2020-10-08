package engine

import "math"

type Plane struct {
	Center   Point
	Material Material
	Normal Vec
}

const epsilon float64 = 0.00001

// implement Hittable
func (p *Plane) hit(ray *Ray, tMin, tMax float64) (*hitRecord, bool) {
	denominator := Dot(p.Normal, ray.Direction)
	// Protect from zero-division
	if math.Abs(denominator) > epsilon {
		t := Dot(p.Center.Subtract(ray.Origin), p.Normal) / denominator
		if (t < tMax && t > tMin) && math.Abs(t) > epsilon {
			hitPoint := ray.Interpolate(t)
			record := &hitRecord{
				p:         hitPoint,
				t:         t,
				material:  p.Material,
			}
			outwardNormal := p.Normal
			record.setFaceNormal(ray, outwardNormal)
			return record, true
		}
	}
	return nil, false
}

// implement Hittable
func (p *Plane) Translate(vec Vec) {
	p.Center = p.Center.Add(vec)
}
