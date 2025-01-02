package heap

// LeftistNode represents a node in the Leftist Heap
type LeftistNode struct {
	key        int
	leftChild  *LeftistNode
	rightChild *LeftistNode
	npl        int // Null Path Length
}

// LeftistHeap represents a leftist heap data structure
type LeftistHeap struct {
	root *LeftistNode
	size int
}

// NewLeftistHeap creates a new empty leftist heap
func NewLeftistHeap() *LeftistHeap {
	return &LeftistHeap{nil, 0}
}

// merge merges two nodes and returns the resulting node
func (lh *LeftistHeap) merge(h1, h2 *LeftistNode) *LeftistNode {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}

	// Ensure h1 has the smaller key (min-heap property)
	if h1.key > h2.key {
		h1, h2 = h2, h1
	}

	// Recursively merge right subtree with h2
	h1.rightChild = lh.merge(h1.rightChild, h2)

	// Update null path length
	if h1.leftChild == nil {
		// If left child is nil, swap with right child
		h1.leftChild = h1.rightChild
		h1.rightChild = nil
		h1.npl = 0
	} else {
		// Ensure leftist property: left child has larger npl
		if h1.leftChild.npl < h1.rightChild.npl {
			h1.leftChild, h1.rightChild = h1.rightChild, h1.leftChild
		}
		h1.npl = h1.rightChild.npl + 1
	}

	return h1
}

// Insert adds a new key to the heap
func (lh *LeftistHeap) Insert(key int) {
	newNode := &LeftistNode{key: key, npl: 0}
	lh.root = lh.merge(lh.root, newNode)
	lh.size++
}

// DeleteMin removes and returns the minimum element
func (lh *LeftistHeap) DeleteMin() (int, bool) {
	if lh.root == nil {
		return 0, false
	}

	minKey := lh.root.key
	lh.root = lh.merge(lh.root.leftChild, lh.root.rightChild)
	lh.size--
	return minKey, true
}

// Merge combines two leftist heaps into one
func (lh *LeftistHeap) Merge(other *LeftistHeap) {
	if other == nil {
		return
	}
	lh.root = lh.merge(lh.root, other.root)
	lh.size += other.size
	other.root = nil
	other.size = 0
}

// FindMin returns the minimum element without removing it
func (lh *LeftistHeap) FindMin() (int, bool) {
	if lh.root == nil {
		return 0, false
	}
	return lh.root.key, true
}

// Size returns the number of elements in the heap
func (lh *LeftistHeap) Size() int {
	return lh.size
}

// IsEmpty returns true if the heap is empty
func (lh *LeftistHeap) IsEmpty() bool {
	return lh.size == 0
}

// Clear removes all elements from the heap
func (lh *LeftistHeap) Clear() {
	lh.root = nil
	lh.size = 0
}
