package _46

type Node struct {
	prev *Node
	next *Node
	key  int
	val  int
}

func NewNode(key, val int) *Node {
	return &Node{
		prev: nil,
		next: nil,
		key:  key,
		val:  val,
	}
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	list := DoublyLinkedList{
		head: &Node{
			key: 0,
			val: 0,
		},
		tail: &Node{
			key: 0,
			val: 0,
		},
		size: 0,
	}
	list.head.next = list.tail
	list.tail.prev = list.head
	return &list
}

func (s *DoublyLinkedList) append(key, val int) *Node {
	node := &Node{
		key: key,
		val: val,
	}
	p := s.tail.prev
	p.next = node
	s.tail.prev = node
	node.prev = p
	node.next = s.tail
	s.size++
	return node
}

func (s *DoublyLinkedList) pop() *Node {
	return s.remove(s.head.next)
}

func (s *DoublyLinkedList) remove(node *Node) *Node {
	node.prev.next = node.next
	node.next.prev = node.prev
	s.size--
	return node
}

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
}``

func (this *LRUCache) Get(key int) int {
	node, ok := this.h[key]
	if ok {
		this.list.remove(node)
		this.h[key] = this.list.append(key, node.val)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.h[key]
	if ok {
		this.list.remove(this.h[key])
	}

	node := this.list.append(key, value)
	this.h[key] = node
	if this.list.size > this.size {
		node = this.list.pop()
		delete(this.h, node.key)
	}
}
