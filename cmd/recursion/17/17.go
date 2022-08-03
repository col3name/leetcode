package _17

func letterCombinations(digits string) []string {
	var ans []string
	if len(digits) == 0 {
		return ans
	}
	helper(digits, 0, "", &ans)

	return ans
}

var mapper = map[byte]string{
	'1': "",
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
	'0': " ",
}

func helper(digits string, pos int, s string, strs *[]string) {
	if pos == len(digits) {
		*strs = append(*strs, s)
		return
	}

	for _, c := range mapper[digits[pos]] {
		helper(digits, pos+1, s+string(c), strs)
	}
}
