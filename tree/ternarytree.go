package tree

// TSTNode represents a node in the Ternary Search Tree
type TSTNode struct {
	char   byte
	isEnd  bool
	value  interface{}
	left   *TSTNode
	middle *TSTNode
	right  *TSTNode
}

// TernarySearchTree represents a ternary search tree data structure
type TernarySearchTree struct {
	root *TSTNode
	size int
}

// NewTernarySearchTree creates a new empty ternary search tree
func NewTernarySearchTree() *TernarySearchTree {
	return &TernarySearchTree{nil, 0}
}

// Insert adds a key-value pair to the tree
func (tst *TernarySearchTree) Insert(key string, value interface{}) {
	if len(key) == 0 {
		return
	}
	tst.root = tst.insert(tst.root, key, 0, value)
}

func (tst *TernarySearchTree) insert(node *TSTNode, key string, pos int, value interface{}) *TSTNode {
	char := key[pos]

	if node == nil {
		node = &TSTNode{char: char}
	}

	if char < node.char {
		node.left = tst.insert(node.left, key, pos, value)
	} else if char > node.char {
		node.right = tst.insert(node.right, key, pos, value)
	} else {
		if pos < len(key)-1 {
			node.middle = tst.insert(node.middle, key, pos+1, value)
		} else {
			if !node.isEnd {
				tst.size++
			}
			node.isEnd = true
			node.value = value
		}
	}

	return node
}

// Search finds a key in the tree and returns its value
func (tst *TernarySearchTree) Search(key string) (interface{}, bool) {
	if len(key) == 0 {
		return nil, false
	}

	node := tst.search(tst.root, key, 0)
	if node != nil && node.isEnd {
		return node.value, true
	}
	return nil, false
}

func (tst *TernarySearchTree) search(node *TSTNode, key string, pos int) *TSTNode {
	if node == nil {
		return nil
	}

	char := key[pos]

	if char < node.char {
		return tst.search(node.left, key, pos)
	} else if char > node.char {
		return tst.search(node.right, key, pos)
	} else {
		if pos == len(key)-1 {
			return node
		}
		return tst.search(node.middle, key, pos+1)
	}
}

// Delete removes a key from the tree
func (tst *TernarySearchTree) Delete(key string) bool {
	if len(key) == 0 {
		return false
	}

	deleted := false
	tst.root, deleted = tst.delete(tst.root, key, 0)
	if deleted {
		tst.size--
	}
	return deleted
}

func (tst *TernarySearchTree) delete(node *TSTNode, key string, pos int) (*TSTNode, bool) {
	if node == nil {
		return nil, false
	}

	deleted := false
	char := key[pos]

	if char < node.char {
		node.left, deleted = tst.delete(node.left, key, pos)
	} else if char > node.char {
		node.right, deleted = tst.delete(node.right, key, pos)
	} else {
		if pos == len(key)-1 {
			if node.isEnd {
				node.isEnd = false
				node.value = nil
				deleted = true
			}
		} else {
			node.middle, deleted = tst.delete(node.middle, key, pos+1)
		}
	}

	// Remove node if it's not marking the end of any string and has no children
	if !node.isEnd && node.left == nil && node.middle == nil && node.right == nil {
		return nil, deleted
	}

	return node, deleted
}

// StartsWith returns all strings in the tree that start with the given prefix
func (tst *TernarySearchTree) StartsWith(prefix string) []string {
	if len(prefix) == 0 {
		return nil
	}

	// Find the node corresponding to the last character of prefix
	node := tst.search(tst.root, prefix, 0)
	if node == nil {
		return nil
	}

	// Collect all strings with this prefix
	result := make([]string, 0)
	if node.isEnd {
		result = append(result, prefix)
	}
	tst.collectStrings(node.middle, prefix, &result)
	return result
}

func (tst *TernarySearchTree) collectStrings(node *TSTNode, prefix string, result *[]string) {
	if node == nil {
		return
	}

	// Traverse left
	tst.collectStrings(node.left, prefix, result)

	// Process current node
	newPrefix := prefix + string(node.char)
	if node.isEnd {
		*result = append(*result, newPrefix)
	}

	// Traverse middle
	tst.collectStrings(node.middle, newPrefix, result)

	// Traverse right
	tst.collectStrings(node.right, prefix, result)
}

// Size returns the number of keys in the tree
func (tst *TernarySearchTree) Size() int {
	return tst.size
}

// IsEmpty returns true if the tree is empty
func (tst *TernarySearchTree) IsEmpty() bool {
	return tst.size == 0
}

// Clear removes all keys from the tree
func (tst *TernarySearchTree) Clear() {
	tst.root = nil
	tst.size = 0
}

// Keys returns all keys in the tree
func (tst *TernarySearchTree) Keys() []string {
	result := make([]string, 0, tst.size)
	tst.collectKeys(tst.root, "", &result)
	return result
}

func (tst *TernarySearchTree) collectKeys(node *TSTNode, prefix string, result *[]string) {
	if node == nil {
		return
	}

	// Traverse left
	tst.collectKeys(node.left, prefix, result)

	// Process current node
	newPrefix := prefix + string(node.char)
	if node.isEnd {
		*result = append(*result, newPrefix)
	}

	// Traverse middle
	tst.collectKeys(node.middle, newPrefix, result)

	// Traverse right
	tst.collectKeys(node.right, prefix, result)
}
