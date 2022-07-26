package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	node := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	solution := Constructor()
	leaves := solution.FindLeaves(&node)
	fmt.Println(leaves)
}
