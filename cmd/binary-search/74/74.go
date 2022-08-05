package _74

func searchMatrix(matrix [][]int, target int) bool {
	rowIndex := 0
	left := 0
	height := len(matrix)
	width := len(matrix[0])
	right := width - 1
	var row []int

	visitedRow := make([]bool, height)
	for rowIndex < height && rowIndex >= 0 && !visitedRow[rowIndex] {
		row = matrix[rowIndex]
		left = 0
		right = width - 1
		visitedRow[rowIndex] = true
		if target < row[0] {
			rowIndex--
			if rowIndex < 0 {
				return false
			}
			continue
		} else if target > row[width-1] {
			rowIndex++
			if rowIndex > height {
				return false
			}
			continue
		} else {
			for left <= right {
				middle := (left + right) / 2
				if row[middle] > target {
					right = middle - 1
				} else if row[middle] < target {
					left = middle + 1
				} else {
					return true
				}
			}
			return false
		}
	}

	return false
}
