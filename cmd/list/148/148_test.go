package _148

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	node := &ListNode{}
	head := node
	for i := 2; i <= 50000; i++ {
		node.Next = &ListNode{
			Val:  i,
			Next: nil,
		}
		node = node.Next
	}
	node.Next = &ListNode{
		Val:  1,
		Next: nil,
	}
	head = head.Next
	node = sortList(head)
	i := 0
	prev := node.Val
	node = node.Next
	for ; node != nil; i++ {
		if prev > node.Val {
			assert.Fail(t, "faile")
		}
		prev = node.Val
		node = node.Next
	}
}
