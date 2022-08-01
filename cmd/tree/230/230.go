package _230

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	ints := make([]int, 0)
	inorder(root, &ints)
	return ints[k-1]
}

func inorder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inorder(root.Right, arr)
}
