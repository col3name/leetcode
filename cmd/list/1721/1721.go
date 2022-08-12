package _1721

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	array := make([]int, 0)
	node := head
	for node != nil {
		array = append(array, node.Val)
		node = node.Next
	}
	length := len(array)
	array[k-1], array[length-k] = array[length-k], array[k-1]
	tmp := &ListNode{}
	node = tmp
	for _, num := range array {
		tmp.Next = &ListNode{Val: num}
		tmp = tmp.Next
	}
	return node.Next
}
