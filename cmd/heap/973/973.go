package _973

import (
	"container/heap"
	"math"
)

type Point struct {
	x int
	y int
}

//√(x1 - x2)2 + (y1 - y2)2)
func distance(left *Point, right *Point) float64 {
	x := left.x - right.x
	y := left.y - right.y

	return math.Sqrt(float64(x*x + y*y))
}

type Heap []Point

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(i, j int) bool {
	point0 := Point{
		x: 0,
		y: 0,
	}
	left := &h[i]
	right := &h[j]
	return distance(&point0, left) < distance(&point0, right)
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(point interface{}) {
	*h = append(*h, point.(Point))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := Heap{}
	heap.Init(&h)
	for _, point := range points {
		heap.Push(&h, Point{point[0], point[1]})
	}
	result := make([][]int, k)
	for i := 0; i < k; i++ {
		point := heap.Pop(&h).(Point)
		result[i] = []int{point.x, point.y}
	}

	return result
}
