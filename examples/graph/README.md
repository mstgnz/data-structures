# Graph Algorithm Examples

This directory contains example programs demonstrating the usage of various graph algorithms.

## Examples

1. `graph_algorithms.go`: Basic graph algorithm examples
   - Topological Sort
   - Strongly Connected Components (Tarjan)
   - Articulation Points and Bridges
   - Euler Path and Circuit
   - Hamiltonian Path and Circuit
   - Shortest Path Algorithms (Bellman-Ford, Floyd-Warshall)
   - Minimum Spanning Tree (Prim, Kruskal)

## Running

To run the example programs:

```bash
# Run graph algorithms example
go run graph_algorithms.go
```

## Expected Output

When you run the program, you should see output similar to the following:

```
Topological Sort Example:
Topological Sort: [5 4 2 3 1 0]

Strongly Connected Components Example:
Strongly Connected Components: [[0 1 2] [3] [4]]

Articulation Points and Bridges Example:
Articulation Points: [0 3]
Bridges: [{0 3} {3 4}]

Euler Path and Circuit Example:
Euler Circuit: [0 1 2 3 0]

Hamiltonian Path and Circuit Example:
Hamiltonian Circuit: [0 1 2 3 0]

Shortest Path Algorithms Example:
Bellman-Ford Results:
Distances: [0 3 2 5 6]

Floyd-Warshall Results:
Distance Matrix: [[0 3 2 5 6] [∞ 0 3 2 3] [∞ 1 0 3 4] [∞ ∞ ∞ 0 ∞] [∞ ∞ ∞ 1 0]]

Minimum Spanning Tree Example:
Prim MST Results:
MST Edges: [{0 1} {1 2} {0 3} {1 4}]
MST Cost: 16

Kruskal MST Results:
MST Edges: [{0 1} {1 2} {1 4} {0 3}]
MST Cost: 16
```

## Notes

- Examples use simplified graphs to demonstrate basic usage of algorithms.
- In real-world applications, graphs are typically larger and more complex.
- Each algorithm is optimized for different use cases:
  - Topological Sort: Dependency graphs, task scheduling
  - Strongly Connected Components: Network analysis, social networks
  - Articulation Points: Network reliability, critical points
  - Euler/Hamilton: Route planning, circuit design
  - Shortest Path: Navigation, network routing
  - MST: Network design, clustering 