package _162

func findPeakElement(nums []int) int {
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := (low + high) / 2
		if nums[mid] > nums[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
