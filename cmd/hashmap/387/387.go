package _387

func firstUniqChar(s string) int {
	h := map[byte]int{}
	for _, char := range []byte(s) {
		h[char]++
	}
	for i, char := range []byte(s) {
		if h[char] == 1 {
			return i
		}
	}
	return -1
}
