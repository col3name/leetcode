package _1386

import (
	"fmt"
	"math"
)

//[
// [false true  true  false false false false true  false false]
// [false false false false false true  false false false false]
// [true  false false false false false false false false true]
//]
func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
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
		countRow = n
	}
	fmt.Println(n, countRow, maxRow, minRow)
	seats := make([][]bool, countRow)
	for i := 0; i < len(seats); i++ {
		seats[i] = make([]bool, 10)
	}
	var seat []int
	for i := 0; i < len(reservedSeats); i++ {
		seat = reservedSeats[i]
		seats[seat[0]-minRow][seat[1]-1] = true
	}
	result := 0

	for i, seat := range seats {
		// 2 3 4 5
		if !seat[1] && !seat[2] && !seat[3] && !seat[4] {
			result++
			seats[i][1] = true
			seats[i][2] = true
			seats[i][3] = true
			seats[i][4] = true
		}
		// 6 7 8 9
		if !seat[5] && !seat[6] && !seat[7] && !seat[8] {
			result++
			seats[i][5] = true
			seats[i][6] = true
			seats[i][7] = true
			seats[i][8] = true
		}
		// 4 5 6 7
		if !seat[3] && !seat[4] && !seat[5] && !seat[6] {
			result++
			seats[i][3] = true
			seats[i][4] = true
			seats[i][5] = true
			seats[i][6] = true
		}

	}
	result += (n - countRow) * 2
	return result
}
