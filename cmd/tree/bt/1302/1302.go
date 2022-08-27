package _1302

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	maxHeight := -1
	sum := 0
	var findSum func(node *TreeNode, height int)
	findSum = func(node *TreeNode, height int) {
		if node == nil {
			return
		}
		if height > maxHeight {
			maxHeight = height
			sum = 0
		}
		if height == maxHeight {
			sum += node.Val
		}
		findSum(node.Left, height+1)
		findSum(node.Right, height+1)
	}
	findSum(root, 0)
	return sum
}

func deepestLeavesSumVersion1(root *TreeNode) int {
	maxHeight := 0
	sum := 0
	var findSum func(node *TreeNode, height int)
	findSum = func(node *TreeNode, height int) {
		if node == nil {
			return
		}
		if height > maxHeight {
			maxHeight = height
			sum = node.Val
		} else {
			sum += node.Val
		}
		findSum(node.Left, height+1)
		findSum(node.Right, height+1)
	}
	findSum(root, 0)
	return sum
}
