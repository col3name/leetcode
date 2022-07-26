package _20

import (
	"math"
)

//func minimumTotal(triangle [][]int) int {
//	min := 0
//	i := 0
//	rhs := math.MaxInt64
//	for row := range triangle {
//		lhs := triangle[row][i]
//		if i + 1 < len(triangle[row]) {
//			rhs = triangle[row][i+1]
//		} else {
//			rhs = math.MaxInt64
//		}
//		if lhs > rhs {
//			min += rhs
//			i++
//		} else {
//			min += lhs
//		}
//	}
//	return min
//}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	memo := make([][]int, 2)
	var row []int
	for i := 0; i < 2; i++ {
		row = make([]int, n)
		for j := 0; j < n; j++ {
			row[j] = math.MaxInt64
		}
		memo[i] = row
	}
	memo[0][0] = triangle[0][0]

	for i := 1; i < n; i++ {
		for j := range triangle[i] {
			lhs := memo[(i-1)%2][j]
			rhs := memo[(i-1)%2][max(j-1, 0)]
			memo[i%2][j] = min(lhs, rhs) + triangle[i][j]
		}
	}
	val2 := memo[(n-1)%2]
	print(val2)
	//val := minVal(val2)
	//print(val)

	minVal := math.MaxInt32
	for _, item := range val2 {
		if item < minVal {
			minVal = item
		}
	}
	return minVal
}

func max(lhs, rhs int) int {
	if lhs > rhs {
		return lhs
	}
	return rhs
}
func min(lhs, rhs int) int {
	if lhs < rhs {
		return lhs
	}
	return rhs
}
