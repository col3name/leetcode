package _125

import "strings"

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if !isValid(s[left]) {
			left++
			continue
		}
		if !isValid(s[right]) {
			right--
			continue
		}
		if !strings.EqualFold(string(s[left]), string(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func isValid(a byte) bool {
	return (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9')
}
