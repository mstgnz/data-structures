package advanced

// Node represents a node in the doubly linked list
type Node struct {
	key   interface{}
	value interface{}
	prev  *Node
	next  *Node
}

// LRUCache represents the LRU cache data structure
type LRUCache struct {
	capacity int
	size     int
	cache    map[interface{}]*Node
	head     *Node // Most recently used
	tail     *Node // Least recently used
}

// NewLRUCache creates a new LRU cache with the given capacity
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		size:     0,
		cache:    make(map[interface{}]*Node),
		head:     nil,
		tail:     nil,
	}
}

// Get retrieves a value from the cache
func (lru *LRUCache) Get(key interface{}) (interface{}, bool) {
	if node, exists := lru.cache[key]; exists {
		// Move to front (most recently used)
		lru.moveToFront(node)
		return node.value, true
	}
	return nil, false
}

// Put adds or updates a value in the cache
func (lru *LRUCache) Put(key, value interface{}) {
	if node, exists := lru.cache[key]; exists {
		// Update existing node
		node.value = value
		lru.moveToFront(node)
		return
	}

	// Create new node
	newNode := &Node{
		key:   key,
		value: value,
	}

	// Add to cache
	lru.cache[key] = newNode

	// Add to linked list
	if lru.size == 0 {
		lru.head = newNode
		lru.tail = newNode
	} else {
		newNode.next = lru.head
		lru.head.prev = newNode
		lru.head = newNode
	}

	lru.size++

	// Remove least recently used if capacity is exceeded
	if lru.size > lru.capacity {
		lru.removeLRU()
	}
}

// moveToFront moves a node to the front of the list (most recently used)
func (lru *LRUCache) moveToFront(node *Node) {
	if node == lru.head {
		return
	}

	// Remove from current position
	if node == lru.tail {
		lru.tail = node.prev
		lru.tail.next = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}

	// Move to front
	node.prev = nil
	node.next = lru.head
	lru.head.prev = node
	lru.head = node
}

// removeLRU removes the least recently used item from the cache
func (lru *LRUCache) removeLRU() {
	if lru.tail == nil {
		return
	}

	// Remove from cache
	delete(lru.cache, lru.tail.key)

	// Remove from linked list
	if lru.head == lru.tail {
		lru.head = nil
		lru.tail = nil
	} else {
		lru.tail = lru.tail.prev
		lru.tail.next = nil
	}

	lru.size--
}

// Clear removes all items from the cache
func (lru *LRUCache) Clear() {
	lru.cache = make(map[interface{}]*Node)
	lru.head = nil
	lru.tail = nil
	lru.size = 0
}

// GetSize returns the current size of the cache
func (lru *LRUCache) GetSize() int {
	return lru.size
}

// GetCapacity returns the capacity of the cache
func (lru *LRUCache) GetCapacity() int {
	return lru.capacity
}

// Contains checks if a key exists in the cache
func (lru *LRUCache) Contains(key interface{}) bool {
	_, exists := lru.cache[key]
	return exists
}
