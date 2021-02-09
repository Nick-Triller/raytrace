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
