package _37

import "sort"

func singleNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	i := 0
	for i < len(nums)-1 {
		if nums[i] != nums[i+1] && nums[i] != nums[i+2] {
			break
		}

		i += 3
	}
	return nums[i]
}
