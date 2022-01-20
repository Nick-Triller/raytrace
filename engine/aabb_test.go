package engine

import (
	"testing"
)

func TestAabbHit(t *testing.T) {
	cases := []struct {
		aabb *aabb
		ray  *Ray
		tMax float64
	}{
		{
			aabb: &aabb{
				min: Point{-1, -1, -1},
				max: Point{1, 1, 1},
			},
			ray: &Ray{
				Direction: Vec{-1, -1, -1},
				Origin:    Point{2, 2, 2},
			},
			tMax: 2,
		},
		// Ray starting on edge
		{
			aabb: &aabb{
				min: Point{-1, -1, -1},
				max: Point{0, 0, 0},
			},
			ray: &Ray{
				Direction: Vec{-1, -1, -1},
				Origin:    Point{0, 0, 0},
			},
			tMax: 1,
		},
		// Ray starting within bounding box
		{
			aabb: &aabb{
				min: Point{-1, -1, -1},
				max: Point{0, 0, 0},
			},
			ray: &Ray{
				Direction: Vec{-1, -1, -1},
				Origin:    Point{-0.5, -0.5, -0.5},
			},
			tMax: 0.1,
		},
	}

	for _, tc := range cases {
		actual := tc.aabb.hit(tc.ray, 0, tc.tMax)
		if true != actual {
			t.Errorf("Expected ray to hit aabb. Test case: %v", tc)
		}
	}
}

func TestAabbMiss(t *testing.T) {
	cases := []struct {
		aabb *aabb
		ray  *Ray
		tMax float64
	}{
		// Too far away
		{
			aabb: &aabb{
				min: Point{-1, -1, -1},
				max: Point{1, 1, 1},
			},
			ray: &Ray{
				Direction: Vec{-1, -1, -1},
				Origin:    Point{2, 2, 2},
			},
			tMax: 0.2,
		},
		// Missing
		{
			aabb: &aabb{
				min: Point{-1, -1, -1},
				max: Point{1, 1, 1},
			},
			ray: &Ray{
				Direction: Vec{-1, -1, -1},
				Origin:    Point{20, 20, 20},
			},
			tMax: 10,
		},
		// Ray behind aabb
		{
			aabb: &aabb{
				min: Point{-1, -1, -1},
				max: Point{1, 1, 1},
			},
			ray: &Ray{
				Direction: Vec{-1, -1, -1},
				Origin:    Point{-2, -2, -2},
			},
			tMax: 10,
		},
	}

	for _, tc := range cases {
		actual := tc.aabb.hit(tc.ray, 0, tc.tMax)
		if false != actual {
			t.Errorf("Expected ray to miss aabb. Test case: %v", tc)
		}
	}
}
