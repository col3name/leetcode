package _109

// 199. Binary Tree Right Side View
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0, 100)
	var helper func(node *TreeNode, depth int)
	helper = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(result) {
			result = append(result, node.Val)
		}
		helper(node.Right, depth+1)
		helper(node.Left, depth+1)
	}
	helper(root, 0)
	return result
}
