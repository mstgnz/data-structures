# Stack Package

This package provides two different stack implementations in Go: Array-based and LinkedList-based. All implementations are designed to be thread-safe and support generic types.

## Features

### Core Structures
- Array Stack (Dynamic array-based implementation with generic type support)
- LinkedList Stack (Linked list-based implementation with generic type support)

### Common Operations
- Push: Add element to the stack
- Pop: Remove element from the stack
- IsEmpty: Check if stack is empty
- List: Get all elements
- Print: Display elements

### Thread Safety
- Safe read/write operations with RWMutex
- Concurrent access support for all structures
- Deadlock prevention mechanisms

### Generic Type Support
- Support for any comparable type
- Type-safe operations
- Custom type support with proper comparison functions

## Usage Examples

### Array Stack
```go
// Create a new integer stack
intStack := NewArrayStack[int]()

// Add elements
intStack.Push(1)
intStack.Push(2)
intStack.Push(3)

// Remove element from top
intStack.Pop()

// Check if empty
isEmpty := intStack.IsEmpty()

// List elements
elements := intStack.List()
intStack.Print()

// Create a new string stack
strStack := NewArrayStack[string]()

// Add elements
strStack.Push("a")
strStack.Push("b")
strStack.Push("c")

// Remove element from top
strStack.Pop()

// List elements
elements = strStack.List()
strStack.Print()

// Create a stack with custom type
type Person struct {
    Name string
    Age  int
}

personStack := NewArrayStack[Person]()
personStack.Push(Person{Name: "John", Age: 30})
personStack.Push(Person{Name: "Jane", Age: 25})
```

### LinkedList Stack
```go
// Create a new integer stack
intStack := NewLinkedListStack[int](0)

// Add elements
intStack.Push(1)
intStack.Push(2)
intStack.Push(3)

// Remove element from top
intStack.Pop()

// Check if empty
isEmpty := intStack.IsEmpty()

// List elements
elements := intStack.List()
intStack.Print()

// Create a new string stack
strStack := NewLinkedListStack[string]("")

// Add elements
strStack.Push("a")
strStack.Push("b")
strStack.Push("c")

// Remove element from top
strStack.Pop()

// List elements
elements = strStack.List()
strStack.Print()

// Create a stack with custom type
type Person struct {
    Name string
    Age  int
}

personStack := NewLinkedListStack[Person](Person{Name: "", Age: 0})
personStack.Push(Person{Name: "John", Age: 30})
personStack.Push(Person{Name: "Jane", Age: 25})
```

## Implementation Details

### Data Structures

#### Array Stack
- Generic type support with comparable constraint
- Dynamic array-based implementation
- Auto-resizing capability (grows and shrinks)
- Efficient memory management
- Index tracking for top element

#### LinkedList Stack
- Generic type support with comparable constraint
- Node-based implementation
- Dynamic memory allocation
- LIFO (Last-In-First-Out) structure
- No size limitations

### Time Complexities

#### Array Stack
- Push: O(1) amortized, O(n) worst case when resizing
- Pop: O(1) amortized, O(n) worst case when shrinking
- IsEmpty: O(1)
- List: O(n)
- Space: O(n)

#### LinkedList Stack
- Push: O(1)
- Pop: O(1)
- IsEmpty: O(1)
- List: O(n)
- Space: O(n)

### Memory Management

#### Array Stack
- Dynamic array resizing (doubles when full)
- Array shrinking (halves when 1/4 full)
- Efficient memory utilization
- Automatic capacity management

#### LinkedList Stack
- Dynamic node allocation
- No pre-allocated memory
- Memory freed on pop
- No explicit size limitations

### Thread Safety Details
- RLock for read operations (IsEmpty, List, Print)
- Lock for write operations (Push, Pop)
- Automatic unlock with defer
- Safe design for concurrent access

### Generic Type Constraints
- Types must satisfy the `comparable` interface
- Support for built-in types (int, string, etc.)
- Support for custom types that implement `comparable`
- Type safety at compile time

## Testing
The package comes with comprehensive test coverage for various types. To run tests:
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
- Generic type support considerations

## License
This package is distributed under the MIT license. See the LICENSE file for more details. 