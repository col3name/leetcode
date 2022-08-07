package _295

import "container/heap"

type FirstItem interface {
	First() interface{}
}
type maxheap []int
type minheap []int

func (h maxheap) Len() int {
	return len(h)
}

func (h maxheap) First() interface{} {
	return h[0]
}
func (h maxheap) Less(i int, j int) bool {
	return h[i] > h[j]
}

func (h maxheap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxheap) Pop() interface{} {
	a := *h
	res := a[len(a)-1]
	*h = a[:len(a)-1]
	return res
}

func (h *maxheap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

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

type MedianFinder struct {
	max *maxheap
	min *minheap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	max := &maxheap{}
	min := &minheap{}
	heap.Init(max)
	heap.Init(min)
	return MedianFinder{max: max,
		min: min,
	}
}

func (this *MedianFinder) AddNum(num int) {
	a := this.max
	b := this.min
	if b.Len() == 0 {
		heap.Push(b, num)
		return
	}
	if a.Len() == b.Len() {
		first := a.First().(int)
		if num < first {
			heap.Push(b, first)
			heap.Pop(a)
			heap.Push(a, num)
			return
		}
		heap.Push(b, num)
		return
	}
	first := b.First().(int)
	if num > first {
		heap.Push(a, first)
		heap.Pop(b)
		heap.Push(b, num)
		return
	}
	heap.Push(a, num)
}

func (this *MedianFinder) FindMedian() float64 {
	a := this.max
	b := this.min
	if len(*b) == 0 {
		return 0
	}
	if a.Len() != b.Len() {
		return float64(b.First().(int))
	}

	return (float64(b.First().(int)) + float64(a.First().(int))) / 2
}
