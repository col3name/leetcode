package _1054

import "container/heap"

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

func rearrangeBarcodes(barcodes []int) []int {
	m := map[int]int{}

	freqHeap := &FreqHeap{}
	heap.Init(freqHeap)
	for _, barcode := range barcodes {
		m[barcode]++
	}
	for value, freq := range m {
		heap.Push(freqHeap, Node{
			value: value,
			freq:  freq,
		})
	}
	var res []int

	last := 0
	for freqHeap.Len() > 0 {
		node := heap.Pop(freqHeap).(Node)
		if node.value == last {
			nextNode := heap.Pop(freqHeap).(Node)
			res = append(res, nextNode.value)
			nextNode.freq--
			if nextNode.freq != 0 {
				heap.Push(freqHeap, nextNode)
			}
			heap.Push(freqHeap, node)
			last = nextNode.value
		} else {
			res = append(res, node.value)
			node.freq--
			last = node.value
			if node.freq != 0 {
				heap.Push(freqHeap, node)
			}
		}
	}
	return res
}
