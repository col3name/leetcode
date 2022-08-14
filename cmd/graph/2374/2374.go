package _2374

import (
	"math"
)

func edgeScore2(edges []int) int {
	scores := make([]int, len(edges))
	maxScore := 0

	ans := math.MaxInt32

	for i := 0; i < len(edges); i++ {
		to := edges[i]
		scores[to] = scores[to] + i

		if scores[to] > maxScore {
			maxScore = scores[to]
			ans = to
		}

		if scores[to] == maxScore {
			if to < ans {
				ans = to
			}
		}
	}

	return ans
}

func edgeScore(edges []int) int {
	adj := make(map[int]int, 0)
	for from, to := range edges {
		adj[to] += from
	}

	max := math.MinInt
	node := 0
	for key, ageScore := range adj {
		if ageScore > max {
			max = ageScore
			node = key
		} else if ageScore == max && key < node {
			node = key
		}
	}
	return node
}
