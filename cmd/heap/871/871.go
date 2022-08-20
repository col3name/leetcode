package _871

import "container/heap"

type maxHeap []int

func (m maxHeap) Len() int           { return len(m) }
func (m maxHeap) Less(i, j int) bool { return m[i] > m[j] }
func (m maxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}
func (m *maxHeap) Pop() interface{} {
	old := *m
	v := old[len(old)-1]
	*m = old[:len(old)-1]
	return v
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	h := &maxHeap{}
	heap.Init(h)
	countStation := 0
	prev := 0
	for _, station := range stations {
		location := station[0]
		capacity := station[1]
		startFuel -= location - prev
		for h.Len() > 0 && startFuel < 0 {
			startFuel += heap.Pop(h).(int)
			countStation++
		}
		if startFuel < 0 {
			return -1
		}
		heap.Push(h, capacity)
		prev = location
	}

	return countStation
}
