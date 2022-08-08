package _912

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

func sortArray(nums []int) []int {
	pq := &Minheap{}
	heap.Init(pq)
	for _, num := range nums {
		heap.Push(pq, num)
	}
	var result []int

	for pq.Len() > 0 {
		result = append(result, heap.Pop(pq).(int))
	}
	return result
}
