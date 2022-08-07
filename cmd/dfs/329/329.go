package _329

import "fmt"

var memory [][]int
var matrix [][]int

func longestIncreasingPath(input [][]int) int {
	width := len(input)
	height := len(input[0])

	fmt.Println("width", width, height)
	matrix = input
	memory = make([][]int, height)
	for i := 0; i < height; i++ {
		tmp := make([]int, width)
		for j := 0; j < width; j++ {
			tmp[j] = -1
		}
		memory[i] = tmp
	}

	maxLength := -1
	var length int
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			length = dfs(i, j, height, width)
			if length > maxLength {
				maxLength = length
			}
		}
	}

	return maxLength
}

type Direction struct {
	x int
	y int
}

var directions = []Direction{
	{x: 1, y: 0},
	{x: -1, y: 0},
	{x: 0, y: 1},
	{x: 0, y: -1},
}

func dfs(row, col, height, width int) int {
	fmt.Println(row, col)
	if memory[col][row] != -1 {
		return memory[col][row]
	}
	result := 1

	var length int

	//for _, direction := range directions {
	//	x = direction.x + row
	//	y = direction.y + col
	if col < 0 || row < 0 || row > width-1 || col > height-1 || matrix[col][row] <= matrix[row][col] {
		return -1
	}
	//	//fmt.Println("dfs")
	//	length = dfs(x, y, height, width) + 1
	//	if length > result {
	//		result = length
	//	}
	//}
	length = dfs(row, col-1, height, width) + 1
	if length > result {
		result = length
	}
	length = dfs(row, col+1, height, width) + 1
	if length > result {
		result = length
	}
	length = dfs(row-1, col, height, width) + 1
	if length > result {
		result = length
	}
	length = dfs(row+1, col, height, width) + 1
	if length > result {
		result = length
	}
	memory[col][row] = result
	return result
}
