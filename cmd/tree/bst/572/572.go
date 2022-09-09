package _572

// 572. Subtree of Another Tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if isEqualTree(root, subRoot) {
		return true
	}
	if root == nil {
		return false
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isEqualTree(root, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == subRoot
	}
	return root.Val == subRoot.Val &&
		isEqualTree(root.Left, subRoot.Left) &&
		isEqualTree(root.Right, subRoot.Right)
}
