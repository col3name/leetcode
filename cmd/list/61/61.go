package _61

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	node := head
	nodes := make([]*ListNode, 0)
	for node != nil {
		nodes = append(nodes, node)
		node = node.Next
	}
	length := len(nodes)
	k = k % length
	if k > 0 {
		nodes[length-1-k].Next = nil
		nodes[length-1].Next = nodes[0]
		head = nodes[length-k]
	}
	return head
}
