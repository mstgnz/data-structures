# Data Structures and Algorithms Examples

This directory contains comprehensive examples demonstrating the practical usage of various data structures and algorithms implemented in this project. Each example is designed to showcase real-world applications and best practices.

## 📁 Directory Structure

```
examples/
├── linkedlist/     # Linked List implementation examples
├── stack/         # Stack implementation examples
├── queue/         # Queue implementation examples
├── tree/          # Tree implementation examples
├── heap/          # Heap implementation examples
├── hash/          # Hash Table implementation examples
├── graph/         # Graph implementation examples
└── algorithms/    # Algorithm implementation examples
```

## 🎯 Available Examples

### 1. Data Structure Examples

#### Linked List Examples (`linkedlist/linkedlist_examples.go`)
- Generic type support for all implementations
- Singly Linked List operations
- Doubly Linked List operations
- Circular Linked List implementations
- Custom comparison functions for sorting and equality
- Examples with different data types (int, string, custom structs)
- Thread-safe operations
- List traversal and manipulation
- Real-world use cases

#### Stack Examples (`stack/stack_examples.go`)
- Array-based stack implementation with generic type support
- Linked List-based stack implementation with generic type support
- LIFO (Last In First Out) operations
- Support for any comparable type
- Examples with different data types (int, string, custom structs)
- Thread-safe operations
- Expression evaluation examples
- Practical stack applications

#### Queue Examples (`queue/queue_examples.go`)
- Array-based queue implementation with generic type support
- Linked List-based queue implementation with generic type support
- FIFO (First In First Out) operations
- Support for any comparable type
- Examples with different data types (int, string, custom structs)
- Thread-safe operations
- Priority queue examples
- Real-world queue scenarios

#### Tree Examples (`tree/tree_examples.go`)
- Binary Search Tree operations
- Tree traversal methods (Inorder, Preorder, Postorder)
- AVL Tree balancing examples
- Tree-based searching and sorting
- Practical tree applications

#### Heap Examples (`heap/heap_examples.go`)
- Min Heap implementation
- Max Heap implementation
- Priority queue applications
- Heap sort demonstrations
- Real-world heap usage

#### Hash Table Examples (`hash/hash_examples.go`)
- Hash function implementations
- Collision resolution strategies
- Key-value pair operations
- Hash table performance examples
- Practical hashing applications

#### Graph Examples (`graph/graph_examples.go`)
- Graph traversal algorithms
- Shortest path implementations
- Graph representation methods
- Real-world graph problems
- Network flow examples

### 2. Algorithm Examples (`algorithms/`)
- Sorting algorithm comparisons
- Searching algorithm implementations
- Dynamic programming examples
- Greedy algorithm demonstrations
- Divide and conquer examples

## 🚀 Running the Examples

Each example can be run independently using the Go command line:

```bash
# Run specific examples
go run examples/linkedlist/linkedlist_examples.go
go run examples/stack/stack_examples.go
go run examples/queue/queue_examples.go
go run examples/tree/tree_examples.go
go run examples/heap/heap_examples.go
go run examples/hash/hash_examples.go
go run examples/graph/graph_examples.go
go run examples/algorithms/sorting_examples.go
```

## 📝 Example Structure

Each example file follows a consistent structure:
1. Package declaration and imports
2. Example struct/type definitions
3. Generic type implementations where applicable
4. Implementation examples with various data types
5. Usage demonstrations
6. Performance considerations
7. Best practices and common pitfalls

## 🎓 Learning Objectives

These examples are designed to help you:
- Understand practical implementations of data structures
- Learn how to use generic types effectively
- Master comparison function implementations
- Explore real-world applications
- Understand performance implications
- Learn concurrent programming patterns
- Master Go programming concepts

## 📚 Additional Resources

- Refer to the main [README.md](../README.md) for complete documentation
- Check individual package documentation for detailed API references
- Review test files for additional usage examples
- See [CONTRIBUTING.md](../CONTRIBUTING.md) for contribution guidelines

## ⚠️ Notes

- All examples include detailed comments explaining the implementation
- Generic type support is demonstrated with various data types
- Thread safety is ensured in all concurrent operations
- Each example demonstrates error handling and best practices
- Performance considerations are documented where relevant
- Examples are designed to be self-contained and educational
- Code is optimized for readability and learning 