# Stack Package

This package provides two different stack implementations in Go: Array-based and LinkedList-based. All implementations are designed to be thread-safe.

## Features

### Core Structures
- Array Stack (Dynamic array-based implementation)
- LinkedList Stack (Linked list-based implementation)

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

## Usage Examples

### Array Stack
```go
// Create a new stack
stack := NewArrayStack()

// Add elements
stack.Push(1)
stack.Push(2)
stack.Push(3)

// Remove element from top
stack.Pop()

// Check if empty
isEmpty := stack.IsEmpty()

// List elements
elements := stack.List()
stack.Print()
```

### LinkedList Stack
```go
// Create a new stack
stack := NewLinkedListStack(0)

// Add elements
stack.Push(1)
stack.Push(2)
stack.Push(3)

// Remove element from top
stack.Pop()

// Check if empty
isEmpty := stack.IsEmpty()

// List elements
elements := stack.List()
stack.Print()
```

## Implementation Details

### Data Structures

#### Array Stack
- Dynamic array-based implementation
- Auto-resizing capability (grows and shrinks)
- Efficient memory management
- Index tracking for top element

#### LinkedList Stack
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