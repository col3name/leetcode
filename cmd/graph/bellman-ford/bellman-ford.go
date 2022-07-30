package bellman_ford

import (
	"errors"
	"math"
)

type Edge struct {
	source      int
	destination int
	weight      int
}

type Graph struct {
	vertices  int
	edges     int
	edgesList []Edge
}

func NewGraph(vertices int, edges int) *Graph {
	data := make([]Edge, edges)
	for i := 0; i < len(data); i++ {
		data[i] = Edge{}
	}
	return &Graph{
		vertices:  vertices,
		edges:     edges,
		edgesList: data,
	}
}

var ErrorCycleGraph = errors.New("cycleGraphError")

func (this *Graph) ShortestPath(from int) error {
	distance := make([]int, this.vertices)
	for i := 0; i < len(distance); i++ {
		distance[i] = math.MaxInt32
	}
	distance[from] = 0

	var (
		source      int
		destination int
		weight      int
	)
	for i := 0; i < this.vertices; i++ {
		for j, pair := range this.edgesList {
			source = pair.source
			destination = pair.destination
			weight = pair.weight
			if distance[source] != math.MaxInt32 && distance[source]+weight < distance[destination] {
				distance[destination] += weight
			}
		}
	}

	for _, edge := range this.edgesList {
		source = edge.source
		destination = edge.destination
		weight = edge.weight
		if distance[source] != math.MaxInt32 && distance[source]+weight < distance[destination] {
			return ErrorCycleGraph
		}
	}
	return nil
}
