package _2373

import "math"

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	result := make([][]int, n-2)
	for i := 0; i < n-2; i++ {
		result[i] = make([]int, n-2)
	}

	for i := 0; i < n-2; i++ {
		for j := 0; j < n-2; j++ {
			result[i][j] = findMax(grid, i, j)
		}
	}
	return result
}

func findMax(grid [][]int, i int, j int) int {
	max := math.MinInt32
	for r := i; r < i+2; r++ {
		for c := j; c < j+2; c++ {
			x := grid[r][c]
			if x > max {
				max = x
			}
		}
	}

	return max
}
