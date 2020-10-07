package main

import (
	"log"
	_ "net/http/pprof"
	"os"
	"raytrace/engine"
	"runtime/pprof"
	"time"
)

func main() {
	start := time.Now()

	settings := engine.DefaultRenderSettings()
	settings.FileName = "./out/demo01.png"

	if settings.CpuProfiling {
		f, err := os.Create("cpu.prof")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	camera := engine.ConstructCamera(
		engine.Point{1.4, 0.7, 1.7},
		engine.Point{0.2, 0.55, -1},
		engine.Vec{0, 1, 0},
		35,
		settings.AspectRatio,
	)
	rendered := engine.Render(settings, createSpheresScene(), camera)
	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("Total render time: %s", elapsed)
	engine.WriteToFile(rendered, settings.FileName)
}

func createSpheresScene() engine.HittableList {
	world := engine.HittableList{
		Objects: make([]engine.Hittable, 0),
	}
	materialGround := &engine.Lambertian{
		Albedo: engine.Color{X: 0.8, Y: 0.8, Z: 0.8},
	}
	materialFront := &engine.Lambertian{
		Albedo: engine.Color{X: 0.7, Y: 0.3, Z: 0.3},
	}
	materialGlass := &engine.Dielectric{
		RefIdx: 1.5,
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

	world.Add(&engine.Plane{
		engine.Point{0, 0, 0},
		materialGround,
		engine.Vec{0, 1, 0},
	})
	world.Add(&engine.Sphere{
		Center:   engine.Point{X: -0.4, Y: 0.1, Z: -0.7},
		Radius:   0.1,
		Material: materialFront,
	})
	world.Add(&engine.Sphere{
		Center:   engine.Point{X: 0.4, Y: 0.1, Z: -0.7},
		Radius:   0.1,
		Material: materialGlass,
	})
	world.Add(&engine.Sphere{
		Center:   engine.Point{Z: -1, Y: 0.5},
		Radius:   0.5,
		Material: materialCenter,
	})
	world.Add(&engine.Sphere{
		Center:   engine.Point{X: -1.07, Y: 0.5, Z: -1},
		Radius:   0.5,
		Material: materialLeft,
	})
	world.Add(&engine.Sphere{
		Center:   engine.Point{X: 1.07, Y: 0.5, Z: -1},
		Radius:   0.5,
		Material: materialRight,
	})
	return world
}
