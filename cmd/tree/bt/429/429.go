package _429

// 429. N-ary Tree Level Order Traversal
type Node struct {
	Val      int
	Children []*Node
}

type Item struct {
	node  *Node
	level int
}

func levelOrder(root *Node) [][]int {
	mp := make([][]int, 0)
	maxLevel := -1
	q := make([]*Item, 0, 10)
	q = append(q, &Item{
		node:  root,
		level: 0,
	})
	for len(q) > 0 {
		item := q[0]
		q = q[1:]
		node := item.node
		level := item.level
		if node == nil {
			continue
		}
		if level > maxLevel {
			maxLevel = level
		}
		if level >= len(mp) {
			mp = append(mp, make([]int, 0, 10))
		}
		mp[level] = append(mp[level], node.Val)

		nextLevel := level + 1

		for _, child := range node.Children {
			q = append(q, &Item{node: child, level: nextLevel})
		}
	}

	return mp
}
