# Linked List Data Structures Package

This package provides comprehensive implementations of various linked list data structures in Go. It includes different types of linked lists optimized for various use cases and requirements.

## Features

### Linear Linked List
- Singly linked list implementation
- Basic operations: insert, delete, search
- Traversal and manipulation utilities
- Memory efficient for sequential access

### Double Linked List
- Doubly linked list implementation
- Bidirectional traversal support
- Enhanced deletion operations
- Efficient for both forward and backward traversal

### Circular Linked List
- Circular linked list implementation
- Continuous cyclic traversal
- Efficient for circular buffer implementations
- Support for both singly and doubly linked variants

## Usage Examples

### Linear Linked List
```go
// Create a new linear linked list
list := NewLinear()

// Add elements
list.Append(1)
list.Append(2)
list.Append(3)

// Insert at specific position
list.Insert(1, 4) // Insert 4 after first element

// Remove elements
list.Remove(2) // Remove element at index 2

// Search for elements
found := list.Search(4)
```

### Double Linked List
```go
// Create a new double linked list
list := NewDouble()

// Add elements at both ends
list.Append(1)
list.Prepend(0)
list.Append(2)

// Traverse in both directions
forward := list.Forward()  // [0, 1, 2]
backward := list.Backward() // [2, 1, 0]

// Remove from either end
list.RemoveFirst()
list.RemoveLast()
```

### Circular Linked List
```go
// Create a new circular linked list
list := NewCircular()

// Add elements
list.Append(1)
list.Append(2)
list.Append(3)

// Rotate the list
list.Rotate(1) // Rotates one position forward

// Traverse the entire circle
elements := list.Traverse() // Returns to starting point
```

## Implementation Details

### Time Complexities

#### Linear Linked List
- Insert at beginning: O(1)
- Insert at end: O(1) with tail pointer, O(n) without
- Insert at position: O(n)
- Delete: O(n)
- Search: O(n)

#### Double Linked List
- Insert at beginning: O(1)
- Insert at end: O(1)
- Insert at position: O(n)
- Delete: O(1) with node reference
- Reverse traversal: O(n)

#### Circular Linked List
- Insert at beginning: O(1)
- Insert at end: O(1) with tail pointer
- Rotate: O(1)
- Search: O(n)

### Space Complexities
- Linear Linked List: O(n)
- Double Linked List: O(n)
- Circular Linked List: O(n)

## Use Cases

### Linear Linked List
- Simple sequential data storage
- Stack and queue implementations
- Memory efficient list operations
- Dynamic size requirements

### Double Linked List
- Browser history implementation
- Undo/Redo functionality
- Music player playlists
- Text editors

### Circular Linked List
- Round-robin scheduling
- Circular buffer implementation
- Game turn management
- Repeating playlist implementation

## Best Practices

### Choosing the Right List Type
- Use Linear List for simple sequential access
- Use Double List when bidirectional traversal is needed
- Use Circular List for cyclic data structures

### Performance Optimization
- Keep track of tail pointer for O(1) append operations
- Use appropriate list type based on traversal requirements
- Consider memory overhead of additional pointers
- Implement proper cleanup to prevent memory leaks

### Memory Management
- Properly handle node deletion
- Clear references in removed nodes
- Implement proper cleanup methods
- Consider garbage collection implications

## Testing
Each list implementation comes with comprehensive test coverage. Run tests using:
```bash
go test ./...
```

## Contributing
Contributions are welcome! Please ensure that any new features or modifications come with:
- Proper documentation
- Time and space complexity analysis
- Comprehensive test cases
- Example usage
- Memory management considerations

## License
This package is distributed under the MIT license. See the LICENSE file for more details. 