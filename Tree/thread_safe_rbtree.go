package tree

import (
	"sync"

	"github.com/mstgnz/data-structures/utils"
)

// RBColor represents the color of a Red-Black tree node
type RBColor bool

const (
	RB_RED   RBColor = true
	RB_BLACK RBColor = false
)

// ThreadSafeRBNode represents a node in the Red-Black tree
type ThreadSafeRBNode[T any] struct {
	Value  T
	Left   *ThreadSafeRBNode[T]
	Right  *ThreadSafeRBNode[T]
	Parent *ThreadSafeRBNode[T]
	Color  RBColor
}

// ThreadSafeRBTree represents a thread-safe generic Red-Black tree
type ThreadSafeRBTree[T utils.Ordered] struct {
	root  *ThreadSafeRBNode[T]
	size  int
	mutex sync.RWMutex
}

// NewThreadSafeRBTree creates a new thread-safe Red-Black tree
func NewThreadSafeRBTree[T utils.Ordered]() *ThreadSafeRBTree[T] {
	return &ThreadSafeRBTree[T]{
		root: nil,
		size: 0,
	}
}

// rotateLeft performs a left rotation
func (t *ThreadSafeRBTree[T]) rotateLeft(x *ThreadSafeRBNode[T]) {
	y := x.Right
	x.Right = y.Left
	if y.Left != nil {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		t.root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

// rotateRight performs a right rotation
func (t *ThreadSafeRBTree[T]) rotateRight(y *ThreadSafeRBNode[T]) {
	x := y.Left
	y.Left = x.Right
	if x.Right != nil {
		x.Right.Parent = y
	}
	x.Parent = y.Parent
	if y.Parent == nil {
		t.root = x
	} else if y == y.Parent.Right {
		y.Parent.Right = x
	} else {
		y.Parent.Left = x
	}
	x.Right = y
	y.Parent = x
}

// Insert adds a new value to the tree
func (t *ThreadSafeRBTree[T]) Insert(value T) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	node := &ThreadSafeRBNode[T]{
		Value: value,
		Color: RB_RED,
	}

	var parent *ThreadSafeRBNode[T]
	current := t.root

	// Find the insertion point
	for current != nil {
		parent = current
		if value < current.Value {
			current = current.Left
		} else if value > current.Value {
			current = current.Right
		} else {
			return // Duplicate values are not allowed
		}
	}

	node.Parent = parent
	if parent == nil {
		t.root = node
	} else if value < parent.Value {
		parent.Left = node
	} else {
		parent.Right = node
	}

	t.size++
	t.fixInsert(node)
}

// fixInsert maintains Red-Black tree properties after insertion
func (t *ThreadSafeRBTree[T]) fixInsert(node *ThreadSafeRBNode[T]) {
	for node != t.root && node.Parent.Color == RB_RED {
		if node.Parent == node.Parent.Parent.Left {
			uncle := node.Parent.Parent.Right
			if uncle != nil && uncle.Color == RB_RED {
				node.Parent.Color = RB_BLACK
				uncle.Color = RB_BLACK
				node.Parent.Parent.Color = RB_RED
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Right {
					node = node.Parent
					t.rotateLeft(node)
				}
				node.Parent.Color = RB_BLACK
				node.Parent.Parent.Color = RB_RED
				t.rotateRight(node.Parent.Parent)
			}
		} else {
			uncle := node.Parent.Parent.Left
			if uncle != nil && uncle.Color == RB_RED {
				node.Parent.Color = RB_BLACK
				uncle.Color = RB_BLACK
				node.Parent.Parent.Color = RB_RED
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Left {
					node = node.Parent
					t.rotateRight(node)
				}
				node.Parent.Color = RB_BLACK
				node.Parent.Parent.Color = RB_RED
				t.rotateLeft(node.Parent.Parent)
			}
		}
	}
	t.root.Color = RB_BLACK
}

// Contains checks if a value exists in the tree
func (t *ThreadSafeRBTree[T]) Contains(value T) bool {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.contains(value)
}

func (t *ThreadSafeRBTree[T]) contains(value T) bool {
	current := t.root
	for current != nil {
		if value == current.Value {
			return true
		}
		if value < current.Value {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return false
}

// FindMin returns the minimum value in the tree
func (t *ThreadSafeRBTree[T]) FindMin() (T, bool) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	if t.root == nil {
		var zero T
		return zero, false
	}

	current := t.root
	for current.Left != nil {
		current = current.Left
	}
	return current.Value, true
}

// FindMax returns the maximum value in the tree
func (t *ThreadSafeRBTree[T]) FindMax() (T, bool) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()

	if t.root == nil {
		var zero T
		return zero, false
	}

	current := t.root
	for current.Right != nil {
		current = current.Right
	}
	return current.Value, true
}

// Size returns the number of nodes in the tree
func (t *ThreadSafeRBTree[T]) Size() int {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.size
}

// IsEmpty returns true if the tree is empty
func (t *ThreadSafeRBTree[T]) IsEmpty() bool {
	return t.Size() == 0
}

// Clear removes all nodes from the tree
func (t *ThreadSafeRBTree[T]) Clear() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.root = nil
	t.size = 0
}
