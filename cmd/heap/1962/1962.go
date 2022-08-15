package _1962

import (
	"container/heap"
)

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) First() interface{} {
	return h[0]
}
func (h maxHeap) Less(i int, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Pop() interface{} {
	a := *h
	res := a[len(a)-1]
	*h = a[:len(a)-1]
	return res
}

func (h *maxHeap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

func minStoneSum(piles []int, k int) int {
	sum := 0
	for _, num := range piles {
		sum += num
	}

	h := maxHeap(piles)
	heap.Init(&h)

	for k > 0 {
		k--
		count := h[0] / 2
		if count == 0 {
			break
		}

		h[0] -= count
		sum -= count
		heap.Fix(&h, 0)
	}

	return sum
}
