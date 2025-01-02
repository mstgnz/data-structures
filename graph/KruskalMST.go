package graph

import (
	"sort"
	"sync"
)

// KruskalMST implements Kruskal's algorithm for finding Minimum Spanning Tree
type KruskalMST struct {
	graph    *Graph
	parent   []int   // Parent array for Union-Find
	rank     []int   // Rank array for Union-Find
	mstEdges []Edge  // Edges in MST
	mstCost  float64 // Total cost of MST
	mutex    sync.RWMutex
}

// NewKruskalMST creates a new Kruskal's MST instance
func NewKruskalMST(g *Graph) *KruskalMST {
	if g.IsDirected() {
		return nil // Kruskal algorithm works for undirected graphs
	}
	return &KruskalMST{
		graph:    g,
		mstEdges: make([]Edge, 0),
		mstCost:  0,
		mutex:    sync.RWMutex{},
	}
}

// initialize prepares the Union-Find data structure
func (k *KruskalMST) initialize() {
	n := k.graph.GetVertices()
	k.parent = make([]int, n)
	k.rank = make([]int, n)
	k.mstEdges = make([]Edge, 0)
	k.mstCost = 0

	// Initialize each node in its own set
	for i := 0; i < n; i++ {
		k.parent[i] = i
		k.rank[i] = 0
	}
}

// find returns the representative of the set containing x
func (k *KruskalMST) find(x int) int {
	if k.parent[x] != x {
		k.parent[x] = k.find(k.parent[x]) // Path compression
	}
	return k.parent[x]
}

// union merges sets containing x and y
func (k *KruskalMST) union(x, y int) {
	rootX := k.find(x)
	rootY := k.find(y)

	if rootX != rootY {
		// Union by rank
		if k.rank[rootX] < k.rank[rootY] {
			k.parent[rootX] = rootY
		} else if k.rank[rootX] > k.rank[rootY] {
			k.parent[rootY] = rootX
		} else {
			k.parent[rootY] = rootX
			k.rank[rootX]++
		}
	}
}

// countComponents counts the number of connected components
func (k *KruskalMST) countComponents() int {
	// Reset Union-Find structure
	k.initialize()

	// Mark vertices with edges
	hasEdge := make([]bool, k.graph.GetVertices())
	for v := 0; v < k.graph.GetVertices(); v++ {
		if len(k.graph.adjList[v]) > 0 {
			hasEdge[v] = true
		}
	}

	// Create adjacency list for vertices with edges
	verticesWithEdges := make([]int, 0)
	for v := 0; v < k.graph.GetVertices(); v++ {
		if hasEdge[v] {
			verticesWithEdges = append(verticesWithEdges, v)
		}
	}

	// Special case: no edges
	if len(verticesWithEdges) == 0 {
		return 1
	}

	// Run union operations on edges
	for v := 0; v < k.graph.GetVertices(); v++ {
		for _, edge := range k.graph.adjList[v] {
			if edge.From < edge.To { // Process each edge once
				k.union(edge.From, edge.To)
			}
		}
	}

	// Count unique components among vertices with edges
	components := make(map[int]bool)
	for _, v := range verticesWithEdges {
		components[k.find(v)] = true
	}

	return len(components)
}

// GetNumComponents returns the number of connected components
func (k *KruskalMST) GetNumComponents() int {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	return k.countComponents()
}

// FindMST finds the Minimum Spanning Tree
func (k *KruskalMST) FindMST() bool {
	k.mutex.Lock()
	defer k.mutex.Unlock()

	k.initialize()

	// Special case: single vertex
	if k.graph.GetVertices() == 1 {
		return true
	}

	// Count vertices with edges
	verticesWithEdges := 0
	hasEdge := make([]bool, k.graph.GetVertices())
	for v := 0; v < k.graph.GetVertices(); v++ {
		if len(k.graph.adjList[v]) > 0 {
			hasEdge[v] = true
			verticesWithEdges++
		}
	}

	// Special case: no edges
	if verticesWithEdges == 0 {
		return true
	}

	// Check if graph is connected
	numComponents := k.countComponents()
	if numComponents > 1 {
		return false
	}

	// Collect all edges
	edges := make([]Edge, 0)
	for v := 0; v < k.graph.GetVertices(); v++ {
		for _, edge := range k.graph.adjList[v] {
			if edge.From < edge.To { // Add each edge once
				edges = append(edges, edge)
			}
		}
	}

	// Sort edges by weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	// Reset for MST construction
	k.initialize()

	// Run Kruskal's algorithm
	edgeCount := 0
	for _, edge := range edges {
		if k.find(edge.From) != k.find(edge.To) {
			k.union(edge.From, edge.To)
			k.mstEdges = append(k.mstEdges, edge)
			k.mstCost += float64(edge.Weight)
			edgeCount++
		}
	}

	return edgeCount == verticesWithEdges-1
}

// GetMSTEdges returns the edges in the MST
func (k *KruskalMST) GetMSTEdges() []Edge {
	k.mutex.RLock()
	defer k.mutex.RUnlock()
	return k.mstEdges
}

// GetMSTCost returns the total cost of the MST
func (k *KruskalMST) GetMSTCost() float64 {
	k.mutex.RLock()
	defer k.mutex.RUnlock()
	return k.mstCost
}

// IsConnected checks if two vertices are connected in the MST
func (k *KruskalMST) IsConnected(x, y int) bool {
	k.mutex.RLock()
	defer k.mutex.RUnlock()
	return k.find(x) == k.find(y)
}
