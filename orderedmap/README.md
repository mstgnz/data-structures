# OrderedMap Package

This package provides a thread-safe implementation of an ordered map data structure in Go, maintaining insertion order while providing fast key-based access.

## Features

### Core Functionality
- Generic key-value storage using `any` type
- Maintains insertion order of elements
- Thread-safe operations with RWMutex
- Fast key-based lookups using internal index map

### Operations
- Set: Add or update key-value pairs
- Get: Retrieve values by key
- Delete: Remove key-value pairs
- Clear: Remove all elements
- Copy: Create a deep copy
- Has: Check key existence

### Order-Aware Operations
- Keys: Get all keys in insertion order
- Values: Get all values in insertion order
- Range: Iterate over pairs in order
- String: Get ordered string representation

### Thread Safety
- All operations are protected with RWMutex
- Read operations use RLock
- Write operations use Lock
- Proper lock/unlock handling with defer

## Usage Examples

### Basic Operations
```go
// Create a new ordered map
om := NewOrderedMap()

// Add key-value pairs
om.Set("first", 1)
om.Set("second", 2)
om.Set("third", 3)

// Get value by key
value, exists := om.Get("second")
if exists {
    fmt.Println(value) // Outputs: 2
}

// Check if key exists
exists = om.Has("first") // returns true

// Delete a key-value pair
om.Delete("second")
```

### Order-Aware Operations
```go
// Get all keys in order
keys := om.Keys() // ["first", "third"]

// Get all values in order
values := om.Values() // [1, 3]

// Iterate over pairs in order
om.Range(func(key, value any) bool {
    fmt.Printf("%v: %v\n", key, value)
    return true // continue iteration
})

// Get string representation
str := om.String() // "{first: 1, third: 3}"
```

### Advanced Operations
```go
// Create a copy
copy := om.Copy()

// Clear all elements
om.Clear()

// Get number of elements
size := om.Len()
```

## Implementation Details

### Data Structure
- Internal slice for maintaining order
- Hash map for fast key lookups
- Pair structure for key-value storage

### Time Complexities
- Set: O(1) average
- Get: O(1)
- Delete: O(1) average
- Has: O(1)
- Keys/Values: O(n)
- Range: O(n)
- Clear: O(1)
- Copy: O(n)
- Len: O(1)

### Space Complexity
- O(n) where n is the number of elements
- Additional O(n) for the index map

### Memory Management
- Efficient memory usage with slice and map combination
- Automatic cleanup of deleted elements
- No memory leaks in circular references

### Thread Safety Details
- Read operations can occur concurrently
- Write operations are serialized
- Safe for concurrent access from multiple goroutines
- Deadlock prevention with deferred mutex unlocks

## Testing
The package comes with comprehensive test coverage. Run tests using:
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