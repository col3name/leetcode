package _606

import (
	"strconv"
	"strings"
)

// 606. Construct String from Binary Tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	var result strings.Builder

	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		result.WriteString(strconv.Itoa(node.Val))

		if node.Left == nil && node.Right != nil {
			result.WriteString("()")
		} else if node.Left != nil {
			result.WriteString("(")
			helper(node.Left)
			result.WriteString(")")
		}
		if node.Right != nil {
			result.WriteString("(")
			helper(node.Right)
			result.WriteString(")")
		}
	}
	helper(root)
	return result.String()
}
