package _278

func isBadVersion(version int) bool {
	return false
}

func firstBadVersion(n int) int {
	left := 1
	right := n
	var middle int
	for left+1 < right {
		middle = (left + right) / 2
		if isBadVersion(middle) {
			right = middle
		} else {
			left = middle
		}
	}
	if isBadVersion(left) {
		return left
	}
	return right
}
