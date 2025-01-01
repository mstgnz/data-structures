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
// Returns true if no negative cycle is reachable from the source
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
				if newDist < bf.dist[edge.To] {
					bf.dist[edge.To] = newDist
					bf.prev[edge.To] = edge.From
					bf.reachable[edge.To] = true
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

	// Update reachability using DFS
	bf.updateReachability()

	return true
}

// updateReachability performs DFS to mark reachable vertices
func (bf *BellmanFord) updateReachability() {
	visited := make([]bool, bf.graph.GetVertices())
	bf.dfs(bf.source, visited)

	// Update reachability based on DFS results
	for i := range bf.reachable {
		bf.reachable[i] = visited[i]
	}
}

// dfs performs depth-first search for reachability
func (bf *BellmanFord) dfs(v int, visited []bool) {
	visited[v] = true
	for _, edge := range bf.graph.adjList[v] {
		if !visited[edge.To] {
			bf.dfs(edge.To, visited)
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

	if bf.hasNegativeCycle || !bf.reachable[to] {
		return nil
	}

	// Check for cycles in the path
	visited := make(map[int]bool)
	path := make([]int, 0)
	curr := to

	for curr != -1 {
		if visited[curr] {
			return nil // Cycle detected
		}
		visited[curr] = true
		path = append([]int{curr}, path...)
		curr = bf.prev[curr]
	}

	// Verify path starts from source
	if len(path) > 0 && path[0] != bf.source {
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
