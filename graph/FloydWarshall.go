package graph

import (
	"math"
	"sync"
)

// FloydWarshall implements the Floyd-Warshall algorithm for all-pairs shortest paths
type FloydWarshall struct {
	graph    *Graph
	dist     [][]float64 // Distance matrix
	next     [][]int     // Path matrix
	infinity float64
	mutex    sync.RWMutex
}

// NewFloydWarshall creates a new Floyd-Warshall instance
func NewFloydWarshall(g *Graph) *FloydWarshall {
	fw := &FloydWarshall{
		graph:    g,
		infinity: math.Inf(1),
		mutex:    sync.RWMutex{},
	}
	fw.initialize()
	return fw
}

// initialize prepares the distance and next matrices
func (fw *FloydWarshall) initialize() {
	fw.mutex.Lock()
	defer fw.mutex.Unlock()

	n := fw.graph.GetVertices()
	fw.dist = make([][]float64, n)
	fw.next = make([][]int, n)

	// Initialize matrices
	for i := 0; i < n; i++ {
		fw.dist[i] = make([]float64, n)
		fw.next[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				fw.dist[i][j] = 0
			} else {
				fw.dist[i][j] = fw.infinity
			}
			fw.next[i][j] = -1
		}
	}

	// Add edge weights
	for v := 0; v < n; v++ {
		for _, edge := range fw.graph.adjList[v] {
			fw.dist[v][edge.To] = float64(edge.Weight)
			fw.next[v][edge.To] = edge.To
		}
	}
}

// ComputeShortestPaths computes all-pairs shortest paths
func (fw *FloydWarshall) ComputeShortestPaths() {
	fw.mutex.Lock()
	defer fw.mutex.Unlock()

	n := fw.graph.GetVertices()

	// Floyd-Warshall algorithm
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if fw.dist[i][k] != fw.infinity && fw.dist[k][j] != fw.infinity {
					newDist := fw.dist[i][k] + fw.dist[k][j]
					if newDist < fw.dist[i][j] {
						fw.dist[i][j] = newDist
						fw.next[i][j] = fw.next[i][k]
					}
				}
			}
		}
	}
}

// GetDistance returns the shortest distance between two vertices
func (fw *FloydWarshall) GetDistance(from, to int) float64 {
	fw.mutex.RLock()
	defer fw.mutex.RUnlock()
	return fw.dist[from][to]
}

// GetPath returns the shortest path between two vertices
func (fw *FloydWarshall) GetPath(from, to int) []int {
	fw.mutex.RLock()
	defer fw.mutex.RUnlock()

	if fw.next[from][to] == -1 {
		return nil
	}

	path := []int{from}
	for from != to {
		from = fw.next[from][to]
		path = append(path, from)
	}

	return path
}

// HasNegativeCycle checks if the graph contains a negative cycle
func (fw *FloydWarshall) HasNegativeCycle() bool {
	fw.mutex.RLock()
	defer fw.mutex.RUnlock()

	n := fw.graph.GetVertices()
	for i := 0; i < n; i++ {
		if fw.dist[i][i] < 0 {
			return true
		}
	}
	return false
}

// GetAllPairsDistances returns the distance matrix
func (fw *FloydWarshall) GetAllPairsDistances() [][]float64 {
	fw.mutex.RLock()
	defer fw.mutex.RUnlock()
	return fw.dist
}

// GetAllPairsNextHops returns the next hop matrix
func (fw *FloydWarshall) GetAllPairsNextHops() [][]int {
	fw.mutex.RLock()
	defer fw.mutex.RUnlock()
	return fw.next
}
