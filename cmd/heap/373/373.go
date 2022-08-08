package _373

import (
	"container/heap"
)

type Pair struct {
	first  int
	second int
	sum    int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) First() interface{} {
	return h[0]
}

func (h MaxHeap) Less(i int, j int) bool {
	return h[i].sum < h[j].sum
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
	*h = append(*h, i.(Pair))
}

type MapSet map[int]struct{}

func (m *MapSet) Has(key int) bool {
	_, ok := (*m)[key]
	return ok
}
func (m *MapSet) Set(key int) {
	(*m)[key] = struct{}{}
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return [][]int{}
	}
	pqLen := min(k, len(nums1))
	pq := make(MaxHeap, pqLen)
	for l := 0; l < k && l < len(nums1); l++ {
		pq[l] = Pair{first: l, second: 0, sum: nums1[l] + nums2[0]}
	}
	heap.Init(&pq)

	var result [][]int
	for ; k > 0 && pq.Len() > 0; k-- {
		curr := heap.Pop(&pq).(Pair)

		result = append(result, []int{nums1[curr.first], nums2[curr.second]})
		if curr.second+1 < len(nums2) {
			heap.Push(&pq, Pair{
				first:  curr.first,
				second: curr.second + 1,
				sum:    nums1[curr.first] + nums2[curr.second+1],
			})
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
