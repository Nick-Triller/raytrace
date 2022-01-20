package engine

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
