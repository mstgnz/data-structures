# Heap Package

This package provides thread-safe implementations of various heap data structures in Go, including binary heaps, advanced heaps, and a priority queue.

## Features

### Basic Heaps
- MinHeap: Root is the minimum element
- MaxHeap: Root is the maximum element
- Common operations:
  - Insert: Add new element
  - Extract: Remove root element
  - Peek: View root element
  - Size/IsEmpty checks

### Advanced Heaps
- Fibonacci Heap:
  - Amortized O(1) for insert and decrease-key
  - Efficient merge operations
  - Marked nodes for cascading cuts
- Binomial Heap:
  - Tree-based implementation
  - Efficient merge operations
  - Logarithmic height guarantee
- Leftist Heap:
  - Self-adjusting structure
  - Efficient merge operations
  - Path length optimization
- Skew Heap:
  - Self-adjusting structure
  - Simpler than leftist heap
  - Probabilistic balance

### Priority Queue
- Generic value storage with interface{} type
- Priority-based ordering
- FIFO behavior for equal priorities
- Thread-safe operations
- Core operations:
  - Enqueue: Add with priority
  - Dequeue: Remove highest priority
  - Peek: View highest priority

## Usage Examples

### Basic Heap Operations
```go
// MinHeap
minHeap := NewMinHeap()
minHeap.Insert(5)
minHeap.Insert(3)
min, _ := minHeap.Extract() // returns 3

// MaxHeap
maxHeap := NewMaxHeap()
maxHeap.Insert(5)
maxHeap.Insert(3)
max, _ := maxHeap.Extract() // returns 5
```

### Fibonacci Heap
```go
fibHeap := NewFibonacciHeap()
fibHeap.Insert(10)
fibHeap.Insert(5)
fibHeap.Insert(15)

min, _ := fibHeap.Extract() // returns 5
value, _ := fibHeap.Peek()  // returns 10
```

### Priority Queue
```go
pq := NewPriorityQueue()

// Add items with priorities
pq.Enqueue("high", 3)
pq.Enqueue("medium", 2)
pq.Enqueue("low", 1)

// Get highest priority item
value, _ := pq.Dequeue() // returns "high"
```

## Implementation Details

### Time Complexities

#### Basic Heaps (Min/Max)
- Insert: O(log n)
- Extract: O(log n)
- Peek: O(1)
- Size/IsEmpty: O(1)

#### Fibonacci Heap
- Insert: O(1)
- Extract-Min: O(log n) amortized
- Decrease-Key: O(1) amortized
- Merge: O(1)

#### Priority Queue
- Enqueue: O(log n)
- Dequeue: O(log n)
- Peek: O(1)

### Space Complexities
- Binary Heaps: O(n)
- Fibonacci Heap: O(n)
- Priority Queue: O(n)

### Thread Safety
- All operations are protected with RWMutex
- Read operations use RLock
- Write operations use Lock
- Proper lock/unlock handling with defer

### Special Features
1. Fibonacci Heap:
   - Lazy merging of trees
   - Marked nodes for optimizing decrease-key
   - Degree-based consolidation

2. Priority Queue:
   - FIFO ordering for equal priorities
   - Generic value storage
   - Index tracking for stable sorting

## Testing
Each heap implementation comes with comprehensive test coverage. Run tests using:
```bash
go test ./...
```

## Contributing
Contributions are welcome! Please ensure that any new features or modifications come with:
- Proper documentation
- Thread safety considerations
- Comprehensive test cases
- Example usage
- Time complexity analysis

## License
This package is distributed under the MIT license. See the LICENSE file for more details. 