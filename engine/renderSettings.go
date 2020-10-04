package engine

import "runtime"

type DiffuseRenderStrategy string

const(
	RandomInUnitSphereStrategy      DiffuseRenderStrategy = "RandomInUnitSphereStrategy"
	RandomUnitVectorStrategy                              = "RandomUnitVectorStrategy"
	HemisphericalScatteringStrategy                       = "HemisphericalScatteringStrategy"
)

type RenderSettings struct {
	DiffuseStrategy DiffuseRenderStrategy
	// Image height is derived from width and aspect ratio
	AspectRatio     float64
	ImageWidth      int
	SamplesPerPixel int
	MaxDepth        int
	FileName        string
	Parallelism		int
}

func DefaultRenderSettings() RenderSettings {
	// Keep one core idle so the machine keeps working smoothly while rendering an image
	parallelism := runtime.NumCPU() - 0
	if parallelism <= 0 {
		parallelism = 1
	}

	return RenderSettings{
		RandomUnitVectorStrategy,
		16. / 9.,
		800,
		500,
		50,
		"out/image.png",
		parallelism,
	}
}
