package orderedmap

import (
	"fmt"
	"sync"
)

// Pair represents a key-value pair
type Pair struct {
	Key   any
	Value any
}

// OrderedMap represents an ordered map data structure
type OrderedMap struct {
	mu       sync.RWMutex // Mutex for thread safety
	pairs    []Pair
	keyIndex map[any]int
}

// New creates a new instance of OrderedMap
func New() *OrderedMap {
	return &OrderedMap{
		pairs:    make([]Pair, 0),
		keyIndex: make(map[any]int),
	}
}

// Set adds a new key-value pair or updates an existing one
func (om *OrderedMap) Set(key, value any) {
	om.mu.Lock()
	defer om.mu.Unlock()

	if idx, exists := om.keyIndex[key]; exists {
		om.pairs[idx].Value = value
	} else {
		om.pairs = append(om.pairs, Pair{Key: key, Value: value})
		om.keyIndex[key] = len(om.pairs) - 1
	}
}

// Get returns the value associated with the given key
func (om *OrderedMap) Get(key any) (any, bool) {
	om.mu.RLock()
	defer om.mu.RUnlock()

	if idx, exists := om.keyIndex[key]; exists {
		return om.pairs[idx].Value, true
	}
	return nil, false
}

// Delete removes the element with the given key
func (om *OrderedMap) Delete(key any) {
	om.mu.Lock()
	defer om.mu.Unlock()

	if idx, exists := om.keyIndex[key]; exists {
		// Move the last element to the position of the deleted element
		lastIdx := len(om.pairs) - 1
		if idx < lastIdx {
			om.pairs[idx] = om.pairs[lastIdx]
			om.keyIndex[om.pairs[idx].Key] = idx
		}
		// Remove the last element
		om.pairs = om.pairs[:lastIdx]
		delete(om.keyIndex, key)
	}
}

// Len returns the number of elements in the map
func (om *OrderedMap) Len() int {
	om.mu.RLock()
	defer om.mu.RUnlock()
	return len(om.pairs)
}

// Keys returns all keys in order
func (om *OrderedMap) Keys() []any {
	om.mu.RLock()
	defer om.mu.RUnlock()

	keys := make([]any, len(om.pairs))
	for i, pair := range om.pairs {
		keys[i] = pair.Key
	}
	return keys
}

// Values returns all values in order
func (om *OrderedMap) Values() []any {
	om.mu.RLock()
	defer om.mu.RUnlock()

	values := make([]any, len(om.pairs))
	for i, pair := range om.pairs {
		values[i] = pair.Value
	}
	return values
}

// String returns the string representation of the map
func (om *OrderedMap) String() string {
	om.mu.RLock()
	defer om.mu.RUnlock()

	result := "{"
	for i, pair := range om.pairs {
		if i > 0 {
			result += ", "
		}
		result += fmt.Sprintf("%v: %v", pair.Key, pair.Value)
	}
	result += "}"
	return result
}

// Range iterates over the map in order and calls the given function for each key-value pair
// If the function returns false, the iteration stops
func (om *OrderedMap) Range(f func(key, value any) bool) {
	om.mu.RLock()
	defer om.mu.RUnlock()

	for _, pair := range om.pairs {
		if !f(pair.Key, pair.Value) {
			break
		}
	}
}

// Clear removes all elements from the map
func (om *OrderedMap) Clear() {
	om.mu.Lock()
	defer om.mu.Unlock()

	om.pairs = make([]Pair, 0)
	om.keyIndex = make(map[any]int)
}

// Copy returns a new OrderedMap with the same elements
func (om *OrderedMap) Copy() *OrderedMap {
	om.mu.RLock()
	defer om.mu.RUnlock()

	newMap := New()
	for _, pair := range om.pairs {
		newMap.Set(pair.Key, pair.Value)
	}
	return newMap
}

// Has returns true if the key exists in the map
func (om *OrderedMap) Has(key any) bool {
	om.mu.RLock()
	defer om.mu.RUnlock()

	_, exists := om.keyIndex[key]
	return exists
}
