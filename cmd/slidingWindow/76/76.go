package _6

func minWindow(s string, t string) string {
	left, res, memo := 0, make([]int, 2), make(map[string]int, 0)
	for i := range t {
		memo[string(t[i])] = 1
	}
	curSize := 0
	for right, ch := range s {
		ch := string(ch)
		if _, exist := memo[ch]; exist {
			memo[ch]--
			if memo[ch] == 0 {
				curSize++
			}
		}
		for curSize == len(memo) {
			curLen := right - left
			prevLen := res[1] - res[0]
			if curLen < prevLen {
				res[0] = left
				res[1] = right
			}
			leftCh := string(s[left])
			if _, exist := memo[leftCh]; exist {
				if memo[leftCh] == 0 {
					curSize--
				}
				memo[leftCh]++
			}
			left++
		}
	}

	return s[res[0] : res[1]+1]
}
