package objloader

import (
	"bufio"
	"log"
	"os"
	"raytrace/engine"
	"strconv"
	"strings"
)

func ReadFromFile(filename string, material engine.Material, ignoreErrors bool) *engine.HittableList {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	mesh := &engine.HittableList{}
	vertices := make([]engine.Point, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line[0] == '#' {
			// Skip empty lines and comment lines
			continue
		}
		fields := strings.Split(line, " ")
		switch fields[0] {
		case "v":
			x := parseFloat64(fields[1])
			y := parseFloat64(fields[2])
			z := parseFloat64(fields[3])
			vertices = append(vertices, engine.Point{x, y, z})
		case "f":
			v1id := parseInt(strings.Split(fields[1], "/")[0])
			v2id := parseInt(strings.Split(fields[2], "/")[0])
			v3id := parseInt(strings.Split(fields[3], "/")[0])
			triangle := engine.NewTriangle(vertices[v1id-1], vertices[v2id-1], vertices[v3id-1], material)
			mesh.Add(triangle)
		default:
			if !ignoreErrors {
				log.Fatalf("Failed to parse .obj file, encountered unknown token \"%s\"\n", fields[0])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return mesh
}

func parseFloat64(s string) float64 {
	if s, err := strconv.ParseFloat(s, 64); err == nil {
		return s
	} else {
		log.Fatalf("Failed to parse .obj file, failed to parse token \"%f\" as float64\n", s)
	}
	return 0
}

func parseInt(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	} else {
		log.Fatalf("Failed to parse .obj file, failed to parse token \"%s\" as float64\n", s)
	}
	return 0
}
