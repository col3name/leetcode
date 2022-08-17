package _78

func subsets(nums []int) [][]int {
	var result [][]int

	for k := 0; k <= len(nums); k++ {
		helper(nums, &result, 0, []int{}, k)
	}
	return result
}

func helper(nums []int, result *[][]int, index int, current []int, k int) {
	if len(current) == k {
		var tmp []int
		for _, num := range current {
			tmp = append(tmp, num)
		}
		*result = append(*result, tmp)
		return
	}
	for i := index; i < len(nums); i++ {
		helper(nums, result, i+1, append(current, nums[i]), k)
	}
}
