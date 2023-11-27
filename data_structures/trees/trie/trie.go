package trie

// TrieNode represents a node in a trie data structure.
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// NewTrieNode creates and returns a new TrieNode.
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

// Insert inserts a word into the trie.
func (t *TrieNode) Insert(word string) {
	current := t
	for _, c := range word {
		if _, ok := current.children[c]; !ok {
			current.children[c] = NewTrieNode()
		}
		current = current.children[c]
	}
	current.isEnd = true
}

// Search searches for a word in the trie.
// Returns true if the word is found, false otherwise.
func (t *TrieNode) Search(word string) bool {
	current := t
	for _, c := range word {
		if _, ok := current.children[c]; !ok {
			return false
		}
		current = current.children[c]
	}
	return current.isEnd
}

// Delete deletes a word from the trie.
func (t *TrieNode) Delete(word string) {
	current := t
	for _, c := range word {
		if _, ok := current.children[c]; !ok {
			return
		}
		current = current.children[c]
	}
	current.isEnd = false
}

// StartsWith checks if there is any word in the trie that starts with the given prefix.
// Returns true if there is at least one word with the given prefix, false otherwise.
func (t *TrieNode) StartsWith(prefix string) bool {
	current := t
	for _, c := range prefix {
		if _, ok := current.children[c]; !ok {
			return false
		}
		current = current.children[c]
	}
	return true
}
