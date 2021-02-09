package engine

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
	parallelism := 1 // runtime.NumCPU()
	return RenderSettings{
		RandomUnitVectorStrategy,
		16. / 10.,
		600,
		3,
		6,
		"out/image.png",
		parallelism,
		false,
	}
}
