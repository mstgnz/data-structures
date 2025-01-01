package heap

import (
	"errors"
	"math"
	"sync"
)

// FibNode represents a node in the Fibonacci heap
type FibNode struct {
	key    int
	degree int
	marked bool
	parent *FibNode
	child  *FibNode
	left   *FibNode
	right  *FibNode
}

// FibonacciHeap represents a Fibonacci heap data structure
type FibonacciHeap struct {
	min   *FibNode
	size  int
	mutex sync.RWMutex
}

// NewFibonacciHeap creates and returns a new instance of FibonacciHeap
func NewFibonacciHeap() *FibonacciHeap {
	return &FibonacciHeap{
		min:   nil,
		size:  0,
		mutex: sync.RWMutex{},
	}
}

// Insert adds a new value to the Fibonacci heap
func (fh *FibonacciHeap) Insert(value int) {
	fh.mutex.Lock()
	defer fh.mutex.Unlock()

	node := &FibNode{
		key:    value,
		marked: false,
	}
	node.left = node
	node.right = node

	fh.addToRootList(node)
	fh.size++
}

// Extract removes and returns the minimum element from the heap
func (fh *FibonacciHeap) Extract() (int, error) {
	fh.mutex.Lock()
	defer fh.mutex.Unlock()

	if fh.IsEmpty() {
		return 0, errors.New("heap is empty")
	}

	min := fh.min
	if min.child != nil {
		// Add all children to root list
		child := min.child
		for {
			next := child.right
			// Remove from child list
			child.parent = nil
			// Add to root list
			child.left = fh.min
			child.right = fh.min.right
			fh.min.right.left = child
			fh.min.right = child

			if next == min.child {
				break
			}
			child = next
		}
	}

	// Remove min from root list
	if min.right == min {
		fh.min = nil
	} else {
		min.left.right = min.right
		min.right.left = min.left
		fh.min = min.right
	}

	if fh.min != nil {
		fh.consolidate()
	}

	fh.size--
	return min.key, nil
}

// consolidate combines trees of the same degree
func (fh *FibonacciHeap) consolidate() {
	if fh.min == nil {
		return
	}

	// Calculate max degree
	maxDegree := int(math.Log2(float64(fh.size))) + 1
	degreeTable := make([]*FibNode, maxDegree)

	// Collect all roots
	var roots []*FibNode
	current := fh.min
	for {
		roots = append(roots, current)
		current = current.right
		if current == fh.min {
			break
		}
	}

	// Process each root
	for _, root := range roots {
		degree := root.degree
		current := root

		for degreeTable[degree] != nil {
			other := degreeTable[degree]

			// Skip if the node has been processed
			if other == current {
				break
			}

			// Ensure current has smaller key
			if current.key > other.key {
				current, other = other, current
			}

			// Link other under current
			other.left.right = other.right
			other.right.left = other.left

			// Make other a child of current
			other.parent = current
			if current.child == nil {
				current.child = other
				other.right = other
				other.left = other
			} else {
				other.left = current.child
				other.right = current.child.right
				current.child.right = other
				other.right.left = other
			}

			current.degree++
			other.marked = false
			degreeTable[degree] = nil
			degree++
		}

		degreeTable[degree] = current
	}

	// Rebuild root list and find new minimum
	fh.min = nil
	for _, root := range degreeTable {
		if root != nil {
			if fh.min == nil {
				fh.min = root
				root.left = root
				root.right = root
			} else {
				root.left = fh.min
				root.right = fh.min.right
				fh.min.right.left = root
				fh.min.right = root
				if root.key < fh.min.key {
					fh.min = root
				}
			}
		}
	}
}

// addToRootList adds a node to the root list
func (fh *FibonacciHeap) addToRootList(node *FibNode) {
	if fh.min == nil {
		fh.min = node
		node.left = node
		node.right = node
	} else {
		node.right = fh.min.right
		node.left = fh.min
		fh.min.right.left = node
		fh.min.right = node
		if node.key < fh.min.key {
			fh.min = node
		}
	}
}

// Peek returns the minimum element without removing it
func (fh *FibonacciHeap) Peek() (int, error) {
	fh.mutex.RLock()
	defer fh.mutex.RUnlock()

	if fh.IsEmpty() {
		return 0, errors.New("heap is empty")
	}
	return fh.min.key, nil
}

// Size returns the number of elements in the heap
func (fh *FibonacciHeap) Size() int {
	fh.mutex.RLock()
	defer fh.mutex.RUnlock()
	return fh.size
}

// IsEmpty returns true if the heap contains no elements
func (fh *FibonacciHeap) IsEmpty() bool {
	return fh.min == nil
}

// DecreaseKey decreases the key of a node
func (fh *FibonacciHeap) DecreaseKey(node *FibNode, newKey int) error {
	fh.mutex.Lock()
	defer fh.mutex.Unlock()

	if newKey > node.key {
		return errors.New("new key is greater than current key")
	}

	node.key = newKey
	parent := node.parent

	if parent != nil && node.key < parent.key {
		fh.cut(node, parent)
		fh.cascadingCut(parent)
	}

	if node.key < fh.min.key {
		fh.min = node
	}

	return nil
}

// cut removes node from its parent's child list and adds it to the root list
func (fh *FibonacciHeap) cut(node, parent *FibNode) {
	parent.degree--
	if parent.child == node {
		parent.child = node.right
	}
	if parent.degree == 0 {
		parent.child = nil
	}

	node.left.right = node.right
	node.right.left = node.left

	fh.addToRootList(node)
	node.parent = nil
	node.marked = false
}

// cascadingCut performs cascading cuts starting from the given node
func (fh *FibonacciHeap) cascadingCut(node *FibNode) {
	parent := node.parent
	if parent != nil {
		if !node.marked {
			node.marked = true
		} else {
			fh.cut(node, parent)
			fh.cascadingCut(parent)
		}
	}
}
