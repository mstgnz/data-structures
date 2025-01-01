package tree

import (
	"sync"

	"github.com/mstgnz/data-structures/utils"
)

// ThreadSafeBTreeNode represents a node in the B-tree
type ThreadSafeBTreeNode[T utils.Ordered] struct {
	keys     []T
	children []*ThreadSafeBTreeNode[T]
	leaf     bool
}

// ThreadSafeBTree represents a thread-safe generic B-tree
type ThreadSafeBTree[T utils.Ordered] struct {
	root   *ThreadSafeBTreeNode[T]
	degree int // Minimum degree (defines minimum and maximum number of keys)
	size   int
	mutex  sync.RWMutex
}

// NewThreadSafeBTree creates a new thread-safe B-tree with specified minimum degree
func NewThreadSafeBTree[T utils.Ordered](degree int) *ThreadSafeBTree[T] {
	if degree < 2 {
		degree = 2 // Minimum degree must be at least 2
	}
	return &ThreadSafeBTree[T]{
		root:   nil,
		degree: degree,
		size:   0,
	}
}

// createNode creates a new B-tree node
func (t *ThreadSafeBTree[T]) createNode(leaf bool) *ThreadSafeBTreeNode[T] {
	return &ThreadSafeBTreeNode[T]{
		keys:     make([]T, 0, 2*t.degree-1),
		children: make([]*ThreadSafeBTreeNode[T], 0, 2*t.degree),
		leaf:     leaf,
	}
}

// Insert adds a new value to the tree
func (t *ThreadSafeBTree[T]) Insert(value T) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// If tree is empty
	if t.root == nil {
		t.root = t.createNode(true)
		t.root.keys = append(t.root.keys, value)
		t.size++
		return
	}

	// If root is full, split it
	if len(t.root.keys) == 2*t.degree-1 {
		newRoot := t.createNode(false)
		newRoot.children = append(newRoot.children, t.root)
		t.splitChild(newRoot, 0)
		t.root = newRoot
	}

	t.insertNonFull(t.root, value)
	t.size++
}

// splitChild splits the child of a node
func (t *ThreadSafeBTree[T]) splitChild(parent *ThreadSafeBTreeNode[T], index int) {
	child := parent.children[index]
	median := (len(child.keys) - 1) / 2
	newNode := t.createNode(child.leaf)

	// Copy the latter half of keys to the new node
	newNode.keys = make([]T, 0, t.degree-1)
	newNode.keys = append(newNode.keys, child.keys[median+1:]...)

	// Save median key
	medianKey := child.keys[median]

	// Truncate child's keys
	child.keys = child.keys[:median]

	// If not leaf, handle children
	if !child.leaf {
		newNode.children = make([]*ThreadSafeBTreeNode[T], 0, t.degree)
		newNode.children = append(newNode.children, child.children[median+1:]...)
		child.children = child.children[:median+1]
	}

	// Make space in parent's keys and insert median key
	parent.keys = append(parent.keys, medianKey)
	if index < len(parent.keys)-1 {
		copy(parent.keys[index+1:], parent.keys[index:len(parent.keys)-1])
		parent.keys[index] = medianKey
	}

	// Make space in parent's children and insert new node
	parent.children = append(parent.children, nil)
	if index+1 < len(parent.children)-1 {
		copy(parent.children[index+2:], parent.children[index+1:len(parent.children)-1])
	}
	parent.children[index+1] = newNode
}

// insertNonFull inserts a value into a non-full node
func (t *ThreadSafeBTree[T]) insertNonFull(node *ThreadSafeBTreeNode[T], value T) {
	i := len(node.keys) - 1

	if node.leaf {
		// Insert into leaf node
		for i >= 0 && value < node.keys[i] {
			i--
		}
		i++
		node.keys = append(node.keys[:i], append([]T{value}, node.keys[i:]...)...)
	} else {
		// Find the child to insert into
		for i >= 0 && value < node.keys[i] {
			i--
		}
		i++

		// If child is full, split it
		if len(node.children[i].keys) == 2*t.degree-1 {
			t.splitChild(node, i)
			if value > node.keys[i] {
				i++
			}
		}
		t.insertNonFull(node.children[i], value)
	}
}

// Contains checks if a value exists in the tree
func (t *ThreadSafeBTree[T]) Contains(value T) bool {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.search(t.root, value) != nil
}

// search searches for a value in the tree
func (t *ThreadSafeBTree[T]) search(node *ThreadSafeBTreeNode[T], value T) *ThreadSafeBTreeNode[T] {
	if node == nil {
		return nil
	}

	i := 0
	for i < len(node.keys) && value > node.keys[i] {
		i++
	}

	if i < len(node.keys) && value == node.keys[i] {
		return node
	}

	if node.leaf {
		return nil
	}

	return t.search(node.children[i], value)
}

// Size returns the number of values in the tree
func (t *ThreadSafeBTree[T]) Size() int {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.size
}

// IsEmpty returns true if the tree is empty
func (t *ThreadSafeBTree[T]) IsEmpty() bool {
	return t.Size() == 0
}

// Clear removes all values from the tree
func (t *ThreadSafeBTree[T]) Clear() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.root = nil
	t.size = 0
}

// FindMin returns the minimum value in the tree
func (t *ThreadSafeBTree[T]) FindMin() (T, bool) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	if t.root == nil {
		var zero T
		return zero, false
	}

	node := t.root
	for !node.leaf {
		node = node.children[0]
	}
	return node.keys[0], true
}

// FindMax returns the maximum value in the tree
func (t *ThreadSafeBTree[T]) FindMax() (T, bool) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	if t.root == nil {
		var zero T
		return zero, false
	}

	node := t.root
	for !node.leaf {
		node = node.children[len(node.children)-1]
	}
	return node.keys[len(node.keys)-1], true
}
