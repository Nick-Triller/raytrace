package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"raytrace/engine"
	"runtime/pprof"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	winWidth := 600
	var winHeight int

	settings := engine.DefaultRenderSettings()
	settings.CpuProfiling = true
	settings.ImageWidth = winWidth
	settings.SamplesPerPixel = 1

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

	winHeight = int(float64(winWidth) / settings.AspectRatio)

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Raytrace", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tex.Destroy()

	quitChan := make(chan struct{})
	doneChan := make(chan struct{})

	go func() {
		bvh := createSpheresScene()
		xx := 0.

		for {
			select {
			case <-quitChan:
				close(doneChan)
				return
			case <-time.After(0 * time.Millisecond):
				xx += 0.1
				vv := math.Sin(xx)
				camera := engine.ConstructCamera(
					engine.Point{1.4 + vv, 0.7, 1.7},
					engine.Point{0.2, 0.55, -1},
					engine.Vec{0, 1, 0},
					35,
					settings.AspectRatio,
				)

				image := engine.Render(settings, bvh, camera)

				err := tex.Update(&sdl.Rect{
					X: 0,
					Y: 0,
					W: int32(settings.ImageWidth),
					H: int32(float64(settings.ImageWidth) / settings.AspectRatio),
				}, image.Pix, settings.ImageWidth*4)
				if err != nil {
					panic(err)
				}
				err = renderer.Copy(tex, nil, nil)
				if err != nil {
					panic(err)
				}
				renderer.Present()
			}
		}
	}()

	waitQuit := func() {
		close(quitChan)
		<-doneChan
	}

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.KeyboardEvent:
				kevent := event.(*sdl.KeyboardEvent)
				if kevent.Keysym.Sym == sdl.K_ESCAPE {
					waitQuit()
					return
				}
			case *sdl.QuitEvent:
				waitQuit()
				return
			}
		}
	}
}

func createSpheresScene() engine.Hittable {
	world := &engine.HittableList{}

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
	r := rand.New(rand.NewSource(int64(0)))
	bvh := engine.NewBvhNode(world.Objects, 0, len(world.Objects), r)
	return bvh
}
