package Tree

// TrieNode represents a node in Trie
type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

// Trie represents a Trie (prefix tree)
type Trie struct {
	root *TrieNode
}

// NewTrie creates a new Trie
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children:    make(map[rune]*TrieNode),
			isEndOfWord: false,
		},
	}
}

// Insert adds a word to the trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = &TrieNode{
				children:    make(map[rune]*TrieNode),
				isEndOfWord: false,
			}
		}
		node = node.children[ch]
	}
	node.isEndOfWord = true
}

// Search returns true if the word is in the trie
func (t *Trie) Search(word string) bool {
	node := t.searchNode(word)
	return node != nil && node.isEndOfWord
}

// StartsWith returns true if there is any word in the trie that starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	return t.searchNode(prefix) != nil
}

// searchNode returns the node at the end of the word/prefix path, or nil if not found
func (t *Trie) searchNode(word string) *TrieNode {
	node := t.root
	for _, ch := range word {
		if next, exists := node.children[ch]; exists {
			node = next
		} else {
			return nil
		}
	}
	return node
}

// Delete removes a word from the trie
func (t *Trie) Delete(word string) bool {
	return t.deleteHelper(t.root, word, 0)
}

func (t *Trie) deleteHelper(node *TrieNode, word string, depth int) bool {
	if node == nil {
		return false
	}

	// Base case: end of word reached
	if depth == len(word) {
		if !node.isEndOfWord {
			return false
		}
		node.isEndOfWord = false
		return len(node.children) == 0
	}

	ch := rune(word[depth])
	child, exists := node.children[ch]
	if !exists {
		return false
	}

	shouldDeleteChild := t.deleteHelper(child, word, depth+1)

	if shouldDeleteChild {
		delete(node.children, ch)
		return len(node.children) == 0 && !node.isEndOfWord
	}

	return false
}

// GetAllWords returns all words stored in the trie
func (t *Trie) GetAllWords() []string {
	var result []string
	t.getAllWordsHelper(t.root, "", &result)
	return result
}

func (t *Trie) getAllWordsHelper(node *TrieNode, prefix string, result *[]string) {
	if node.isEndOfWord {
		*result = append(*result, prefix)
	}

	for ch, child := range node.children {
		t.getAllWordsHelper(child, prefix+string(ch), result)
	}
}

// GetWordsWithPrefix returns all words that start with the given prefix
func (t *Trie) GetWordsWithPrefix(prefix string) []string {
	node := t.searchNode(prefix)
	if node == nil {
		return nil
	}

	var result []string
	t.getAllWordsHelper(node, prefix, &result)
	return result
}
