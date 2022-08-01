package _735

type Stack []int

func (s *Stack) Pop() {
	*s = (*s)[:(*s).Size()-1]
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

func handleCollision(ast int, stack *Stack) {
	if ast > 0 {
		(*stack).Push(ast)
		return
	}
	for (*stack).Size() > 0 && (*stack).Last() > 0 && -ast > (*stack).Last() {
		(*stack).Pop()
	}
	if (*stack).Size() <= 0 || (*stack).Last() < 0 {
		(*stack).Push(ast)
		return
	}
	if (*stack).Last() == -ast {
		(*stack).Pop()
	}
}

func asteroidCollision(asteroids []int) []int {
	stack := &Stack{}

	for _, ast := range asteroids {
		handleCollision(ast, stack)
	}

	return *stack
}
