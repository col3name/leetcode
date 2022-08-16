package _387

func firstUniqChar(s string) int {
	h := [26]int{}
	for i := 0; i < 26; i++ {
		h[i] = -1
	}
	for _, char := range []byte(s) {
		h[char-byte('a')]++
	}
	for i, char := range []byte(s) {
		i2 := h[char-byte('a')]
		if i2 == -1 {
			continue
		}
		if i2 == 1 {
			return i
		}
	}
	return -1
}
