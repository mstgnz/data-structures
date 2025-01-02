package tree

const (
	// Maximum number of keys in a node
	maxKeys = 3
	// Minimum number of keys in a node
	minKeys = (maxKeys + 1) / 2
)

// BPlusNode represents a node in the B+ tree
type BPlusNode struct {
	keys     []int
	children []*BPlusNode
	next     *BPlusNode
	isLeaf   bool
}

// BPlusTree represents a B+ tree
type BPlusTree struct {
	root *BPlusNode
	size int
}

// NewBPlusTree creates a new B+ tree
func NewBPlusTree() *BPlusTree {
	return &BPlusTree{
		root: &BPlusNode{
			keys:     make([]int, 0),
			children: make([]*BPlusNode, 0),
			isLeaf:   true,
		},
		size: 0,
	}
}

// newNode creates a new node
func newNode(isLeaf bool) *BPlusNode {
	return &BPlusNode{
		keys:     make([]int, 0),
		children: make([]*BPlusNode, 0),
		isLeaf:   isLeaf,
	}
}

// Insert adds a new key to the tree
func (bt *BPlusTree) Insert(key int) {
	root := bt.root
	if len(root.keys) == maxKeys {
		newRoot := newNode(false)
		bt.root = newRoot
		newRoot.children = append(newRoot.children, root)
		bt.splitChild(newRoot, 0)
		bt.insertNonFull(newRoot, key)
	} else {
		bt.insertNonFull(root, key)
	}
	bt.size++
}

// insertNonFull inserts a key into a non-full node
func (bt *BPlusTree) insertNonFull(node *BPlusNode, key int) {
	i := len(node.keys) - 1
	if node.isLeaf {
		// Insert into leaf node
		node.keys = append(node.keys, 0)
		for i >= 0 && key < node.keys[i] {
			node.keys[i+1] = node.keys[i]
			i--
		}
		node.keys[i+1] = key
	} else {
		// Find the child to insert into
		for i >= 0 && key < node.keys[i] {
			i--
		}
		i++
		if len(node.children[i].keys) == maxKeys {
			bt.splitChild(node, i)
			if key > node.keys[i] {
				i++
			}
		}
		bt.insertNonFull(node.children[i], key)
	}
}

// splitChild splits a full child node
func (bt *BPlusTree) splitChild(parent *BPlusNode, index int) {
	child := parent.children[index]
	newNode := newNode(child.isLeaf)
	parent.keys = append(parent.keys, 0)
	copy(parent.keys[index+1:], parent.keys[index:])
	parent.keys[index] = child.keys[minKeys-1]

	if !child.isLeaf {
		newNode.keys = append(newNode.keys, child.keys[minKeys:]...)
		child.keys = child.keys[:minKeys-1]

		newNode.children = append(newNode.children, child.children[minKeys:]...)
		child.children = child.children[:minKeys]
	} else {
		newNode.keys = append(newNode.keys, child.keys[minKeys-1:]...)
		child.keys = child.keys[:minKeys-1]

		// Link the leaf nodes
		newNode.next = child.next
		child.next = newNode
	}

	parent.children = append(parent.children, nil)
	copy(parent.children[index+2:], parent.children[index+1:])
	parent.children[index+1] = newNode
}

// Search finds a key in the tree
func (bt *BPlusTree) Search(key int) bool {
	return bt.searchNode(bt.root, key)
}

// searchNode searches for a key in a node and its children
func (bt *BPlusTree) searchNode(node *BPlusNode, key int) bool {
	i := 0
	for i < len(node.keys) && key > node.keys[i] {
		i++
	}

	if node.isLeaf {
		return i < len(node.keys) && node.keys[i] == key
	}

	if i < len(node.keys) && key == node.keys[i] {
		return bt.searchNode(node.children[i+1], key)
	}
	return bt.searchNode(node.children[i], key)
}

// Size returns the number of keys in the tree
func (bt *BPlusTree) Size() int {
	return bt.size
}

// IsEmpty returns true if the tree is empty
func (bt *BPlusTree) IsEmpty() bool {
	return bt.size == 0
}

// Clear removes all keys from the tree
func (bt *BPlusTree) Clear() {
	bt.root = &BPlusNode{
		keys:     make([]int, 0),
		children: make([]*BPlusNode, 0),
		isLeaf:   true,
	}
	bt.size = 0
}
