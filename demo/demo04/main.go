package main

import (
	"log"
	"raytrace/engine"
	"time"
)

func main() {
	start := time.Now()
	settings := engine.DefaultRenderSettings()
	settings.FileName = "./out/demo04.png"

	camera := engine.ConstructCamera(
		engine.Point{0, 0.75, 3},
		engine.Point{0, 0.65, 0},
		engine.Vec{0, 1, 0},
		35,
		settings.AspectRatio,
	)
	rendered := engine.Render(settings, createTriangleScene(), camera)
	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("Total render time: %s", elapsed)
	engine.WriteToFile(rendered, settings.FileName)
}

func createTriangleScene() engine.HittableList {
	// Scene
	world := engine.HittableList{
		Objects: make([]engine.Hittable, 0),
	}
	groundMaterial := &engine.Lambertian{
		Albedo: engine.Color{0.5, 0.5, 0.5},
	}
	metalMaterial := &engine.Metal{
		Albedo: engine.Color{0.2, 0.5, 0.7},
	}
	world.Add(&engine.Plane{
		engine.Point{0, 0, 0},
		groundMaterial,
		engine.Vec{0, 1, 0},
	})
	world.Add(engine.NewTriangle(
		engine.Point{0, 0.1, 0},
		engine.Point{2, 0.1, -10},
		engine.Point{-2, 0.1, -10},
		metalMaterial),
	)
	return world
}
