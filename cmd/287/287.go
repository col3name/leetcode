package _87

import "math"

func findDuplicate(nums []int) int {
	duplicate := -1
	for i := 0; i < len(nums); i++ {
		cur := int(math.Abs(float64(nums[i])))
		if nums[cur] < 1 {
			duplicate = cur
			break
		}
		nums[cur] *= -1
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = int(math.Abs(float64(i)))
	}

	return duplicate
}

func findDuplicate2(nums []int) int {
	duplicate := -1

	left := 0
	right := len(nums) - 1
	cur := 0
	count := 0
	for left <= right {
		cur = (left + right) / 2

		for _, num := range nums {
			if num <= cur {
				count++
			}
		}

		if count <= cur {
			left = cur + 1
		} else {
			right = cur - 1
			duplicate = cur
		}
		count = 0
	}
	return duplicate
}
