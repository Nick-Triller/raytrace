package engine

import (
	"math"
	"math/rand"
)

type Point = Vec
type Color = Vec

type Vec struct {
	X, Y, Z float64
}

func (v Vec) Add(v2 Vec) Vec {
	v.X += v2.X
	v.Y += v2.Y
	v.Z += v2.Z
	return v
}

func (v Vec) AddScalar(val float64) Vec {
	v.X += val
	v.Y += val
	v.Z += val
	return v
}

func (v Vec) AddScalars(x, y, z float64) Vec {
	v.X += x
	v.Y += y
	v.Z += z
	return v
}

func (v Vec) Subtract(v2 Vec) Vec {
	v.X -= v2.X
	v.Y -= v2.Y
	v.Z -= v2.Z
	return v
}

func (v Vec) SubtractScalar(val float64) Vec {
	v.X -= val
	v.Y -= val
	v.Z -= val
	return v
}

func (v Vec) SubtractScalars(x, y, z float64) Vec {
	v.X -= x
	v.Y -= y
	v.Z -= z
	return v
}

func (v Vec) Multiply(v2 Vec) Vec {
	v.X *= v2.X
	v.Y *= v2.Y
	v.Z *= v2.Z
	return v
}

func (v Vec) MultiplyScalar(val float64) Vec {
	v.X *= val
	v.Y *= val
	v.Z *= val
	return v
}

func (v Vec) MultiplyScalars(x, y, z float64) Vec {
	v.X *= x
	v.Y *= y
	v.Z *= z
	return v
}

func (v Vec) Divide(v2 Vec) Vec {
	v.X /= v2.X
	v.Y /= v2.Y
	v.Z /= v2.Z
	return v
}

func (v Vec) DivideScalar(val float64) Vec {
	v.X /= val
	v.Y /= val
	v.Z /= val
	return v
}

func (v Vec) DivideScalars(x, y, z float64) Vec {
	v.X /= x
	v.Y /= y
	v.Z /= z
	return v
}

func (v Vec) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec) LengthSquared() float64 {
	return v.X * v.X + v.Y * v.Y + v.Z * v.Z
}

func Dot(v1 Vec, v2 Vec) float64 {
	return v1.X * v2.X + v1.Y * v2.Y + v1.Z * v2.Z
}

func Random() Vec {
	return Vec{
		rand.Float64(),
		rand.Float64(),
		rand.Float64(),
	}
}

func RandomInRange(min, max float64, r *rand.Rand) Vec {
	return Vec{
		randFloat64(min, max, r),
		randFloat64(min, max, r),
		randFloat64(min, max, r),
	}
}

func RandomInUnitSphere(r *rand.Rand) Vec {
	for {
		p := RandomInRange(-1, 1, r)
		if p.LengthSquared() >= 1 {
			// Random point is not in unit Sphere, try again
			continue
		}
		return p
	}
}

func RandomUnitVector(ra *rand.Rand) Vec {
	// Used to calculate true Lambertian Reflection
	// a is random float between 0 and 2 Pi
	a := randFloat64(0, 2 * math.Pi, ra)
	// z is random float between -1 and 1
	z := randFloat64(-1, 1, ra)
	r := math.Sqrt(1 - z * z)
	return Vec{r * math.Cos(a), r * math.Sin(a), z}
}

func RandomInHemisphere(normal Vec, r *rand.Rand) Vec {
	inUnitSphere := RandomInUnitSphere(r)
	if Dot(inUnitSphere, normal) > 0.0 {
		// In the same hemisphere as the normal
		return inUnitSphere
	}
	return inUnitSphere.MultiplyScalar(-1)
}

func Cross(u Vec, v Vec) Vec {
	return Vec{
		u.Y * v.Z - u.Z * v.Y,
		u.Z * v.X - u.X * v.Z,
		u.X * v.Y - u.Y * v.X,
	}
}

func (v Vec) Normalize() Vec {
	return v.DivideScalar(v.Length())
}
