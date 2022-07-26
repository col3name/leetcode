package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func testHelper(t *testing.T, cache *LRUCache, tests []NodeItem) {
	node := cache.list.head
	for i := 0; i < cache.size; i++ {
		fmt.Println(node)
		if node.next == nil {
			break
		}
		expected := tests[i]
		assert.Equal(t, expected.Key, node.Key)
		assert.Equal(t, expected.Val, node.Val)
		node = node.next
	}
}

func TestLRUCacheTest(t *testing.T) {
	cache := Constructor(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4)
	tests := []NodeItem{
		{Key: 2, Val: 2, Size: 0},
		{Key: 3, Val: 3, Size: 0},
		{Key: 4, Val: 4, Size: 0},
	}
	testHelper(t, &cache, tests)

	fmt.Println("test")
}

func get(t *testing.T, cache *LRUCache, key, expect int) int {
	val := cache.Get(key)
	assert.Equal(t, expect, val)
	return val
}

func TestLRUCacheTestWithGet(t *testing.T) {
	result := make([]int, 0)
	cache := Constructor(2)
	cache.Put(1, 1)
	dumpList(cache.list)
	cache.Put(2, 2)
	dumpList(cache.list)
	result = append(result, get(t, &cache, 1, 1))
	cache.Put(3, 3)
	dumpList(cache.list)
	result = append(result, get(t, &cache, 2, -1))
	cache.Put(4, 4)
	dumpList(cache.list)

	result = append(result, get(t, &cache, 1, -1))
	result = append(result, get(t, &cache, 3, 3))
	result = append(result, get(t, &cache, 4, 4))
	//
	//tests := []NodeItem{
	//	{Key: 2, Val: 2, size: 0},
	//	{Key: 3, Val: 3, size: 0},
	//	{Key: 4, Val: 4, size: 0},
	//}

	//testHelper(t, cache, tests)
	dumpList(cache.list)
	fmt.Println("test")
	fmt.Println(result)
}

func dumpList(list *DoublyLinkedList) {
	node := list.head
	for node.next != nil {
		fmt.Print(node, ",")
		node = node.next
	}
	fmt.Println()
}
