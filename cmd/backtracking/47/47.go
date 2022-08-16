package _47

func permuteUnique(nums []int) [][]int {
	var result [][]int
	backtracking(&result, nums, 0)
	return result
}

func backtracking(result *[][]int, nums []int, index int) {
	if index == len(nums) {
		if !existPermutation(result, nums) {
			tmp := make([]int, 0, len(nums))
			for _, num := range nums {
				tmp = append(tmp, num)
			}
			*result = append(*result, tmp)
		}
	}
	for i := index; i < len(nums); i++ {
		nums[i], nums[index] = nums[index], nums[i]
		backtracking(result, nums, index+1)
		nums[i], nums[index] = nums[index], nums[i]
	}
}

func existPermutation(result *[][]int, nums []int) bool {
	for _, permutation := range *result {
		if equalArray(permutation, nums) {
			return true
		}
	}

	return false
}

func equalArray(left, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}
