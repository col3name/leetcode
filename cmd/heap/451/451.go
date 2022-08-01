package _451

import (
	"container/heap"
	"fmt"
	"strings"
)

type Node struct {
	char byte
	freq int
}

type NodeHeap []Node

func (h NodeHeap) Len() int { return len(h) }

func (h NodeHeap) Less(i, j int) bool {
	if h[i].freq == h[j].freq {
		return h[i].char < h[j].char
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

func frequencySort(s string) string {
	if len(s) == 0 {
		return ""
	}
	dp := make(map[byte]int, 0)
	for _, char := range []byte(s) {
		if _, ok := dp[char]; ok {
			dp[char]++
		} else {
			dp[char] = 1
		}
	}

	if len(dp) == 1 {
		var result strings.Builder
		l := dp[s[0]]
		for i := 0; i < l; i++ {
			result.WriteByte(s[0])
		}
		return result.String()
	}

	h := NodeHeap{}
	for key, value := range dp {
		heap.Push(&h, Node{char: key, freq: value})
	}
	fmt.Println(h)
	var result strings.Builder
	length := h.Len()
	for i := 0; i < length; i++ {
		node := heap.Pop(&h).(Node)
		for j := 0; j < node.freq; j++ {
			result.WriteByte(node.char)
		}
	}
	return result.String()
}
