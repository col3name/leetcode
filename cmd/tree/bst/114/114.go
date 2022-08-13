package _114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	helper(root)
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}
	helper(root.Left)
	temp := root.Right
	root.Right = root.Left
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = temp

	helper(root.Right)
}
