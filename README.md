# Data Structures and Algorithms in Go

This repository contains implementations of various data structures and algorithms in Go programming language. It serves as both a learning resource and a practical reference for developers.

## 📚 Data Structures

### Basic Data Structures
- **Linked List**
  - Singly Linked List
  - Doubly Linked List
  - Circular Linked List
- **Stack**
  - Array-based implementation
  - Linked List-based implementation
- **Queue**
  - Simple Queue
  - Priority Queue
  - Circular Queue
- **Tree**
  - Binary Tree
  - Binary Search Tree (BST)
  - AVL Tree
- **Heap**
  - Min Heap
  - Max Heap
- **Hash**
  - Hash Table
  - Hash Map implementations
- **OrderedMap**
  - Thread-safe implementation
  - Order preservation
  - Concurrent operations support
  - Advanced features (Copy, Clear, Range iteration)

### Advanced Data Structures
- **Graph**
  - Adjacency Matrix
  - Adjacency List
  - Graph Algorithms
    - Depth First Search (DFS)
    - Breadth First Search (BFS)
    - Topological Sort
    - Shortest Path Algorithms

## 🔧 Algorithms

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

```go
// Example: Creating and using a Binary Search Tree
import "github.com/mstgnz/data-structures/tree"

bst := tree.NewBST()
bst.Insert(5)
bst.Insert(3)
bst.Insert(7)
```

```go
// Example: Using OrderedMap with thread-safety and order preservation
import "github.com/mstgnz/data-structures/orderedmap"

om := orderedmap.New()

// Adding elements (order is preserved)
om.Set("first", 1)
om.Set("second", 2)
om.Set("third", 3)

// Reading values
if val, exists := om.Get("second"); exists {
    fmt.Printf("Value: %v\n", val)
}

// Iterating in order
om.Range(func(key, value any) bool {
    fmt.Printf("%v: %v\n", key, value)
    return true
})

// Thread-safe operations
go func() {
    om.Set("concurrent", 42)
}()
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
├── hash/          # Hash table implementations
├── heap/          # Heap implementations
├── linkedlist/    # Linked list implementations
├── queue/         # Queue implementations
├── stack/         # Stack implementations
├── tree/          # Tree implementations
└── utils/         # Utility functions
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