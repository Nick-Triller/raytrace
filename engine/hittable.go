package engine

type Hittable interface {
	hit(ray *Ray, tMin float64, tMax float64) (*hitRecord, bool)
	// boundingBox return bool because not all geometries have a bounding box (e. g. infinite plane)
	boundingBox() (*aabb, bool)
	Translate(vec Vec)
}
