package _1386

import (
	"math"
)

func maxNumberOfFamilies(rowNumbers int, reservedSeats [][]int) int {
	maxRow := 0
	minRow := math.MaxInt32
	for _, pair := range reservedSeats {
		if pair[0] > maxRow {
			maxRow = pair[0]
		}
		if pair[0] < minRow {
			minRow = pair[0]
		}
	}
	countRow := maxRow - minRow + 1
	if maxRow == minRow {
		countRow = rowNumbers
	}

	mapper := map[int]int16{
		1:  0b1000000000,
		2:  0b0100000000,
		3:  0b0010000000,
		4:  0b0001000000,
		5:  0b0000100000,
		6:  0b0000010000,
		7:  0b0000001000,
		8:  0b0000000100,
		9:  0b0000000010,
		10: 0b0000000001,
	}
	seats := make([]int16, countRow)
	for i := 0; i < len(reservedSeats); i++ {
		seats[reservedSeats[i][0]-minRow] |= mapper[reservedSeats[i][1]]
	}
	result := 0

	masks := []int16{0b0111100000, 0b0001111000, 0b0000011110}

	var left, middle, right bool
	for _, row := range seats {
		left = row&masks[0] == 0
		middle = row&masks[1] == 0
		right = row&masks[2] == 0
		if left && right {
			result += 2
		} else if left || middle || right {
			result++
		}
	}

	result += (rowNumbers - countRow) * 2

	return result
}
