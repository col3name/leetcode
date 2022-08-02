package _1

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)

	for index, num := range nums {
		if requiredIndex, exist := indexMap[target-num]; exist {
			return []int{requiredIndex, index}
		}
		indexMap[num] = index
	}

	return []int{}
}
