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
