package _621

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

func leastInterval(tasks []byte, n int) int {
	//     init
	counter := map[byte]int{}
	for _, task := range tasks {
		counter[task]++
	}
	pq := maxHeap{}
	for _, count := range counter {
		heap.Push(&pq, count)
	}
	result := 0

	for pq.Len() > 0 {
		k := n + 1

		var freqs []int
		for k > 0 && pq.Len() > 0 {
			freq := heap.Pop(&pq).(int)
			freq--
			k--
			result++
			if freq > 0 {
				freqs = append(freqs, freq)
			}
		}
		for _, freq := range freqs {
			heap.Push(&pq, freq)
		}
		if pq.Len() == 0 {
			break
		}

		result += k
	}

	return result
}
