package trie

import (
	"strings"
)

// Node is an implementation of prefix tree.
type Node struct {
	word     string
	children map[string]*Node
}

// Init returns pointer to the new Node.
func Init() *Node {
	return newNode()
}

func newNode() *Node {
	return &Node{children: make(map[string]*Node)}
}

// Insert populates trie data structure with the new word.
func (t *Node) Insert(word string) {
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

// GetWord returns word from specific Node.
func (t *Node) GetWord() string {
	return t.word
}

// GetChildren returns map of children from specific Node.
func (t *Node) GetChildren() map[string]*Node {
	return t.children
}
