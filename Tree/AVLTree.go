package tree

import "sync"

// AVLNode represents a node in AVL tree
type AVLNode struct {
	Key    int
	Height int
	Left   *AVLNode
	Right  *AVLNode
}

// AVLTree represents an AVL tree
type AVLTree struct {
	Root  *AVLNode
	mutex sync.RWMutex
}

// NewAVLTree creates a new AVL tree
func NewAVLTree() *AVLTree {
	return &AVLTree{nil, sync.RWMutex{}}
}

// Height returns the height of the node
func (n *AVLNode) height() int {
	if n == nil {
		return 0
	}
	return n.Height
}

// getBalance calculates the balance factor of the node
func (n *AVLNode) getBalance() int {
	if n == nil {
		return 0
	}
	return n.Left.height() - n.Right.height()
}

// maxInt returns the maximum of two integers
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// updateHeight updates the height of the node
func (n *AVLNode) updateHeight() {
	n.Height = maxInt(n.Left.height(), n.Right.height()) + 1
}

// rightRotate performs right rotation
func rightRotate(y *AVLNode) *AVLNode {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	y.updateHeight()
	x.updateHeight()

	return x
}

// leftRotate performs left rotation
func leftRotate(x *AVLNode) *AVLNode {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	x.updateHeight()
	y.updateHeight()

	return y
}

// Insert adds a new node to the tree
func (t *AVLTree) Insert(key int) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.Root = t.insert(t.Root, key)
}

func (t *AVLTree) insert(node *AVLNode, key int) *AVLNode {
	// Normal BST insertion
	if node == nil {
		return &AVLNode{Key: key, Height: 1}
	}

	if key < node.Key {
		node.Left = t.insert(node.Left, key)
	} else if key > node.Key {
		node.Right = t.insert(node.Right, key)
	} else {
		return node // Duplicate keys not allowed
	}

	// Update height
	node.updateHeight()

	// Check balance factor
	balance := node.getBalance()

	// Left Left Case
	if balance > 1 && key < node.Left.Key {
		return rightRotate(node)
	}

	// Right Right Case
	if balance < -1 && key > node.Right.Key {
		return leftRotate(node)
	}

	// Left Right Case
	if balance > 1 && key > node.Left.Key {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	// Right Left Case
	if balance < -1 && key < node.Right.Key {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}

// Search looks for a value in the tree
func (t *AVLTree) Search(key int) bool {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.search(t.Root, key)
}

func (t *AVLTree) search(node *AVLNode, key int) bool {
	if node == nil {
		return false
	}

	if key == node.Key {
		return true
	}

	if key < node.Key {
		return t.search(node.Left, key)
	}

	return t.search(node.Right, key)
}

// InOrderTraversal performs inorder traversal of the tree
func (t *AVLTree) InOrderTraversal(result *[]int) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	t.inOrderTraversal(t.Root, result)
}

func (t *AVLTree) inOrderTraversal(node *AVLNode, result *[]int) {
	if node != nil {
		t.inOrderTraversal(node.Left, result)
		*result = append(*result, node.Key)
		t.inOrderTraversal(node.Right, result)
	}
}
