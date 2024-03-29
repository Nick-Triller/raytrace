package main

import (
	"log"
	"math/rand"
	"os"
	"raytrace/engine"
	"raytrace/objloader"
	"runtime/pprof"
	"time"
)

func main() {
	// Render time width=1000, spp=3000, maxDepth=12, cores=4, without bvh was 6 min 51 seconds
	// With same settings and bvh, it was ?

	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	go func() {
		time.Sleep(20 * time.Second)
		pprof.StopCPUProfile()
	}()

	start := time.Now()
	settings := engine.DefaultRenderSettings()
	settings.FileName = "./out/demo06.png"

	camera := engine.ConstructCamera(
		engine.Point{20, 35, 220},
		engine.Point{15, -5, 0},
		engine.Vec{0, 1, 0},
		35,
		settings.AspectRatio,
	)
	rendered := engine.Render(settings, createTriangleSphereScene(), camera)
	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("Total render time: %s", elapsed)
	engine.WriteToFile(rendered, settings.FileName)
}

func createTriangleSphereScene() engine.Hittable {
	material := &engine.Metal{
		Albedo: engine.Color{X: 0.65, Y: 0.6, Z: 0.7},
	}
	mesh := objloader.ReadFromFile("demo/demo06/models/dragon.obj", material, false)
	log.Printf("Building BVH\n")
	r := rand.New(rand.NewSource(int64(0)))
	bvhTree := engine.NewBvhNode(mesh.Objects, 0, len(mesh.Objects), r)
	return bvhTree
}
