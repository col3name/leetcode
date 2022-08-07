package _378

import "container/heap"

type Minheap []int

func (h Minheap) Len() int {
	return len(h)
}
func (h Minheap) First() interface{} {
	return h[0]
}

func (h Minheap) Less(i int, j int) bool {
	return h[i] < h[j]
}

func (h Minheap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Minheap) Pop() interface{} {
	a := *h
	res := a[len(a)-1]
	*h = a[:len(a)-1]
	return res
}

func (h *Minheap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

func kthSmallest(matrix [][]int, k int) int {
	pq := &Minheap{}
	heap.Init(pq)
	for _, row := range matrix {
		for _, num := range row {
			heap.Push(pq, num)
		}
	}
	if pq.Len() == 0 {
		return 0
	}
	if pq.Len() == 1 {
		return heap.Pop(pq).(int)
	}

	for i := 1; i < k; i++ {
		heap.Pop(pq)
	}
	return heap.Pop(pq).(int)
}
