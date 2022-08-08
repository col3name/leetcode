package _695

func maxAreaOfIsland(grid [][]int) int {
	rows := len(grid)
	columns := len(grid[0])
	seen := make([][]bool, rows)
	for i := range seen {
		seen[i] = make([]bool, columns)
	}
	var area func(int, int) int
	area = func(row, col int) int {
		if row < 0 || col < 0 || row >= rows || col >= columns || seen[row][col] || grid[row][col] == 0 {
			return 0
		}
		seen[row][col] = true
		return 1 + area(row+1, col) + area(row-1, col) + area(row, col+1) + area(row, col-1)
	}
	result := 0
	for i, row := range grid {
		for j := range row {
			tmp := area(i, j)
			if tmp > result {
				result = tmp
			}
		}
	}
	return result
}
