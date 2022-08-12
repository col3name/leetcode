package _21

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	root := &ListNode{}
	node := root
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			node.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		node = node.Next
	}
	if list1 != nil {
		node.Next = list1
	} else if list2 != nil {
		node.Next = list2
	}
	return root.Next
}
