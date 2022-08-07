package _658

import (
	"container/heap"
	"sort"
)

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

var X int

func (h Heap) Less(i, j int) bool {
	a := h[i]
	b := h[j]
	left := abs(a - X)
	right := abs(b - X)
	if left >= right {
		return true
	}
	return !(left == right && a < b)
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	last := h.Len() - 1
	value := old[last]
	(*h) = old[:last]
	return value
}

func findClosestElements(arr []int, k int, x int) []int {
	X = x
	pq := Heap{}
	heap.Init(&pq)

	for _, item := range arr {
		heap.Push(&pq, item)
	}

	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(&pq).(int))
	}
	sort.Ints(result)
	return result
}

func version2(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k
	for left < right {
		mid := (left + right) / 2
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}

	}
	return arr[left : left+k]
}
