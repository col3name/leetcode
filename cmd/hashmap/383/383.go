package _383

func canConstruct(ransomNote string, magazine string) bool {
	memory := make(map[byte]int, 0)
	for _, ch := range []byte(magazine) {
		if _, ok := memory[ch]; ok {
			memory[ch]++
		} else {
			memory[ch] = 1
		}
	}
	for _, ch := range []byte(ransomNote) {
		if _, ok := memory[ch]; !ok {
			return false
		} else {
			memory[ch]--
			if memory[ch] <= 0 {
				delete(memory, ch)
			}
		}
	}
	return true
}
