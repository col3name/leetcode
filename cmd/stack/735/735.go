package _735

func handleCollision(ast int, stack *[]int) {
	var getSize func() int
	getSize = func() int {
		return len(*stack) - 1
	}
	if ast > 0 {
		*stack = append(*stack, ast)
		return
	}
	for len(*stack) > 0 && (*stack)[getSize()] > 0 && -ast > (*stack)[getSize()] {
		*stack = (*stack)[:getSize()]
	}

	if len(*stack) <= 0 || (*stack)[getSize()] < 0 {
		*stack = append(*stack, ast)
		return
	}
	if (*stack)[getSize()] == -ast {
		*stack = (*stack)[:getSize()]
	}
}

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)

	for _, ast := range asteroids {
		handleCollision(ast, &stack)
	}

	return stack
}
