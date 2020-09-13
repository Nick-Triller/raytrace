package main

import (
	"log"
	"raytrace/engine"
	"time"
)

func main() {
	start := time.Now()
	settings := engine.DefaultRenderSettings()
	camera := engine.ConstructCamera(settings.CameraPos, settings.VFoVDegrees, settings.AspectRatio)
	rendered := engine.Render(settings, createScene(), camera)
	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("Total render time: %s", elapsed)
	engine.WriteToFile(rendered, settings.FileName)
}

func createScene() engine.HittableList {
	// Scene
	world := engine.HittableList{
		Objects: make([]engine.Hittable, 0),
	}
	materialGround := &engine.Lambertian{
		Albedo: engine.Color{X: 0.8, Y: 0.8, Z: 0.8},
	}
	materialFront := &engine.Lambertian{
		Albedo: engine.Color{X: 0.7, Y: 0.3, Z: 0.3},
	}
	materialCenter := &engine.Metal{
		Albedo: engine.Color{X: 0.2, Y: 0.3, Z: 0.7},
		Fuzz:   0.3,
	}
	materialLeft := &engine.Metal{
		Albedo: engine.Color{X: 0.8, Y: 0.8, Z: 0.8},
	}
	materialRight := &engine.Metal{
		Albedo: engine.Color{X: 0.8, Y: 0.6, Z: 0.2},
		Fuzz:   0.5,
	}

	world.Add(&engine.Sphere{
		Center:   engine.Point{Y: -100.5, Z: -1},
		Radius:   100,
		Material: materialGround,
	})
	world.Add(&engine.Sphere{
		Center:   engine.Point{X: -0.4, Y: -0.4, Z: -0.7},
		Radius:   0.1,
		Material: materialFront,
	})
	world.Add(&engine.Sphere{
		Center:   engine.Point{Z: -1},
		Radius:   0.5,
		Material: materialCenter,
	})
	world.Add(&engine.Sphere{
		Center:   engine.Point{X: -1.07, Z: -1},
		Radius:   0.5,
		Material: materialLeft,
	})
	world.Add(&engine.Sphere{
		Center:   engine.Point{X: 1.07, Z: -1},
		Radius:   0.5,
		Material: materialRight,
	})
	return world
}
