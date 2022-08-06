package _1129

type Node struct {
	To    int
	IsRed bool
}

type Adj map[int][]int

type Queue []Node

func (q *Queue) Push(value Node) {
	*q = append(*q, value)
}

func (q *Queue) Ppp() Node {
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

type Visit [2][]int
type Result []int

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	adj := buildAdj(redEdges, blueEdges)
	visit := dfs(adj, n)
	return calculateResult(visit, n)
}

func buildAdj(redEdges [][]int, blueEdges [][]int) []Adj {
	adj := make([]Adj, 2)
	adj[0] = buildAdjColor(redEdges)
	adj[1] = buildAdjColor(blueEdges)
	return adj
}

func calculateResult(visit *Visit, n int) Result {
	result := make(Result, n)
	for i := range result {
		if v1, v2 := visit[0][i], visit[1][i]; v1 < 0 || v2 < 0 {
			result[i] = max(v1, v2)
		} else {
			result[i] = min(v1, v2)
		}
	}
	return result
}

func dfs(adj []Adj, n int) *Visit {
	visit := buildVisit(n)

	q := Queue{Node{To: 0, IsRed: false}, Node{To: 0, IsRed: true}}

	for !q.Empty() {
		node := q.Ppp()
		walkEdges(&node, &q, &adj, visit)
	}

	return visit
}

func walkEdges(node *Node, q *Queue, adj *[]Adj, visit *Visit) {
	nextColor := colorToInt(!node.IsRed)
	nextDistance := 1 + visit[colorToInt(node.IsRed)][node.To]
	edges := (*adj)[nextColor][node.To]
	visiting := visit[nextColor]

	for _, edge := range edges {
		if visiting[edge] < 0 {
			visiting[edge] = nextDistance
			q.Push(Node{To: edge, IsRed: nextColor == 0})
		}
	}
}

func buildVisit(n int) *Visit {
	var visit Visit
	visit[0] = make([]int, n)
	visit[1] = make([]int, n)
	for i := range visit[0] {
		visit[0][i] = -1
		visit[1][i] = -1
	}
	visit[0][0] = 0
	visit[1][0] = 0
	return &visit
}

func buildAdjColor(edges [][]int) Adj {
	adj := make(Adj)
	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		if _, ok := adj[from]; !ok {
			adj[from] = make([]int, 0)
		}
		adj[from] = append(adj[from], to)
	}

	return adj
}

func max(lhs, rhs int) int {
	if lhs > rhs {
		return lhs
	}
	return rhs
}

func min(lhs, rhs int) int {
	if lhs < rhs {
		return lhs
	}
	return rhs
}

func colorToInt(isRed bool) int {
	if isRed {
		return 0
	}
	return 1
}
