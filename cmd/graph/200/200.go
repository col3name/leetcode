package _200

type Point struct {
	i int
	j int
}

type Queue []*Point

func (q *Queue) Push(point *Point) {
	*q = append(*q, point)
}

func (q *Queue) Pop() *Point {
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

func numIslands(grid [][]byte) int {
	count := 0
	q := Queue{}
	var point *Point
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				q = Queue{}
				q.Push(&Point{i, j})
				for !q.Empty() {
					point = q.Pop()
					row := point.i
					column := point.j
					if row < 0 || column < 0 || row > len(grid)-1 || column > len((grid)[0])-1 {
						continue
					}
					if (grid)[row][column] == '0' {
						continue
					}
					grid[row][column] = '0'
					q.Push(&Point{row, column - 1})
					q.Push(&Point{row, column + 1})
					q.Push(&Point{row - 1, column})
					q.Push(&Point{row + 1, column})
				}
				//dfsMark(&grid, i, j)
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
