package _23

import "container/heap"

/**
 * Definition for singly-linked list.
 *
 *
 *
 *
 */
type ListNode struct {
	Val  int
	Next *ListNode
}
type Minheap []int

func (h Minheap) Len() int {
	return len(h)
}
func (h Minheap) First() interface{} {
	return h[0]
}

func (h Minheap) Less(i int, j int) bool {
	return h[i] < h[j]
}

func (h Minheap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Minheap) Pop() interface{} {
	a := *h
	res := a[len(a)-1]
	*h = a[:len(a)-1]
	return res
}

func (h *Minheap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := &Minheap{}
	heap.Init(pq)

	for _, root := range lists {
		node := root
		for node != nil {
			heap.Push(pq, node.Val)
			node = node.Next
		}
	}
	node := &ListNode{}
	head := node
	for pq.Len() > 0 {
		node.Next = &ListNode{
			Val:  heap.Pop(pq).(int),
			Next: nil,
		}
		node = node.Next
	}

	return head.Next
}
