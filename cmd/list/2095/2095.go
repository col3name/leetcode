package _2095

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		head = nil
		return head
	}
	slow := head
	fast := head
	prev := head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	next := slow.Next
	prev.Next = next

	return head
}
