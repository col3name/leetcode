package _96

func numTrees(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1

	for level := 2; level <= n; level++ {
		for root := 1; root <= level; root++ {
			dp[level] += dp[level-root] * dp[root-1]
		}
	}

	return dp[n]
}
