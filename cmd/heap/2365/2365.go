package _2365

func taskSchedulerII(tasks []int, space int) int64 {
	last := map[int]int64{}
	result := int64(0)
	spac := int64(space)
	for _, task := range tasks {
		if value, ok := last[task]; ok {
			if result < value+spac {
				result = value + spac
			}
			result++
			last[task] = result
		} else {
			result++
			last[task] = result
		}
	}

	return result
}
