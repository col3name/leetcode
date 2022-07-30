package _50

//https://leetcode.com/problems/powx-n/
import "math"

func myPow(x float64, n int) float64 {
	if n == 0 || x == 1.0 {
		return 1.0
	}
	if x == -1.0 {
		if n%2 == 0 {
			return 1.0
		}
		return -1.0
	}

	var result float64
	result = x
	isLessThanZero := n < 0
	if isLessThanZero {
		n = -n
	}
	for i := 1; i < n; i++ {
		if result > math.MaxFloat64-result*x {
			return math.MaxFloat64
		} else if result < -math.MaxFloat64-1+result*x {
			return -math.MaxFloat64 - 1
		}
		if isLessThanZero && 1/result < 0.00000001 {
			return 0.0
		}
		result *= x
	}
	if isLessThanZero {
		return 1 / result
	}

	return result
}
