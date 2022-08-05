package _45

func jump(nums []int) int {
	maxJumpPosition := 0
	count := 0
	currentJumpPosition := 0
	for i := 0; i < len(nums)-1; i++ {
		if maxJumpPosition < i+nums[i] {
			maxJumpPosition = i + nums[i]
		}
		if i == currentJumpPosition {
			count++
			currentJumpPosition = maxJumpPosition
		}
	}
	return count
}
