package _857

import (
	"math"
)

func Max(lhs, rhs int) int {
	return int(math.Max(float64(lhs), float64(rhs)))
}

func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	graph := make([][]int, n)

	for _, pair := range edges {
		key := pair[0]
		val := pair[1]
		if graph[key] == nil {
			graph[key] = make([]int, 0)
		}
		graph[key] = append(graph[key], val)
	}

	resultFrequencyMap := make(map[int][]int, n)

	var dfs func(to int, status []int) []int
	dfs = func(to int, status []int) []int {
		val, ok := resultFrequencyMap[to]
		if ok {
			return val
		}
		frequencies := make([]int, 27)
		if status[to] == 1 {
			frequencies[26] = -1
			return frequencies
		}
		status[to] = 1
		for _, child := range graph[to] {
			localMax := dfs(child, status)
			if localMax[26] == -1 {
				return localMax
			}
			for i := 0; i < 26; i++ {
				frequencies[i] = Max(frequencies[i], localMax[i])
			}
		}
		status[to] = 2
		color := colors[to] - 'a'
		frequencies[color]++
		resultFrequencyMap[to] = frequencies
		return frequencies
	}

	frequencies := make([]int, 26)

	status := make([]int, n)
	for i := 0; i < n; i++ {
		if status[i] != 0 {
			continue
		}
		localMax := dfs(i, status)
		if localMax[26] == -1 {
			for j := 0; j < 26; j++ {
				frequencies[i] = -1
			}
			break
		}
		for color := 0; i < 26; i++ {
			frequencies[color] = Max(frequencies[color], localMax[color])
		}
	}

	max := math.MaxInt64
	for _, freq := range frequencies {
		max = Max(freq, max)
	}

	return max
}
