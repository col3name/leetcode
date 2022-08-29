package _1143

func longestCommonSubsequence(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}
	lengthOne := len(text1)
	lengthTwo := len(text2)
	dp := make([]int, lengthTwo+1)

	for i := 1; i <= lengthOne; i++ {
		tmp := make([]int, lengthTwo+1)
		for j := 1; j <= lengthTwo; j++ {
			if text1[i-1] == text2[j-1] {
				tmp[j] = dp[j-1] + 1
			} else if dp[j] < tmp[j-1] {
				tmp[j] = tmp[j-1]
			} else {
				tmp[j] = dp[j]
			}
		}
		dp = tmp
	}
	return dp[lengthTwo]
}

func longestCommonSubsequenceVersion1(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}
	length1 := len(text1)
	length2 := len(text2)
	dp := make([][]int, length1+1)
	for i := range dp {
		dp[i] = make([]int, length2+1)
	}

	for i := 1; i <= length1; i++ {
		for j := 1; j <= length2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				continue
			}
			lhs := dp[i-1][j]
			rhs := dp[i][j-1]
			if lhs > rhs {
				dp[i][j] = lhs
			} else {
				dp[i][j] = rhs
			}
		}
	}
	return dp[length1][length2]
}
