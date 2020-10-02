package engine

import (
	"math"
	"math/rand"
)

type Dielectric struct {
	// Refractive index
	RefIdx float64
}

func (d *Dielectric) scatter(rayIn *Ray, hitRecord *hitRecord, r *rand.Rand) (bool, *Ray, Color) {
	// attenuation is always 1 because the material doesn't absorb light
	attenuation := Color{1,1,1}
	etaiOverEtat := d.RefIdx
	if hitRecord.frontFace {
		etaiOverEtat = 1.0 / d.RefIdx
	}
	unitDirection := rayIn.Direction.Normalize()

	cosTheta := math.Min(Dot(unitDirection.MultiplyScalar(-1), hitRecord.normal), 1.)
	sinTheta := math.Sqrt(1. - cosTheta * cosTheta)
	if etaiOverEtat * sinTheta > 1. {
		// Reflect
		reflected := reflect(unitDirection, hitRecord.normal)
		return true, &Ray{
			Origin:    hitRecord.p,
			Direction: reflected,
		}, attenuation
	}
	reflectProbability := schlickApproximation(cosTheta, etaiOverEtat)
	if r.Float64() < reflectProbability {
		reflected := reflect(unitDirection, hitRecord.normal)
		return true, &Ray{
			Origin:    hitRecord.p,
			Direction: reflected,
		}, attenuation
	}
	// Refract
	refracted := refract(unitDirection, hitRecord.normal, etaiOverEtat)
	return true, &Ray{
		Origin:    hitRecord.p,
		Direction: refracted,
	}, attenuation
}

func refract(uv Vec, n Vec, etaiOverEtat float64) Vec {
	cosTheta := Dot(uv.MultiplyScalar(-1), n)
	rOutPerp := uv.Add(n.MultiplyScalar(cosTheta)).MultiplyScalar(etaiOverEtat)
	rOutParallel := n.MultiplyScalar(-math.Sqrt(math.Abs(1. - rOutPerp.LengthSquared())))
	return rOutPerp.Add(rOutParallel)
}

func schlickApproximation(cosine, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1 - r0) * math.Pow(1 - cosine, 5)
}
