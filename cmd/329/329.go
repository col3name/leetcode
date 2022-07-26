package _29

import (
	"fmt"
	"math"
)

func longestIncreasingPath(matrix [][]int) int {
	maxLength := -1
	n := len(matrix) - 1
	m := len(matrix[0]) - 1

	memory := make([][]int, n+1)
	for i := 0; i < len(matrix[0]); i++ {
		memory[i] = make([]int, m+1)
		for j := 0; j < m; j++ {
			memory[i][j] = 0
		}
	}

	//  val := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			maxLength = max(maxLength, move(i, j, n, m, matrix, memory))
		}
	}

	fmt.Println(memory)

	return maxLength
}

func move(row, col, n, m int, matrix, memory [][]int) int {
	if memory[row][col] == 0 {
		count := 1

		// down
		if row < n && matrix[row+1][col] > matrix[row][col] {
			count = max(count, move(row+1, col, n, m, matrix, memory)+1)
		}
		// right
		if col < m && matrix[row][col+1] > matrix[row][col] {
			count = max(count, move(row, col+1, n, m, matrix, memory)+1)
		}
		//	up
		if col < m && matrix[row-1][col] > matrix[row][col] {
			count = max(count, move(row-1, col, n, m, matrix, memory)+1)
		}
		// right
		if col < m && matrix[row][col-1] > matrix[row][col] {
			count = max(count, move(row, col-1, n, m, matrix, memory)+1)
		}

		memory[row][col] = count
	}
	return memory[row][col]
}

func max(lhs, rhs int) int {
	return int(math.Max(float64(lhs), float64(rhs)))
}
