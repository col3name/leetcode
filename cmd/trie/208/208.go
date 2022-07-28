package _08

type TrieNode struct {
	childs   map[byte]*TrieNode
	isEndVal bool
}

func NewNode() *TrieNode {
	return &TrieNode{childs: make(map[byte]*TrieNode, 0)}
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
}

func Constructor() Trie {
	return Trie{root: NewNode(), result: make([]string, 0)}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, currentChar := range []byte(word) {
		if !node.containsKey(currentChar) {
			node.put(currentChar, NewNode())
		}
		node = node.get(currentChar)
	}
	node.setKey(true)
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
