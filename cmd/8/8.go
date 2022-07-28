package t8

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	pairs := map[string]int64{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	s = strip(s)
	s = strings.TrimSpace(s)
	s = strings.TrimLeft(s, "0")

	if len(s) == 0 {
		return 0
	}
	var result int64

	isNegative := false
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			isNegative = true
		}
		s = s[1:]
	}
	s = strings.TrimLeft(s, "0")
	if len(s) > len("2147483647") {
		if isNegative {
			return math.MinInt32
		}
		return math.MaxInt32
	}
	shift := len(s) - 1
	for _, ch := range s {
		value, ok := pairs[string(ch)]
		if !ok {
			break
		}
		for i := 0; i < shift; i++ {
			value *= 10
		}

		if result > math.MaxInt32-value {
			if isNegative {
				return math.MinInt32
			}
			return math.MaxInt32
		}

		result += value
		shift--
	}

	if isNegative {
		result = -result
	}
	return int(result)
}

func strip(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == '.' || ch == ' ' {
			break
		}

		if i == 0 && ('a' <= ch && ch <= 'z') ||
			('A' <= ch && ch <= 'Z') {
			return ""
		}
		if ('0' <= ch && ch <= '9') ||
			(ch == '-' && i == 0) || (ch == '+' && i == 0) ||
			ch == ' ' {
			result.WriteByte(ch)
		} else {
			break
		}
	}
	return result.String()
}
