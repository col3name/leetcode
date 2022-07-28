package _11

type TrieNode struct {
	childs map[byte]*TrieNode
	isWord bool
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

func NewNode() *TrieNode {
	return &TrieNode{childs: make(map[byte]*TrieNode, 0)}
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: NewNode()}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, currentChar := range []byte(word) {
		if !node.containsKey(currentChar) {
			node.put(currentChar, NewNode())
		}
		node = node.get(currentChar)
	}
	node.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.search(this.root, word, 0)
}

func (this *WordDictionary) search(node *TrieNode, word string, index int) bool {
	pointer := node
	for i := index; i < len(word); i++ {
		ch := word[i]
		if ch == '.' {
			for _, child := range pointer.childs {
				if this.search(child, word, i+1) {
					return true
				}
			}

			return false
		} else {
			nextNode, ok := pointer.childs[ch]
			if !ok {
				return false
			}
			pointer = nextNode
		}
	}
	return pointer.isWord
}
