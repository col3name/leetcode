package _377

func combinationSum4(nums []int, target int) int {
	memory := make(map[int]int, 0)
	var backtrack func(int) int

	backtrack = func(total int) int {
		if total > target {
			return 0
		}
		if total == target {
			return 1
		}
		if value, ok := memory[total]; ok {
			return value
		}
		result := 0
		for _, num := range nums {
			result += backtrack(total + num)
		}
		memory[total] = result

		return result
	}

	return backtrack(0)
}
