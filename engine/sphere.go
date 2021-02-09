package engine

import "math"

type Sphere struct {
	Center   Point
	Radius   float64
	Material Material
}

// implement Hittable
func (s *Sphere) hit(ray *Ray, tMin float64, tMax float64) (*hitRecord, bool) {
	oc := ray.Origin.Subtract(s.Center)
	a := ray.Direction.LengthSquared()
	halfB := Dot(oc, ray.Direction)
	c := oc.LengthSquared() - s.Radius* s.Radius
	discriminant := halfB * halfB - a * c

	if discriminant > 0 {
		root := math.Sqrt(discriminant)
		temp := (-halfB - root) / a
		if temp < tMax && temp > tMin {
			p := ray.Interpolate(temp)
			outwardNormal := p.Subtract(s.Center).DivideScalar(s.Radius)
			record := &hitRecord{
				p:      p,
				t:      temp,
			}
			record.setFaceNormal(ray, outwardNormal)
			record.material = s.Material
			return record, true
		}
	}

	return nil, false
}

// implement Hittable
func (s *Sphere) boundingBox() (*aabb, bool) {
	return &aabb{
		min: s.Center.SubtractScalars(s.Radius, s.Radius, s.Radius),
		max: s.Center.AddScalars(s.Radius, s.Radius, s.Radius),
	}, true
}

// implement Hittable
func (s *Sphere) Translate(vec Vec) {
	s.Center = s.Center.Add(vec)
}
