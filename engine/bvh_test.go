package engine

import (
	"math"
	"math/rand"
	"testing"
)

func TestBvhSingleObjectHit(t *testing.T) {
	r := rand.New(rand.NewSource(int64(0)))
	bvh := NewBvhNode([]Hittable{
		NewTriangle(Point{0, 0, 0}, Point{1, 0, 0}, Point{0, 1, 0}, nil),
	}, 0, 1, r)
	hitRecord, hasHit := bvh.hit(&Ray{
		Origin:    Point{0.6, 0.1, 0.5},
		Direction: Vec{0, 0, -1},
	}, 0, 10)
	if hasHit == false {
		t.Fatalf("Expected ray to hit object")
	}
	expectedHitPoint := Point{0.6, 0.1, 0}
	if ! pointsEqual(hitRecord.p, expectedHitPoint, 0.001) {
		t.Errorf("Unexpected hit point. Expected: %v, actual: %v", expectedHitPoint, hitRecord.p)
	}
}

func TestBvhSingleObjectMiss(t *testing.T) {
	r := rand.New(rand.NewSource(int64(0)))
	bvh := NewBvhNode([]Hittable{
		NewTriangle(Point{0, 0, 0}, Point{1, 0, 0}, Point{0, 1, 0}, nil),
	}, 0, 1, r)
	_, hasHit := bvh.hit(&Ray{
		Origin:    Point{-0.6, 0.1, 0.5},
		Direction: Vec{0, 0, -1},
	}, 0, 10)
	if hasHit == true {
		t.Fatalf("Expected ray to miss object")
	}
}

func pointsEqual(p1, p2 Point, epsilon float64) bool {
	if math.Abs(p1.X - p2.X) >= epsilon {
		return false
	}
	if math.Abs(p1.Y - p2.Y) >= epsilon {
		return false
	}
	if math.Abs(p1.Z - p2.Z) >= epsilon {
		return false
	}
	return true
}
