package _56

// 56. Merge Intervals
import "sort"

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
