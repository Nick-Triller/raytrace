# Ray Tracer

This is a CPU based ray tracer implemented in Go with no dependencies.
It was created for educational purposes as part of 
[Hacktoberfest 2020](https://hacktoberfest.digitalocean.com/).
The code was heavily inspired by the C++ code from the books 
["Ray Tracing in One Weekend"](https://raytracing.github.io/books/RayTracingInOneWeekend.html) and
["Ray Tracing: The Next Week"](https://raytracing.github.io/books/RayTracingTheNextWeek.html).

## Configuration

Checkout [renderSettings.go](./engine/renderSettings.go).

## Usage

Rendering can take a long time depending on number of CPU cores, image resolution, samples per pixel and scene 
complexity.

```
go run demo/demo01/main.go
```

## Demos

### demo01

![](./img/demo01.png)

### demo02

![](./img/demo02.png)

### demo03

![](./img/demo03.png)

### demo04

![](./img/demo04.png)

### demo05

![](./img/demo05.png)
