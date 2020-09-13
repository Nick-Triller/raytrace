package engine

import "math/rand"

type Material interface {
	scatter(rayIn *Ray, hitRecord *hitRecord, r *rand.Rand) (bool, *Ray, Color)
}
