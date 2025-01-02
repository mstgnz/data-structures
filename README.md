# Data Structures and Algorithms in Go

This repository contains implementations of various data structures and algorithms in Go programming language. It serves as both a learning resource and a practical reference for developers.

## 📚 Data Structures

### Basic Data Structures
- **Linked List** - [Documentation](linkedlist/README.md)
  - Singly Linked List
  - Doubly Linked List
  - Circular Linked List
  - Iterator Pattern Implementation
  - Thread-safe Operations
- **Stack** - [Documentation](stack/README.md)
  - Array-based Implementation
  - Linked List-based Implementation
  - Thread-safe Operations
  - Generic Type Support
- **Queue** - [Documentation](queue/README.md)
  - Array Queue Implementation
  - Linked List Queue Implementation
  - Priority Queue
  - Circular Queue
  - Thread-safe Operations
  - Generic Type Support
- **Tree** - [Documentation](tree/README.md)
  - Binary Tree
  - Binary Search Tree (BST)
  - AVL Tree (Self-balancing)
  - Red-Black Tree
  - B-Tree
  - Trie (Prefix Tree)
  - N-ary Tree
  - Expression Tree
  - Thread-safe Operations
- **Heap** - [Documentation](heap/README.md)
  - Binary Heap
  - Min Heap
  - Max Heap
  - Fibonacci Heap
  - Binomial Heap
  - Priority Queue Implementation
  - Thread-safe Operations
- **Hash** - [Documentation](hash/README.md)
  - Hash Table
  - Hash Map
  - Hash Set
  - Consistent Hashing
  - Linear Probing
  - Quadratic Probing
  - Double Hashing
  - Separate Chaining
  - Thread-safe Operations
- **OrderedMap** - [Documentation](orderedmap/README.md)
  - Thread-safe Implementation
  - Order Preservation
  - Concurrent Operations Support
  - Skip List Implementation
  - Advanced Features (Copy, Clear, Range iteration)
  - Generic Type Support

### Advanced Data Structures
- **Graph** - [Documentation](graph/README.md)
  - Adjacency Matrix
  - Adjacency List
  - Weighted Graph
  - Directed Graph
  - Undirected Graph
  - Graph Algorithms
    - Depth First Search (DFS)
    - Breadth First Search (BFS)
    - Dijkstra's Algorithm
    - Bellman-Ford Algorithm
    - Floyd-Warshall Algorithm
    - Kruskal's Algorithm
    - Prim's Algorithm
    - Topological Sort
    - Strongly Connected Components
    - Cycle Detection

## 🔧 Algorithms - [Documentation](algorithms/README.md)

### Sorting Algorithms
- Bubble Sort
- Selection Sort
- Insertion Sort
- Quick Sort
- Merge Sort
- Heap Sort

### Searching Algorithms
- Linear Search
- Binary Search
- Interpolation Search

## 🚀 Getting Started

### Prerequisites
- Go 1.23 or higher

### Installation
```bash
git clone https://github.com/mstgnz/data-structures.git
cd data-structures
go mod download
```

### Running Tests
```bash
go test ./...
```

## 📖 Usage Examples

You can find example implementations in the `examples` directory. Each data structure and algorithm includes its own test files demonstrating usage patterns.

### Example: Using a Linked List
```go
import "github.com/mstgnz/data-structures/linkedlist"

// Create a new doubly linked list
list := linkedlist.NewDouble()
list.Add(1)
list.Add(2)
list.Add(3)
```

### Example: Using a Stack
```go
import "github.com/mstgnz/data-structures/stack"

// Create a new array-based stack
stack := stack.NewArrayStack()
stack.Push(1)
stack.Push(2)
value, _ := stack.Pop()
```

### Example: Using a Queue
```go
import "github.com/mstgnz/data-structures/queue"

// Create a new array-based queue
queue := queue.NewArrayQueue()
queue.Enqueue(1)
queue.Enqueue(2)
value, _ := queue.Dequeue()
```

## 🤝 Contributing

Contributions are welcome! Please read our [Contributing Guidelines](CONTRIBUTING.md) for details on how to submit pull requests.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔍 Project Structure
```
.
├── algorithms/    # Basic algorithm implementations
├── advanced/      # Advanced data structures
├── examples/      # Usage examples
├── graph/         # Graph implementations
├── hash/         # Hash table implementations
├── heap/         # Heap implementations
├── linkedlist/   # Linked list implementations
├── orderedmap/   # Ordered map implementations
├── queue/        # Queue implementations
├── stack/        # Stack implementations
├── tree/         # Tree implementations
└── utils/        # Utility functions
```

## ✨ Features

- Clean and efficient implementations
- Comprehensive test coverage
- Well-documented code
- Generic implementations where applicable
- Performance optimized
- Thread-safe implementations where necessary
- Order preservation in map operations
- Concurrent access support with proper synchronization
- Advanced data structure features (Copy, Clear, Range operations)

## 📊 Performance

Each implementation includes performance considerations and Big O notation analysis in its respective documentation.

## 🔄 Version History

See [CHANGELOG.md](CHANGELOG.md) for release history and version details.