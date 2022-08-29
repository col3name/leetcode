package _322

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	var minCoin int
	for i := 1; i <= amount; i++ {
		minCoin = math.MaxInt32
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 && dp[i-coin]+1 < minCoin {
				minCoin = dp[i-coin] + 1
			}
		}
		if minCoin == math.MaxInt32 {
			dp[i] = -1
		} else {
			dp[i] = minCoin
		}
	}

	return dp[amount]
}
