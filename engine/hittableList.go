package engine

import "math"

type HittableList struct {
	Objects []Hittable
}

// implement Hittable on HittableList
func (hl *HittableList) hit(ray *Ray, tMin float64, tMax float64) (*hitRecord, bool) {
	var closestRec *hitRecord = nil
	var hitAnything bool
	closest := tMax

	for _, object := range hl.Objects {
		record, hit := object.hit(ray, tMin, closest)
		if hit {
			hitAnything = true
			closest = record.t
			closestRec = record
		}
	}
	return closestRec, hitAnything
}

// implement Hittable on HittableList
func (hl *HittableList) boundingBox() (*aabb, bool) {
	if len(hl.Objects) == 0 {
		return nil, false
	}
	var tmpBox *aabb
	for _, object := range hl.Objects {
		if box, hasBox := object.boundingBox(); hasBox {
			tmpBox = surroundingBox(tmpBox, box)
		} else {
			// No bounding box exists, e. g. because of infinite plane object
			return nil, false
		}
	}
	return tmpBox, true
}

// implement Hittable on HittableList
func (hl *HittableList) Translate(vec Vec) {
	for _, object := range hl.Objects {
		object.Translate(vec)
	}
}

func (hl *HittableList) Clear() {
	hl.Objects = make([]Hittable, 0)
}

func (hl *HittableList) Add(object Hittable) {
	hl.Objects = append(hl.Objects, object)
}

func (hl *HittableList) Length() int {
	return len(hl.Objects)
}

func surroundingBox(box0, box1 *aabb) *aabb {
	if box0 == nil && box1 == nil {
		return nil
	}
	if box0 == nil {
		return box1
	}
	if box1 == nil {
		return box0
	}
	small := Point{
		X: math.Min(box0.min.X, box1.min.X),
		Y: math.Min(box0.min.Y, box1.min.Y),
		Z: math.Min(box0.min.Z, box1.min.Z),
	}
	big := Point{
		X: math.Max(box0.max.X, box1.max.X),
		Y: math.Max(box0.max.Y, box1.max.Y),
		Z: math.Max(box0.max.Z, box1.max.Z),
	}
	return &aabb{
		min: small,
		max: big,
	}
}
