package _98

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var dfs func(*TreeNode, int, int) bool
	dfs = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if min >= node.Val || max <= node.Val {
			return false
		}

		return dfs(node.Left, min, node.Val) && dfs(node.Right, node.Val, max)
	}

	return dfs(root, math.MinInt, math.MaxInt)
}
