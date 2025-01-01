package Graph

import "math"

// BellmanFord implements the Bellman-Ford algorithm for single-source shortest paths
type BellmanFord struct {
	graph    *Graph
	source   int
	dist     []float64
	prev     []int
	infinity float64
}

// NewBellmanFord creates a new Bellman-Ford instance
func NewBellmanFord(g *Graph, source int) *BellmanFord {
	bf := &BellmanFord{
		graph:    g,
		source:   source,
		infinity: math.Inf(1),
	}
	bf.initialize()
	return bf
}

// initialize prepares the distance and predecessor arrays
func (bf *BellmanFord) initialize() {
	n := bf.graph.GetVertices()
	bf.dist = make([]float64, n)
	bf.prev = make([]int, n)

	// Initialize all distances to infinity
	for i := 0; i < n; i++ {
		bf.dist[i] = bf.infinity
		bf.prev[i] = -1
	}

	// Set distance of source node to 0
	bf.dist[bf.source] = 0
}

// ComputeShortestPaths computes single-source shortest paths
// Returns true if no negative cycle is reachable from the source
func (bf *BellmanFord) ComputeShortestPaths() bool {
	n := bf.graph.GetVertices()

	// Repeat for each node n-1 times
	for i := 0; i < n-1; i++ {
		// For each edge
		for v := 0; v < n; v++ {
			for _, edge := range bf.graph.adjList[v] {
				bf.relax(v, edge)
			}
		}
	}

	// Check for negative cycle
	for v := 0; v < n; v++ {
		for _, edge := range bf.graph.adjList[v] {
			if bf.dist[v] != bf.infinity &&
				bf.dist[v]+float64(edge.Weight) < bf.dist[edge.To] {
				return false // Negative cycle found
			}
		}
	}

	return true
}

// relax performs edge relaxation
func (bf *BellmanFord) relax(from int, edge Edge) {
	if bf.dist[from] != bf.infinity {
		newDist := bf.dist[from] + float64(edge.Weight)
		if newDist < bf.dist[edge.To] {
			bf.dist[edge.To] = newDist
			bf.prev[edge.To] = from
		}
	}
}

// GetDistance returns the shortest distance to a vertex
func (bf *BellmanFord) GetDistance(to int) float64 {
	return bf.dist[to]
}

// GetPath returns the shortest path to a vertex
func (bf *BellmanFord) GetPath(to int) []int {
	if bf.dist[to] == bf.infinity {
		return nil
	}

	path := make([]int, 0)
	for curr := to; curr != -1; curr = bf.prev[curr] {
		path = append([]int{curr}, path...)
	}

	return path
}

// GetAllDistances returns all computed distances
func (bf *BellmanFord) GetAllDistances() []float64 {
	return bf.dist
}

// GetPredecessors returns the predecessor array
func (bf *BellmanFord) GetPredecessors() []int {
	return bf.prev
}

// IsReachable checks if a vertex is reachable from the source
func (bf *BellmanFord) IsReachable(to int) bool {
	return bf.dist[to] != bf.infinity
}
