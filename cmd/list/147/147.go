package _147

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList1(head *ListNode) *ListNode {
	newHead := &ListNode{Val: math.MinInt}

	node := head
	resultNode := newHead.Next
	tail := newHead
	var next *ListNode
	var prev *ListNode

	for i := 0; node != nil; i++ {
		resultNode = newHead
		if node.Val <= resultNode.Val {
			next = newHead.Next
			newHead.Next = &ListNode{Val: node.Val}
			newHead.Next.Next = next
			node = node.Next
			continue
		}
		if node.Val >= tail.Val {
			tail.Next = &ListNode{Val: node.Val, Next: nil}
			tail = tail.Next
			node = node.Next
			continue
		}
		prev = resultNode
		for node.Val > resultNode.Val {
			prev = resultNode
			resultNode = resultNode.Next
		}
		if resultNode == nil {
			prev.Next = &ListNode{Val: node.Val, Next: nil}
		} else {
			next = prev.Next
			prev.Next = &ListNode{Val: node.Val, Next: next}
		}
		node = node.Next
	}

	return newHead.Next
}

func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	node := dummy
	for head != nil {
		for node.Next != nil && node.Next.Val < head.Val {
			node = node.Next
		}
		nextHead := head.Next
		node.Next, head.Next = head, node.Next

		head = nextHead

		if head != nil && head.Val < node.Next.Val {
			node = dummy
		}
	}

	return dummy.Next
}
