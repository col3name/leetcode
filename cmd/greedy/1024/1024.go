package _1024

import "sort"

func videoStitching(clips [][]int, time int) int {
	sort.Slice(clips, func(i, j int) bool {
		return clips[i][0] < clips[j][0]
	})
	clip := clips[0]
	end := clip[1]
	result := 0

	for i := 1; i < len(clips); i++ {
		clip = clips[i]
		if clip[0] == end {
			result++
			end = clip[1]
		}
		if clip[1] == time {
			return result
		}
	}

	return -1
}
