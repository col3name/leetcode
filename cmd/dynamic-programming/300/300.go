package _300

func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return 1
	}
	result := 0
	dp := make([]int, length)
	for i := range dp {
		dp[i] = 1
	}
	for i := length - 1; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			if nums[j] > nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}
