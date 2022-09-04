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

type ItemQ struct {
	Row  int
	Col  int
	node *TreeNode
}

func verticalTraversal(root *TreeNode) [][]int {
	m := make(map[int][]*Item)
	q := make([]*ItemQ, 0, 10)
	q = append(q, &ItemQ{
		Row:  0,
		Col:  0,
		node: root,
	})
	for len(q) > 0 {
		item := q[0]
		q = q[1:]
		node := item.node
		if node == nil {
			continue
		}
		col := item.Col
		row := item.Row
		if m[col] == nil {
			m[col] = make([]*Item, 0, 4)
		}
		m[col] = append(m[col], &Item{Value: node.Val, Row: row, Col: col})
		q = append(q, &ItemQ{Row: row + 1, Col: col - 1, node: node.Left})
		q = append(q, &ItemQ{Row: row + 1, Col: col + 1, node: node.Right})
	}

	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	sort.Ints(keys)
	result := make([][]int, len(keys))
	var ints []*Item
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
