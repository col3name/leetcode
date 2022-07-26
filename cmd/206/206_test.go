package _06

import (
	"github.com/col3name/algo/pkg/util"
	"github.com/stretchr/testify/assert"
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

	head = reverseList(head)
	res := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		res = append(res, cur.Val)
	}

	assert.True(t, util.Equal(res, []int{5, 4, 3, 2, 1}))
}
