package utils

// Iterator interface defines the basic iterator operations
type Iterator[T any] interface {
	HasNext() bool
	Next() T
	Reset()
}

// SliceIterator implements Iterator for slices
type SliceIterator[T any] struct {
	data  []T
	index int
}

// NewSliceIterator creates a new slice iterator
func NewSliceIterator[T any](data []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		data:  data,
		index: 0,
	}
}

func (it *SliceIterator[T]) HasNext() bool {
	return it.index < len(it.data)
}

func (it *SliceIterator[T]) Next() T {
	if !it.HasNext() {
		panic("no more elements")
	}
	value := it.data[it.index]
	it.index++
	return value
}

func (it *SliceIterator[T]) Reset() {
	it.index = 0
}

// MapIterator implements Iterator for maps
type MapIterator[K comparable, V any] struct {
	data     map[K]V
	keys     []K
	index    int
	keyOrder func([]K)
}

// NewMapIterator creates a new map iterator with optional key ordering function
func NewMapIterator[K comparable, V any](data map[K]V, keyOrder func([]K)) *MapIterator[K, V] {
	keys := make([]K, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}

	it := &MapIterator[K, V]{
		data:     data,
		keys:     keys,
		index:    0,
		keyOrder: keyOrder,
	}

	if keyOrder != nil {
		keyOrder(it.keys)
	}

	return it
}

func (it *MapIterator[K, V]) HasNext() bool {
	return it.index < len(it.keys)
}

func (it *MapIterator[K, V]) Next() (K, V) {
	if !it.HasNext() {
		panic("no more elements")
	}
	key := it.keys[it.index]
	value := it.data[key]
	it.index++
	return key, value
}

func (it *MapIterator[K, V]) Reset() {
	it.index = 0
	if it.keyOrder != nil {
		it.keyOrder(it.keys)
	}
}

// LinkedListIterator implements Iterator for linked lists
type LinkedListIterator[T any] struct {
	head    *Node[T]
	current *Node[T]
}

// NewLinkedListIterator creates a new linked list iterator
func NewLinkedListIterator[T any](head *Node[T]) *LinkedListIterator[T] {
	return &LinkedListIterator[T]{
		head:    head,
		current: head,
	}
}

func (it *LinkedListIterator[T]) HasNext() bool {
	return it.current != nil
}

func (it *LinkedListIterator[T]) Next() T {
	if !it.HasNext() {
		panic("no more elements")
	}
	value := it.current.Value
	it.current = it.current.Next
	return value
}

func (it *LinkedListIterator[T]) Reset() {
	it.current = it.head
}

// TreeIterator implements Iterator for binary trees with different traversal orders
type TraversalOrder int

const (
	PreOrder TraversalOrder = iota
	InOrder
	PostOrder
	LevelOrder
)

type TreeIterator[T any] struct {
	root    *TreeNode[T]
	order   TraversalOrder
	stack   []*TreeNode[T]
	queue   []*TreeNode[T]
	visited map[*TreeNode[T]]bool
	current *TreeNode[T]
}

// NewTreeIterator creates a new tree iterator with specified traversal order
func NewTreeIterator[T any](root *TreeNode[T], order TraversalOrder) *TreeIterator[T] {
	it := &TreeIterator[T]{
		root:    root,
		order:   order,
		stack:   make([]*TreeNode[T], 0),
		queue:   make([]*TreeNode[T], 0),
		visited: make(map[*TreeNode[T]]bool),
	}
	it.Reset()
	return it
}

func (it *TreeIterator[T]) HasNext() bool {
	switch it.order {
	case PreOrder:
		return len(it.stack) > 0
	case InOrder:
		return it.current != nil || len(it.stack) > 0
	case PostOrder:
		return it.current != nil || len(it.stack) > 0
	case LevelOrder:
		return len(it.queue) > 0
	default:
		return false
	}
}

func (it *TreeIterator[T]) Next() T {
	if !it.HasNext() {
		panic("no more elements")
	}

	var node *TreeNode[T]

	switch it.order {
	case PreOrder:
		node = it.stack[len(it.stack)-1]
		it.stack = it.stack[:len(it.stack)-1]

		// Push right first so left is processed first
		if node.Right != nil {
			it.stack = append(it.stack, node.Right)
		}
		if node.Left != nil {
			it.stack = append(it.stack, node.Left)
		}

	case InOrder:
		// Travel to leftmost node
		for it.current != nil {
			it.stack = append(it.stack, it.current)
			it.current = it.current.Left
		}

		// Process current node and move to right subtree
		node = it.stack[len(it.stack)-1]
		it.stack = it.stack[:len(it.stack)-1]
		it.current = node.Right

	case PostOrder:
		for {
			if it.current != nil {
				it.stack = append(it.stack, it.current)
				it.current = it.current.Left
			} else {
				peek := it.stack[len(it.stack)-1]
				if peek.Right != nil && !it.visited[peek.Right] {
					it.current = peek.Right
				} else {
					node = peek
					it.stack = it.stack[:len(it.stack)-1]
					it.visited[node] = true
					break
				}
			}
		}

	case LevelOrder:
		node = it.queue[0]
		it.queue = it.queue[1:]

		if node.Left != nil {
			it.queue = append(it.queue, node.Left)
		}
		if node.Right != nil {
			it.queue = append(it.queue, node.Right)
		}
	}

	return node.Value
}

func (it *TreeIterator[T]) Reset() {
	it.stack = make([]*TreeNode[T], 0)
	it.queue = make([]*TreeNode[T], 0)
	it.visited = make(map[*TreeNode[T]]bool)
	it.current = nil

	if it.root == nil {
		return
	}

	switch it.order {
	case PreOrder:
		it.stack = append(it.stack, it.root)
	case InOrder:
		it.current = it.root
	case PostOrder:
		it.current = it.root
	case LevelOrder:
		it.queue = append(it.queue, it.root)
	}
}
