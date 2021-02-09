package engine

import (
	"image"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
)

func clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func randFloat64(min, max float64, r *rand.Rand) float64 {
	return r.Float64() * (max - min) + min
}

func randInt(min, max int, r *rand.Rand) int {
	return int(randFloat64(float64(min), float64(max), r))
}

func WriteToFile(img *image.RGBA, fileName string) {
	// Encode as PNG.
	f, _ := os.Create(fileName)
	err := png.Encode(f, img)
	if err != nil {
		log.Fatalf("Encoding error: %v", err)
	}
	_ = f.Close()
}

func degreesToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}
