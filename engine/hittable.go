package engine

type hitRecord struct {
	p         Point
	normal    Vec
	t         float64
	frontFace bool
	material  Material
}

func (r *hitRecord) setFaceNormal(ray *Ray, outwardNormal Vec) {
	r.frontFace = Dot(ray.Direction, outwardNormal) < 0
	if r.frontFace {
		r.normal = outwardNormal
	} else {
		r.normal = outwardNormal.MultiplyScalar(-1)
	}
}

type Hittable interface {
	hit(ray *Ray, tMin float64, tMax float64) (*hitRecord, bool)
	Translate(vec Vec)
}

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
func (hl *HittableList) Translate(vec Vec) {
	for _, object := range hl.Objects {
		object.Translate(vec)
	}
}

func (hl *HittableList) clear() {
	hl.Objects = hl.Objects[:0]
}

func (hl *HittableList) Add(object Hittable) {
	hl.Objects = append(hl.Objects, object)
}
