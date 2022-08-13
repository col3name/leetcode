package _2326

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = -1
		}
	}
	row, column := 0, 0
	direction := 1
	node := head
	for node != nil && result[row][column] == -1 {
		result[row][column] = node.Val
		node = node.Next
		switch direction {
		case 1:
			if column == n-1 || result[row][column+1] != -1 {
				direction = 2
				row++
			} else {
				column++
			}
		case 2:
			if row == m-1 || result[row+1][column] != -1 {
				direction = 3
				column--
			} else {
				row++
			}
		case 3:
			if column == 0 || result[row][column-1] != -1 {
				direction = 4
				row--
			} else {
				column--
			}
		case 4:
			if row == 0 || result[row-1][column] != -1 {
				direction = 1
				column++
			} else {
				row--
			}
		}
	}
	return result
}
