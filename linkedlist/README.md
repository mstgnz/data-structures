# LinkedList Package

This package provides three different linked list implementations in Go: Linear (Singly), Double (Doubly), and Circular. All implementations are designed to be thread-safe.

## Features

### Core Structures
- Linear (Singly) Linked List
- Double (Doubly) Linked List
- Circular Linked List

### Common Operations
- AddToStart: Add element to the beginning
- AddToSequentially: Add element in sorted order
- AddToAfter: Add element after a specific value
- AddToEnd: Add element to the end
- Delete: Remove element
- Search: Find element
- List: Get all elements
- Print: Display elements

### Thread Safety
- Safe read/write operations with RWMutex
- Concurrent access support for all structures
- Deadlock prevention mechanisms

## Usage Examples

### Linear (Singly) List
```go
// Create a new list
list := NewLinear(1)

// Add elements
list.AddToStart(0)      // Add to beginning
list.AddToEnd(3)        // Add to end
list.AddToAfter(2, 1)   // Add 2 after 1
list.AddToSequentially(4) // Add in sorted order

// Delete element
list.Delete(2)

// Search element
exists := list.Search(3)

// List elements
elements := list.List()
list.Print()
```

### Double (Doubly) List
```go
// Create a new list
list := NewDouble(1)

// Add elements
list.AddToStart(0)
list.AddToEnd(2)
list.AddToSequentially(1.5)

// Forward and backward listing
forward := list.List(false)  // [0, 1, 1.5, 2]
backward := list.List(true)  // [2, 1.5, 1, 0]

// Print
list.Print(false)  // Forward direction
list.Print(true)   // Backward direction
```

### Circular List
```go
// Create a new list
list := NewCircular(1)

// Add elements
list.AddToStart(0)
list.AddToEnd(2)
list.AddToSequentially(1.5)

// List in circular structure
elements := list.List()  // Last element connects back to first

// Print
list.Print()
```

## Implementation Details

### Data Structures
- Linear: Single next pointer
- Double: Both next and prev pointers
- Circular: Last element connected to first with next pointer

### Time Complexities

#### Linear and Double List
- AddToStart: O(1)
- AddToEnd: O(n)
- AddToSequentially: O(n)
- AddToAfter: O(n)
- Delete: O(n)
- Search: O(n)
- List: O(n)

#### Circular List
- AddToStart: O(1)
- AddToEnd: O(n)
- AddToSequentially: O(n)
- AddToAfter: O(n)
- Delete: O(n)
- List: O(n)

### Thread Safety Details
- RLock for read operations
- Lock for write operations
- Automatic unlock with defer
- Safe design for concurrent access

### Memory Management
- Dynamic node creation and deletion
- Pointer management
- Circular reference prevention
- Memory leak prevention mechanisms

## Testing
The package comes with comprehensive test coverage. To run tests:
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