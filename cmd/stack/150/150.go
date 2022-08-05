package _150

import (
	"strconv"
)

type Stack []int

func (s *Stack) push(value int) {
	*s = append(*s, value)
}

func (s *Stack) pop() int {
	n := len(*s) - 1
	result := (*s)[n]
	*s = (*s)[:n]
	return result
}

func evalRPN(tokens []string) int {
	stack := Stack{}
	var (
		left  int
		right int
	)
	for _, token := range tokens {
		if isOperation(token) {
			right = stack.pop()
			left = stack.pop()
			stack.push(evaluate(left, right, token))
		} else {
			value, err := strconv.Atoi(token)
			if err != nil {
				return 0
			}
			stack.push(value)
		}
	}
	return stack.pop()
}

func evaluate(left, right int, operand string) int {
	switch operand {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	}
	return 0
}

func isOperation(token string) bool {
	return token == "-" || token == "+" || token == "*" || token == "/"
}
