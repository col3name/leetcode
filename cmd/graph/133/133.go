package _133

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	mapping := make(map[*Node]*Node)
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		newNode := new(Node)
		newNode.Val = node.Val
		mapping[node] = newNode
		for _, child := range node.Neighbors {
			if _, ok := mapping[child]; !ok {
				dfs(child)
			}
			newNode.Neighbors = append(newNode.Neighbors, mapping[child])
		}
	}

	dfs(node)

	return mapping[node]
}
