package Tree

// BTreeNode represents a node in B-tree
type BTreeNode struct {
	keys     []int
	children []*BTreeNode
	leaf     bool
}

// BTree represents a B-tree
type BTree struct {
	root *BTreeNode
	t    int // Minimum degree (defines the range for number of keys)
}

// NewBTree creates a new B-tree with minimum degree t
func NewBTree(t int) *BTree {
	return &BTree{
		root: &BTreeNode{
			keys:     make([]int, 0),
			children: make([]*BTreeNode, 0),
			leaf:     true,
		},
		t: t,
	}
}

// Search searches for a key in the B-tree
func (tree *BTree) Search(k int) bool {
	return tree.searchNode(tree.root, k)
}

func (tree *BTree) searchNode(x *BTreeNode, k int) bool {
	i := 0
	for i < len(x.keys) && k > x.keys[i] {
		i++
	}

	if i < len(x.keys) && k == x.keys[i] {
		return true
	}

	if x.leaf {
		return false
	}

	return tree.searchNode(x.children[i], k)
}

// Insert inserts a key into the B-tree
func (tree *BTree) Insert(k int) {
	root := tree.root
	if len(root.keys) == 2*tree.t-1 {
		// Create new root
		newRoot := &BTreeNode{
			keys:     make([]int, 0),
			children: make([]*BTreeNode, 0),
			leaf:     false,
		}
		tree.root = newRoot
		newRoot.children = append(newRoot.children, root)
		tree.splitChild(newRoot, 0)
		tree.insertNonFull(newRoot, k)
	} else {
		tree.insertNonFull(root, k)
	}
}

func (tree *BTree) insertNonFull(x *BTreeNode, k int) {
	i := len(x.keys) - 1

	if x.leaf {
		// Yaprak düğüme ekleme
		pos := 0
		// Anahtarın ekleneceği pozisyonu bul
		for pos < len(x.keys) && x.keys[pos] < k {
			pos++
		}
		// Yeni anahtarı ekle
		x.keys = append(x.keys, 0)
		copy(x.keys[pos+1:], x.keys[pos:])
		x.keys[pos] = k
	} else {
		// İç düğüme ekleme
		// Uygun çocuk düğümü bul
		for i >= 0 && k < x.keys[i] {
			i--
		}
		i++

		// Eğer çocuk düğüm doluysa, böl
		if len(x.children[i].keys) == 2*tree.t-1 {
			tree.splitChild(x, i)
			if k > x.keys[i] {
				i++
			}
		}
		tree.insertNonFull(x.children[i], k)
	}
}

func (tree *BTree) splitChild(x *BTreeNode, i int) {
	t := tree.t
	y := x.children[i]

	// Yeni düğüm oluştur
	z := &BTreeNode{
		keys:     make([]int, 0, t-1),
		children: make([]*BTreeNode, 0, t),
		leaf:     y.leaf,
	}

	// Orta anahtarı belirle
	midKey := y.keys[t-1]

	// z'ye sağdaki anahtarları kopyala
	z.keys = append(z.keys, y.keys[t:]...)

	// Eğer yaprak değilse, çocukları da kopyala
	if !y.leaf {
		z.children = append(z.children, y.children[t:]...)
	}

	// y'nin boyutunu küçült
	y.keys = y.keys[:t-1]
	if !y.leaf {
		y.children = y.children[:t]
	}

	// x'e yeni çocuğu ekle
	x.children = append(x.children, nil)
	copy(x.children[i+2:], x.children[i+1:])
	x.children[i+1] = z

	// x'e orta anahtarı ekle
	x.keys = append(x.keys, 0)
	copy(x.keys[i+1:], x.keys[i:])
	x.keys[i] = midKey
}

// Delete removes a key from the B-tree
func (tree *BTree) Delete(k int) {
	tree.deleteNode(tree.root, k)
	if len(tree.root.keys) == 0 && !tree.root.leaf {
		tree.root = tree.root.children[0]
	}
}

func (tree *BTree) deleteNode(x *BTreeNode, k int) {
	t := tree.t
	i := 0
	for i < len(x.keys) && k > x.keys[i] {
		i++
	}

	if i < len(x.keys) && k == x.keys[i] {
		if x.leaf {
			// Case 1: Key is in leaf node
			x.keys = append(x.keys[:i], x.keys[i+1:]...)
		} else {
			// Case 2: Key is in internal node
			if len(x.children[i].keys) >= t {
				// Case 2a: Left child has at least t keys
				pred := tree.getPredecessor(x, i)
				x.keys[i] = pred
				tree.deleteNode(x.children[i], pred)
			} else if len(x.children[i+1].keys) >= t {
				// Case 2b: Right child has at least t keys
				succ := tree.getSuccessor(x, i)
				x.keys[i] = succ
				tree.deleteNode(x.children[i+1], succ)
			} else {
				// Case 2c: Both children have t-1 keys
				tree.mergeChildren(x, i)
				tree.deleteNode(x.children[i], k)
			}
		}
	} else {
		if x.leaf {
			return
		}
		lastI := i
		if i == len(x.keys) {
			lastI--
		}

		if len(x.children[i].keys) == t-1 {
			tree.fill(x, i)
		}

		if i > len(x.keys) {
			tree.deleteNode(x.children[lastI], k)
		} else {
			tree.deleteNode(x.children[i], k)
		}
	}
}

func (tree *BTree) getPredecessor(x *BTreeNode, idx int) int {
	curr := x.children[idx]
	for !curr.leaf {
		curr = curr.children[len(curr.children)-1]
	}
	return curr.keys[len(curr.keys)-1]
}

func (tree *BTree) getSuccessor(x *BTreeNode, idx int) int {
	curr := x.children[idx+1]
	for !curr.leaf {
		curr = curr.children[0]
	}
	return curr.keys[0]
}

func (tree *BTree) fill(x *BTreeNode, idx int) {
	if idx != 0 && len(x.children[idx-1].keys) >= tree.t {
		tree.borrowFromPrev(x, idx)
	} else if idx != len(x.keys) && len(x.children[idx+1].keys) >= tree.t {
		tree.borrowFromNext(x, idx)
	} else {
		if idx != len(x.keys) {
			tree.mergeChildren(x, idx)
		} else {
			tree.mergeChildren(x, idx-1)
		}
	}
}

func (tree *BTree) borrowFromPrev(x *BTreeNode, idx int) {
	child := x.children[idx]
	sibling := x.children[idx-1]

	// Move keys and children
	child.keys = append([]int{x.keys[idx-1]}, child.keys...)
	if !child.leaf {
		child.children = append([]*BTreeNode{sibling.children[len(sibling.children)-1]}, child.children...)
	}

	x.keys[idx-1] = sibling.keys[len(sibling.keys)-1]
	sibling.keys = sibling.keys[:len(sibling.keys)-1]
	if !sibling.leaf {
		sibling.children = sibling.children[:len(sibling.children)-1]
	}
}

func (tree *BTree) borrowFromNext(x *BTreeNode, idx int) {
	child := x.children[idx]
	sibling := x.children[idx+1]

	// Move keys and children
	child.keys = append(child.keys, x.keys[idx])
	if !child.leaf {
		child.children = append(child.children, sibling.children[0])
	}

	x.keys[idx] = sibling.keys[0]
	sibling.keys = sibling.keys[1:]
	if !sibling.leaf {
		sibling.children = sibling.children[1:]
	}
}

func (tree *BTree) mergeChildren(x *BTreeNode, idx int) {
	child := x.children[idx]
	sibling := x.children[idx+1]

	// Move key from x to child
	child.keys = append(child.keys, x.keys[idx])
	child.keys = append(child.keys, sibling.keys...)
	if !child.leaf {
		child.children = append(child.children, sibling.children...)
	}

	// Remove key and child pointer from x
	x.keys = append(x.keys[:idx], x.keys[idx+1:]...)
	x.children = append(x.children[:idx+1], x.children[idx+2:]...)
}

// GetInOrder returns all keys in the B-tree in sorted order
func (tree *BTree) GetInOrder() []int {
	var result []int
	tree.inOrderTraversal(tree.root, &result)
	return result
}

func (tree *BTree) inOrderTraversal(x *BTreeNode, result *[]int) {
	if x == nil {
		return
	}

	i := 0
	for i < len(x.keys) {
		if !x.leaf {
			tree.inOrderTraversal(x.children[i], result)
		}
		*result = append(*result, x.keys[i])
		i++
	}
	if !x.leaf {
		tree.inOrderTraversal(x.children[i], result)
	}
}
