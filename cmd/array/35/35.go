package _35

func searchInsert(nums []int, target int) int {
	if len(nums) > 0 && nums[0] >= target {
		return 0
	}
	for i, num := range nums {
		if num >= target {
			return i
		}
	}
	return len(nums)
}
