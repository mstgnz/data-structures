package tree

import (
	"sync"
)

// ThreadSafeTrieNode represents a node in the Trie
type ThreadSafeTrieNode struct {
	children map[rune]*ThreadSafeTrieNode
	isEnd    bool
}

// ThreadSafeTrie represents a thread-safe Trie
type ThreadSafeTrie struct {
	root  *ThreadSafeTrieNode
	size  int
	mutex sync.RWMutex
}

// NewThreadSafeTrie creates a new thread-safe Trie
func NewThreadSafeTrie() *ThreadSafeTrie {
	return &ThreadSafeTrie{
		root: &ThreadSafeTrieNode{
			children: make(map[rune]*ThreadSafeTrieNode),
			isEnd:    false,
		},
		size: 0,
	}
}

// Insert adds a new word to the Trie
func (t *ThreadSafeTrie) Insert(word string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	current := t.root
	for _, ch := range word {
		if _, exists := current.children[ch]; !exists {
			current.children[ch] = &ThreadSafeTrieNode{
				children: make(map[rune]*ThreadSafeTrieNode),
				isEnd:    false,
			}
		}
		current = current.children[ch]
	}

	if !current.isEnd {
		current.isEnd = true
		t.size++
	}
}

// Contains checks if a word exists in the Trie
func (t *ThreadSafeTrie) Contains(word string) bool {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	current := t.root
	for _, ch := range word {
		if _, exists := current.children[ch]; !exists {
			return false
		}
		current = current.children[ch]
	}
	return current.isEnd
}

// StartsWith checks if any word in the Trie starts with the given prefix
func (t *ThreadSafeTrie) StartsWith(prefix string) bool {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	current := t.root
	for _, ch := range prefix {
		if _, exists := current.children[ch]; !exists {
			return false
		}
		current = current.children[ch]
	}
	return true
}

// Remove removes a word from the Trie
func (t *ThreadSafeTrie) Remove(word string) bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// Check if word exists without using Contains method
	current := t.root
	for _, ch := range word {
		if next, exists := current.children[ch]; !exists {
			return false
		} else {
			current = next
		}
	}
	if !current.isEnd {
		return false
	}

	// Word exists, proceed with removal
	t.remove(t.root, word, 0)
	return true
}

func (t *ThreadSafeTrie) remove(node *ThreadSafeTrieNode, word string, depth int) bool {
	if node == nil {
		return false
	}

	// If we've reached the end of the word
	if depth == len(word) {
		if !node.isEnd {
			return false
		}
		node.isEnd = false
		t.size--
		return len(node.children) == 0
	}

	ch := rune(word[depth])
	child, exists := node.children[ch]
	if !exists {
		return false
	}

	shouldDeleteChild := t.remove(child, word, depth+1)

	if shouldDeleteChild {
		delete(node.children, ch)
		return len(node.children) == 0 && !node.isEnd
	}

	return false
}

// Size returns the number of words in the Trie
func (t *ThreadSafeTrie) Size() int {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.size
}

// IsEmpty returns true if the Trie is empty
func (t *ThreadSafeTrie) IsEmpty() bool {
	return t.Size() == 0
}

// Clear removes all words from the Trie
func (t *ThreadSafeTrie) Clear() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.root = &ThreadSafeTrieNode{
		children: make(map[rune]*ThreadSafeTrieNode),
		isEnd:    false,
	}
	t.size = 0
}

// GetAllWords returns all words in the Trie
func (t *ThreadSafeTrie) GetAllWords() []string {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	result := make([]string, 0, t.size)
	t.getAllWords(t.root, "", &result)
	return result
}

func (t *ThreadSafeTrie) getAllWords(node *ThreadSafeTrieNode, prefix string, result *[]string) {
	if node.isEnd {
		*result = append(*result, prefix)
	}

	for ch, child := range node.children {
		t.getAllWords(child, prefix+string(ch), result)
	}
}

// GetWordsWithPrefix returns all words that start with the given prefix
func (t *ThreadSafeTrie) GetWordsWithPrefix(prefix string) []string {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	current := t.root
	for _, ch := range prefix {
		if _, exists := current.children[ch]; !exists {
			return []string{}
		}
		current = current.children[ch]
	}

	result := make([]string, 0)
	t.getAllWords(current, prefix, &result)
	return result
}
