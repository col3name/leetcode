package _24

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node, next := head, head.Next

	for next != nil {
		node.Val, next.Val = next.Val, node.Val
		if next.Next == nil {
			break
		}
		next = next.Next.Next
		node = node.Next.Next
	}

	return head
}
