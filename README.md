# Data Structures With Go

This repository contains various data structures implemented in the Go programming language.

## Table of Contents
- [Installation](#installation)
- [Data Structures](#data-structures)
  - [Linked List](#linked-list)
  - [Queue](#queue)
  - [Stack](#stack)
  - [Tree](#tree)
- [Performance Comparisons](#performance-comparisons)
- [Benchmark Results](#benchmark-results)
- [Big O Complexity](#big-o-complexity)
- [Usage Examples](#usage-examples)
- [Contributing](#contributing)
- [License](#license)

## Installation

```bash
go get github.com/yourusername/data-structures
```

## Data Structures

### Linked List

Three different Linked List implementations are available:

#### 1. Linear Linked List
![Linear Linked List](img/LinearLinkedList.png)
```go
type linear struct {
    Data int
    Next *linear
}
```

**Usage Example:**
```go
list := &Linear{}
list.Add(1)
list.Add(2)
list.Add(3)
value := list.Get(1) // returns 2
list.Remove(1)       // removes 2 from the list
```

**Time Complexity:**
- Insertion: O(1) - at the beginning
- Deletion: O(n)
- Search: O(n)

#### 2. Circular Linked List
![Circular Linked List](img/CircularLinkedList.png)
```go
type circular struct {
    Data int
    Next *circular
}
```

**Usage Example:**
```go
list := &Circular{}
list.Add(1)
list.Add(2)
list.Add(3)
list.Display() // 1 -> 2 -> 3 -> 1
```

#### 3. Double Linked List
![Double Linked List](img/DoubleLinkedList.png)
```go
type double struct {
    Data int
    Next *double
    Prev *double
}
```

**Usage Example:**
```go
list := &Double{}
list.Add(1)
list.Add(2)
list.AddToEnd(3)
list.DisplayReverse() // 3 -> 2 -> 1
```

### Queue

Two different Queue implementations are available:

#### 1. Array Queue
```go
type arrayQueue struct {
    Arr []int
    ArrSize int
    FirstIndex int
    LastIndex int
}
```

**Usage Example:**
```go
queue := NewArrayQueue(5)
queue.Enqueue(1)
queue.Enqueue(2)
value := queue.Dequeue() // returns 1
```

#### 2. Linked List Queue
```go
type linkedListQueue struct {
    X int
    Next *linkedListQueue
}
```

**Usage Example:**
```go
queue := NewLinkedListQueue()
queue.Enqueue(1)
queue.Enqueue(2)
value := queue.Dequeue() // returns 1
```

### Stack

Two different Stack implementations are available:

#### 1. Array Stack
```go
type arrayStack struct {
    Arr []int
    ArrSize int
    Index int
}
```

**Usage Example:**
```go
stack := NewArrayStack(5)
stack.Push(1)
stack.Push(2)
value := stack.Pop() // returns 2
```

#### 2. Linked List Stack
```go
type linkedListStack struct {
    X int
    Next *linkedListStack
}
```

**Usage Example:**
```go
stack := NewLinkedListStack()
stack.Push(1)
stack.Push(2)
value := stack.Pop() // returns 2
```

## Performance Comparisons

### Queue Performance Comparison
| Operation | Array Queue | Linked List Queue |
|-----------|-------------|-------------------|
| Enqueue   | O(1)        | O(1)              |
| Dequeue   | O(1)        | O(1)              |
| Peek      | O(1)        | O(1)              |
| Memory    | Fixed       | Dynamic           |

### Stack Performance Comparison
| Operation | Array Stack | Linked List Stack |
|-----------|-------------|-------------------|
| Push      | O(1)        | O(1)              |
| Pop       | O(1)        | O(1)              |
| Peek      | O(1)        | O(1)              |
| Memory    | Fixed       | Dynamic           |

## Benchmark Results

### Linked List Benchmarks
```
BenchmarkLinear_AddToStart-8    36861304     39.80 ns/op    16 B/op    1 allocs/op
BenchmarkLinear_AddToEnd-8        229119    115234 ns/op    16 B/op    1 allocs/op
BenchmarkLinear_Delete-8        14885366     81.15 ns/op    37 B/op    2 allocs/op
BenchmarkLinear_Search-8         2304028    525.5 ns/op      0 B/op    0 allocs/op
```

These benchmark results show:

1. **AddToStart**
   - Fastest operation (39.80 ns/op)
   - Constant memory usage (16 B/op)
   - Single allocation

2. **AddToEnd**
   - Slowest operation (115234 ns/op)
   - Constant memory usage (16 B/op)
   - Single allocation
   - O(n) complexity to reach the end of the list

3. **Delete**
   - Medium speed operation (81.15 ns/op)
   - Slightly higher memory usage (37 B/op)
   - Requires two allocations

4. **Search**
   - Medium-slow operation (525.5 ns/op)
   - No additional memory usage
   - No allocations required
   - O(n) complexity

### Queue Benchmarks
```
BenchmarkArrayQueue_Enqueue-8      220398138     6.193 ns/op    19 B/op    0 allocs/op
BenchmarkArrayQueue_Dequeue-8      189323732     6.082 ns/op    16 B/op    0 allocs/op
BenchmarkLinkedListQueue_Enqueue-8    209122   104090 ns/op     15 B/op    0 allocs/op
BenchmarkLinkedListQueue_Dequeue-8    812548     1459 ns/op     16 B/op    1 allocs/op
```

These benchmark results show:

1. **Array Queue**
   - Enqueue: Very fast (6.193 ns/op), low memory usage (19 B/op), no allocations
   - Dequeue: Fastest operation (6.082 ns/op), low memory usage (16 B/op), no allocations
   - Fast access due to fixed-size array usage
   - Has size limitation

2. **Linked List Queue**
   - Enqueue: Slower (104090 ns/op), low memory usage (15 B/op), no allocations
   - Dequeue: Medium speed (1459 ns/op), low memory usage (16 B/op), one allocation
   - Dynamic size advantage
   - Higher memory usage due to node structure

**Comparison:**
- Array Queue is fixed-size but faster and memory efficient
- Linked List Queue is dynamic-sized but significantly slower
- Prefer Array Queue for frequent add/remove operations
- Prefer Linked List Queue when size is variable and unknown beforehand

### Stack Benchmarks
```
BenchmarkArrayStack_Push-8          255688695     4.151 ns/op    16 B/op    0 allocs/op
BenchmarkArrayStack_Pop-8           271091098     4.384 ns/op     0 B/op    0 allocs/op
BenchmarkLinkedListStack_Push-8      34971062    41.36 ns/op    16 B/op    1 allocs/op
BenchmarkLinkedListStack_Pop-8       48903242    22.98 ns/op    16 B/op    1 allocs/op
```

These benchmark results show:

1. **Array Stack**
   - Push: Fastest operation (4.151 ns/op), low memory usage (16 B/op), no allocations
   - Pop: Very fast operation (4.384 ns/op), no memory usage, no allocations
   - Fast access due to fixed-size array usage
   - Has size limitation

2. **Linked List Stack**
   - Push: Slower (41.36 ns/op), low memory usage (16 B/op), one allocation
   - Pop: Medium speed (22.98 ns/op), low memory usage (16 B/op), one allocation
   - Dynamic size advantage
   - Higher memory usage due to node structure

**Comparison:**
- Array Stack is faster and memory efficient in all operations
- Linked List Stack is dynamic-sized but slower
- Prefer Array Stack for frequent push/pop operations
- Prefer Linked List Stack when size is variable and unknown beforehand

## Big O Complexity

### Linked List
| Operation           | Linear | Circular | Double |
|--------------------|--------|-----------|---------|
| Insert at Start    | O(1)   | O(1)      | O(1)    |
| Insert at End      | O(n)   | O(1)      | O(1)    |
| Delete at Start    | O(1)   | O(1)      | O(1)    |
| Delete at End      | O(n)   | O(n)      | O(1)    |
| Search             | O(n)   | O(n)      | O(n)    |
| Space Complexity   | O(n)   | O(n)      | O(n)    |

### Queue
| Operation           | Array Queue | Linked List Queue |
|--------------------|-------------|-------------------|
| Enqueue            | O(1)*       | O(1)              |
| Dequeue            | O(1)*       | O(1)              |
| Peek               | O(1)        | O(1)              |
| Space Complexity   | O(n)        | O(n)              |
| Memory Usage       | Contiguous  | Scattered         |

*Amortized time complexity. May require O(n) for resizing.

### Stack
| Operation           | Array Stack | Linked List Stack |
|--------------------|-------------|-------------------|
| Push               | O(1)*       | O(1)              |
| Pop                | O(1)        | O(1)              |
| Peek               | O(1)        | O(1)              |
| Space Complexity   | O(n)        | O(n)              |
| Memory Usage       | Contiguous  | Scattered         |

*Amortized time complexity. May require O(n) for resizing.

## Contributing

This project is open-source, and contributions are welcome. Feel free to contribute or provide feedback of any kind.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.