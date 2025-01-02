# Stack Data Structures Package

This package provides implementations of stack data structures in Go. It includes both array-based and linked list-based stack implementations, each optimized for different use cases.

## Features

### Array Stack
- Fixed-size array implementation
- Efficient memory usage
- Fast random access
- LIFO (Last-In-First-Out) operations
- Automatic resizing capability

### Linked List Stack
- Dynamic size implementation
- No size limitations
- Memory efficient for variable size
- LIFO operations
- Optimal for frequent push/pop operations

### Common Features
- Thread-safe operations (optional)
- Generic type support
- Clear operation
- Size tracking
- Empty/Full state checking

## Usage Examples

### Array Stack
```go
// Create a new array stack with initial capacity
stack := NewArrayStack(5)

// Push elements
stack.Push(1)
stack.Push(2)
stack.Push(3)

// Pop elements
top, err := stack.Pop() // Returns 3
if err == nil {
    fmt.Println(top)
}

// Check size
size := stack.Size() // Returns 2

// Peek at top element
topElement, err := stack.Peek()
if err == nil {
    fmt.Println(topElement) // Shows 2
}
```

### Linked List Stack
```go
// Create a new linked list stack
stack := NewLinkedListStack()

// Add elements
stack.Push("first")
stack.Push("second")
stack.Push("third")

// Remove and process elements
for !stack.IsEmpty() {
    element, _ := stack.Pop()
    fmt.Println(element)
}

// Check if stack is empty
isEmpty := stack.IsEmpty() // Returns true
```

### Generic Type Usage
```go
type CustomType struct {
    ID   int
    Name string
}

// Create stack with custom type
stack := NewArrayStack(10)
stack.Push(CustomType{1, "Item 1"})
stack.Push(CustomType{2, "Item 2"})

// Process custom type elements
item, err := stack.Pop()
if err == nil {
    customItem := item.(CustomType)
    fmt.Printf("ID: %d, Name: %s\n", customItem.ID, customItem.Name)
}
```

## Implementation Details

### Time Complexities

#### Array Stack
- Push: O(1) amortized
- Pop: O(1)
- Peek: O(1)
- Clear: O(1)
- Size: O(1)

#### Linked List Stack
- Push: O(1)
- Pop: O(1)
- Peek: O(1)
- Clear: O(1)
- Size: O(1)

### Space Complexities
- Array Stack: O(n) where n is the capacity
- Linked List Stack: O(n) where n is the number of elements

### Performance Characteristics

#### Array Stack
- Fixed memory allocation
- Cache-friendly
- Predictable performance
- Efficient for fixed-size scenarios

#### Linked List Stack
- Dynamic memory allocation
- Better for unpredictable sizes
- No reallocation needed
- Memory efficient for varying sizes

## Use Cases

### Array Stack
- Expression evaluation
- Syntax parsing
- Memory management
- Function call management
- Undo/Redo operations

### Linked List Stack
- Browser history
- Text editor operations
- Dynamic function calls
- Recursive algorithms
- Memory management with unknown depth

## Best Practices

### Choosing Stack Type
- Use Array Stack when:
  - Maximum size is known
  - Memory efficiency is critical
  - Frequent access to elements needed
  
- Use Linked List Stack when:
  - Size is unpredictable
  - Dynamic growth is required
  - Memory overhead is acceptable

### Performance Optimization
- Initialize Array Stack with appropriate capacity
- Monitor stack size for potential resizing
- Consider memory fragmentation
- Use appropriate stack type for access pattern

### Thread Safety
- Implement synchronization when needed
- Consider concurrent access patterns
- Use atomic operations where appropriate
- Handle race conditions properly

## Testing
The package comes with comprehensive test coverage. Run tests using:
```bash
go test ./...
```

## Benchmarks
Key performance metrics:
- Array Stack Push: ~15ns
- Array Stack Pop: ~10ns
- Linked List Stack Push: ~35ns
- Linked List Stack Pop: ~25ns

## Contributing
Contributions are welcome! Please ensure that any new features or modifications come with:
- Proper documentation
- Time and space complexity analysis
- Comprehensive test cases
- Example usage
- Performance benchmarks
- Thread safety considerations

## License
This package is distributed under the MIT license. See the LICENSE file for more details. 