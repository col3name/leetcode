package _143

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	arr := make([]*ListNode, 0, 100)
	node := head
	for node != nil {
		arr = append(arr, node)
		node = node.Next
	}
	length := len(arr)
	node = head
	for i := len(arr) - 1; i > (length-1)/2; i-- {
		next := node.Next
		node.Next = arr[i]
		node.Next.Next = next
		node = next
	}
	node.Next = nil
}
