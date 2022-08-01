package _215

import "container/heap"

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	length := len(*h) - 1
	last := (*h)[length]
	*h = (*h)[0:length]
	return last
}

func findKthLargest(nums []int, k int) int {
	intHeap := IntHeap(nums)
	heap.Init(&intHeap)
	if k == 1 {
		return intHeap[0]
	}
	for i := 1; i < k; i++ {
		heap.Remove(&intHeap, 0)
	}
	return intHeap[0]
}
