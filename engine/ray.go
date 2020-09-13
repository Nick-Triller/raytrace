package engine

type Ray struct {
	Origin    Point
	Direction Vec
}

func (r Ray) Interpolate(time float64) Point {
	return r.Origin.Add(r.Direction.MultiplyScalar(time))
}
