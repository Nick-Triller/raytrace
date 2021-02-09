package engine

import (
	langReflect "reflect"
	"testing"
)

func TestBvh(t *testing.T) {
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