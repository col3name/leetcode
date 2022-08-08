package _1046

import "container/heap"

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) First() interface{} {
	return h[0]
}
func (h MaxHeap) Less(i int, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Pop() interface{} {
	a := *h
	res := a[len(a)-1]
	*h = a[:len(a)-1]
	return res
}

func (h *MaxHeap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

func lastStoneWeight(stones []int) int {
	pq := &MaxHeap{}
	heap.Init(pq)
	for _, stone := range stones {
		heap.Push(pq, stone)
	}
	var (
		left  int
		right int
	)

	for pq.Len() > 1 {
		left = heap.Pop(pq).(int)
		right = heap.Pop(pq).(int)
		if left != right {
			heap.Push(pq, left-right)
		}
	}
	if pq.Len() == 0 {
		return 0
	}
	return heap.Pop(pq).(int)
}
