package _142

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := slow
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			slow = head
			for fast != slow {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}
