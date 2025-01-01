package Tree

// Node AVL ağacı için düğüm yapısı
type AVLNode struct {
	Key    int
	Height int
	Left   *AVLNode
	Right  *AVLNode
}

// AVLTree AVL ağacı yapısı
type AVLTree struct {
	Root *AVLNode
}

// NewAVLTree yeni bir AVL ağacı oluşturur
func NewAVLTree() *AVLTree {
	return &AVLTree{nil}
}

// Height düğümün yüksekliğini döndürür
func (n *AVLNode) height() int {
	if n == nil {
		return 0
	}
	return n.Height
}

// getBalance düğümün denge faktörünü hesaplar
func (n *AVLNode) getBalance() int {
	if n == nil {
		return 0
	}
	return n.Left.height() - n.Right.height()
}

// maxInt iki sayıdan büyük olanı döndürür
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// updateHeight düğümün yüksekliğini günceller
func (n *AVLNode) updateHeight() {
	n.Height = maxInt(n.Left.height(), n.Right.height()) + 1
}

// rightRotate sağa döndürme işlemi
func rightRotate(y *AVLNode) *AVLNode {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	y.updateHeight()
	x.updateHeight()

	return x
}

// leftRotate sola döndürme işlemi
func leftRotate(x *AVLNode) *AVLNode {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	x.updateHeight()
	y.updateHeight()

	return y
}

// Insert ağaca yeni bir düğüm ekler
func (t *AVLTree) Insert(key int) {
	t.Root = t.insert(t.Root, key)
}

func (t *AVLTree) insert(node *AVLNode, key int) *AVLNode {
	// Normal BST ekleme işlemi
	if node == nil {
		return &AVLNode{Key: key, Height: 1}
	}

	if key < node.Key {
		node.Left = t.insert(node.Left, key)
	} else if key > node.Key {
		node.Right = t.insert(node.Right, key)
	} else {
		return node // Aynı anahtarlar kabul edilmez
	}

	// Yüksekliği güncelle
	node.updateHeight()

	// Denge faktörünü kontrol et
	balance := node.getBalance()

	// Sol Sol Durumu
	if balance > 1 && key < node.Left.Key {
		return rightRotate(node)
	}

	// Sağ Sağ Durumu
	if balance < -1 && key > node.Right.Key {
		return leftRotate(node)
	}

	// Sol Sağ Durumu
	if balance > 1 && key > node.Left.Key {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	// Sağ Sol Durumu
	if balance < -1 && key < node.Right.Key {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}

// Search ağaçta bir değer arar
func (t *AVLTree) Search(key int) bool {
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

// InOrderTraversal ağacı inorder gezer
func (t *AVLTree) InOrderTraversal(node *AVLNode, result *[]int) {
	if node != nil {
		t.InOrderTraversal(node.Left, result)
		*result = append(*result, node.Key)
		t.InOrderTraversal(node.Right, result)
	}
}
