package _2233

import "container/heap"

type minheap []int

func (h minheap) Len() int {
	return len(h)
}
func (h minheap) First() interface{} {
	return h[0]
}

func (h minheap) Less(i int, j int) bool {
	return h[i] < h[j]
}

func (h minheap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minheap) Pop() interface{} {
	a := *h
	res := a[len(a)-1]
	*h = a[:len(a)-1]
	return res
}

func (h *minheap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

func maximumProduct(nums []int, k int) int {
	h := minheap(nums)
	heap.Init(&h)
	for i := 0; i < k; i++ {
		value := heap.Pop(&h).(int)
		heap.Push(&h, value+1)
	}
	result := heap.Pop(&h).(int)
	for h.Len() > 0 {
		result = (result * heap.Pop(&h).(int)) % (1e9 + 7)

	}

	return result
}
