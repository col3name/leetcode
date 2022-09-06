package _814

// 814. Binary Tree Pruning
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	var helper func(node *TreeNode) bool
	helper = func(node *TreeNode) bool {
		if node == nil {
			return false
		}

		containsLeft := helper(node.Left)
		containsRight := helper(node.Right)
		if !containsLeft {
			node.Left = nil
		}
		if !containsRight {
			node.Right = nil
		}
		return node.Val == 1 || containsLeft || containsRight
	}

	if !helper(root) {
		return nil
	}
	return root
}
