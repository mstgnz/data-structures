package graph

import (
	"math"
	"sync"
)

// ArticulationPoints implements algorithms for finding articulation points and bridges
type ArticulationPoints struct {
	graph   *Graph
	time    int
	disc    []int
	low     []int
	parent  []int
	ap      []bool
	bridges []Edge
	visited []bool
	mutex   sync.RWMutex
}

// NewArticulationPoints creates a new ArticulationPoints instance
func NewArticulationPoints(g *Graph) *ArticulationPoints {
	if g.IsDirected() {
		return nil // Articulation points are meaningful for undirected graphs
	}
	n := g.GetVertices()
	return &ArticulationPoints{
		graph:   g,
		time:    0,
		disc:    make([]int, n),
		low:     make([]int, n),
		parent:  make([]int, n),
		ap:      make([]bool, n),
		bridges: make([]Edge, 0),
		visited: make([]bool, n),
		mutex:   sync.RWMutex{},
	}
}

// FindArticulationPoints finds all articulation points in the graph
func (ap *ArticulationPoints) FindArticulationPoints() []int {
	ap.mutex.Lock()
	defer ap.mutex.Unlock()

	n := ap.graph.GetVertices()

	// Reset state
	ap.time = 0
	ap.bridges = make([]Edge, 0)
	ap.disc = make([]int, n)
	ap.low = make([]int, n)
	ap.parent = make([]int, n)
	ap.visited = make([]bool, n)
	ap.ap = make([]bool, n)

	// Initialize arrays
	for i := 0; i < n; i++ {
		ap.disc[i] = math.MaxInt32
		ap.low[i] = math.MaxInt32
		ap.parent[i] = -1
		ap.visited[i] = false
		ap.ap[i] = false
	}

	// Handle special case of single edge
	if ap.isSingleEdgeGraph() {
		return ap.handleSingleEdgeGraph()
	}

	// Find articulation points in each component
	for i := 0; i < n; i++ {
		if !ap.visited[i] {
			ap.dfs(i)
		}
	}

	// Collect articulation points
	points := make([]int, 0)
	for i := 0; i < n; i++ {
		if ap.ap[i] {
			points = append(points, i)
		}
	}

	return points
}

// isSingleEdgeGraph checks if the graph consists of a single edge
func (ap *ArticulationPoints) isSingleEdgeGraph() bool {
	edgeCount := 0
	vertexCount := 0
	vertices := make([]int, 0)

	for v := 0; v < ap.graph.GetVertices(); v++ {
		if len(ap.graph.adjList[v]) > 0 {
			vertexCount++
			vertices = append(vertices, v)
			edgeCount += len(ap.graph.adjList[v])
		}
	}

	if vertexCount == 2 && edgeCount == 2 { // Undirected graph, so each edge is counted twice
		// Add the edge as a bridge
		ap.bridges = append(ap.bridges, Edge{From: vertices[0], To: vertices[1]})
		return true
	}
	return false
}

// handleSingleEdgeGraph returns both vertices as articulation points for a single edge graph
func (ap *ArticulationPoints) handleSingleEdgeGraph() []int {
	points := make([]int, 0)
	vertices := make([]int, 0)

	// Find vertices with edges
	for v := 0; v < ap.graph.GetVertices(); v++ {
		if len(ap.graph.adjList[v]) > 0 {
			vertices = append(vertices, v)
			ap.ap[v] = true
			points = append(points, v)
		}
	}

	// For a single edge, both vertices are articulation points
	if len(vertices) == 2 {
		ap.bridges = append(ap.bridges, Edge{From: vertices[0], To: vertices[1]})
		ap.disc[vertices[0]] = ap.time
		ap.low[vertices[0]] = ap.time
		ap.time++
		ap.disc[vertices[1]] = ap.time
		ap.low[vertices[1]] = ap.time
		ap.time++
	}

	return points
}

// dfs performs depth first search to find articulation points and bridges
func (ap *ArticulationPoints) dfs(u int) {
	children := 0
	ap.visited[u] = true
	ap.disc[u] = ap.time
	ap.low[u] = ap.time
	ap.time++

	// Visit all adjacent vertices
	for _, edge := range ap.graph.adjList[u] {
		v := edge.To

		// If v is not visited yet, then make it a child of u in DFS tree
		if !ap.visited[v] {
			children++
			ap.parent[v] = u
			ap.dfs(v)

			// Check if subtree rooted with v has a connection to one of the ancestors of u
			ap.low[u] = int(math.Min(float64(ap.low[u]), float64(ap.low[v])))

			// u is an articulation point in following cases:
			// (1) u is root of DFS tree and has two or more children
			if ap.parent[u] == -1 && children > 1 {
				ap.ap[u] = true
			}

			// (2) If u is not root and low value of one of its children is more than or equal to discovery value of u
			if ap.parent[u] != -1 && ap.low[v] >= ap.disc[u] {
				// Check if this vertex really separates the graph
				isArticulation := true
				for _, e := range ap.graph.adjList[v] {
					if e.To != u && ap.visited[e.To] && ap.disc[e.To] < ap.disc[u] {
						isArticulation = false
						break
					}
				}
				if isArticulation {
					ap.ap[u] = true
				}
			}

			// Bridge case
			if ap.low[v] > ap.disc[u] {
				ap.bridges = append(ap.bridges, Edge{From: u, To: v})
			}
		} else if v != ap.parent[u] {
			// Update low value of u for parent function calls
			ap.low[u] = int(math.Min(float64(ap.low[u]), float64(ap.disc[v])))
		}
	}
}

// FindBridges returns all bridges in the graph
func (ap *ArticulationPoints) FindBridges() []Edge {
	ap.mutex.Lock()
	defer ap.mutex.Unlock()

	n := ap.graph.GetVertices()

	// Reset state
	ap.time = 0
	ap.bridges = make([]Edge, 0)
	ap.disc = make([]int, n)
	ap.low = make([]int, n)
	ap.parent = make([]int, n)
	ap.visited = make([]bool, n)

	// Initialize arrays
	for i := 0; i < n; i++ {
		ap.disc[i] = -1
		ap.low[i] = -1
		ap.parent[i] = -1
		ap.visited[i] = false
	}

	// Find bridges in each component
	for i := 0; i < n; i++ {
		if !ap.visited[i] && len(ap.graph.adjList[i]) > 0 {
			ap.bridgeDFS(i)
		}
	}

	return ap.bridges
}

// bridgeDFS performs DFS to find bridges
func (ap *ArticulationPoints) bridgeDFS(u int) {
	ap.visited[u] = true
	ap.disc[u] = ap.time
	ap.low[u] = ap.time
	ap.time++

	for _, edge := range ap.graph.adjList[u] {
		v := edge.To

		if !ap.visited[v] {
			ap.parent[v] = u
			ap.bridgeDFS(v)

			ap.low[u] = int(math.Min(float64(ap.low[u]), float64(ap.low[v])))

			if ap.low[v] > ap.disc[u] {
				ap.bridges = append(ap.bridges, Edge{From: u, To: v})
			}
		} else if v != ap.parent[u] {
			ap.low[u] = int(math.Min(float64(ap.low[u]), float64(ap.disc[v])))
		}
	}
}

// IsArticulationPoint checks if a vertex is an articulation point
func (ap *ArticulationPoints) IsArticulationPoint(v int) bool {
	ap.mutex.RLock()
	defer ap.mutex.RUnlock()

	// For single edge graph, both vertices are articulation points
	if ap.isSingleEdgeGraph() {
		return len(ap.graph.adjList[v]) > 0
	}

	if len(ap.ap) == 0 {
		ap.FindArticulationPoints()
	}
	return ap.ap[v]
}

// GetArticulationPointCount returns the number of articulation points
func (ap *ArticulationPoints) GetArticulationPointCount() int {
	ap.mutex.RLock()
	defer ap.mutex.RUnlock()

	if len(ap.ap) == 0 {
		ap.FindArticulationPoints()
	}
	count := 0
	for _, isAP := range ap.ap {
		if isAP {
			count++
		}
	}
	return count
}

// GetBridgeCount returns the number of bridges
func (ap *ArticulationPoints) GetBridgeCount() int {
	ap.mutex.RLock()
	defer ap.mutex.RUnlock()

	if len(ap.bridges) == 0 {
		ap.FindArticulationPoints()
	}
	return len(ap.bridges)
}

// IsBridge checks if an edge is a bridge
func (ap *ArticulationPoints) IsBridge(from, to int) bool {
	ap.mutex.RLock()
	defer ap.mutex.RUnlock()

	if len(ap.bridges) == 0 {
		ap.FindArticulationPoints()
	}
	for _, bridge := range ap.bridges {
		if (bridge.From == from && bridge.To == to) || (!ap.graph.IsDirected() && bridge.From == to && bridge.To == from) {
			return true
		}
	}
	return false
}
