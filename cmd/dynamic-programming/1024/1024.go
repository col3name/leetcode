package _1024

import (
	"fmt"
	"sort"
)

// 1024. Video Stitching
func videoStitching(clips [][]int, time int) int {
	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] == clips[j][0] {
			return clips[i][1] < clips[j][1]
		}
		return clips[i][0] < clips[j][0]
	})
	fmt.Println(clips)
	dp := make([]int, time+1)
	for i := range dp {
		dp[i] = time + 1
	}

	dp[0] = 0
	// DP[pos, limit], pos - position of the current interval, limit - current covered area
	for currentTime := 0; currentTime <= time; currentTime++ {
		updateCurrentCoveredArea(currentTime, dp, clips)
		if dp[currentTime] == time+1 {
			return -1
		}
	}

	return dp[time]
}

func updateCurrentCoveredArea(currentTime int, dp []int, clips [][]int) {
	for _, clip := range clips {
		start := clip[0]
		end := clip[1]
		// current time inside clip time range
		if currentTime >= start && end >= currentTime {
			dp[currentTime] = min(dp[currentTime], dp[start]+1)
		}
	}
}

func min(lhs, rhs int) int {
	if lhs < rhs {
		return lhs
	}
	return rhs
}
