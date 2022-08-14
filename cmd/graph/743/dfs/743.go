package dfs

import (
	"math"
)

type Pair struct {
	Destination int
	TravelTime  int
}

func NewPair(destination int, travelTime int) Pair {
	return Pair{Destination: destination, TravelTime: travelTime}
}

func networkDelayTime(times [][]int, n int, k int) int {
	adj := make(map[int][]Pair, 0)

	var ok bool
	for _, time := range times {
		source := time[0]
		dest := time[1]
		travelTime := time[2]

		if _, ok = adj[source]; !ok {
			adj[source] = make([]Pair, 0)
		}
		adj[source] = append(adj[source], NewPair(dest, travelTime))
	}

	signalReceivedAt := make([]int, n+1)
	for i := 0; i < len(signalReceivedAt); i++ {
		signalReceivedAt[i] = math.MaxInt32
	}

	var bfs func([]int, int)
	bfs = func(signalReceivedAt []int, sourceNode int) {
		q := []int{sourceNode}
		signalReceivedAt[sourceNode] = 0

		var (
			childs       []Pair
			time         int
			neighborNode int
			currentNode  int
			arrivalTime  int
		)
		for len(q) > 0 {
			currentNode = q[0]
			q = q[1:]
			childs, ok = adj[currentNode]
			if !ok {
				continue
			}
			for _, edge := range childs {
				time = edge.TravelTime
				neighborNode = edge.Destination

				arrivalTime = signalReceivedAt[currentNode] + time
				if signalReceivedAt[neighborNode] > arrivalTime {
					signalReceivedAt[neighborNode] = arrivalTime
					q = append(q, neighborNode)
				}
			}
		}
	}
	bfs(signalReceivedAt, k)
	result := math.MinInt32
	for i := 1; i <= n; i++ {
		if signalReceivedAt[i] > result {
			result = signalReceivedAt[i]
		}
	}
	if result == math.MaxInt32 {
		return -1
	}
	return result
}
