package search

import "github.com/pavlov-tony/alGOtask/trie"

type result struct {
	word     string
	distance int
}

// Distance returns the result of search in Trie.
func Distance(node *trie.Node, word string, ch chan int) {
	currentRow := make([]int, len(word)+1)
	for k := range currentRow {
		currentRow[k] = k
	}
	results := &result{distance: len(word)}
	for letter, node := range node.GetChildren() {
		deepSearch(node, letter, word, currentRow, results)
	}
	ch <- results.distance
}

func deepSearch(node *trie.Node, letter string, word string, previousRow []int, results *result) {
	// Preallocate slice with size of len(word) + 1.
	// Add 1 to the capacity to avoid reallocation.
	currentRow := make([]int, 0, len(word)+1)
	currentRow = append(currentRow, previousRow[0]+1)
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
	currentRowDistance := currentRow[len(currentRow)-1]
	maxChanges := results.distance
	if currentRowDistance <= maxChanges && node.GetWord() != "" {
		if currentRowDistance < results.distance {
			results.distance = currentRowDistance
		}
	}
	if minIntElement(currentRow) <= maxChanges {
		for l, n := range node.GetChildren() {
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
	if len(s) == 0 {
		return 0
	}
	answer := s[0]
	for _, v := range s {
		if v < answer {
			answer = v
		}
	}
	return answer
}
