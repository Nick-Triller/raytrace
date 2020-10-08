package engine

import (
	"image"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"
)

func rayColor(ray *Ray, world Hittable, depth int, r *rand.Rand) Color {
	if depth <= 0 {
		// If we've exceeded the ray bounce limit, no more light is gathered.
		return Color{0, 0, 0}
	}

	record, hit := world.hit(ray, 0.001, math.Inf(1))
	if hit {
		// Color based on normal vector
		scattered, scatteredRay, attenuation := record.material.scatter(ray, record, r)
		if scattered {
			return attenuation.Multiply(rayColor(scatteredRay, world, depth - 1, r))
		}
		// ALl light was absorbed
		return Color{0, 0, 0}
	}

	// No hit, use background gradient
	unitDirection := ray.Direction.Normalize()
	t := 0.5 * (unitDirection.Y + 1.)
	endVal := Color{1, 1, 1}
	startVal := Color{0.5, 0.7, 1.0}
	a := endVal.MultiplyScalar(1 - t)
	b := startVal.MultiplyScalar(t)
	return a.Add(b)
}

func Render(settings RenderSettings, world *HittableList, camera *Camera) *image.RGBA {
	diffuseStrategy = settings.DiffuseStrategy
	images := make([]*image.RGBA, 0, settings.Parallelism)
	resultChan := make(chan *image.RGBA)
	for a := 0; a < settings.Parallelism; a++ {
		log.Printf("Starting worker %d\n", a)
		go renderWorker(settings, camera, world, a, resultChan)
	}
	for a := 0; a < settings.Parallelism; a++ {
		images = append(images, <- resultChan)
		log.Printf("Collected %d images\n", a + 1)
	}
	log.Printf("Combinding images...\n")
	return combineImages(images...)
}

func renderWorker(settings RenderSettings, camera *Camera, world *HittableList, workerId int, resultChan chan *image.RGBA) {
	start := time.Now()

	r := rand.New(rand.NewSource(int64(workerId)))
	samplesPerWorker := float64(settings.SamplesPerPixel) / float64(settings.Parallelism)
	if samplesPerWorker < 1 {
		samplesPerWorker = 1
	}
	imageHeight := int(float64(settings.ImageWidth) / settings.AspectRatio)

	// Create image object
	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: settings.ImageWidth, Y: imageHeight}
	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight })

	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < settings.ImageWidth; i++ {
			var pixelColor Color
			for s := 0; s < int(samplesPerWorker); s++ {
				u := (float64(i) + r.Float64()) / float64(settings.ImageWidth- 1)
				v := (float64(j) + r.Float64()) / float64(imageHeight - 1)
				ray := camera.getRay(u, v)
				pixelColor = pixelColor.Add(rayColor(ray, world, settings.MaxDepth, r))
			}
			pixelColor = pixelColor.DivideScalar(samplesPerWorker)
			// https://www.cambridgeincolour.com/tutorials/gamma-correction.htm
			// Sqrt is approximate gamma correction
			pixelColor.X = math.Sqrt(pixelColor.X)
			pixelColor.Y = math.Sqrt(pixelColor.Y)
			pixelColor.Z = math.Sqrt(pixelColor.Z)
			invertY := imageHeight - j - 1
			img.SetRGBA(i, invertY, color.RGBA{
				R: uint8(256 * clamp(pixelColor.X,0, 0.9999999)),
				G: uint8(256 * clamp(pixelColor.Y,0, 0.9999999)),
				B: uint8(256 * clamp(pixelColor.Z,0, 0.9999999)),
				A: 255,
			})
		}
	}

	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("Render time worker %d: %s", workerId, elapsed)

	resultChan <- img
}

func combineImages(inputs ...*image.RGBA) *image.RGBA {
	width := inputs[0].Rect.Max.X
	height := inputs[0].Rect.Max.Y
	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: width, Y: height}
	result := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight })

	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			var pixelColor = Color{0, 0, 0}
			for img := 0; img < len(inputs); img++ {
				pixel := inputs[img].RGBAAt(i, j)
				pixelColor = pixelColor.AddScalars(float64(pixel.R), float64(pixel.G), float64(pixel.B))
			}
			pixelColor = pixelColor.DivideScalar(float64(len(inputs)))
			result.SetRGBA(i, j, color.RGBA{
				R: uint8(pixelColor.X),
				G: uint8(pixelColor.Y),
				B: uint8(pixelColor.Z),
				A: 255,
			})
		}
	}

	return result
}
