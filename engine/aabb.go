package engine

import "math"

// aabb = axis-aligned bounding boxes
type aabb struct {
	min Point
	max Point
}

func (a *aabb) hit(ray *Ray, tMin, tMax float64) bool {
	// X dimension
	t0x := math.Min((a.min.X - ray.Origin.X) / ray.Direction.X, (a.max.X - ray.Origin.X) / ray.Direction.X)
	t1x := math.Max((a.min.X - ray.Origin.X) / ray.Direction.X, (a.max.X - ray.Origin.X) / ray.Direction.X)

	tMin = math.Max(t0x, tMin)
	tMax = math.Min(t1x, tMax)

	if tMax <= tMin {
		return false
	}
	// Y dimension
	t0y:= math.Min((a.min.Y - ray.Origin.Y) / ray.Direction.Y, (a.max.Y - ray.Origin.Y) / ray.Direction.Y)
	t1y := math.Max((a.min.Y - ray.Origin.Y) / ray.Direction.Y, (a.max.Y - ray.Origin.Y) / ray.Direction.Y)

	tMin = math.Max(t0y, tMin)
	tMax = math.Min(t1y, tMax)

	if tMax <= tMin {
		return false
	}
	// Z dimension
	t0z := math.Min((a.min.Z - ray.Origin.Z) / ray.Direction.Z, (a.max.Z - ray.Origin.Z) / ray.Direction.Z)
	t1z := math.Max((a.min.Z - ray.Origin.Z) / ray.Direction.Z, (a.max.Z - ray.Origin.Z) / ray.Direction.Z)

	tMin = math.Max(t0z, tMin)
	tMax = math.Min(t1z, tMax)

	if tMax <= tMin {
		return false
	}

	return true
}
