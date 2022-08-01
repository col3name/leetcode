package _172

func trailingZeroes(n int) int {
	num := 0

	for n > 0 {
		n /= 5
		num += n
	}

	return num
}
