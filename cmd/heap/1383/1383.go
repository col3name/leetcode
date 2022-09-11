package _1383

import (
	"container/heap"
	"math"
	"sort"
)

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Pop() interface{} {
	a := *h
	res := a[len(a)-1]
	*h = a[:len(a)-1]
	return res
}

func (h *minHeap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

func maxPerformance(n int, speeds []int, efficiencies []int, k int) int {
	engineers := make([][2]int, n)

	for i, speed := range speeds {
		engineers[i] = [2]int{efficiencies[i], speed}
	}

	sort.Slice(engineers, func(i, j int) bool {
		return engineers[i][0] > engineers[j][0]
	})
	pq := &minHeap{}
	heap.Init(pq)
	var totalSpeed int
	result := math.MinInt32
	var efficiency int
	for _, engineer := range engineers {
		heap.Push(pq, engineer[1])
		totalSpeed += engineer[1]
		if pq.Len() > k {
			totalSpeed -= heap.Pop(pq).(int)
		}
		efficiency = engineer[0]
		if totalSpeed*efficiency > result {
			result = totalSpeed * efficiency
		}
	}
	return result % int(1e9+7)
}
