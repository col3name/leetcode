package _435

import (
	"math"
	"sort"
)

// 435. Non-overlapping Intervals
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	result := 0
	currentEnd := math.MinInt32

	for _, interval := range intervals {
		if interval[0] >= currentEnd {
			currentEnd = interval[1]
		} else {
			result++
		}
	}

	return result
}
