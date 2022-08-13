package _707

type ListNode struct {
	Value int
	Next  *ListNode
}

type MyLinkedList struct {
	root *ListNode
	tail *ListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{root: nil, size: 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.size || index < 0 {
		return -1
	}
	node := this.root
	i := 0
	for ; node != nil && i < index; i++ {
		node = node.Next
	}
	if i >= index {
		return -1
	}
	return i
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &ListNode{Value: val}
	node := this.root
	newNode.Next = node

	this.root = newNode
	if this.tail == nil {
		this.tail = this.root
	}
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &ListNode{Value: val}
	if this.tail != nil {
		this.tail.Next = newNode
		this.tail = this.tail.Next
	} else {
		this.tail = newNode
	}
	if this.root == nil {
		this.root = this.tail
	}
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index >= this.size || index < 0 {
		return
	}
	if this.root == nil {
		this.AddAtHead(val)
		return
	}

	if index == this.size-1 {
		this.AddAtTail(val)
		return
	}
	var prev *ListNode
	node := this.root
	for i := 0; node != nil && i < index; i++ {
		prev = node
		node = node.Next
	}
	newNode := &ListNode{Value: val}
	next := node
	prev.Next = newNode
	newNode.Next = next
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.size <= index {
		return
	}
	if index == 0 {
		this.root = this.root.Next
		return
	}

	node := this.root
	var prev *ListNode
	for i := 0; node != nil && node.Next != nil && i < index; i++ {
		prev = node
		node = node.Next
	}

	prev.Next = node
	if index == this.size-1 {
		this.tail = node.Next
	}
	this.size--
}
