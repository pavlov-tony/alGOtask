package trie

import (
	"strings"
	"sync"
)

type TrieNode struct {
	lock     sync.RWMutex
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

func (t *TrieNode) SearchDistance(word string, ch chan int) {
	currentRow := make([]int, len(word)+1)
	for k := range currentRow {
		currentRow[k] = k
	}
	var results []int

	for letter, node := range t.children {
		deepSearch(node, letter, word, currentRow, &results)
	}
	ch <- minIntElement(results)
}

func deepSearch(node *TrieNode, letter string, word string, previousRow []int, results *[]int) {
	currentRow := []int{previousRow[0] + 1}
	cols := len(word) + 1
	for i := 1; i < cols; i++ {
		repCost := 0
		if string(word[i-1]) != letter {
			repCost = previousRow[i-1] + 1
		} else {
			repCost = previousRow[i-1]
		}
		delCost := previousRow[i] + 1
		insCost := currentRow[i-1] + 1
		currentRow = append(currentRow, min(insCost, delCost, repCost))
	}
	if currentRow[len(currentRow)-1] <= len(word) && node.word != "" {
		*results = append(*results, currentRow[len(currentRow)-1])
	}

	if minIntElement(currentRow) <= len(word) {
		for l, n := range node.children {
			deepSearch(n, l, word, currentRow, results)
		}
	}
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

func minIntElement(s []int) int {
	answer := s[0]
	for _, v := range s {
		if v < answer {
			answer = v
		}
	}
	return answer
}
