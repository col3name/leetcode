package _101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var dfs func(*TreeNode, *TreeNode) bool
	dfs = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		return left.Val == right.Val && dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
	}
	if root == nil {
		return true
	}

	return dfs(root.Left, root.Right)
}
