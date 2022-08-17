package _57

import "sort"

//57. Insert Interval
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return false
	})
	var (
		merged = make([][]int, 0, len(intervals))
		start  = 0
		end    = 0
		last   = 0
	)
	for _, interval := range intervals {
		start = interval[0]
		end = interval[1]
		last = len(merged) - 1
		if len(merged) == 0 || merged[last][1] < start {
			merged = append(merged, interval)
		} else {
			if end > merged[last][1] {
				merged[last][1] = end
			}
		}
	}
	return merged
}

func insert1(intervals [][]int, newInterval []int) [][]int {
	return merge(append(intervals, newInterval))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	l, r := [][]int{}, [][]int{}
	start, end := newInterval[0], newInterval[1]
	for _, interval := range intervals {
		if interval[1] < start {
			l = append(l, interval)
		} else if interval[0] > end {
			r = append(r, interval)
		} else {
			if interval[0] < start {
				start = interval[0]
			}
			if interval[1] > end {
				end = interval[1]
			}
		}
	}
	return append(append(l, []int{start, end}), r...)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
