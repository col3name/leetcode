package _1669

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	start := list1
	end := list1
	for i := 0; i <= b && start != nil && end != nil; i++ {
		if i < a-1 {
			start = start.Next
		}
		if i <= b {
			end = end.Next
		}

	}
	tail := list2
	for tail != nil && tail.Next != nil {
		tail = tail.Next
	}
	start.Next = list2
	tail.Next = end
	return list1
}
