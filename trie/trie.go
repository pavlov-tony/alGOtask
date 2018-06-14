package trie

// Node is an implementation of prefix tree.
type Node struct {
	word     []rune
	children map[rune]*Node
}

// Init returns pointer to the new Node.
func Init() *Node {
	return newNode()
}

func newNode() *Node {
	return &Node{children: make(map[rune]*Node)}
}

// Insert populates trie data structure with the new word.
func (t *Node) Insert(word []rune) {
	node := t
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			newNode := newNode()
			node.children[char] = newNode
			node = newNode
		} else {
			childNode := node.children[char]
			node = childNode
		}
	}
	node.word = word
}

// GetWord returns word from specific Node.
func (t *Node) GetWord() []rune {
	return t.word
}

// GetChildren returns map of children from specific Node.
func (t *Node) GetChildren() map[rune]*Node {
	return t.children
}
