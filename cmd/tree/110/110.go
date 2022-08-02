package _110

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	result := true

	var walk func(*TreeNode) int
	walk = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := walk(node.Left)
		right := walk(node.Right)
		if math.Abs(float64(left-right)) > 1 {
			result = false
		}
		if left > right {
			return left + 1
		}
		return right + 1
	}

	walk(root)

	return result
}
