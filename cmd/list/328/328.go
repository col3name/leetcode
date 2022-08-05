package _328

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	evenHead := head.Next
	even := head.Next
	odd := head
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head
}
