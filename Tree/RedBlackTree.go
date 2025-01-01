package tree

import "sync"

// Color represents the color of a node in Red-Black tree
type Color bool

const (
	RED   Color = true
	BLACK Color = false
)

// RBNode represents a node in Red-Black tree
type RBNode struct {
	Key                 int
	Color               Color
	Left, Right, Parent *RBNode
}

// RedBlackTree represents a Red-Black tree
type RedBlackTree struct {
	Root  *RBNode
	NIL   *RBNode // Sentinel node
	mutex sync.RWMutex
}

// NewRedBlackTree creates a new Red-Black tree
func NewRedBlackTree() *RedBlackTree {
	nil_node := &RBNode{Color: BLACK}
	return &RedBlackTree{
		NIL:   nil_node,
		Root:  nil_node,
		mutex: sync.RWMutex{},
	}
}

// Insert adds a new key to the tree
func (t *RedBlackTree) Insert(key int) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	node := &RBNode{
		Key:    key,
		Color:  RED,
		Left:   t.NIL,
		Right:  t.NIL,
		Parent: t.NIL,
	}

	var y *RBNode = t.NIL
	x := t.Root

	// Binary Search Tree insertion
	for x != t.NIL {
		y = x
		if node.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	node.Parent = y
	if y == t.NIL {
		t.Root = node
	} else if node.Key < y.Key {
		y.Left = node
	} else {
		y.Right = node
	}

	// Fix Red-Black tree properties
	t.insertFixup(node)
}

// insertFixup fixes the Red-Black tree properties after insertion
func (t *RedBlackTree) insertFixup(z *RBNode) {
	for z.Parent.Color == RED {
		if z.Parent == z.Parent.Parent.Left {
			y := z.Parent.Parent.Right
			if y.Color == RED {
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Right {
					z = z.Parent
					t.leftRotate(z)
				}
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				t.rightRotate(z.Parent.Parent)
			}
		} else {
			y := z.Parent.Parent.Left
			if y.Color == RED {
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Left {
					z = z.Parent
					t.rightRotate(z)
				}
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				t.leftRotate(z.Parent.Parent)
			}
		}
		if z == t.Root {
			break
		}
	}
	t.Root.Color = BLACK
}

// leftRotate performs a left rotation
func (t *RedBlackTree) leftRotate(x *RBNode) {
	y := x.Right
	x.Right = y.Left
	if y.Left != t.NIL {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == t.NIL {
		t.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

// rightRotate performs a right rotation
func (t *RedBlackTree) rightRotate(x *RBNode) {
	y := x.Left
	x.Left = y.Right
	if y.Right != t.NIL {
		y.Right.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == t.NIL {
		t.Root = y
	} else if x == x.Parent.Right {
		x.Parent.Right = y
	} else {
		x.Parent.Left = y
	}
	y.Right = x
	x.Parent = y
}

// Search looks for a key in the tree
func (t *RedBlackTree) Search(key int) bool {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.searchNode(t.Root, key) != t.NIL
}

func (t *RedBlackTree) searchNode(node *RBNode, key int) *RBNode {
	if node == t.NIL || key == node.Key {
		return node
	}
	if key < node.Key {
		return t.searchNode(node.Left, key)
	}
	return t.searchNode(node.Right, key)
}

// InOrderTraversal performs an inorder traversal of the tree
func (t *RedBlackTree) InOrderTraversal(result *[]int) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	t.inOrderTraversal(t.Root, result)
}

func (t *RedBlackTree) inOrderTraversal(node *RBNode, result *[]int) {
	if node != t.NIL {
		t.inOrderTraversal(node.Left, result)
		*result = append(*result, node.Key)
		t.inOrderTraversal(node.Right, result)
	}
}
