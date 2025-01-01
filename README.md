# Data Structures and Algorithms in Go

This project contains implementations of various data structures and algorithms in the Go programming language.

## Data Structures

1. **Stack**
   - Push, Pop, Peek operations
   - Size control
   - Empty check

2. **Queue**
   - Enqueue, Dequeue operations
   - Front, Rear access
   - Size control
   - Empty check

3. **LinkedList**
   - Singly and doubly linked
   - Insert, delete, search operations
   - Head/tail access
   - List traversal

4. **Tree**
   - Binary Tree
   - Binary Search Tree
   - AVL Tree
   - Red-Black Tree
   - B-Tree
   - Trie (Prefix Tree)
   - Segment Tree

5. **Graph**
   - Directed and undirected graphs
   - Weighted and unweighted edges
   - Adjacency List implementation
   - Graph algorithms:
     - Depth First Search (DFS)
     - Breadth First Search (BFS)
     - Dijkstra's Algorithm
     - Bellman-Ford Algorithm
     - Floyd-Warshall Algorithm
     - Prim's Algorithm
     - Kruskal's Algorithm
     - Topological Sort
     - Tarjan's Algorithm (Strongly Connected Components)
     - Articulation Points and Bridges
     - Euler Path/Circuit
     - Hamiltonian Path/Circuit

## Examples

See the `example/` directory for usage examples:

- `graph_algorithms.go`: Graph algorithms examples

## Installation

```bash
# Clone the repository
git clone https://github.com/mstgnz/data-structures.git

# Go to project directory
cd data-structures

# Install dependencies
go mod download
```

## Testing

```bash
# Run all tests
go test ./...

# Run tests for a specific package
go test ./Stack
go test ./Queue
go test ./LinkedList
go test ./Tree
go test ./Graph
```

## Contributing

See the `CONTRIBUTING.md` file for details.

## License

This project is licensed under the MIT License - see the `LICENSE` file for details.