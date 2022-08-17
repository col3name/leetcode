package _31

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	i++
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
