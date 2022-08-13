package _707

import (
	"fmt"
	"testing"
)

func printList(head *ListNode) {
	node := head
	for node != nil {
		fmt.Print(node.Value, ",")
		node = node.Next
	}
	fmt.Println()
}

func TestName(t *testing.T) {
	list := Constructor()
	list.AddAtHead(1)
	printList(list.root)
	list.AddAtHead(2)
	printList(list.root)
	list.AddAtTail(4)
	printList(list.root)
	list.AddAtIndex(4, 10)
	printList(list.root)
	list.AddAtIndex(1, 10)
	list.AddAtIndex(0, -10)
	printList(list.root)
}

func TestName2(t *testing.T) {
	list := Constructor()
	list.AddAtIndex(10, 4)
	list.AddAtIndex(0, 4)
	printList(list.root)
}
