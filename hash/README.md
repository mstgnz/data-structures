# Hash Package

This package provides thread-safe implementations of hash-based data structures in Go, including hash tables with different collision resolution strategies and Bloom filters.

## Features

### Hash Table
- Generic key-value storage with interface{} type
- Multiple collision resolution strategies:
  - Linear probing
  - Chaining with linked lists
- Thread-safe operations with RWMutex
- Core operations:
  - Put: Insert or update key-value pairs
  - Get: Retrieve values by key
  - Remove: Delete key-value pairs
- Dynamic sizing and load factor management
- Support for different key types (string, int)

### Bloom Filter
- Space-efficient probabilistic data structure
- Configurable parameters:
  - Expected number of elements
  - Desired false positive rate
- Multiple hash functions using FNV-1a
- Operations:
  - Add: Insert elements
  - Contains: Test membership
  - Clear: Reset filter
  - EstimateFalsePositiveRate: Calculate current FPR
- Automatic optimization:
  - Optimal bit array size calculation
  - Optimal hash function count calculation

## Usage Examples

### Hash Table
```go
// Create a new hash table with linear probing
ht := NewHashTable(100, "linear")

// Insert key-value pairs
ht.Put("key1", "value1")
ht.Put("key2", "value2")
ht.Put(42, "numeric key")

// Retrieve values
value, exists := ht.Get("key1")
if exists {
    fmt.Println(value) // Outputs: value1
}

// Remove entries
removed := ht.Remove("key2")

// Create a hash table with chaining
htChain := NewHashTable(100, "chain")
htChain.Put("key1", "value1")
```

### Bloom Filter
```go
// Create a new Bloom filter
// Expected elements: 1000, False positive rate: 0.01
bf := NewBloomFilter(1000, 0.01)

// Add elements
bf.Add([]byte("element1"))
bf.Add([]byte("element2"))

// Check membership
exists := bf.Contains([]byte("element1")) // true
exists = bf.Contains([]byte("element3"))  // false (probably)

// Get current false positive rate
fpr := bf.EstimateFalsePositiveRate(500) // for 500 inserted elements

// Reset the filter
bf.Clear()
```

## Implementation Details

### Hash Table

#### Collision Resolution
1. Linear Probing:
   - Open addressing with linear search
   - Good cache performance
   - Susceptible to clustering

2. Chaining:
   - Separate chaining with linked lists
   - Better for high load factors
   - More memory overhead

#### Time Complexities
- Average Case:
  - Insert: O(1)
  - Lookup: O(1)
  - Delete: O(1)
- Worst Case (with collisions):
  - Linear Probing: O(n)
  - Chaining: O(n)

### Bloom Filter

#### Design Considerations
- Bit array size optimization
- Number of hash functions optimization
- Thread-safe operations
- FNV-1a hash function implementation

#### Space and Time Complexities
- Space: O(m) where m is the bit array size
- Time:
  - Add: O(k) where k is the number of hash functions
  - Contains: O(k)
  - Clear: O(m)

#### False Positive Rate
The false positive rate (p) is approximately:
```
p = (1 - e^(-kn/m))^k
```
where:
- k is the number of hash functions
- n is the number of inserted elements
- m is the size of the bit array

## Thread Safety
- All operations are protected with RWMutex
- Read operations use RLock
- Write operations use Lock
- Proper lock/unlock handling with defer

## Testing
Each data structure comes with comprehensive test coverage. Run tests using:
```bash
go test ./...
```

## Contributing
Contributions are welcome! Please ensure that any new features or modifications come with:
- Proper documentation
- Thread safety considerations
- Comprehensive test cases
- Example usage
- Performance analysis

## License
This package is distributed under the MIT license. See the LICENSE file for more details. 