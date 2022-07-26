package _004

func longestOnes(nums []int, k int) int {
	left := 0
	maxLen := 0
	flipped := 0
	var num int
	for right := 0; right < len(nums); right++ {
		num = nums[right]
		if num == 0 {
			for flipped >= k && left < len(nums) {
				if nums[left] == 0 {
					flipped--
				}
				left++
			}
			flipped++
		}
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}

func max(lhs, rhs int) int {
	if rhs > lhs {
		return rhs
	}
	return lhs
}
