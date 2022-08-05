package _11

func maxArea(height []int) int {
	result := 0
	var tmpArea int
	left := 0
	right := len(height) - 1
	for left < right {
		tmpArea = getRectangleArea(left, right, height[left], height[right])
		if tmpArea > result {
			result = tmpArea
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return result
}

func maxAreaVersionOne(height []int) int {
	result := 0
	var tmpArea int

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			tmpArea = getRectangleArea(i, j, height[i], height[j])
			if tmpArea > result {
				result = tmpArea
			}
		}
	}
	return result
}

func getRectangleArea(leftIndex, rightIndex, leftHeight, rightHeight int) int {
	return min(leftHeight, rightHeight) * (rightIndex - leftIndex)
}

func min(left, right int) int {
	if left < right {
		return left
	}
	return right
}
