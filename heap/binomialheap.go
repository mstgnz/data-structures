package heap

import (
	"errors"
	"sync"
)

// BinomialNode represents a node in the binomial heap
type BinomialNode struct {
	key     int
	degree  int
	parent  *BinomialNode
	child   *BinomialNode
	sibling *BinomialNode
}

// BinomialHeap represents a binomial heap data structure
type BinomialHeap struct {
	head  *BinomialNode
	size  int
	mutex sync.RWMutex
}

// NewBinomialHeap creates and returns a new instance of BinomialHeap
func NewBinomialHeap() *BinomialHeap {
	return &BinomialHeap{
		head:  nil,
		size:  0,
		mutex: sync.RWMutex{},
	}
}

// Insert adds a new value to the binomial heap
func (bh *BinomialHeap) Insert(value int) {
	bh.mutex.Lock()
	defer bh.mutex.Unlock()

	node := &BinomialNode{key: value}
	if bh.head == nil {
		bh.head = node
	} else {
		bh.union(NewBinomialHeap().withNode(node))
	}
	bh.size++
}

// withNode is a helper function to create a heap with a single node
func (bh *BinomialHeap) withNode(node *BinomialNode) *BinomialHeap {
	bh.head = node
	return bh
}

// Extract removes and returns the minimum element from the heap
func (bh *BinomialHeap) Extract() (int, error) {
	bh.mutex.Lock()
	defer bh.mutex.Unlock()

	if bh.IsEmpty() {
		return 0, errors.New("heap is empty")
	}

	min := bh.head
	prevMin := (*BinomialNode)(nil)
	curr := bh.head
	prev := (*BinomialNode)(nil)

	// Find minimum
	for curr != nil {
		if curr.key < min.key {
			min = curr
			prevMin = prev
		}
		prev = curr
		curr = curr.sibling
	}

	// Remove minimum node
	if prevMin == nil {
		bh.head = min.sibling
	} else {
		prevMin.sibling = min.sibling
	}

	// Create a new heap from the children
	newHeap := NewBinomialHeap()
	child := min.child
	for child != nil {
		next := child.sibling
		child.sibling = newHeap.head
		child.parent = nil
		newHeap.head = child
		child = next
	}

	bh.union(newHeap)
	bh.size--

	return min.key, nil
}

// union combines two binomial heaps
func (bh *BinomialHeap) union(other *BinomialHeap) {
	if other.head == nil {
		return
	}

	bh.head = bh.merge(bh.head, other.head)

	if bh.head == nil {
		return
	}

	var prev *BinomialNode
	curr := bh.head
	next := curr.sibling

	for next != nil {
		if curr.degree != next.degree ||
			(next.sibling != nil && next.sibling.degree == curr.degree) {
			prev = curr
			curr = next
		} else if curr.key <= next.key {
			curr.sibling = next.sibling
			bh.link(next, curr)
		} else {
			if prev == nil {
				bh.head = next
			} else {
				prev.sibling = next
			}
			bh.link(curr, next)
			curr = next
		}
		next = curr.sibling
	}
}

// merge combines two binomial trees of the same degree
func (bh *BinomialHeap) merge(h1, h2 *BinomialNode) *BinomialNode {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}

	var head *BinomialNode
	curr := &head

	for h1 != nil && h2 != nil {
		if h1.degree <= h2.degree {
			*curr = h1
			h1 = h1.sibling
		} else {
			*curr = h2
			h2 = h2.sibling
		}
		curr = &((*curr).sibling)
	}

	if h1 != nil {
		*curr = h1
	} else {
		*curr = h2
	}

	return head
}

// link makes node y a child of node x
func (bh *BinomialHeap) link(y, x *BinomialNode) {
	y.parent = x
	y.sibling = x.child
	x.child = y
	x.degree++
}

// Peek returns the minimum element without removing it
func (bh *BinomialHeap) Peek() (int, error) {
	bh.mutex.RLock()
	defer bh.mutex.RUnlock()

	if bh.IsEmpty() {
		return 0, errors.New("heap is empty")
	}

	min := bh.head.key
	curr := bh.head.sibling

	for curr != nil {
		if curr.key < min {
			min = curr.key
		}
		curr = curr.sibling
	}

	return min, nil
}

// Size returns the number of elements in the heap
func (bh *BinomialHeap) Size() int {
	bh.mutex.RLock()
	defer bh.mutex.RUnlock()
	return bh.size
}

// IsEmpty returns true if the heap contains no elements
func (bh *BinomialHeap) IsEmpty() bool {
	return bh.head == nil
}
