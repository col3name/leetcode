package aho_corasik

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	dictionary := []string{"he", "she", "his", "hers"}
	trie := Constructor()
	for _, word := range dictionary {
		trie.Insert(word)
	}
	fmt.Println(trie.Size())
	for _, word := range dictionary {
		fmt.Println(trie.Search(word))
	}

	fmt.Println(trie.Search("se"))
	fmt.Println(trie.Search("hi"))

	text := "hashfjksahflkjasfhlks"
	node := trie.root
	for i, character := range []byte(text) {
		node = trie.gotoLink(node, character)
		if node.isEndVal {
			fmt.Println(character, text[i-3:i+1])
		}
	}
}
