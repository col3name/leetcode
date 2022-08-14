package bellman_ford

import (
	"fmt"
	"testing"
)

func calculate(data [][]int, vertices, from int) {
	graph := NewGraph(vertices, data)
	shortestPath, err := graph.ShortestPath(from)
	if err != nil {
		fmt.Println(err)
		return
	}
	for edge, weight := range shortestPath {
		fmt.Println(edge, weight)
	}

}
func TestName(t *testing.T) {
	data := [][]int{{0, 1, -1}, {0, 2, 4},
		{1, 2, 3}, {1, 3, 2},
		{1, 4, 2}, {3, 2, 5},
		{3, 1, 1}, {4, 3, -3}}
	calculate(data, 5, 0)
}

func TestName2(t *testing.T) {
	data := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	calculate(data, 4, 2)
}
