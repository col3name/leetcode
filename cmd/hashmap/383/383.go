package _383

func canConstruct(ransomNote string, magazine string) bool {
	memory := make(map[byte]int, 0)
	for _, ch := range []byte(magazine) {
		memory[ch]++
	}
	for _, ch := range []byte(ransomNote) {
		if _, ok := memory[ch]; !ok {
			return false
		}
		memory[ch]--
		if memory[ch] <= 0 {
			delete(memory, ch)
		}
	}
	return true
}

func canConstruct2(ransomNote string, magazine string) bool {
	memory := make([]rune, 26)
	for _, ch := range []byte(magazine) {
		memory[ch-'a']++
	}
	for _, ch := range []byte(ransomNote) {
		memory[ch-'a']--
		if memory[ch-'a'] < 0 {
			return false
		}
	}
	return true
}
