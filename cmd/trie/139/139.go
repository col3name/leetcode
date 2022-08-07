package _139

type TrieNode struct {
	isWord bool
	Childs map[byte]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{Childs: make(map[byte]*TrieNode, 0)}
}

type Trie struct {
	root *TrieNode
	memo map[int]bool
}

func (t *Trie) Insert(text string) {
	node := t.root
	for _, char := range []byte(text) {
		_, ok := node.Childs[char]
		if !ok {
			node.Childs[char] = NewTrieNode()
		}
		node = node.Childs[char]
	}
	node.isWord = true
}

func (t *Trie) workBreakDetail(text string, startIndex int) bool {
	if startIndex == len(text) {
		return true
	}

	if value, ok := t.memo[startIndex]; ok {
		return value
	}

	resultFound := false
	node := t.root
	var ok bool
	for i := startIndex; i < len(text); i++ {
		node, ok = node.Childs[text[i]]
		if !ok {
			break
		}
		if node.isWord {
			resultFound = t.workBreakDetail(text, i+1)
		}
		if resultFound {
			break
		}
	}

	t.memo[startIndex] = resultFound

	return t.memo[startIndex]
}

func NewTrie() Trie {
	return Trie{
		root: NewTrieNode(),
		memo: make(map[int]bool),
	}
}

func wordBreak(text string, wordDict []string) bool {
	trie := NewTrie()
	for _, word := range wordDict {
		trie.Insert(word)
	}

	trie.workBreakDetail(text, 0)
	return trie.memo[0]
}
