package _560

func subarraySum(nums []int, k int) int {
	sum := 0
	count := 0
	for i := range nums {
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
		sum = 0
	}
	return count
}
