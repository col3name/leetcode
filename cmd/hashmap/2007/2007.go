package _2007

import "sort"

// 2007. Find Original Array From Doubled Array
func findOriginalArraySuboptimal(changed []int) []int {
	sort.Ints(changed)
	result := make([]int, 0, 100_001)

	memo := make(map[int]int, 100_001)
	for _, num := range changed {
		if memo[num] == 0 {
			memo[num*2]++
			result = append(result, num)
		} else {
			memo[num]--
		}
	}
	for _, v := range memo {
		if v != 0 {
			return []int{}
		}
	}

	return result
}

func findOriginalArray(changed []int) []int {
	memo := [100_001]int{}
	for _, num := range changed {
		memo[num]++
	}
	var result []int

	if memo[0]%2 == 1 {
		return result
	}
	for i := 0; i < memo[0]/2; i++ {
		result = append(result, 0)
	}
	var j int
	for i := 1; i < len(memo); i++ {
		if memo[i] == 0 {
			continue
		}
		if memo[i] > 50_000 || memo[i] > memo[i*2] {
			return []int{}
		}
		memo[i*2] -= memo[i]
		for j = 0; j < memo[i]; j++ {
			result = append(result, i)
		}
	}

	return result
}
