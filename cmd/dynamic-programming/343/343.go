package _343

func integerBreak(n int) int {
	cases := []int{0, 0, 1, 2, 4, 6, 9}
	if n < 7 {
		return cases[n]
	}
	dp := append(cases, make([]int, n-6)...)
	for i := 7; i <= n; i++ {
		dp[i] = 3 * dp[i-3]
	}
	return dp[n]
}
