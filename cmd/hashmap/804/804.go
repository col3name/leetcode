package _804

func uniqueMorseRepresentations(words []string) int {
	MORSE := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.",
		"....", "..", ".---", "-.-", ".-..", "--", "-.",
		"---", ".--.", "--.-", ".-.", "...", "-", "..-",
		"...-", ".--", "-..-", "-.--", "--.."}
	seen := map[string]bool{}
	for _, word := range words {
		code := ""
		for _, char := range []byte(word) {
			code += MORSE[char-'a']
		}
		seen[code] = true
	}
	return len(seen)
}
