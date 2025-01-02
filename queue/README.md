# Queue Package

This package provides two different queue implementations in Go: Array-based and LinkedList-based. All implementations are designed to be thread-safe and support generic types.

## Features

### Core Structures
- Array Queue (Dynamic array-based implementation with generic type support)
- LinkedList Queue (Linked list-based implementation with generic type support)

### Common Operations
- Enqueue: Add element to the queue
- Dequeue: Remove element from the queue
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

### Array Queue
```go
// Create a new integer queue
intQueue := NewArrayQueue[int]()

// Add elements
intQueue.Enqueue(1)
intQueue.Enqueue(2)
intQueue.Enqueue(3)

// Remove element from front
intQueue.Dequeue()

// List elements
elements := intQueue.List()
intQueue.Print()

// Create a new string queue
strQueue := NewArrayQueue[string]()

// Add elements
strQueue.Enqueue("a")
strQueue.Enqueue("b")
strQueue.Enqueue("c")

// Remove element from front
strQueue.Dequeue()

// List elements
elements = strQueue.List()
strQueue.Print()

// Create a queue with custom type
type Person struct {
    Name string
    Age  int
}

personQueue := NewArrayQueue[Person]()
personQueue.Enqueue(Person{Name: "John", Age: 30})
personQueue.Enqueue(Person{Name: "Jane", Age: 25})
```

### LinkedList Queue
```go
// Create a new integer queue
intQueue := NewLinkedListQueue[int](0)

// Add elements
intQueue.Enqueue(1)
intQueue.Enqueue(2)
intQueue.Enqueue(3)

// Remove element from front
intQueue.Dequeue()

// List elements
elements := intQueue.List()
intQueue.Print()

// Create a new string queue
strQueue := NewLinkedListQueue[string]("")

// Add elements
strQueue.Enqueue("a")
strQueue.Enqueue("b")
strQueue.Enqueue("c")

// Remove element from front
strQueue.Dequeue()

// List elements
elements = strQueue.List()
strQueue.Print()

// Create a queue with custom type
type Person struct {
    Name string
    Age  int
}

personQueue := NewLinkedListQueue[Person](Person{Name: "", Age: 0})
personQueue.Enqueue(Person{Name: "John", Age: 30})
personQueue.Enqueue(Person{Name: "Jane", Age: 25})
```

## Implementation Details

### Data Structures

#### Array Queue
- Generic type support with comparable constraint
- Dynamic array-based implementation
- Auto-resizing capability (grows and shrinks)
- Efficient memory management with reordering
- First and last index tracking

#### LinkedList Queue
- Generic type support with comparable constraint
- Node-based implementation
- Dynamic memory allocation
- Single direction linking
- No size limitations

### Time Complexities

#### Array Queue
- Enqueue: O(1) amortized, O(n) worst case when resizing
- Dequeue: O(1) amortized, O(n) worst case when reordering
- List: O(n)
- Space: O(n)

#### LinkedList Queue
- Enqueue: O(n) - needs to traverse to end
- Dequeue: O(1)
- List: O(n)
- Space: O(n)

### Memory Management

#### Array Queue
- Dynamic array resizing (doubles when full)
- Array shrinking (halves when 1/4 full)
- Automatic reordering to optimize space
- Efficient memory utilization

#### LinkedList Queue
- Dynamic node allocation
- No pre-allocated memory
- Memory freed on dequeue
- No explicit size limitations

### Thread Safety Details
- RLock for read operations (List, Print)
- Lock for write operations (Enqueue, Dequeue)
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