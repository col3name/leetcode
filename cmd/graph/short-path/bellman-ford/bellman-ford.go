package bellman_ford

import (
	"errors"
	"math"
)

var ErrorExistCycle = errors.New("errExistCycle")

type Edge struct {
	From   uint
	To     uint
	Weight int
}

type Graph struct {
	edges    []*Edge
	vertices int
}

func NewGraph(vertices int, graph [][]int) *Graph {
	edges := make([]*Edge, 0, vertices)

	for _, edge := range graph {
		edges = append(edges, &Edge{
			From:   uint(edge[0] - 1),
			To:     uint(edge[1] - 1),
			Weight: edge[2],
		})
	}
	return &Graph{
		edges:    edges,
		vertices: vertices,
	}
}

func (g *Graph) ShortestPath(source int) ([]int, error) {
	source--
	distance := g.bellmanFord(source)
	return distance, g.FindNegativeCycle(distance)
}

func (g *Graph) bellmanFord(source int) []int {
	distance := make([]int, g.vertices)
	vertices := g.vertices
	for i := 0; i < vertices; i++ {
		distance[i] = math.MaxInt32
	}
	distance[source] = 0
	var from uint
	var to uint
	var weight int
	for i := 0; i < vertices-1; i++ {
		for _, edge := range g.edges {
			from = edge.From
			to = edge.To
			weight = edge.Weight
			newDistance := distance[from] + weight
			if distance[from] != math.MinInt32 && newDistance < distance[to] {
				distance[to] = newDistance
			}
		}
	}
	return distance
}

func (g *Graph) FindNegativeCycle(distance []int) error {
	var (
		from   uint
		to     uint
		weight int
	)
	for _, edge := range g.edges {
		from = edge.From
		to = edge.To
		weight = edge.Weight
		if distance[from]+weight < distance[to] {
			return ErrorExistCycle
		}
	}
	return nil
}
