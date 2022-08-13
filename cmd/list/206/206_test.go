package _206

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
	result := reverseList(head)
	printList(head)
	printList(result)
}

func printList(head *ListNode) {
	node := head
	for node != nil {
		fmt.Print(node.Val, ",")
		node = node.Next
	}
	fmt.Println()
}

func reverseListRange(head, end *ListNode) *ListNode {
	var prev *ListNode
	node := head
	var next *ListNode
	for node != nil && node != end {
		// fmt.Println(node.Val)
		next = node.Next
		node.Next = prev
		prev = node
		node = next
	}

	return prev
}

func TestName2(t *testing.T) {
	end := &ListNode{
		Val:  5,
		Next: nil,
	}
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: end,
				},
			},
		},
	}
	list := reverseListRange(head, end)
	printList(list)
}
