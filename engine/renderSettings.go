package engine

import "runtime"

type DiffuseRenderStrategy string

const (
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
	CpuProfiling	bool
}

func DefaultRenderSettings() RenderSettings {
	parallelism := runtime.NumCPU()
	return RenderSettings{
		RandomUnitVectorStrategy,
		16. / 9.,
		900,
		100,
		50,
		"out/image.png",
		parallelism,
		false,
	}
}
