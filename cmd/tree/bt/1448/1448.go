package _1448

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	count := 0
	var helper func(node *TreeNode, max int)
	helper = func(node *TreeNode, max int) {
		if node == nil {
			return
		}

		if node.Val >= max {
			count++
		}
		m := maxFN(node.Val, max)
		helper(node.Left, m)
		helper(node.Right, m)
	}
	helper(root, root.Val)

	return count
}

func maxFN(lhs, rhs int) int {
	if lhs > rhs {
		return lhs
	}

	return rhs
}
