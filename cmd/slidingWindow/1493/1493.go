package _1493

func longestSubarray(nums []int) int {
	previous := 0
	count := 0
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			if count+previous > result {
				result = count + previous
			}
			previous = count
			count = 0
		}
	}
	if count+previous > result {
		result = count + previous
	}

	if result == len(nums) {
		result--
	}
	return result
}
