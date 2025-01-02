# Graph Package

This package provides comprehensive graph data structure implementations and algorithms in Go. All implementations are thread-safe and support concurrent operations.

## Features

### Core Graph Structure
- Weighted graph implementation supporting both directed and undirected graphs
- Thread-safe operations with RWMutex
- Adjacency list representation
- Basic operations:
  - Add edge
  - Get neighbors
  - Get vertex count
  - Check if directed

### Graph Traversal
- Breadth-First Search (BFS)
- Depth-First Search (DFS)
- Support for custom traversal orders

### Shortest Path Algorithms
- Dijkstra's Algorithm:
  - Single-source shortest paths
  - Priority queue optimization
  - Distance and path reconstruction
- Bellman-Ford Algorithm:
  - Negative weight support
  - Negative cycle detection
  - Path reconstruction
  - Reachability checking
- Floyd-Warshall Algorithm:
  - All-pairs shortest paths
  - Path reconstruction

### Minimum Spanning Trees
- Kruskal's Algorithm:
  - Union-Find data structure
  - Edge sorting optimization
- Prim's Algorithm:
  - Priority queue implementation
  - Efficient edge selection

### Graph Analysis
- Tarjan's Strongly Connected Components:
  - Component identification
  - Component size analysis
  - Connectivity checking
- Articulation Points:
  - Cut vertex detection
  - Bridge identification
- Euler Path:
  - Path existence checking
  - Path construction
- Hamiltonian Path:
  - Path existence checking
  - Path construction

## Usage Examples

### Basic Graph Operations
```go
// Create a new graph
graph := NewGraph(5, false) // 5 vertices, undirected

// Add edges
graph.AddEdge(0, 1, 4)  // edge from 0 to 1 with weight 4
graph.AddEdge(1, 2, 3)
graph.AddEdge(2, 3, 2)

// Get neighbors
neighbors := graph.GetNeighbors(1)

// Graph traversal
bfsOrder := graph.BFS(0)
dfsOrder := graph.DFS(0)
```

### Shortest Path Algorithms
```go
// Dijkstra's Algorithm
distances := graph.Dijkstra(0)

// Bellman-Ford Algorithm
bf := NewBellmanFord(graph, 0)
if bf.ComputeShortestPaths() {
    distance := bf.GetDistance(3)
    path := bf.GetPath(3)
}

// Floyd-Warshall Algorithm
fw := NewFloydWarshall(graph)
fw.ComputeAllPairs()
distance := fw.GetDistance(0, 3)
```

### Minimum Spanning Tree
```go
// Kruskal's Algorithm
mstEdges := graph.Kruskal()

// Prim's Algorithm
mstEdges = graph.Prim(0)
```

### Graph Analysis
```go
// Strongly Connected Components
tarjan := NewTarjanSCC(graph)
components := tarjan.FindComponents()
isStronglyConnected := tarjan.IsStronglyConnected()

// Articulation Points
ap := NewArticulationPoints(graph)
cutVertices := ap.FindArticulationPoints()
bridges := ap.FindBridges()

// Euler Path
euler := NewEulerPath(graph)
if euler.HasEulerPath() {
    path := euler.FindEulerPath()
}
```

## Implementation Details

### Time Complexities

#### Basic Operations
- Add Edge: O(1)
- Get Neighbors: O(1)
- BFS/DFS: O(V + E)

#### Shortest Path Algorithms
- Dijkstra: O((V + E) log V)
- Bellman-Ford: O(VE)
- Floyd-Warshall: O(V³)

#### Minimum Spanning Tree
- Kruskal: O(E log E)
- Prim: O((V + E) log V)

#### Graph Analysis
- Tarjan's SCC: O(V + E)
- Articulation Points: O(V + E)
- Euler Path: O(E)
- Hamiltonian Path: O(2^N * N^2)

Where:
- V is the number of vertices
- E is the number of edges
- N is the size of the graph

### Thread Safety
- All operations are protected with RWMutex
- Read operations use RLock
- Write operations use Lock
- Proper lock/unlock handling with defer

## Testing
Each component comes with comprehensive test coverage. Run tests using:
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