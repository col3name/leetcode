package t7

import "strconv"

func reverse(x int) int {
	max := 2147483647
	min := -2147483648
	isMinus := false
	if x < 0 {
		isMinus = true
		x = -x
	}
	itoa := strconv.Itoa(x)

	itoa = reverseStr(itoa)
	val, _ := strconv.Atoi(itoa)
	if isMinus {
		val = -val
	}
	if val < min || val > max {
		return 0
	}
	return val
}

func reverseStr(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
