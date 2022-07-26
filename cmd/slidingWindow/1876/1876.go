package _876

func countGoodSubstrings1(s string) int {
	if len(s) < 3 {
		return 0
	}
	if len(s) == 3 && s[0] != s[1] && s[0] != s[2] && s[1] != s[2] {
		return 1
	}
	res := 0
	var substr string
	for right := 3; right <= len(s); right++ {
		substr = s[right-3 : right]
		if substr[0] != substr[1] && substr[0] != substr[2] &&
			substr[1] != substr[2] {
			res++
		}
	}
	return res
}

//"aababcabc"
func countGoodSubstrings(s string) int {
	memory := make(map[string]int, 0)
	left := 0
	res := 0
	for right, curChar := range s {
		memory[string(curChar)]++
		if right-left == 3 {
			leftCh := s[left]
			memory[string(leftCh)]--
			if memory[string(leftCh)] == 0 {
				delete(memory, string(leftCh))
			}
			left++
		}
		if len(memory) == 3 {
			res++
		}
	}
	return res
}
