package _92

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	head = dummy
	for i := 0; i < left-1; i++ {
		head = head.Next
	}
	var (
		prev    *ListNode
		current = head.Next
	)
	for i := 0; i <= right-left; i++ {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	head.Next.Next = current
	head.Next = prev
	return dummy.Next
}
