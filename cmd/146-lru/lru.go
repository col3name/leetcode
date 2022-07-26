package main

type LRUCache struct {
	size int
	list *DoublyLinkedList
	h    map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		size: capacity,
		list: NewDoublyLinkedList(),
		h:    make(map[int]*Node, 0),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.h[key]
	if ok {
		this.list.remove(node)
		this.h[key] = this.list.append(key, node.Val)
		return node.Val
	}

	return -1
}

func (this *LRUCache) Put(key, value int) {
	node, ok := this.h[key]
	if ok {
		this.list.remove(node)
	}
	this.h[key] = this.list.append(key, value)
	if this.list.Size > this.size {
		node = this.list.pop()
		delete(this.h, node.Key)
	}
}
