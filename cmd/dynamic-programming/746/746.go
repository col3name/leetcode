package _746

func minCostClimbingStairs(cost []int) int {
	a, b := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		if b < a {
			a, b = b, cost[i]+b
		} else {
			a, b = b, cost[i]+a
		}
	}

	if a < b {
		return a
	}
	return b
}
