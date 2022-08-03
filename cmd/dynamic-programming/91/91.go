package _91

import (
	"strconv"
)

func numDecodings(s string) int {
	memory := make(map[string]int, 0)

	pairs := make(map[byte]string, 26)
	j := 1
	for i := 'A'; i < 'Z'; i++ {
		pairs[byte(i)] = strconv.Itoa(j)
		j++
	}

	var backtrack func(string) int
	backtrack = func(current string) int {
		if value, ok := memory[current]; ok {
			return value
		}
		if len(s) != 0 && s[0] == '0' {
			return 0
		}
		if current == "" || len(current) == 1 {
			return 1
		}

		atoi, err := strconv.Atoi(s[:2])
		if err == nil && 10 <= atoi && atoi <= 26 {
			memory[s] = backtrack(s[1:]) + backtrack(s[2:])
		} else {
			memory[s] = backtrack(s[1:])
		}
		return memory[current]
	}
	return backtrack(s)
}
