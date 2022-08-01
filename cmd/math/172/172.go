package _172

import "fmt"

func trailingZeroes(n int) int {
	num := 0

	for n > 0 {
		n /= 5
		fmt.Println(num)
		num += n
	}

	return num
}
