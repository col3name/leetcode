package _1334

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
		from := edge[0]
		to := edge[1]
		weight := edge[2]
		edges = append(edges, &Edge{
			From:   uint(from),
			To:     uint(to),
			Weight: weight,
		})
		edges = append(edges, &Edge{
			From:   uint(to),
			To:     uint(from),
			Weight: weight,
		})
	}
	return &Graph{
		edges:    edges,
		vertices: vertices,
	}
}

func (g *Graph) ShortestPath(source, distanceThreshold int) ([]int, error) {
	distance := g.bellmanFord(source, distanceThreshold)
	return distance, nil
}

func (g *Graph) bellmanFord(source, distanceThreshold int) []int {
	vertices := g.vertices
	distance := make([]int, vertices)
	for i := 0; i < vertices; i++ {
		distance[i] = math.MaxInt32
	}
	distance[source] = 0
	var (
		from        uint
		to          uint
		weight      int
		newDistance int
	)

	for i := 0; i < vertices-1; i++ {
		for _, edge := range g.edges {
			from = edge.From
			to = edge.To
			weight = edge.Weight
			newDistance = distance[from] + weight
			if distance[from] != math.MinInt32 && newDistance < distance[to] && newDistance <= distanceThreshold {
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

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	g := NewGraph(n, edges)
	minCount := math.MaxInt32
	cities := make([]int, n)
	for from := 0; from < n; from++ {
		distance := g.bellmanFord(from, distanceThreshold)
		reachableCities := getReachableCities(distance)

		cities[from] = reachableCities
		if reachableCities < minCount {
			minCount = reachableCities
		}
	}
	for i := n - 1; i >= 0; i-- {
		if minCount == cities[i] {
			return i
		}
	}
	return 0
}

func getReachableCities(distance []int) int {
	reachableCity := 0

	for _, weight := range distance {
		if weight != math.MaxInt32 && weight != 0 {
			reachableCity++
		}
	}

	return reachableCity
}
