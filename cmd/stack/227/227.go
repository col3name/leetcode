package _227

import (
	"strconv"
	"unicode"
)

type Stack []int

func (s *Stack) Pop() int {
	last := (*s).Size() - 1
	val := (*s)[last]
	*s = (*s)[:last]
	return val
}

func (s *Stack) Push(value int) {
	*s = append(*s, value)
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Last() int {
	return (*s)[(*s).Size()-1]
}

func calculate(input string) int {
	if len(input) == 0 {
		return 0
	}
	stack := &Stack{}
	result := 0
	operation := "+"
	lastIndex := len(input) - 1
	for i, char := range input {
		isDigit := unicode.IsDigit(char)
		if isDigit {
			val, _ := strconv.Atoi(string(char))
			result = (result * 10) + val
		}
		if (!isDigit && char != ' ') || i == lastIndex {
			switch operation {
			case "+":
				stack.Push(result)
			case "-":
				stack.Push(-result)
			case "*":
				stack.Push(stack.Pop() * result)
			case "/":
				stack.Push(stack.Pop() / result)
			}
			operation = string(char)
			result = 0
		}
	}
	result = 0
	for stack.Size() > 0 {
		result += stack.Pop()
	}
	return result
}
