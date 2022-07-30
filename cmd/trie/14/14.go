package _14

import (
	"math"
	"strings"
)

type TrieNode struct {
	isWord bool
	childs map[byte]*TrieNode
}

func NewNode() *TrieNode {
	return &TrieNode{childs: make(map[byte]*TrieNode)}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: NewNode()}
}

func (t *Trie) insert(text string) {
	node := t.root
	for _, ch := range []byte(text) {
		_, ok := node.childs[ch]
		if !ok {
			node.childs[ch] = NewNode()
		}
		node = node.childs[ch]
	}
	node.isWord = true
}

func (t *Trie) findLongestCommonPrefix(word string) string {
	if len(word) == 0 {
		return ""
	}
	node := t.root
	var result strings.Builder
	for _, ch := range []byte(word) {
		next, ok := node.childs[ch]
		if ok && len(node.childs) == 1 && !node.isWord {
			result.WriteString(string(ch))
			node = next
		} else {
			return result.String()
		}
	}
	return result.String()
}

func longestCommonPrefix(stringList []string) string {
	trie := NewTrie()
	longestWord := ""

	for _, str := range stringList {
		if len(str) > len(longestWord) {
			longestWord = str
		}
		if len(str) == 0 {
			return ""
		}
		trie.insert(str)
	}

	return trie.findLongestCommonPrefix(longestWord)
}

func longestCommonPrefixVersion2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLength := math.MaxInt32
	for _, str := range strs {
		if minLength > len(str) {
			minLength = len(str)
		}
	}
	low := 1
	high := minLength
	middle := 0
	for low <= high {
		middle = (low + high) / 2
		if !isCommonPrefix(strs, middle) {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}

	return strs[0][0 : (low+high)/2]
}

func isCommonPrefix(strs []string, length int) bool {
	str := strs[0][:length]
	for i := 1; i < len(strs); i++ {
		if !strings.HasPrefix(strs[i], str) {
			return false
		}
	}
	return true
}
