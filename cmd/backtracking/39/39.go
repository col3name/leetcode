package _39

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	helper(&result, target, []int{}, candidates)
	return result
}

func helper(result *[][]int, target int, combine, candidates []int) {
	for i, candidate := range candidates {
		temp := target
		temp -= candidate
		if temp == 0 {
			tmp := make([]int, len(combine))
			copy(tmp, combine)
			*result = append(*result, append(tmp, candidate))
		} else if temp > 0 {
			helper(result, temp, append(combine, candidate), candidates[i:])
		}
	}
}
