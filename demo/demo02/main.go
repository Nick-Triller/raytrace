package main

import (
	"log"
	"math/rand"
	"raytrace/engine"
	"raytrace/objloader"
	"time"
)

func main() {
	start := time.Now()
	settings := engine.DefaultRenderSettings()
	settings.FileName = "./out/demo02.png"

	camera := engine.ConstructCamera(
		engine.Point{0.8, 0.4, 1},
		engine.Point{0.2, 0.1, 0},
		engine.Vec{0, 1, 0},
		35,
		settings.AspectRatio,
	)
	rendered := engine.Render(settings, createCubeScene(), camera)
	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("Total render time: %s", elapsed)
	engine.WriteToFile(rendered, settings.FileName)
}

func createCubeScene() engine.Hittable {
	world := &engine.HittableList{}
	materialLeft := &engine.Metal{
		Albedo: engine.Color{X: 0.7, Y: 0.8, Z: 0.8},
	}
	//materialGround := &engine.Lambertian{
	//	Albedo: engine.Color{X: 1, Y: 1, Z: 1},
	//}
	materialSphere := &engine.Lambertian{
		Albedo: engine.Color{X: 0.3, Y: 0.1, Z: 0.8},
	}
	//world.Add(&engine.Plane{
	//	engine.Point{0, -0.1, 0},
	//	materialGround,
	//	engine.Vec{0, 1, 0},
	//})
	world.Add(&engine.Sphere{
		Center:   engine.Point{X: 0.35, Y: 0.1, Z: -0.25},
		Radius:   0.2,
		Material: materialSphere,
	})
	mesh := objloader.ReadFromFile("demo/demo02/models/cube.obj", materialLeft, false)
	for _, obj := range mesh.Objects {
		world.Add(obj)
	}
	world.Add(mesh)

	r := rand.New(rand.NewSource(int64(0)))
	bvhTree := engine.NewBvhNode(world.Objects, 0, len(world.Objects), r)

	return bvhTree
}
