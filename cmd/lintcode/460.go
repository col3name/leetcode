package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solution struct {
	hash map[int][]int
}

func Constructor() *Solution {
	return &Solution{
		hash: make(map[int][]int, 0),
	}
}

func (s *Solution) FindLeaves(root *TreeNode) [][]int {
	s.iterate(root)
	var res [][]int
	for i := 0; len(s.hash[i]) > 0; i++ {
		res = append(res, s.hash[i])
	}
	return res
}

func (s *Solution) iterate(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := s.iterate(node.Left)
	right := s.iterate(node.Right)
	max := int(math.Max(float64(left), float64(right)))
	s.hash[max] = append(s.hash[max], node.Val)
	return max + 1
}
