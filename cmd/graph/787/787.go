package _787

import (
	"fmt"
	"math"
)

type Flight struct {
	Destination int
	Price       int
}

func findCheapestPrice(n int, flights [][]int, src int, destination int, k int) int {
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		prices[i] = math.MaxInt32
	}
	prices[src] = 0

	flightsMap := make(map[int][]*Flight, 0)
	var (
		ok     bool
		source int
	)
	for _, flight := range flights {
		source = flight[0]
		if _, ok = flightsMap[source]; !ok {
			flightsMap[source] = make([]*Flight, 0)
		}
		flightsMap[source] = append(flightsMap[source], &Flight{Destination: flight[1], Price: flight[2]})
	}

	queue := make([]int, 0)
	queue = append(queue, src)

	var (
		tempPrices    []int
		size          int
		currentSource int
		currentPrice  int
		i             int
		flight        *Flight
	)

	for k >= 0 && len(queue) > 0 {
		tempPrices = make([]int, n)
		for i = 0; i < n; i++ {
			tempPrices[i] = prices[i]
		}

		size = len(queue)
		for i = 0; i < size; i++ {
			currentSource = queue[0]
			queue = queue[1:]
			currentPrice = prices[currentSource]
			for _, flight = range flightsMap[currentSource] {
				if currentPrice+flight.Price < tempPrices[flight.Destination] {
					tempPrices[flight.Destination] = currentPrice + flight.Price

					queue = append(queue, flight.Destination)
				}
			}
		}

		for i = 0; i < n; i++ {
			prices[i] = tempPrices[i]
		}
		k--
	}
	fmt.Println(prices)
	totalPrice := prices[destination]
	if totalPrice == math.MaxInt32 {
		return -1
	}
	return totalPrice
}
