package _692

import "container/heap"

type Node struct {
	word string
	freq int
}

type NodeHeap []Node

func (h NodeHeap) Len() int { return len(h) }

func (h NodeHeap) Less(i, j int) bool {
	if h[i].freq == h[j].freq {
		return h[i].word < h[j].word
	}
	return h[i].freq > h[j].freq
}
func (h NodeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	lastEle := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return lastEle
}

func topKFrequent(words []string, k int) []string {
	dp := map[string]int{}
	for _, word := range words {
		if _, ok := dp[word]; ok {
			dp[word]++
		} else {
			dp[word] = 1
		}
	}

	h := NodeHeap{}
	for key, value := range dp {
		heap.Push(&h, Node{
			word: key,
			freq: value,
		})
	}

	result := make([]string, 0)
	for itr := 0; len(h) > 0 && itr < k; itr++ {
		lst := heap.Pop(&h).(Node).word
		result = append(result, lst)
	}
	return result
}
