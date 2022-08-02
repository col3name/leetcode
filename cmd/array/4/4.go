package _4

import "math"

func maxProfit(prices []int) int {
	min := math.MaxInt32
	profit := 0
	for i, price := range prices {
		if price < min {
			min = price
		}
		if prices[i]-min > profit {
			profit = prices[i] - min
		}
	}

	return profit
}
