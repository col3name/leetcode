package _345

func reverseVowels(s string) string {
	b := []rune(s)
	vowelIndex := make([]int, 0)
	vowels := map[rune]interface{}{
		'A': nil,
		'a': nil,
		'E': nil,
		'e': nil,
		'I': nil,
		'i': nil,
		'O': nil,
		'o': nil,
		'U': nil,
		'u': nil,
	}
	for i, v := range s {
		if _, exist := vowels[v]; exist {
			vowelIndex = append(vowelIndex, i)
		}
	}
	for i := 0; i < len(vowelIndex)/2; i++ {
		b[vowelIndex[i]], b[vowelIndex[len(vowelIndex)-1-i]] = b[vowelIndex[len(vowelIndex)-1-i]], b[vowelIndex[i]]
	}
	return string(b)
}
