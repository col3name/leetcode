package _1338

import (
	"container/heap"
)

type Node struct {
	value int
	freq  int
}

type FreqHeap []Node

func (h FreqHeap) Len() int { return len(h) }

func (h FreqHeap) Less(i, j int) bool {
	if h[i].freq == h[j].freq {
		return h[i].value < h[j].value
	}
	return h[i].freq > h[j].freq
}
func (h FreqHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *FreqHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *FreqHeap) Pop() interface{} {
	old := *h
	lastEle := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return lastEle
}

func minSetSize1(arr []int) int {
	ma := map[int]int{}
	for _, num := range arr {
		ma[num]++
	}
	freqList := make([][]int, len(arr)+1)

	for num, freq := range ma {
		if freqList[freq] == nil {
			freqList[freq] = make([]int, 0)
		}
		freqList[freq] = append(freqList[freq], num)
	}

	count := 0
	result := 0
	sum := 0
	for freq := len(freqList) - 1; freq > 0; freq-- {
		nums := freqList[freq]
		if nums == nil || len(nums) == 0 {
			continue
		}

		for _, num := range nums {
			sum += num
			count += freq
			result++
			if count >= len(arr)/2 {
				return result
			}
		}
	}
	return len(arr)
}

func minSetSize(arr []int) int {
	ma := map[int]int{}
	for _, num := range arr {
		ma[num]++
	}
	freqHeap := make(FreqHeap, 0, len(ma))
	heap.Init(&freqHeap)

	for num, freq := range ma {
		heap.Push(&freqHeap, Node{value: num, freq: freq})
	}
	result := 0
	node := Node{}
	for freq := 0; freq < len(arr)/2 && freqHeap.Len() > 0; {
		node = heap.Pop(&freqHeap).(Node)
		result++
		freq += node.freq
	}
	return result
}
