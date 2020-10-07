# Ray Tracer

This is a CPU based ray tracer implemented in Go with no dependencies.
It was created for educational purposes as part of 
[Hacktoberfest 2020](https://hacktoberfest.digitalocean.com/).
It's code is heavily inspired by the C++ code from the book 
["Ray Tracing in One Weekend"](https://raytracing.github.io/books/RayTracingInOneWeekend.html).

## Usage

Rendering can take a long time depending on number of CPU cores, image resolution 
and samples per pixel.

```
go run demo/demo01.go
```

![](./img/demo01.png)

![](./img/demo02.png)

![](./img/demo03.png)

![](./img/demo04.png)
