package _86

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	beforeHead := &ListNode{}
	before := beforeHead
	afterHead := &ListNode{}
	after := afterHead
	node := head
	for node != nil {
		if node.Val < x {
			before.Next = &ListNode{Val: node.Val}
			before = before.Next
		} else {
			after.Next = &ListNode{Val: node.Val}
			after = after.Next
		}
		node = node.Next
	}
	before.Next = afterHead.Next
	return beforeHead.Next
}
