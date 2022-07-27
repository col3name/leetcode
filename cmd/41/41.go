package _1

func firstMissingPositive(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] >= 1 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			j := nums[i] - 1
			nums[i], nums[j] = nums[j], nums[i]
			//swap(nums, i, j)
		} else {
			i++
		}
	}
	missingNum := 0
	for missingNum < len(nums) && nums[missingNum] == missingNum+1 {
		missingNum++
	}
	return missingNum + 1
}
