package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type NodeItem struct {
	Key  int
	Val  int
	Size int
}

func TestDoubleLinkedList_Append(t *testing.T) {
	list := NewDoublyLinkedList()
	list.append(0, 0)
	list.append(1, 1)
	list.append(2, 2)
	size := list.Size
	assert.Equal(t, 3, size)

	tests := []NodeItem{
		//{Key: 0, Val: 0, size: 2},
		{Key: 1, Val: 1, Size: 1},
		{Key: 2, Val: 2, Size: 0},
		{Key: 2, Val: 2, Size: 0},
	}
	for i := 0; i < size; i++ {
		node := list.pop()

		if node != nil {
			fmt.Println(node.Val, node.Key, list.Size)
			expected := tests[i]
			assert.Equal(t, expected.Key, node.Key)
			assert.Equal(t, expected.Val, node.Val)
			//assert.Equal(t, expected.size, list.size)
		} else {
			fmt.Println("empty")
			//assert.Fail(t, "test failed")
		}

	}
	fmt.Println(list.Size)
}
