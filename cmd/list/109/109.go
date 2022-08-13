package _109

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var prev *ListNode

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		prev = slow
		fast = fast.Next.Next
		slow = slow.Next
	}
	if prev != nil {
		prev.Next = nil
	}

	tree := &TreeNode{Val: slow.Val}
	if head == slow {
		return tree
	}
	tree.Left = sortedListToBST(head)
	tree.Right = sortedListToBST(slow.Next)
	return tree
}
