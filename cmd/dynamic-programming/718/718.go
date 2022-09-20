package _718

func findLength(nums1 []int, nums2 []int) int {
	result := 0

	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	for i := len(nums1) - 1; i >= 0; i-- {
		for j := len(nums2) - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				if result < dp[i][j] {
					result = dp[i][j]
				}
			}
		}
	}

	return result
}
