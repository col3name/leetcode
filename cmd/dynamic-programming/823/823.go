package _823

import "sort"

const MODULE = 1_000_000_007

func numFactoredBinaryTrees(A []int) int {
	MOD := int64(1_000_000_007)
	N := len(A)
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	dp := make([]int64, N)
	for i := range dp {
		dp[i] = 1
	}
	index := make(map[int]int)
	for i := 0; i < N; i++ {
		index[A[i]] = i
	}
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			if A[i]%A[j] == 0 {
				right := A[i] / A[j]
				if value, ok := index[right]; ok {
					i2 := dp[i] + dp[j]*dp[value]
					dp[i] = i2 % MOD
				}
			}
		}
	}
	ans := int64(0)
	for _, x := range dp {
		ans += x
	}
	return (int)(ans % MOD)
}
