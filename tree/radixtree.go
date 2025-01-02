package tree

import "strings"

// RadixNode represents a node in the Radix Tree
type RadixNode struct {
	prefix   string
	isEnd    bool
	children map[byte]*RadixNode
	value    interface{}
}

// RadixTree represents a radix tree (compact trie) data structure
type RadixTree struct {
	root *RadixNode
	size int
}

// NewRadixTree creates a new empty radix tree
func NewRadixTree() *RadixTree {
	return &RadixTree{
		root: &RadixNode{
			prefix:   "",
			children: make(map[byte]*RadixNode),
		},
	}
}

// longestCommonPrefix finds the longest common prefix of two strings
func longestCommonPrefix(a, b string) string {
	i := 0
	for i < len(a) && i < len(b) && a[i] == b[i] {
		i++
	}
	return a[:i]
}

// Insert adds a key-value pair to the tree
func (rt *RadixTree) Insert(key string, value interface{}) {
	current := rt.root

	for len(key) > 0 {
		// Find matching child
		firstChar := key[0]
		child, exists := current.children[firstChar]

		if !exists {
			// No matching child, create a new one
			newNode := &RadixNode{
				prefix:   key,
				isEnd:    true,
				children: make(map[byte]*RadixNode),
				value:    value,
			}
			current.children[firstChar] = newNode
			rt.size++
			return
		}

		// Find common prefix
		commonPrefix := longestCommonPrefix(child.prefix, key)

		if commonPrefix == child.prefix {
			// Move to child node and remove common prefix from key
			current = child
			key = key[len(commonPrefix):]
			if len(key) == 0 {
				child.isEnd = true
				child.value = value
				return
			}
		} else {
			// Split the node
			newNode := &RadixNode{
				prefix:   child.prefix[len(commonPrefix):],
				isEnd:    child.isEnd,
				children: child.children,
				value:    child.value,
			}

			child.prefix = commonPrefix
			child.children = make(map[byte]*RadixNode)
			child.children[newNode.prefix[0]] = newNode
			child.isEnd = len(key) == len(commonPrefix)
			if child.isEnd {
				child.value = value
			}

			if len(key) > len(commonPrefix) {
				// Insert remaining part as a new node
				remainingKey := key[len(commonPrefix):]
				newChild := &RadixNode{
					prefix:   remainingKey,
					isEnd:    true,
					children: make(map[byte]*RadixNode),
					value:    value,
				}
				child.children[remainingKey[0]] = newChild
			}
			rt.size++
			return
		}
	}
}

// Search finds a key in the tree and returns its value
func (rt *RadixTree) Search(key string) (interface{}, bool) {
	current := rt.root

	for len(key) > 0 {
		firstChar := key[0]
		child, exists := current.children[firstChar]

		if !exists {
			return nil, false
		}

		if !strings.HasPrefix(key, child.prefix) {
			return nil, false
		}

		key = key[len(child.prefix):]
		current = child
	}

	if current.isEnd {
		return current.value, true
	}
	return nil, false
}

// Delete removes a key from the tree
func (rt *RadixTree) Delete(key string) bool {
	return rt.delete(rt.root, key)
}

func (rt *RadixTree) delete(node *RadixNode, key string) bool {
	if len(key) == 0 {
		if !node.isEnd {
			return false
		}
		node.isEnd = false
		node.value = nil
		rt.size--
		return true
	}

	firstChar := key[0]
	child, exists := node.children[firstChar]
	if !exists {
		return false
	}

	if !strings.HasPrefix(key, child.prefix) {
		return false
	}

	deleted := rt.delete(child, key[len(child.prefix):])

	// Merge nodes if possible
	if deleted && len(child.children) == 1 && !child.isEnd {
		for _, grandChild := range child.children {
			newPrefix := child.prefix + grandChild.prefix
			node.children[firstChar] = &RadixNode{
				prefix:   newPrefix,
				isEnd:    grandChild.isEnd,
				children: grandChild.children,
				value:    grandChild.value,
			}
		}
	}

	return deleted
}

// Size returns the number of keys in the tree
func (rt *RadixTree) Size() int {
	return rt.size
}

// IsEmpty returns true if the tree is empty
func (rt *RadixTree) IsEmpty() bool {
	return rt.size == 0
}

// Clear removes all keys from the tree
func (rt *RadixTree) Clear() {
	rt.root = &RadixNode{
		prefix:   "",
		children: make(map[byte]*RadixNode),
	}
	rt.size = 0
}

// Keys returns all keys in the tree
func (rt *RadixTree) Keys() []string {
	keys := make([]string, 0, rt.size)
	rt.collectKeys(rt.root, "", &keys)
	return keys
}

func (rt *RadixTree) collectKeys(node *RadixNode, prefix string, keys *[]string) {
	currentPrefix := prefix + node.prefix
	if node.isEnd {
		*keys = append(*keys, currentPrefix)
	}
	for _, child := range node.children {
		rt.collectKeys(child, currentPrefix, keys)
	}
}
