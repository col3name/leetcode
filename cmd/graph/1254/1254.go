package _1254

func closedIsland(grid [][]int) int {
	rows := len(grid)
	columns := len(grid[0])
	seen := make([][]bool, rows)
	for i := range seen {
		seen[i] = make([]bool, columns)
	}

	var dfs func(int, int) bool
	dfs = func(row, col int) bool {
		if row == 0 || col == 0 || row == rows-1 || col == columns-1 {
			if grid[row][col] == 1 {
				return true
			}
			return false
		}
		if grid[row][col] == 1 {
			return true
		}
		grid[row][col] = 1
		val1 := dfs(row+1, col)
		val2 := dfs(row, col+1)
		val3 := dfs(row-1, col)
		val4 := dfs(row, col-1)

		return val1 && val2 && val3 && val4
	}

	result := 0
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == 0 && dfs(i, j) {
				result++
			}
		}
	}
	return result
}
