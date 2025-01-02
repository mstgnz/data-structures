package graph

import (
	"math"
	"sync"
)

// BellmanFord implements the Bellman-Ford algorithm for single-source shortest paths
type BellmanFord struct {
	graph            *Graph
	source           int
	dist             []float64
	prev             []int
	infinity         float64
	hasNegativeCycle bool
	reachable        []bool
	mutex            sync.RWMutex
}

// NewBellmanFord creates a new Bellman-Ford instance
func NewBellmanFord(g *Graph, source int) *BellmanFord {
	bf := &BellmanFord{
		graph:            g,
		source:           source,
		infinity:         math.Inf(1),
		hasNegativeCycle: false,
		mutex:            sync.RWMutex{},
	}
	bf.initialize()
	return bf
}

// initialize prepares the distance and predecessor arrays
func (bf *BellmanFord) initialize() {
	bf.mutex.Lock()
	defer bf.mutex.Unlock()

	n := bf.graph.GetVertices()
	bf.dist = make([]float64, n)
	bf.prev = make([]int, n)
	bf.reachable = make([]bool, n)

	// Initialize all distances to infinity and predecessors to -1
	for i := 0; i < n; i++ {
		bf.dist[i] = bf.infinity
		bf.prev[i] = -1
		bf.reachable[i] = false
	}

	// Set distance of source node to 0
	bf.dist[bf.source] = 0
	bf.reachable[bf.source] = true
}

// ComputeShortestPaths computes single-source shortest paths
func (bf *BellmanFord) ComputeShortestPaths() bool {
	bf.mutex.Lock()
	defer bf.mutex.Unlock()

	n := bf.graph.GetVertices()
	edges := bf.getAllEdges()

	// First pass: Relax all edges |V|-1 times
	for i := 0; i < n-1; i++ {
		for _, edge := range edges {
			if bf.dist[edge.From] != bf.infinity {
				newDist := bf.dist[edge.From] + float64(edge.Weight)
				if bf.dist[edge.To] == bf.infinity || newDist < bf.dist[edge.To] {
					bf.dist[edge.To] = newDist
					bf.prev[edge.To] = edge.From
				}
			}
		}
	}

	// Second pass: Check for negative cycles
	for _, edge := range edges {
		if bf.dist[edge.From] != bf.infinity {
			newDist := bf.dist[edge.From] + float64(edge.Weight)
			if newDist < bf.dist[edge.To] {
				bf.hasNegativeCycle = true
				return false
			}
		}
	}

	// Update reachability
	bf.updateReachability()
	return true
}

// updateReachability performs BFS to mark reachable vertices
func (bf *BellmanFord) updateReachability() {
	n := bf.graph.GetVertices()
	visited := make([]bool, n)
	queue := []int{bf.source}
	visited[bf.source] = true
	bf.reachable[bf.source] = true

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		for _, edge := range bf.graph.adjList[v] {
			if !visited[edge.To] {
				visited[edge.To] = true
				bf.reachable[edge.To] = true
				queue = append(queue, edge.To)
			}
		}
	}

	// Reset distances for unreachable vertices
	for i := 0; i < n; i++ {
		if !visited[i] {
			bf.dist[i] = bf.infinity
			bf.prev[i] = -1
			bf.reachable[i] = false
		}
	}
}

// getAllEdges returns all edges in the graph
func (bf *BellmanFord) getAllEdges() []Edge {
	edges := make([]Edge, 0)
	n := bf.graph.GetVertices()

	for v := 0; v < n; v++ {
		edges = append(edges, bf.graph.adjList[v]...)
	}

	return edges
}

// GetDistance returns the shortest distance to a vertex
func (bf *BellmanFord) GetDistance(to int) float64 {
	bf.mutex.RLock()
	defer bf.mutex.RUnlock()

	if bf.hasNegativeCycle {
		return math.Inf(-1)
	}
	if !bf.reachable[to] {
		return bf.infinity
	}
	return bf.dist[to]
}

// GetPath returns the shortest path to a vertex
func (bf *BellmanFord) GetPath(to int) []int {
	bf.mutex.RLock()
	defer bf.mutex.RUnlock()

	if bf.hasNegativeCycle || !bf.reachable[to] || bf.dist[to] == bf.infinity {
		return nil
	}

	path := make([]int, 0)
	curr := to
	visited := make(map[int]bool)

	for curr != -1 {
		if visited[curr] {
			return nil // Cycle detected
		}
		visited[curr] = true
		path = append([]int{curr}, path...)
		curr = bf.prev[curr]
	}

	if path[0] != bf.source {
		return nil
	}

	return path
}

// GetAllDistances returns all computed distances
func (bf *BellmanFord) GetAllDistances() []float64 {
	bf.mutex.RLock()
	defer bf.mutex.RUnlock()
	return bf.dist
}

// GetPredecessors returns the predecessor array
func (bf *BellmanFord) GetPredecessors() []int {
	bf.mutex.RLock()
	defer bf.mutex.RUnlock()
	return bf.prev
}

// IsReachable checks if a vertex is reachable from the source
func (bf *BellmanFord) IsReachable(to int) bool {
	bf.mutex.RLock()
	defer bf.mutex.RUnlock()
	return bf.reachable[to]
}
