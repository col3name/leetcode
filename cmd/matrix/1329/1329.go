package _1329

import "sort"

func diagonalSort(mat [][]int) [][]int {
	diagonalMap := make(map[int][]int)

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			diagonalMap[i-j] = append(diagonalMap[i-j], mat[i][j])
		}
	}

	for _, val := range diagonalMap {
		sort.Ints(val)
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			mat[i][j] = diagonalMap[i-j][0]
			diagonalMap[i-j] = diagonalMap[i-j][1:]
		}
	}
	return mat
}
