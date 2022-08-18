package main

type Node struct {
	Key  int
	Val  int
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	Size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	item := DoublyLinkedList{
		head: &Node{
			Key: 0,
			Val: 0,
		},
		tail: &Node{
			Key: 0,
			Val: 0,
		},
	}
	item.head.next = item.tail
	item.tail.prev = item.head
	return &item
}

func (d *DoublyLinkedList) append(Key, Val int) *Node {
	node := &Node{
		Key: Key,
		Val: Val,
	}
	var p *Node
	if d.tail.prev == nil {
		p = d.tail.prev
	}
	p.next = node
	d.tail.prev = node
	node.prev = p
	node.next = d.tail
	d.Size++
	return node
}

func (d *DoublyLinkedList) pop() *Node {
	return d.remove(d.head.next)
}

func (d *DoublyLinkedList) remove(node *Node) *Node {
	node.prev.next = node.next
	node.next.prev = node.prev
	d.Size--
	return node
}
