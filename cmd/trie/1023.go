package trie

import "fmt"

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

func (this *Trie) GetWordsByPrefix(prefix string, maxCount int) []string {
	node := this.root

	for _, ch := range []byte(prefix) {
		value, ok := node.childs[ch]
		if !ok {
			return this.result
		}
		node = value
	}

	this.dfs(node, prefix, maxCount)

	result := this.result
	this.result = make([]string, 0)
	return result
}

func (this *Trie) dfs(current *TrieNode, word string, maxCount int) {
	if len(this.result) == maxCount {
		return
	}
	if current.isEnd() {
		this.result = append(this.result, word)
	}
	for ch := 'a'; ch <= 'z'; ch++ {
		node, ok := current.childs[byte(ch)]
		if ok {
			this.dfs(node, word+string(ch), maxCount)
		}
	}
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

func suggestedProducts(products []string, searchWord string) [][]string {
	result := make([][]string, 0)

	trie := Constructor()
	for _, product := range products {
		trie.Insert(product)
	}
	for i := 1; i <= len(searchWord); i++ {
		prefix := searchWord[0:i]
		fmt.Println(prefix)
		result = append(result, trie.GetWordsByPrefix(prefix, 3))
	}
	return result
}

func camelMatch(queries []string, pattern string) []bool {

}
