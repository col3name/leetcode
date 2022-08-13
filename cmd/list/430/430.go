package _430

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	current := root
	tail := root
	stack := make([]*Node, 0, 1000)
	for current != nil {
		if current.Child != nil {
			child := current.Child
			if current.Next != nil {
				stack = append(stack, current.Next)
				current.Next.Prev = nil
			}
			current.Next = child
			child.Prev = current
			current.Child = nil
		}
		tail = current
		current = current.Next
	}

	for len(stack) > 0 {
		last := len(stack) - 1
		current = stack[last]
		stack = stack[:last]

		tail.Next = current
		current.Prev = tail
		for current != nil {
			tail = current
			current = current.Next
		}
	}
	return root
}
