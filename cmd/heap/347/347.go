package _347

import "container/heap"

type Item struct {
	Num  int
	Freq int
}
type MaxHeap []Item

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) First() interface{} {
	return h[0]
}

func (h MaxHeap) Less(i int, j int) bool {
	return h[i].Freq > h[j].Freq
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
	*h = append(*h, i.(Item))
}
func topKFrequent(nums []int, k int) []int {
	freqs := map[int]int{}
	for _, num := range nums {
		freqs[num]++
	}
	pq := &MaxHeap{}
	heap.Init(pq)
	for num, freq := range freqs {
		heap.Push(pq, Item{
			Num:  num,
			Freq: freq,
		})
	}
	length := pq.Len()
	var result []int
	for i := 0; i < k && i < length; i++ {
		result = append(result, heap.Pop(pq).(Item).Num)
	}
	return result
}
