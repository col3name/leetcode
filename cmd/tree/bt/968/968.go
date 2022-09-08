package _968

import "fmt"

// 968. Binary Tree Cameras
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	result := 0
	var helper func(node *TreeNode, level int) int
	helper = func(node *TreeNode, level int) int {
		if node == nil {
			return 2
		}
		left := helper(node.Left, level+1) - 1
		right := helper(node.Right, level+1) - 1
		if left == 0 || right == 0 {
			result++
			return 3
		}
		if left > right {
			return left
		}
		return right
	}

	if helper(root, 0) < 2 {
		result++
	}
	return result
}
