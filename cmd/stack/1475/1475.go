package _1475

func finalPrices(prices []int) []int {
	stack := make([]int, 0, len(prices))

	for i, v := range prices {
		for len(stack) != 0 && prices[stack[len(stack)-1]] >= v {
			prices[stack[len(stack)-1]] -= v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return prices
}
