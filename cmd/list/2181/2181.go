package _2181

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var sum int
	node := head.Next
	dummy := &ListNode{}
	slider := dummy

	for node != nil {
		sum = 0
		for node != nil && node.Val != 0 {
			sum += node.Val
			node = node.Next
		}
		slider.Next = &ListNode{Val: sum}
		slider = slider.Next
		if node == nil {
			return dummy.Next
		}
		node = node.Next
	}

	return dummy.Next
}
