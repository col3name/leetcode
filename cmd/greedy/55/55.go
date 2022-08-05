package _5

func canJump(nums []int) bool {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if i > max {
			return false
		}

		if max < i+nums[i] {
			max = i + nums[i]
		}
	}
	return true
}
