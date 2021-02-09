package engine

import (
	langReflect "reflect"
	"testing"
)

func TestBoundingBox(t *testing.T) {
	expected := &aabb{
		min: Point{0, -10, 0,},
		max: Point{10, 0, 15,},
	}
	triangle := &Triangle{
		V1:       Point{10, 0,0},
		V2:       Point{0, -10, 0},
		V3:       Point{0, 0, 15},
		Material: nil,
		Normal:   Vec{},
	}
	actual, _ := triangle.boundingBox()
	if ! langReflect.DeepEqual(expected, actual) {
		t.Errorf("Expected bounding box not equal to actual bounding box. Expected: %v, actual: %v",
			expected, actual)
	}
}

func TestBoundingBoxZero(t *testing.T) {
	expected := &aabb{min: Point{}, max: Point{}}
	triangle := &Triangle{
		V1:       Point{0, 0,0},
		V2:       Point{0, 0, 0},
		V3:       Point{0, 0, 0},
		Material: nil,
		Normal:   Vec{},
	}
	actual, _ := triangle.boundingBox()
	if ! langReflect.DeepEqual(expected, actual) {
		t.Errorf("Expected bounding box not equal to actual bounding box. Expected: %v, actual: %v",
			expected, actual)
	}
}
