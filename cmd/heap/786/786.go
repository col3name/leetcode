package _786

import (
	"container/heap"
)

type Item struct {
	Left  int16
	Right int16
	Value float32
}
type MinHeap []Item

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) First() interface{} {
	return h[0]
}

func (h MinHeap) Less(i int, j int) bool {
	return h[i].Value < h[j].Value
}

func (h MinHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Pop() interface{} {
	a := *h
	res := a[len(a)-1]
	*h = a[:len(a)-1]
	return res
}

func (h *MinHeap) Push(i interface{}) {
	*h = append(*h, i.(Item))
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	pq := make(MinHeap, 0, len(arr))
	heap.Init(&pq)

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			heap.Push(&pq, Item{
				Left:  int16(arr[i]),
				Right: int16(arr[j]),
				Value: float32(arr[i]) / float32(arr[j]),
			})
		}
	}
	for pq.Len() > 0 && k > 1 {
		heap.Pop(&pq)
		k--
	}
	item := heap.Pop(&pq).(Item)
	return []int{int(item.Left), int(item.Right)}
}
