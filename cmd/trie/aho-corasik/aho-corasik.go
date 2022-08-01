package aho_corasik

type TrieNode struct {
	id           int
	childs       map[byte]*TrieNode
	isEndVal     bool
	suffixLink   *TrieNode
	symbol       byte
	parent       *TrieNode
	previousChar byte
	gos          []*TrieNode
	next         []*TrieNode
}

func NewNode(id int, parent *TrieNode, previousChar byte, _ int) *TrieNode {
	return &TrieNode{
		id:           id,
		childs:       make(map[byte]*TrieNode, 0),
		parent:       parent,
		previousChar: previousChar,
		suffixLink:   nil,
		gos:          make([]*TrieNode, 26),
		next:         make([]*TrieNode, 26),
	}
}

func (t *TrieNode) containsKey(ch byte) bool {
	_, ok := t.childs[ch]
	return ok
}

func (t *TrieNode) get(ch byte) *TrieNode {
	value, ok := t.childs[ch]
	if !ok {
		return nil
	}
	return value
}

func (t *TrieNode) put(ch byte, node *TrieNode) {
	t.childs[ch] = node
}

func (t *TrieNode) setKey(key bool) {
	t.isEndVal = key
}

func (t *TrieNode) isEnd() bool {
	return t.isEndVal
}

type Trie struct {
	root   *TrieNode
	result []string
	size   int
}

func Constructor() Trie {
	return Trie{
		root:   NewNode(0, nil, 0, 0),
		result: make([]string, 0),
		size:   0,
	}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, currentChar := range []byte(word) {
		if !node.containsKey(currentChar) {
			newNode := NewNode(0, node, currentChar, 0)
			node.put(currentChar, newNode)
			node.next[num(int(currentChar))] = newNode
			this.size++
		}
		node = node.get(currentChar)
	}
	node.setKey(true)
}

func (this *Trie) Size() int {
	return this.size
}
func (this *Trie) Search(word string) bool {
	node := this.searchPrefix(word)
	return node != nil && node.isEnd()
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.searchPrefix(prefix) != nil
}

func (this *Trie) searchPrefix(prefix string) *TrieNode {
	node := this.root
	for _, ch := range []byte(prefix) {
		value, ok := node.childs[ch]
		if !ok {
			return nil
		}
		node = value
	}
	return node
}

func (this *Trie) getLink(node *TrieNode) *TrieNode {
	if node.suffixLink == nil {
		if node == this.root || node.parent == this.root {
			node.suffixLink = this.root
		} else {
			node.suffixLink = this.gotoLink(this.getLink(node.parent), node.previousChar)
		}
	}

	return node.suffixLink
}

func (this *Trie) gotoLink(node *TrieNode, previousChar byte) *TrieNode {
	previousIndex := num(int(previousChar))
	if node.gos[previousIndex] == nil {
		if node.next[previousIndex] != nil {
			node.gos[previousIndex] = node.next[previousChar]
		} else if node == this.root {
			node.gos[previousIndex] = this.root
		} else {
			node.gos[previousIndex] = this.gotoLink(this.getLink(node), previousChar)
		}
	}

	return node.gos[previousIndex]
}

func num(i int) int {
	return i - 'a'
}
