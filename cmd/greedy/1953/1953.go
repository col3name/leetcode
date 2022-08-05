package _1953

func numberOfWeeks(milestones []int) int64 {
	acc := int64(0)
	maxE := int64(0)

	for _, m := range milestones {
		acc += int64(m)
		if int64(m) > maxE {
			maxE = int64(m)
		}
	}

	var left int64
	left = acc - maxE
	if left >= maxE {
		return acc
	}
	return 2*left + 1
}
