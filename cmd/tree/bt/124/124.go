package _124

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	result := math.MinInt
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := 0
		right := 0
		val := dfs(node.Left)
		if val > left {
			left = val
		}
		val = dfs(node.Right)
		if val > right {
			right = val
		}
		value := left + right + node.Val
		if value > result {
			result = value
		}
		if left >= right {
			return left + node.Val
		}
		return right + node.Val
	}

	dfs(root)
	return result
}
