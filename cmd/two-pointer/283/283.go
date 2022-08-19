package _283

// 283. Move Zeroes

func moveZeroes(nums []int) {
	notZero := 0
	for i, num := range nums {
		if num != 0 {
			nums[notZero], nums[i] = num, nums[notZero]
			notZero++
		}
	}
}
