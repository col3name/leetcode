package _438

func findAnagrams(s string, p string) []int {
	result := make([]int, 0)
	length := len(p)
	maps := map[byte]int{}

	for _, r := range []byte(p) {
		maps[r]++
	}
	for id, r := range []byte(s) {
		maps[r]--
		if maps[r] == 0 {
			delete(maps, r)
		}
		if id-length >= 0 {
			char := s[id-length]
			maps[char]++
			if maps[char] == 0 {
				delete(maps, char)
			}
		}
		if len(maps) == 0 {
			result = append(result, id-length+1)
		}
	}
	return result
}
