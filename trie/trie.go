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

func (t *TrieNode) GetWord() string {
	return t.word
}

func (t *TrieNode) GetChildren() map[string]*TrieNode {
	return t.children
}
