package engine

import (
	"log"
	"math/rand"
	"sort"
)

type bvhNode struct {
	Left  Hittable
	right Hittable
	box   *aabb
}

type hittableComparator = func(a, b Hittable) bool

// 0 = X, 1 = Y, 2 = Z
var comparators = []hittableComparator{
	boxCompareX, boxCompareY, boxCompareZ,
}

func NewBvhNode(objects []Hittable, start, end int, r *rand.Rand) *bvhNode {
	node := &bvhNode{}

	axis := randInt(0, 2, r)
	comparator := comparators[axis]
	objectCount := end - start

	if objectCount == 1 {
		node.Left = objects[start]
		node.right = objects[start]
	} else if objectCount == 2 {
		if comparator(objects[start], objects[start + 1]) {
			// First is smaller
			node.Left = objects[start]
			node.right = objects[start + 1]
		} else {
			// Second is smaller
			node.Left = objects[start + 1]
			node.right = objects[start]
		}
	} else {
		sortRange(objects, start, end, comparator)
		mid := start + objectCount / 2
		node.Left = NewBvhNode(objects, start, mid, r)
		node.right = NewBvhNode(objects, mid, end, r)
	}

	boxLeft, hasBoxLeft := node.Left.boundingBox()
	boxRight, hasBoxRight := node.right.boundingBox()

	if !hasBoxLeft || !hasBoxRight {
		log.Fatal("No bounding box in function newBvhNode")
	}
	node.box = surroundingBox(boxLeft, boxRight)

	return node
}

// Implement Hittable
func (n *bvhNode) hit(ray *Ray, tMin float64, tMax float64) (*hitRecord, bool) {
	var hitRecord *hitRecord
	recLeft, hitLeft := n.Left.hit(ray, tMin, tMax)
	if hitLeft {
		hitRecord = recLeft
		tMax = recLeft.t
	}
	recRight, hitRight := n.right.hit(ray, tMin, tMax)
	if hitRight {
		// Right hit is closer
		hitRecord = recRight
	}
	if hitLeft || hitRight {
		return hitRecord, true
	}
	return nil, false
}

// implement Hittable
func (n *bvhNode) boundingBox() (*aabb, bool) {
	return n.box, true
}

// implement Hittable
func (n *bvhNode) Translate(vec Vec) {
	panic("Translate is not implemented for bvhNode")
}

func boxCompareX(a, b Hittable) bool {
	// Return true if a is less than b
	boxA, hasBoxA := a.boundingBox()
	boxB, hasBoxB := b.boundingBox()
	if !hasBoxA || !hasBoxB {
		log.Fatal("Missing box in boxCompare")
	}
	return boxA.min.X < boxB.min.X
}

func boxCompareY(a, b Hittable) bool {
	// Return true if a is less than b
	boxA, hasBoxA := a.boundingBox()
	boxB, hasBoxB := b.boundingBox()
	if !hasBoxA || !hasBoxB {
		log.Fatal("Missing box in boxCompare")
	}
	return boxA.min.Y < boxB.min.Y
}

func boxCompareZ(a, b Hittable) bool {
	// Return true if a is less than b
	boxA, hasBoxA := a.boundingBox()
	boxB, hasBoxB := b.boundingBox()
	if !hasBoxA || !hasBoxB {
		log.Fatal("Missing box in boxCompare")
	}
	return boxA.min.Z < boxB.min.Z
}

func sortRange(objects []Hittable, start, end int, comparator hittableComparator) {
	rangeSlice := objects[start : end]
	sort.Slice(rangeSlice, func(i, j int) bool {
		return comparator(rangeSlice[i], rangeSlice[j])
	})
}
