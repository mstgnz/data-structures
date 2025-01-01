package hash

// HashNode represents a key-value pair in the hash table
type HashNode struct {
	Key   interface{}
	Value interface{}
	Next  *HashNode // Used for chaining collision resolution
}

// HashTable represents the hash table data structure
type HashTable struct {
	Size     int
	Capacity int
	Table    []*HashNode
	Strategy string // "linear" for linear probing, "chain" for chaining
}

// NewHashTable creates a new hash table with specified capacity and collision resolution strategy
func NewHashTable(capacity int, strategy string) *HashTable {
	return &HashTable{
		Size:     0,
		Capacity: capacity,
		Table:    make([]*HashNode, capacity),
		Strategy: strategy,
	}
}

// hash generates a hash value for the given key
func (h *HashTable) hash(key interface{}) int {
	// Simple hash function for demonstration
	switch v := key.(type) {
	case string:
		hash := 0
		for _, ch := range v {
			hash = (hash*31 + int(ch)) % h.Capacity
		}
		return hash
	case int:
		return v % h.Capacity
	default:
		return 0
	}
}

// Put inserts a key-value pair into the hash table
func (h *HashTable) Put(key, value interface{}) bool {
	hashIndex := h.hash(key)

	if h.Strategy == "linear" {
		return h.putLinearProbing(hashIndex, key, value)
	}
	return h.putChaining(hashIndex, key, value)
}

// putLinearProbing handles insertion using linear probing
func (h *HashTable) putLinearProbing(hashIndex int, key, value interface{}) bool {
	if h.Size >= h.Capacity {
		return false // Table is full
	}

	index := hashIndex
	for {
		if h.Table[index] == nil {
			h.Table[index] = &HashNode{Key: key, Value: value}
			h.Size++
			return true
		}
		if h.Table[index].Key == key {
			h.Table[index].Value = value // Update existing key
			return true
		}
		index = (index + 1) % h.Capacity // Linear probing
		if index == hashIndex {
			return false // Went full circle, table is full
		}
	}
}

// putChaining handles insertion using chaining
func (h *HashTable) putChaining(hashIndex int, key, value interface{}) bool {
	newNode := &HashNode{Key: key, Value: value}

	if h.Table[hashIndex] == nil {
		h.Table[hashIndex] = newNode
	} else {
		// Check if key exists and update value
		current := h.Table[hashIndex]
		for current != nil {
			if current.Key == key {
				current.Value = value
				return true
			}
			if current.Next == nil {
				break
			}
			current = current.Next
		}
		// Add new node to the chain
		current.Next = newNode
	}
	h.Size++
	return true
}

// Get retrieves a value by key from the hash table
func (h *HashTable) Get(key interface{}) (interface{}, bool) {
	hashIndex := h.hash(key)

	if h.Strategy == "linear" {
		return h.getLinearProbing(hashIndex, key)
	}
	return h.getChaining(hashIndex, key)
}

// getLinearProbing retrieves a value using linear probing
func (h *HashTable) getLinearProbing(hashIndex int, key interface{}) (interface{}, bool) {
	index := hashIndex
	for {
		if h.Table[index] == nil {
			return nil, false
		}
		if h.Table[index].Key == key {
			return h.Table[index].Value, true
		}
		index = (index + 1) % h.Capacity
		if index == hashIndex {
			return nil, false
		}
	}
}

// getChaining retrieves a value using chaining
func (h *HashTable) getChaining(hashIndex int, key interface{}) (interface{}, bool) {
	current := h.Table[hashIndex]
	for current != nil {
		if current.Key == key {
			return current.Value, true
		}
		current = current.Next
	}
	return nil, false
}

// Remove removes a key-value pair from the hash table
func (h *HashTable) Remove(key interface{}) bool {
	hashIndex := h.hash(key)

	if h.Strategy == "linear" {
		return h.removeLinearProbing(hashIndex, key)
	}
	return h.removeChaining(hashIndex, key)
}

// removeLinearProbing removes an entry using linear probing
func (h *HashTable) removeLinearProbing(hashIndex int, key interface{}) bool {
	index := hashIndex
	for {
		if h.Table[index] == nil {
			return false
		}
		if h.Table[index].Key == key {
			h.Table[index] = nil
			h.Size--
			return true
		}
		index = (index + 1) % h.Capacity
		if index == hashIndex {
			return false
		}
	}
}

// removeChaining removes an entry using chaining
func (h *HashTable) removeChaining(hashIndex int, key interface{}) bool {
	if h.Table[hashIndex] == nil {
		return false
	}

	if h.Table[hashIndex].Key == key {
		h.Table[hashIndex] = h.Table[hashIndex].Next
		h.Size--
		return true
	}

	current := h.Table[hashIndex]
	for current.Next != nil {
		if current.Next.Key == key {
			current.Next = current.Next.Next
			h.Size--
			return true
		}
		current = current.Next
	}
	return false
}
