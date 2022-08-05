package _186

func rotateVersion1(nums []int, k int) {
	i := int32(len(nums) - k%len(nums))
	copy(nums, append(nums[i:], nums[:i]...))
}

func rotateVersion2(nums []int, k int) {
	length := len(nums)
	k = k % length
	tmp := make([]int, length)

	for j := 0; j < length; j++ {
		if j < length-k {
			tmp[j+k] = nums[j]
		} else {
			tmp[j-length+k] = nums[j]
		}
	}

	copy(nums, tmp)
}
