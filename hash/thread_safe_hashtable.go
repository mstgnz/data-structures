package hash

import (
	"fmt"
	"sync"

	"github.com/mstgnz/data-structures/utils"
)

// bucket represents a bucket in the hash table
type bucket[K comparable, V utils.Any] struct {
	key   K
	value V
	next  *bucket[K, V]
}

// ThreadSafeHashTable represents a thread-safe generic hash table
type ThreadSafeHashTable[K comparable, V utils.Any] struct {
	buckets []*bucket[K, V]
	size    int
	mutex   sync.RWMutex
}

// NewThreadSafeHashTable creates a new thread-safe hash table with initial capacity
func NewThreadSafeHashTable[K comparable, V utils.Any](capacity int) *ThreadSafeHashTable[K, V] {
	if capacity < 1 {
		capacity = 16
	}
	return &ThreadSafeHashTable[K, V]{
		buckets: make([]*bucket[K, V], capacity),
	}
}

// hash generates a hash value for a key
func (h *ThreadSafeHashTable[K, V]) hash(key K) int {
	// Convert the key to its default string representation
	keyStr := fmt.Sprintf("%v", key)
	hash := uint64(0)
	for i := 0; i < len(keyStr); i++ {
		hash = hash*31 + uint64(keyStr[i])
	}
	return int(hash % uint64(len(h.buckets)))
}

// Put adds or updates a key-value pair
func (h *ThreadSafeHashTable[K, V]) Put(key K, value V) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	index := h.hash(key)
	newBucket := &bucket[K, V]{key: key, value: value}

	if h.buckets[index] == nil {
		h.buckets[index] = newBucket
		h.size++
		return
	}

	current := h.buckets[index]
	if current.key == key {
		current.value = value
		return
	}

	for current.next != nil {
		if current.next.key == key {
			current.next.value = value
			return
		}
		current = current.next
	}

	current.next = newBucket
	h.size++
}

// Get retrieves a value by key
func (h *ThreadSafeHashTable[K, V]) Get(key K) (V, bool) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	var zero V
	index := h.hash(key)
	current := h.buckets[index]

	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}

	return zero, false
}

// Remove removes a key-value pair
func (h *ThreadSafeHashTable[K, V]) Remove(key K) bool {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	index := h.hash(key)
	if h.buckets[index] == nil {
		return false
	}

	if h.buckets[index].key == key {
		h.buckets[index] = h.buckets[index].next
		h.size--
		return true
	}

	current := h.buckets[index]
	for current.next != nil {
		if current.next.key == key {
			current.next = current.next.next
			h.size--
			return true
		}
		current = current.next
	}

	return false
}

// Contains checks if a key exists
func (h *ThreadSafeHashTable[K, V]) Contains(key K) bool {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	index := h.hash(key)
	current := h.buckets[index]

	for current != nil {
		if current.key == key {
			return true
		}
		current = current.next
	}

	return false
}

// Size returns the number of key-value pairs
func (h *ThreadSafeHashTable[K, V]) Size() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return h.size
}

// IsEmpty returns true if the hash table is empty
func (h *ThreadSafeHashTable[K, V]) IsEmpty() bool {
	return h.Size() == 0
}

// Clear removes all key-value pairs
func (h *ThreadSafeHashTable[K, V]) Clear() {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.buckets = make([]*bucket[K, V], len(h.buckets))
	h.size = 0
}
