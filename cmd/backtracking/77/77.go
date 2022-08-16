package _77

func combine(n int, k int) [][]int {
	var result [][]int
	backtracking(n, k, &result, []int{}, 1)
	return result
}

func backtracking(n, k int, result *[][]int, permutation []int, index int) {
	if len(permutation) == k {
		tmp := make([]int, 0, k)
		for _, num := range permutation {
			tmp = append(tmp, num)
		}
		*result = append(*result, tmp)
		return
	}
	for i := index; i <= n; i++ {
		permutation = append(permutation, i)
		backtracking(n, k, result, permutation, i+1)
		permutation = permutation[:len(permutation)-1]
	}
}
