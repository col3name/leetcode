package _725

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	length := 0
	node := head
	for node != nil {
		node = node.Next
		length++
	}
	countElements := length / k
	result := make([]*ListNode, k)
	node = head
	for i := 0; i < k; i++ {
		current := &ListNode{}
		partHead := current
		for j := 0; j < countElements && node != nil; j++ {
			current.Next = &ListNode{Val: node.Val}
			node = node.Next
			current = current.Next
		}
		if i < length%k && node != nil {
			current.Next = &ListNode{Val: node.Val}
			current = current.Next
			node = node.Next
		}
		result[i] = partHead.Next
	}

	return result
}
