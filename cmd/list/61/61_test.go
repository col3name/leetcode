package _61

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	head = rotateRight(head, 3)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
