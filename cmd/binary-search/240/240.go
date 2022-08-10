package _240

func searchMatrix(matrix [][]int, target int) bool {
	rowIndex := 0
	left := 0
	height := len(matrix)
	width := len(matrix[0])
	right := width - 1
	var row []int

	for rowIndex < height {
		row = matrix[rowIndex]
		left = 0
		right = width - 1
		for left <= right {
			middle := (left + right) / 2
			if row[middle] == target {
				return true
			}
			if row[middle] > target {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
		rowIndex++
	}

	return false
}
