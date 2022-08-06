package _542

import "math"

type Cell struct {
	Row    int
	Column int
}

type Queue []*Cell

func (q *Queue) Push(value *Cell) {
	*q = append(*q, value)
}

func (q *Queue) Pop() *Cell {
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) Empty() bool {
	return q.Size() == 0
}

type Matrix [][]int

func updateMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	if rows == 0 {
		return matrix
	}
	cols := len(matrix[0])

	distance := make(Matrix, rows)
	for i := range distance {
		distance[i] = make([]int, cols)
		for j := range distance[i] {
			distance[i][j] = math.MaxInt32
		}
	}

	q := make(Queue, 0)
	for i := range distance {
		for j := range distance[i] {
			if matrix[i][j] == 0 {
				distance[i][j] = 0
				q.Push(&Cell{Row: i, Column: j})
			}
		}
	}

	directions := []Cell{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	size := &Cell{
		Row:    rows,
		Column: cols,
	}
	for !q.Empty() {
		current := q.Pop()

		handleDirection(current, &q, &directions, &distance, size)
	}

	return distance
}

func handleDirection(node *Cell, queue *Queue, directions *[]Cell, distance *Matrix, size *Cell) {
	for _, direction := range *directions {
		newRow := node.Row + direction.Row
		newColumn := node.Column + direction.Column

		if newRow < 0 || newColumn < 0 || newRow >= size.Row || newColumn >= size.Column {
			continue
		}

		newDistance := (*distance)[node.Row][node.Column] + 1
		if (*distance)[newRow][newColumn] > newDistance {
			(*distance)[newRow][newColumn] = newDistance
			queue.Push(&Cell{Row: newRow, Column: newColumn})
		}
	}
}
