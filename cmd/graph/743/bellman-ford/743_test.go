package bellman_ford

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	edges := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	time := networkDelayTime(edges, 4, 2)
	fmt.Println(time)
	//graph := NewGraph(4, edges)
	//distance, err := graph.ShortestPath(2)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for edge, weight := range distance {
	//	fmt.Println(edge, weight)
	//}
}
