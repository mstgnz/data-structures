package tree

import (
	"sync"

	"github.com/mstgnz/data-structures/utils"
)

// ThreadSafeAVLNode represents a node in the AVL tree
type ThreadSafeAVLNode[T any] struct {
	Value  T
	Left   *ThreadSafeAVLNode[T]
	Right  *ThreadSafeAVLNode[T]
	Height int
}

// ThreadSafeAVL represents a thread-safe generic AVL tree
type ThreadSafeAVL[T utils.Ordered] struct {
	root  *ThreadSafeAVLNode[T]
	size  int
	mutex sync.RWMutex
}

// NewThreadSafeAVL creates a new thread-safe AVL tree
func NewThreadSafeAVL[T utils.Ordered]() *ThreadSafeAVL[T] {
	return &ThreadSafeAVL[T]{
		root: nil,
		size: 0,
	}
}

// getNodeHeight returns the height of a node
func getNodeHeight[T any](node *ThreadSafeAVLNode[T]) int {
	if node == nil {
		return -1
	}
	return node.Height
}

// getNodeBalance returns the balance factor of a node
func getNodeBalance[T any](node *ThreadSafeAVLNode[T]) int {
	if node == nil {
		return 0
	}
	return getNodeHeight(node.Left) - getNodeHeight(node.Right)
}

// getMaxInt returns the maximum of two integers
func getMaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// rotateRight performs a right rotation
func rotateRight[T any](y *ThreadSafeAVLNode[T]) *ThreadSafeAVLNode[T] {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	y.Height = 1 + getMaxInt(getNodeHeight(y.Left), getNodeHeight(y.Right))
	x.Height = 1 + getMaxInt(getNodeHeight(x.Left), getNodeHeight(x.Right))

	return x
}

// rotateLeft performs a left rotation
func rotateLeft[T any](x *ThreadSafeAVLNode[T]) *ThreadSafeAVLNode[T] {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	x.Height = 1 + getMaxInt(getNodeHeight(x.Left), getNodeHeight(x.Right))
	y.Height = 1 + getMaxInt(getNodeHeight(y.Left), getNodeHeight(y.Right))

	return y
}

// Insert adds a new value to the tree
func (t *ThreadSafeAVL[T]) Insert(value T) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.root = t.insert(t.root, value)
	t.size++
}

func (t *ThreadSafeAVL[T]) insert(node *ThreadSafeAVLNode[T], value T) *ThreadSafeAVLNode[T] {
	// Normal BST insertion
	if node == nil {
		return &ThreadSafeAVLNode[T]{Value: value, Height: 0}
	}

	if value < node.Value {
		node.Left = t.insert(node.Left, value)
	} else if value > node.Value {
		node.Right = t.insert(node.Right, value)
	} else {
		return node // Duplicate values are not allowed
	}

	// Update height
	node.Height = 1 + getMaxInt(getNodeHeight(node.Left), getNodeHeight(node.Right))

	// Get balance factor
	balance := getNodeBalance(node)

	// Left Left Case
	if balance > 1 && value < node.Left.Value {
		return rotateRight(node)
	}

	// Right Right Case
	if balance < -1 && value > node.Right.Value {
		return rotateLeft(node)
	}

	// Left Right Case
	if balance > 1 && value > node.Left.Value {
		node.Left = rotateLeft(node.Left)
		return rotateRight(node)
	}

	// Right Left Case
	if balance < -1 && value < node.Right.Value {
		node.Right = rotateRight(node.Right)
		return rotateLeft(node)
	}

	return node
}

// Remove removes a value from the tree
func (t *ThreadSafeAVL[T]) Remove(value T) bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if !t.contains(t.root, value) {
		return false
	}

	t.root = t.remove(t.root, value)
	t.size--
	return true
}

func (t *ThreadSafeAVL[T]) remove(node *ThreadSafeAVLNode[T], value T) *ThreadSafeAVLNode[T] {
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

	// Update height
	node.Height = 1 + getMaxInt(getNodeHeight(node.Left), getNodeHeight(node.Right))

	// Get balance factor
	balance := getNodeBalance(node)

	// Left Left Case
	if balance > 1 && getNodeBalance(node.Left) >= 0 {
		return rotateRight(node)
	}

	// Left Right Case
	if balance > 1 && getNodeBalance(node.Left) < 0 {
		node.Left = rotateLeft(node.Left)
		return rotateRight(node)
	}

	// Right Right Case
	if balance < -1 && getNodeBalance(node.Right) <= 0 {
		return rotateLeft(node)
	}

	// Right Left Case
	if balance < -1 && getNodeBalance(node.Right) > 0 {
		node.Right = rotateRight(node.Right)
		return rotateLeft(node)
	}

	return node
}

// Contains checks if a value exists in the tree
func (t *ThreadSafeAVL[T]) Contains(value T) bool {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.contains(t.root, value)
}

func (t *ThreadSafeAVL[T]) contains(node *ThreadSafeAVLNode[T], value T) bool {
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

// FindMin returns the minimum value in the tree
func (t *ThreadSafeAVL[T]) FindMin() (T, bool) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	if t.root == nil {
		var zero T
		return zero, false
	}

	return t.findMin(t.root), true
}

func (t *ThreadSafeAVL[T]) findMin(node *ThreadSafeAVLNode[T]) T {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current.Value
}

// FindMax returns the maximum value in the tree
func (t *ThreadSafeAVL[T]) FindMax() (T, bool) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	if t.root == nil {
		var zero T
		return zero, false
	}

	return t.findMax(t.root), true
}

func (t *ThreadSafeAVL[T]) findMax(node *ThreadSafeAVLNode[T]) T {
	current := node
	for current.Right != nil {
		current = current.Right
	}
	return current.Value
}

// Size returns the number of nodes in the tree
func (t *ThreadSafeAVL[T]) Size() int {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.size
}

// IsEmpty returns true if the tree is empty
func (t *ThreadSafeAVL[T]) IsEmpty() bool {
	return t.Size() == 0
}

// Clear removes all nodes from the tree
func (t *ThreadSafeAVL[T]) Clear() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.root = nil
	t.size = 0
}

// GetHeight returns the height of the tree
func (t *ThreadSafeAVL[T]) GetHeight() int {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return getNodeHeight(t.root)
}
