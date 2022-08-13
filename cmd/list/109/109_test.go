package _109

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	head := &ListNode{
		Val: -10,
		Next: &ListNode{
			Val: -3,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
		},
	}
	bst := sortedListToBST(head)
	fmt.Println(bst)
}
