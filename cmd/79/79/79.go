package _9

var height int
var width int
var boardIt [][]byte
var Word string

func exist(board [][]byte, word string) bool {
	height = len(board)
	width = len(board[0])
	boardIt = board
	Word = word
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if dfs(x, y, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(x, y, index int) bool {
	if index == len(Word)-1 && boardIt[y][x] == Word[index] {
		return true
	}

	cacheCh := boardIt[y][x]
	if boardIt[y][x] == Word[index] {
		boardIt[y][x] = ' '
		if x+1 < width && dfs(x+1, y, index+1) {
			return true
		}
		if x-1 >= 0 && dfs(x-1, y, index+1) {
			return true
		}
		if y+1 < height && dfs(x, y+1, index+1) {
			return true
		}
		if y-1 >= 0 && dfs(x, y-1, index+1) {
			return true
		}
	}

	boardIt[y][x] = cacheCh
	return false
}
