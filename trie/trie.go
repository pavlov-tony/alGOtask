package trie

import (
	"strings"
)

type TrieNode struct {
	word     string
	children map[string]*TrieNode
}

func InitTrie() *TrieNode {
	return newNode()
}

func newNode() *TrieNode {
	return &TrieNode{children: make(map[string]*TrieNode)}
}

func (t *TrieNode) Insert(word string) {
	node := t
	for _, letter := range strings.Split(word, "") {
		if _, ok := node.children[letter]; !ok {
			newNode := newNode()
			node.children[letter] = newNode
			node = newNode
		} else {
			childNode := node.children[letter]
			node = childNode
		}
	}
	node.word = word
}

func (t *TrieNode) SearchDistance(word string) int {
	currentRow := make([]int, len(word)+1)
	for k := range currentRow {
		currentRow[k] = k
	}
	var results []int

	for letter, node := range t.children {
		deepSearch(node, letter, word, currentRow, &results)
	}
}

func deepSearch(node *TrieNode, letter string, word string, currentRow []int, results *[]int) {

}
