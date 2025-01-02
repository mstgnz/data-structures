# Advanced Data Structures Package

This package provides implementations of advanced data structures in Go, focusing on specialized use cases and optimized performance characteristics.

## Features

### LRU Cache
- Least Recently Used (LRU) cache implementation
- Generic key-value storage with interface{} type
- Operations:
  - Put: Add or update key-value pairs
  - Get: Retrieve values with LRU update
  - Contains: Check key existence
  - Clear: Remove all entries
- Internal structure:
  - Doubly linked list for order maintenance
  - Hash map for O(1) lookups
  - Automatic capacity management

### Skip List
- Probabilistic data structure for efficient searching
- Configurable:
  - Maximum level (default: 16)
  - Probability factor (default: 0.5)
- Operations:
  - Insert: Add key-value pairs
  - Search: Find values by key
  - Delete: Remove key-value pairs
- Features:
  - Random level generation
  - Multiple layer links
  - O(log n) average time complexity

### Disjoint Set (Union-Find)
- Implementation with path compression and union by rank
- Operations:
  - Find: Get set representative
  - Union: Merge two sets
  - Connected: Check if elements are in same set
- Features:
  - Path compression optimization
  - Union by rank optimization
  - Set size tracking
  - Set count tracking

## Usage Examples

### LRU Cache
```go
// Create a new LRU cache with capacity 3
cache := NewLRUCache(3)

// Add elements
cache.Put("key1", "value1")
cache.Put("key2", "value2")
cache.Put("key3", "value3")

// Get element (moves to most recently used)
value, exists := cache.Get("key1")
if exists {
    fmt.Println(value) // Outputs: value1
}

// Add new element when at capacity (removes least recently used)
cache.Put("key4", "value4") // "key2" will be removed if it's least recently used

// Check if key exists
exists = cache.Contains("key1") // returns: true
```

### Skip List
```go
// Create a new skip list
list := NewSkipList()

// Insert key-value pairs
list.Insert(10, "value1")
list.Insert(20, "value2")
list.Insert(5, "value3")

// Search for values
value, found := list.Search(20)
if found {
    fmt.Println(value) // Outputs: value2
}

// Delete elements
deleted := list.Delete(10) // returns: true
```

### Disjoint Set
```go
// Create a new disjoint set with size 5
ds := NewDisjointSet(5)

// Merge sets containing elements
ds.Union(0, 1)
ds.Union(2, 3)
ds.Union(1, 2)

// Check if elements are connected
connected := ds.Connected(0, 3) // returns: true

// Get total number of sets
setCount := ds.GetSetCount()

// Find set representative
root := ds.Find(0)
```

## Implementation Details

### Performance Characteristics

#### LRU Cache
- Get: O(1)
- Put: O(1)
- Space complexity: O(capacity)

#### Skip List
- Search: O(log n) average case
- Insert: O(log n) average case
- Delete: O(log n) average case
- Space complexity: O(n log n) average case

#### Disjoint Set
- Find: O(α(n)) amortized
- Union: O(α(n)) amortized
- Connected: O(α(n)) amortized
- Space complexity: O(n)

Where α(n) is the inverse Ackermann function, which grows extremely slowly.

### Memory Usage
- LRU Cache: Uses doubly linked list nodes and hash map entries
- Skip List: Uses nodes with variable-length forward arrays
- Disjoint Set: Uses two arrays for parent pointers and ranks

## Testing
Each data structure comes with comprehensive test coverage. Run tests using:
```bash
go test ./...
```

## Contributing
Contributions are welcome! Please ensure that any new features or modifications come with:
- Proper documentation
- Comprehensive test cases
- Example usage
- Performance considerations

## License
This package is distributed under the MIT license. See the LICENSE file for more details. 