package s3

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	res := 0
	for l, r := 0, 0; r < len(s); r++ {
		if index, ok := m[s[r]]; ok {
			l = max(l, index)
		}
		res = max(res, r-l+1)
		m[s[r]] = r + 1
	}
	return res
}

//"abcabcbb"
func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}
