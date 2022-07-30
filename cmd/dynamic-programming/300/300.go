package _300

import "fmt"

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	max := 1
	for i := 1; i < len(dp); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				fmt.Println(dp)
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}
