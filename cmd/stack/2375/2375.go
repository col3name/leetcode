package _2375

import (
	"strconv"
)

func smallestNumber(pattern string) string {
	stack := make([]int, 0, len(pattern))
	result := ""
	for i := 0; i <= len(pattern); i++ {
		stack = append(stack, i+1)
		if i == len(pattern) || pattern[i] == 'I' {
			for len(stack) > 0 {
				result += strconv.Itoa(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		}
	}
	return result
}
