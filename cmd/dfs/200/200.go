package _200

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfsMark(&grid, i, j)
				count++
			}
		}
	}
	return count
}

func dfsMark(grid *[][]byte, row, column int) {
	/*
	   1. Check necessary conditions
	*/
	if row < 0 || column < 0 || row > len(*grid)-1 || column > len((*grid)[0])-1 {
		return
	}
	if (*grid)[row][column] == '0' {
		return
	}
	// 2. Process Cell
	(*grid)[row][column] = '0'

	// 3. Call dfs as needed
	dfsMark(grid, row, column-1)
	dfsMark(grid, row, column+1)
	dfsMark(grid, row-1, column)
	dfsMark(grid, row+1, column)
}
