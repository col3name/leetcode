package _148

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	middle := getMiddle(head)
	left := sortList(head)
	right := sortList(middle)
	return mergeTwoLists(left, right)
}

func getMiddle(head *ListNode) *ListNode {
	slow := head
	fast := slow.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	middle := slow.Next
	slow.Next = nil
	return middle
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	node := &ListNode{}
	root := node
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
	}
	if list2 != nil {
		node.Next = list2
	}
	return root.Next
}
