package _653

// 653. Two Sum IV - Input is a BST
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	set := make(map[int]int)

	var helper func(node *TreeNode, set map[int]int) bool
	helper = func(node *TreeNode, set map[int]int) bool {
		if node == nil {
			return false
		}

		if _, ok := set[k-node.Val]; ok {
			return true
		}
		set[node.Val]++
		return helper(node.Left, set) || helper(node.Right, set)
	}

	return helper(root, set)
}
