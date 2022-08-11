package _242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	memory := make(map[byte]int, 0)
	for _, char := range []byte(s) {
		memory[char]++
	}
	for _, char := range []byte(t) {
		memory[char]--
		if memory[char] <= 0 {
			delete(memory, char)
		}
	}

	return len(memory) == 0
}
