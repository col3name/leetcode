package _4

import "math"

func maxProfit(prices []int) int {
	min := math.MaxInt32
	index := 0
	for i, price := range prices {
		if price < min {
			min = price
			index = i
		}
	}

	profit := 0

	for i := index; i < len(prices); i++ {
		if prices[i]-min > profit {
			profit = prices[i] - min
		}
	}

	return profit
}
