package trie

// Node is an implementation of prefix tree.
type Node struct {
	word     []byte
	children map[byte]*Node
}

// Init returns pointer to the new Node.
func Init() *Node {
	return newNode()
}

func newNode() *Node {
	return &Node{children: make(map[byte]*Node)}
}

// Insert populates trie data structure with the new word.
func (t *Node) Insert(word []byte) {
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
func (t *Node) GetWord() []byte {
	return t.word
}

// GetChildren returns map of children from specific Node.
func (t *Node) GetChildren() map[byte]*Node {
	return t.children
}
