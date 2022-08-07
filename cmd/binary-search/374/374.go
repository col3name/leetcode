package _374

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 *
 */
var guess func(num int) int

func guessNumber(n int) int {
	var (
		low  = 1
		mid1 int
		mid2 int
		res1 int
		res2 int
		high = n
	)

	for low <= high {
		mid1 = low + (high-low)/3
		mid2 = high - (high-low)/3
		res1 = guess(mid1)
		res2 = guess(mid2)
		if res1 == 0 {
			return mid1
		}
		if res2 == 0 {
			return mid2
		}
		if res1 < 0 {
			high = mid1 - 1
		} else if res2 > 0 {
			low = mid2 + 1
		} else {
			low = mid1 + 1
			high = mid2 - 1
		}
	}

	return -1
}
