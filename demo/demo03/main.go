package main

import (
	"log"
	"raytrace/engine"
	"time"
)

func main() {
	start := time.Now()
	settings := engine.DefaultRenderSettings()
	settings.FileName = "./out/demo03.png"

	camera := engine.ConstructCamera(
		engine.Point{0, 0.75, 3},
		engine.Point{0, 0.65, 0},
		engine.Vec{0, 1, 0},
		35,
		settings.AspectRatio,
	)
	rendered := engine.Render(settings, createPlaneScene(), camera)
	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("Total render time: %s", elapsed)
	engine.WriteToFile(rendered, settings.FileName)
}

func createPlaneScene() *engine.HittableList {
	// Scene
	world := &engine.HittableList{}
	groundMaterial := &engine.Lambertian{
		Albedo: engine.Color{0.5, 0.5, 0.5},
	}
	metalMaterial := &engine.Metal{
		Albedo: engine.Color{0.2, 0.5, 0.7},
		Fuzz:   0.,
	}
	glassMaterial := &engine.Dielectric{
		RefIdx: 1.5,
	}
	world.Add(&engine.Plane{
		engine.Point{0, 0, 0},
		groundMaterial,
		engine.Vec{0, 1, 0},
	})
	world.Add(&engine.Sphere{
		Center:   engine.Point{-0.6, 0.5, 0},
		Radius:   0.5,
		Material: metalMaterial,
	})
	world.Add(&engine.Sphere{
		Center:   engine.Point{0.6, 0.5, 0},
		Radius:   0.5,
		Material: glassMaterial,
	})
	return world
}
