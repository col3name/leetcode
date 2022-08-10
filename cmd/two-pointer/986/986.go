package _986

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	result := make([][]int, 0)
	first := 0
	second := 0

	var (
		f []int
		s []int
	)

	for first < len(firstList) && second < len(secondList) {
		f = firstList[first]
		s = secondList[second]
		low := max(f[0], s[0])
		high := min(f[1], s[1])
		if low <= high {
			result = append(result, []int{low, high})
		}
		if s[1] < f[1] {
			second++
		} else {
			first++
		}
	}
	return result
}

func min(left, right int) int {
	if right < left {
		return right
	}
	return left
}

func max(left, right int) int {
	if right > left {
		return right
	}
	return left
}
