# Queue Package

This package provides two different queue implementations in Go: Array-based and LinkedList-based. All implementations are designed to be thread-safe.

## Features

### Core Structures
- Array Queue (Dynamic array-based implementation)
- LinkedList Queue (Linked list-based implementation)

### Common Operations
- Enqueue: Add element to the queue
- Dequeue: Remove element from the queue
- List: Get all elements
- Print: Display elements

### Thread Safety
- Safe read/write operations with RWMutex
- Concurrent access support for all structures
- Deadlock prevention mechanisms

## Usage Examples

### Array Queue
```go
// Create a new queue
queue := NewArrayQueue()

// Add elements
queue.Enqueue(1)
queue.Enqueue(2)
queue.Enqueue(3)

// Remove element from front
queue.Dequeue()

// List elements
elements := queue.List()
queue.Print()
```

### LinkedList Queue
```go
// Create a new queue
queue := NewLinkedListQueue(0)

// Add elements
queue.Enqueue(1)
queue.Enqueue(2)
queue.Enqueue(3)

// Remove element from front
queue.Dequeue()

// List elements
elements := queue.List()
queue.Print()
```

## Implementation Details

### Data Structures

#### Array Queue
- Dynamic array-based implementation
- Auto-resizing capability (grows and shrinks)
- Efficient memory management with reordering
- First and last index tracking

#### LinkedList Queue
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