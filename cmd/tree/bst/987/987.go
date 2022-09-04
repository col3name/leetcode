package _987

import "sort"

// 987. Vertical Order Traversal of a Binary Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Item struct {
	Value int
	Row   int
	Col   int
}

func verticalTraversal(root *TreeNode) [][]int {
	m := make(map[int][]Item)
	var helper func(node *TreeNode, row, col int)
	helper = func(node *TreeNode, row, col int) {
		if node == nil {
			return
		}
		_, ok := m[col]
		if !ok {
			m[col] = make([]Item, 0, 4)
		}
		m[col] = append(m[col], Item{Value: node.Val, Row: row, Col: col})
		helper(node.Left, row+1, col-1)
		helper(node.Right, row+1, col+1)
	}

	helper(root, 0, 0)
	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	sort.Ints(keys)
	result := make([][]int, len(keys))
	var ints []Item
	for i, key := range keys {
		ints = m[key]
		sort.Slice(ints, func(i, j int) bool {
			left := ints[i]
			right := ints[j]
			if left.Row == right.Row && left.Col == right.Col {
				return left.Value < right.Value
			}
			if left.Col == right.Col {
				return left.Row < right.Row
			}
			if left.Row == right.Row {
				return left.Col < right.Col
			}
			return false
		})
		result[i] = make([]int, len(ints))
		for j, item := range ints {
			result[i][j] = item.Value
		}
	}
	return result
}
