package _22

//https://leetcode.com/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	result := make([]string, 0)
	backtrack(n, n, &result, "")
	return result
}

func backtrack(left, right int, result *[]string, current string) {
	if left == 0 && right == 0 {
		*result = append(*result, current)
	}
	if right < left {
		return
	}
	if left > 0 {
		backtrack(left-1, right, result, current+"(")
	}
	if right > 0 {
		backtrack(left, right-1, result, current+")")
	}
}
