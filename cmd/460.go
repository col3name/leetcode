package cmd

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

type LFUCache struct {
	size int
	list *DoublyLinkedList
	//h          map[int]*Node
	//counterMap map[int]int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		size: capacity,
		list: NewDoublyLinkedList(),
		//h:    map[int]*Node,
	}
}

func (this *LFUCache) Get(key int) int {
	return 0
}

func (this *LFUCache) Put(key int, value int) {

}
