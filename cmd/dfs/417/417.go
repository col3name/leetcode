package _417

import (
	"fmt"
	"math"
)

func PacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return [][]int{}
	}
	height := len(matrix)
	width := len(matrix[0])
	pacific := make([][]int, height)
	atlantic := make([][]int, height)
	for i := 0; i < height; i++ {
		pacific[i] = make([]int, width)
		atlantic[i] = make([]int, width)
	}

	// top and bottom
	for col := 0; col < width; col++ {
		dfs(&matrix, 0, col, math.MinInt32, &pacific)
		dfs(&matrix, len(matrix)-1, col, math.MinInt32, &atlantic)
	}

	// left and right
	for row := 0; row < height; row++ {
		dfs(&matrix, row, 0, math.MinInt32, &pacific)
		dfs(&matrix, row, len(matrix[0])-1, math.MinInt32, &atlantic)
	}

	result := make([][]int, 0)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if pacific[i][j] == 1 && atlantic[i][j] == 1 {
				tmp := make([]int, 0)
				tmp = append(tmp, i)
				tmp = append(tmp, j)
				result = append(result, tmp)
			}
		}
	}

	return result
}

func dfs(matrix *[][]int, row, col, previousValue int, ocean *[][]int) {
	if row < 0 || col < 0 || row > len(*matrix)-1 || col > len((*matrix)[0])-1 {
		return
	}
	if (*matrix)[row][col] < previousValue {
		return
	}
	if (*ocean)[row][col] == 1 {
		return
	}

	(*ocean)[row][col] = 1

	value := (*matrix)[row][col]
	dfs(matrix, row-1, col, value, ocean)
	dfs(matrix, row+1, col, value, ocean)
	dfs(matrix, row, col-1, value, ocean)
	dfs(matrix, row, col+1, value, ocean)
}
