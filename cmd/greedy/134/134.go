package _134

func canCompleteCircuit(gas []int, cost []int) int {
	current := 0
	total := 0
	start := 0
	diff := 0
	for i := 0; i < len(gas); i++ {
		diff = gas[i] - cost[i]
		total += diff
		current += diff

		if current < 0 {
			start = i + 1
			current = 0
		}
	}
	if total < 0 {
		return -1
	}
	return start
}
