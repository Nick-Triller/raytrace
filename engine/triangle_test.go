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
	compare(t, expected, actual)
}

func TestBoundingBoxZero(t *testing.T) {
	expected := &aabb{}
	triangle := &Triangle{}
	actual, _ := triangle.boundingBox()
	compare(t, expected, actual)
}

func compare(t *testing.T, expected, actual *aabb) {
	if ! langReflect.DeepEqual(expected, actual) {
		t.Errorf("Expected bounding box is not equal to actual bounding box. Expected: %v, actual: %v",
			expected, actual)
	}
}