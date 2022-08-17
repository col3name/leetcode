package _90

import (
	"sort"
)

//90. Subsets II
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	helper(nums, []int{}, &result)
	return result
}

func helper(nums []int, next []int, result *[][]int) {
	*result = append(*result, next)
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			helper(nums[i+1:], append(append([]int{}, next...), nums[i]), result)
		}
	}
}

func subsetsWithDupVersionOne(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i <= len(nums); i++ {
		helperVersionOne(nums, &result, 0, []int{}, i)
	}
	return result
}

func helperVersionOne(nums []int, result *[][]int, start int, current []int, targetLength int) {
	if targetLength == len(current) && !existPermutation(result, current) {
		var tmp []int
		for _, num := range current {
			tmp = append(tmp, num)
		}
		*result = append(*result, tmp)

		return
	}
	for i := start; i < len(nums); i++ {
		helperVersionOne(nums, result, i+1, append(current, nums[i]), targetLength)
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
	sort.Ints(left)
	sort.Ints(right)
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
