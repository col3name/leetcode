package _1457

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Item struct {
	node *TreeNode
	path int
}

// 1457. Pseudo-Palindromic Paths in a Binary Tree
func pseudoPalindromicPathsRecursion(root *TreeNode) int {
	count := 0
	var preorder func(node *TreeNode, path int)
	preorder = func(node *TreeNode, path int) {
		if node == nil {
			return
		}
		path ^= 1 << node.Val

		if node.Left == nil && node.Right == nil {
			if path&(path-1) == 0 {
				count++
			}
		}

		preorder(node.Left, path)
		preorder(node.Right, path)
	}
	preorder(root, 0)
	return count
}

func pseudoPalindromicPaths(root *TreeNode) int {
	count := 0
	q := make([]*Item, 0, 100000)
	q = append(q, &Item{node: root, path: 0})
	var item *Item
	var node *TreeNode
	var path int
	for len(q) > 0 {
		item = q[0]
		q = q[1:]

		if item.node == nil {
			continue
		}
		node = item.node
		path = item.path

		path ^= 1 << node.Val

		if node.Left == nil && node.Right == nil {
			if path&(path-1) == 0 {
				count++
			}
		}
		q = append(q, &Item{node: node.Left, path: path}, &Item{node: node.Right, path: path})
	}

	return count
}
