package _41

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	p2 := head
	p1 := p2
	for p2 != nil && p2.Next != nil {
		p2 = p2.Next.Next
		p1 = p1.Next
		if p1 == p2 {
			return true
		}
	}

	return false
}
