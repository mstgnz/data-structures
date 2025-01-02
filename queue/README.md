# Queue Data Structures Package

This package provides implementations of queue data structures in Go. It includes both array-based and linked list-based queue implementations, each optimized for different use cases.

## Features

### Array Queue
- Fixed-size circular array implementation
- Efficient memory usage
- Fast random access
- FIFO (First-In-First-Out) operations
- Automatic resizing capability

### Linked List Queue
- Dynamic size implementation
- No size limitations
- Memory efficient for variable size
- FIFO operations
- Optimal for frequent enqueue/dequeue

### Common Features
- Thread-safe operations (optional)
- Generic type support
- Clear operation
- Size tracking
- Empty/Full state checking

## Usage Examples

### Array Queue
```go
// Create a new array queue with initial capacity
queue := NewArrayQueue(5)

// Enqueue elements
queue.Enqueue(1)
queue.Enqueue(2)
queue.Enqueue(3)

// Dequeue elements
first, err := queue.Dequeue() // Returns 1
if err == nil {
    fmt.Println(first)
}

// Check size
size := queue.Size() // Returns 2

// Peek at front element
front, err := queue.Peek()
if err == nil {
    fmt.Println(front) // Shows 2
}
```

### Linked List Queue
```go
// Create a new linked list queue
queue := NewLinkedListQueue()

// Add elements
queue.Enqueue("first")
queue.Enqueue("second")
queue.Enqueue("third")

// Remove and process elements
for !queue.IsEmpty() {
    element, _ := queue.Dequeue()
    fmt.Println(element)
}

// Check if queue is empty
isEmpty := queue.IsEmpty() // Returns true
```

### Generic Type Usage
```go
type CustomType struct {
    ID   int
    Name string
}

// Create queue with custom type
queue := NewArrayQueue(10)
queue.Enqueue(CustomType{1, "Item 1"})
queue.Enqueue(CustomType{2, "Item 2"})

// Process custom type elements
item, err := queue.Dequeue()
if err == nil {
    customItem := item.(CustomType)
    fmt.Printf("ID: %d, Name: %s\n", customItem.ID, customItem.Name)
}
```

## Implementation Details

### Time Complexities

#### Array Queue
- Enqueue: O(1) amortized
- Dequeue: O(1)
- Peek: O(1)
- Clear: O(1)
- Size: O(1)

#### Linked List Queue
- Enqueue: O(1)
- Dequeue: O(1)
- Peek: O(1)
- Clear: O(1)
- Size: O(1)

### Space Complexities
- Array Queue: O(n) where n is the capacity
- Linked List Queue: O(n) where n is the number of elements

### Performance Characteristics

#### Array Queue
- Fixed memory allocation
- Cache-friendly
- Predictable performance
- Efficient for fixed-size scenarios

#### Linked List Queue
- Dynamic memory allocation
- Better for unpredictable sizes
- No reallocation needed
- Memory efficient for varying sizes

## Use Cases

### Array Queue
- Fixed-size buffers
- Circular buffers
- Resource pools
- Message queues with known bounds
- Real-time systems

### Linked List Queue
- Task scheduling
- Event handling
- Message queues with unknown bounds
- Stream processing
- Dynamic workload management

## Best Practices

### Choosing Queue Type
- Use Array Queue when:
  - Maximum size is known
  - Memory efficiency is critical
  - Frequent random access needed
  
- Use Linked List Queue when:
  - Size is unpredictable
  - Dynamic growth is required
  - Memory overhead is acceptable

### Performance Optimization
- Initialize Array Queue with appropriate capacity
- Monitor queue size for potential resizing
- Consider memory fragmentation
- Use appropriate queue type for access pattern

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
- Array Queue Enqueue: ~20ns
- Array Queue Dequeue: ~15ns
- Linked List Queue Enqueue: ~40ns
- Linked List Queue Dequeue: ~30ns

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