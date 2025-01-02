package tree

// Node represents a node in the Splay Tree
type SplayNode struct {
	key         int
	left, right *SplayNode
	parent      *SplayNode
}

// SplayTree represents a self-adjusting binary search tree
type SplayTree struct {
	root *SplayNode
	size int
}

// NewSplayTree creates and returns a new Splay Tree
func NewSplayTree() *SplayTree {
	return &SplayTree{nil, 0}
}

// rotateRight performs a right rotation on the given node
func (st *SplayTree) rotateRight(x *SplayNode) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		st.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

// rotateLeft performs a left rotation on the given node
func (st *SplayTree) rotateLeft(x *SplayNode) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		st.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

// splay brings the given node to the root through a series of rotations
func (st *SplayTree) splay(x *SplayNode) {
	for x.parent != nil {
		if x.parent.parent == nil { // Zig step
			if x == x.parent.left {
				st.rotateRight(x.parent)
			} else {
				st.rotateLeft(x.parent)
			}
		} else if x == x.parent.left && x.parent == x.parent.parent.left { // Zig-zig step
			st.rotateRight(x.parent.parent)
			st.rotateRight(x.parent)
		} else if x == x.parent.right && x.parent == x.parent.parent.right { // Zig-zig step
			st.rotateLeft(x.parent.parent)
			st.rotateLeft(x.parent)
		} else if x == x.parent.right && x.parent == x.parent.parent.left { // Zig-zag step
			st.rotateLeft(x.parent)
			st.rotateRight(x.parent)
		} else { // Zig-zag step
			st.rotateRight(x.parent)
			st.rotateLeft(x.parent)
		}
	}
}

// Insert adds a new key to the tree
func (st *SplayTree) Insert(key int) {
	var y *SplayNode
	x := st.root
	newNode := &SplayNode{key: key}

	for x != nil {
		y = x
		if key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	newNode.parent = y
	if y == nil {
		st.root = newNode
	} else if key < y.key {
		y.left = newNode
	} else {
		y.right = newNode
	}

	st.size++
	st.splay(newNode)
}

// Search finds a node with the given key and splays it to the root
func (st *SplayTree) Search(key int) bool {
	x := st.root
	var lastNode *SplayNode

	for x != nil {
		lastNode = x
		if key < x.key {
			x = x.left
		} else if key > x.key {
			x = x.right
		} else {
			st.splay(x)
			return true
		}
	}

	if lastNode != nil {
		st.splay(lastNode)
	}
	return false
}

// Delete removes a node with the given key from the tree
func (st *SplayTree) Delete(key int) bool {
	if !st.Search(key) {
		return false
	}

	leftTree := st.root.left
	if leftTree != nil {
		leftTree.parent = nil
	}
	rightTree := st.root.right
	if rightTree != nil {
		rightTree.parent = nil
	}

	if leftTree == nil {
		st.root = rightTree
	} else {
		// Find the maximum element in the left subtree
		maxNode := leftTree
		for maxNode.right != nil {
			maxNode = maxNode.right
		}
		st.root = leftTree
		st.splay(maxNode)
		st.root.right = rightTree
		if rightTree != nil {
			rightTree.parent = st.root
		}
	}

	st.size--
	return true
}

// Size returns the number of nodes in the tree
func (st *SplayTree) Size() int {
	return st.size
}

// IsEmpty returns true if the tree is empty
func (st *SplayTree) IsEmpty() bool {
	return st.size == 0
}

// Clear removes all nodes from the tree
func (st *SplayTree) Clear() {
	st.root = nil
	st.size = 0
}
