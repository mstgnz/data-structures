package tree

import (
	"sync"

	"github.com/mstgnz/data-structures/utils"
)

// Node represents a node in the binary search tree
type Node[T utils.Ordered] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

// ThreadSafeBST represents a thread-safe generic binary search tree
type ThreadSafeBST[T utils.Ordered] struct {
	root  *Node[T]
	size  int
	mutex sync.RWMutex
}

// NewThreadSafeBST creates a new thread-safe binary search tree
func NewThreadSafeBST[T utils.Ordered]() *ThreadSafeBST[T] {
	return &ThreadSafeBST[T]{
		root: nil,
		size: 0,
	}
}

// Insert adds a new value to the tree
func (t *ThreadSafeBST[T]) Insert(value T) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.root = t.insert(t.root, value)
	t.size++
}

func (t *ThreadSafeBST[T]) insert(node *Node[T], value T) *Node[T] {
	if node == nil {
		return &Node[T]{Value: value}
	}

	if value < node.Value {
		node.Left = t.insert(node.Left, value)
	} else if value > node.Value {
		node.Right = t.insert(node.Right, value)
	}

	return node
}

// Contains checks if a value exists in the tree
func (t *ThreadSafeBST[T]) Contains(value T) bool {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.contains(t.root, value)
}

func (t *ThreadSafeBST[T]) contains(node *Node[T], value T) bool {
	if node == nil {
		return false
	}

	if value == node.Value {
		return true
	}

	if value < node.Value {
		return t.contains(node.Left, value)
	}
	return t.contains(node.Right, value)
}

// Remove removes a value from the tree
func (t *ThreadSafeBST[T]) Remove(value T) bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if !t.contains(t.root, value) {
		return false
	}

	t.root = t.remove(t.root, value)
	t.size--
	return true
}

func (t *ThreadSafeBST[T]) remove(node *Node[T], value T) *Node[T] {
	if node == nil {
		return nil
	}

	if value < node.Value {
		node.Left = t.remove(node.Left, value)
	} else if value > node.Value {
		node.Right = t.remove(node.Right, value)
	} else {
		// Node with only one child or no child
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		// Node with two children
		minValue := t.findMin(node.Right)
		node.Value = minValue
		node.Right = t.remove(node.Right, minValue)
	}

	return node
}

// FindMin returns the minimum value in the tree
func (t *ThreadSafeBST[T]) FindMin() (T, bool) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	if t.root == nil {
		var zero T
		return zero, false
	}

	return t.findMin(t.root), true
}

func (t *ThreadSafeBST[T]) findMin(node *Node[T]) T {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current.Value
}

// FindMax returns the maximum value in the tree
func (t *ThreadSafeBST[T]) FindMax() (T, bool) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	if t.root == nil {
		var zero T
		return zero, false
	}

	return t.findMax(t.root), true
}

func (t *ThreadSafeBST[T]) findMax(node *Node[T]) T {
	current := node
	for current.Right != nil {
		current = current.Right
	}
	return current.Value
}

// Size returns the number of nodes in the tree
func (t *ThreadSafeBST[T]) Size() int {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.size
}

// IsEmpty returns true if the tree is empty
func (t *ThreadSafeBST[T]) IsEmpty() bool {
	return t.Size() == 0
}

// Clear removes all nodes from the tree
func (t *ThreadSafeBST[T]) Clear() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.root = nil
	t.size = 0
}
