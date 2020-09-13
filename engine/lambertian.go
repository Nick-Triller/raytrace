package engine

import "math/rand"

var diffuseStrategy DiffuseRenderStrategy

type Lambertian struct {
	// Measure of how much light that hits a surface is reflected without being absorbed
	Albedo Color
}

func (l *Lambertian) scatter(rayIn *Ray, hitRecord *hitRecord, r *rand.Rand) (bool, *Ray, Color) {
	var scatterDirection Vec
	switch diffuseStrategy {
	case RandomInUnitSphereStrategy:
		scatterDirection = hitRecord.normal.Add(RandomInUnitSphere(r))
	case RandomUnitVectorStrategy:
		scatterDirection = hitRecord.normal.Add(RandomUnitVector(r))
	case HemisphericalScatteringStrategy:
		scatterDirection = hitRecord.p.Add(RandomInHemisphere(hitRecord.normal, r))
	}
	scatteredRay := &Ray{hitRecord.p, scatterDirection}
	return true, scatteredRay, l.Albedo
}
