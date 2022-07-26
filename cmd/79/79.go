package main

type Point struct {
	row   int
	col   int
	level int
}

var height int
var width int

func exist(board [][]byte, word string) bool {
	height := len(board)
	width := len(board[0])
	var walk func(x, y, index int) bool
	walk = func(x, y, index int) bool {
		if index == len(word)-1 && board[y][x] == word[index] {
			return true
		}
		cacheChar := board[y][x]
		if board[y][x] == word[index] {
			board[y][x] = ' '
			if x+1 < width && walk(x+1, y, index+1) {
				return true
			}
			if x-1 >= 0 && walk(x-1, y, index+1) {
				return true
			}
			if y+1 < height && walk(x, y+1, index+1) {
				return true
			}
			if y-1 >= 0 && walk(x, y-1, index+1) {
				return true
			}
		}
		board[y][x] = cacheChar
		return false
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if walk(x, y, 0) {
				return true
			}
		}
	}
	return false
}

func exist1(board [][]byte, word string) bool {
	//var visited [][]bool
	//visited = initVisited(height, width)
	height := len(board)
	width := len(board[0])
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			ch := word[0]
			boardCh := board[y][x]
			if boardCh == ch {
				//visited = initVisited(width, height)
				if dfs(board, word, y, x, 0) {
					return true
				}

			}
		}
	}
	return false
}

var dir = [][]int{
	{0, 1}, {0, -1}, {1, 0}, {-1, 0},
}

func dfs(board [][]byte, word string, x, y, index int) bool {
	if index == len(word)-1 {
		return true
	}

	cacheCh := board[y][x]
	tempX, tempY := 0, 0

	for d := 0; d < len(dir); d++ {
		tempX = x + dir[d][0]
		tempY = y + dir[d][1]
		if tempX >= 0 &&
			tempX < len(board)-1 &&
			tempY >= 0 &&
			tempY < len(board[0])-1 &&
			board[tempX][tempY] == word[index+1] {
			if dfs(board, word, tempX, tempY, index+1) {
				return true
			}
		}
	}
	board[x][y] = cacheCh
	return false
}
func initVisited(columns, rows int) [][]bool {
	visited := make([][]bool, columns)
	for i := 0; i < columns; i++ {
		row := make([]bool, rows)
		for j := 0; j < rows; j++ {
			row[j] = false
		}
		visited[i] = row
	}
	return visited
}
