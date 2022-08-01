package _032

import "strings"

type TrieNode struct {
	next   []*TrieNode
	isWord bool
}

func NewNode() *TrieNode {
	return &TrieNode{next: make([]*TrieNode, 26)}
}

type StreamChecker struct {
	root *TrieNode
	sb   strings.Builder
}

func Constructor(words []string) StreamChecker {
	root := NewNode()
	var length int
	var node *TrieNode
	var character byte
	var i int
	for _, word := range words {
		node = root
		length = len(word)
		for i = length - 1; i >= 0; i-- {
			character = word[i] - 'a'
			if node.next[character] == nil {
				node.next[character] = NewNode()
			}
			node = node.next[character]
		}
		node.isWord = true
	}
	return StreamChecker{root: root}
}

func (this *StreamChecker) Query(letter byte) bool {
	this.sb.WriteByte(letter)
	node := this.root
	value := this.sb.String()
	for i := this.sb.Len() - 1; i >= 0 && node != nil; i-- {
		node = node.next[value[i]-'a']
		if node != nil && node.isWord {
			return true
		}
	}
	return false
}
