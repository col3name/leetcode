package _869

// 869. Reordered Power of 2
// https://leetcode.com/problems/reordered-power-of-2
func reorderedPowerOf2(n int) bool {
	result := count(n)
	for i := 0; i < 31; i++ {
		if equals(result, count(1<<i)) {
			return true
		}
	}
	return false
}

func count(n int) []int {
	result := make([]int, 10)
	for n > 0 {
		result[n%10]++
		n /= 10
	}
	return result
}

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, num := range a {
		if num != b[i] {
			return false
		}
	}
	return true
}
