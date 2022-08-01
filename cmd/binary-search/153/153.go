package _153

func findMin(nums []int) int {
	end := len(nums) - 1
	if nums[end] > nums[0] {
		return nums[0]
	}
	var (
		start  int
		middle int
	)
	for start < end {
		middle = start + (end-start)/2

		if nums[middle] < nums[end] {
			end = middle
		} else {
			start = middle + 1
		}
	}
	return nums[start]
}
