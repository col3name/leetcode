package _704

import "fmt"

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	var middle int
	for start <= end {
		middle = start + (end-start)/2
		fmt.Println(middle)
		if nums[middle] == target {
			return middle
		}
		if nums[middle] < target {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}

	return -1
}
