package _224

func calculate(s string) int {
	result := 0
	num := 0
	operation := add
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
		case '+':
			result = operation(result, num)
			num = 0
			operation = add
		case '-':
			result = operation(result, num)
			num = 0
			operation = sub
		case '(':
			end := findMatch(s, i)
			num = calculate(s[i+1 : end])
			i = end
		default:
			num *= 10
			num += int(s[i] - '0')
		}
	}

	result = operation(result, num)
	num = 0

	return result
}

func findMatch(s string, start int) int {
	level := 0
	for i := start; i < len(s); i++ {
		switch s[i] {
		case '(':
			level++
		case ')':
			level--
			if level == 0 {
				return i
			}
		}
	}
	return len(s) - 1
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
