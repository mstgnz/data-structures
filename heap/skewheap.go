package heap

// SkewNode represents a node in the Skew Heap
type SkewNode struct {
	key         int
	left, right *SkewNode
}

// SkewHeap represents a skew heap data structure
type SkewHeap struct {
	root *SkewNode
	size int
}

// NewSkewHeap creates a new empty skew heap
func NewSkewHeap() *SkewHeap {
	return &SkewHeap{nil, 0}
}

// merge merges two nodes and returns the resulting node
// The main difference from leftist heap is that we always swap the children
func (sh *SkewHeap) merge(h1, h2 *SkewNode) *SkewNode {
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
	h1.right = sh.merge(h1.right, h2)

	// Swap children (this is what makes it a skew heap)
	h1.left, h1.right = h1.right, h1.left

	return h1
}

// Insert adds a new key to the heap
func (sh *SkewHeap) Insert(key int) {
	newNode := &SkewNode{key: key}
	sh.root = sh.merge(sh.root, newNode)
	sh.size++
}

// DeleteMin removes and returns the minimum element
func (sh *SkewHeap) DeleteMin() (int, bool) {
	if sh.root == nil {
		return 0, false
	}

	minKey := sh.root.key
	sh.root = sh.merge(sh.root.left, sh.root.right)
	sh.size--
	return minKey, true
}

// Merge combines two skew heaps into one
func (sh *SkewHeap) Merge(other *SkewHeap) {
	if other == nil {
		return
	}
	sh.root = sh.merge(sh.root, other.root)
	sh.size += other.size
	other.root = nil
	other.size = 0
}

// FindMin returns the minimum element without removing it
func (sh *SkewHeap) FindMin() (int, bool) {
	if sh.root == nil {
		return 0, false
	}
	return sh.root.key, true
}

// Size returns the number of elements in the heap
func (sh *SkewHeap) Size() int {
	return sh.size
}

// IsEmpty returns true if the heap is empty
func (sh *SkewHeap) IsEmpty() bool {
	return sh.size == 0
}

// Clear removes all elements from the heap
func (sh *SkewHeap) Clear() {
	sh.root = nil
	sh.size = 0
}

// Internal method to check heap property
func (sh *SkewHeap) checkHeapProperty(node *SkewNode) bool {
	if node == nil {
		return true
	}

	// Check left child
	if node.left != nil && node.left.key < node.key {
		return false
	}

	// Check right child
	if node.right != nil && node.right.key < node.key {
		return false
	}

	// Recursively check children
	return sh.checkHeapProperty(node.left) && sh.checkHeapProperty(node.right)
}
