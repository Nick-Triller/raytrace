package engine

import "math/rand"

type Metal struct {
	// Measure of how much light that hits a surface is reflected without being absorbed
	Albedo Color
	Fuzz   float64
}

func (m *Metal) scatter(rayIn *Ray, hitRecord *hitRecord, r *rand.Rand) (bool, *Ray, Color) {
	reflected := reflect(rayIn.Direction.Normalize(), hitRecord.normal)
	direction := reflected.Add(RandomInUnitSphere(r).MultiplyScalar(m.Fuzz))
	scattered := &Ray{hitRecord.p, direction}
	return true, scattered, m.Albedo
}

func reflect(v Vec, normal Vec) Vec {
	a := normal.MultiplyScalar(2).MultiplyScalar(Dot(v, normal))
	return v.Subtract(a)
}
