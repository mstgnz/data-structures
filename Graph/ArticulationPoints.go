package graph

import "math"

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
	}
}

// FindArticulationPoints finds all articulation points in the graph
func (ap *ArticulationPoints) FindArticulationPoints() []int {
	n := ap.graph.GetVertices()

	// Reset state
	ap.time = 0
	ap.bridges = make([]Edge, 0)

	// Initialize arrays
	for i := 0; i < n; i++ {
		ap.disc[i] = -1
		ap.low[i] = -1
		ap.parent[i] = -1
		ap.visited[i] = false
		ap.ap[i] = false
	}

	// Handle special case of single edge
	if ap.isSingleEdgeGraph() {
		return ap.handleSingleEdgeGraph()
	}

	// Call DFS for each unvisited vertex
	for i := 0; i < n; i++ {
		if !ap.visited[i] {
			ap.dfs(i)
		}
	}

	// Post-process: check for vertices that connect different components
	for v := 0; v < n; v++ {
		if len(ap.graph.adjList[v]) > 1 {
			// Find all neighbors that are not directly connected
			for i, edge1 := range ap.graph.adjList[v] {
				for j := i + 1; j < len(ap.graph.adjList[v]); j++ {
					edge2 := ap.graph.adjList[v][j]
					// Check if these neighbors are connected through any other path
					if !ap.hasAlternatePath(edge1.To, edge2.To, v) {
						ap.ap[v] = true
						break
					}
				}
				if ap.ap[v] {
					break
				}
			}
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

// hasAlternatePath checks if there is a path between u and v that doesn't go through exclude
func (ap *ArticulationPoints) hasAlternatePath(u, v, exclude int) bool {
	if u == v {
		return true
	}

	visited := make([]bool, ap.graph.GetVertices())
	visited[exclude] = true
	return ap.dfsPath(u, v, visited)
}

// dfsPath performs DFS to find a path between u and v
func (ap *ArticulationPoints) dfsPath(u, v int, visited []bool) bool {
	if u == v {
		return true
	}

	visited[u] = true
	for _, edge := range ap.graph.adjList[u] {
		if !visited[edge.To] {
			if ap.dfsPath(edge.To, v, visited) {
				return true
			}
		}
	}
	return false
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
		}
	}

	// For a single edge, both vertices are articulation points
	if len(vertices) == 2 {
		ap.bridges = append(ap.bridges, Edge{From: vertices[0], To: vertices[1]})
		for _, v := range vertices {
			ap.ap[v] = true
			points = append(points, v)
			ap.disc[v] = ap.time
			ap.low[v] = ap.time
			ap.time++
		}
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
				ap.ap[u] = true
			}

			// Bridge case
			if ap.low[v] > ap.disc[u] {
				ap.bridges = append(ap.bridges, Edge{From: u, To: v})
				// Mark bridge endpoints as articulation points if they have more than one neighbor
				if len(ap.graph.adjList[u]) > 1 {
					ap.ap[u] = true
				}
				if len(ap.graph.adjList[v]) > 1 {
					ap.ap[v] = true
				}
			}
		} else if v != ap.parent[u] {
			// Update low value of u for parent function calls
			ap.low[u] = int(math.Min(float64(ap.low[u]), float64(ap.disc[v])))
		}
	}

	// Special case: check if this vertex is a cut vertex
	if len(ap.graph.adjList[u]) > 1 {
		// Count the number of biconnected components through this vertex
		components := make(map[int]bool)
		for _, edge := range ap.graph.adjList[u] {
			v := edge.To
			if v != ap.parent[u] {
				if ap.low[v] >= ap.disc[u] {
					components[v] = true
				}
			}
		}
		if len(components) > 1 {
			ap.ap[u] = true
		}
	}
}

// FindBridges finds all bridges in the graph
func (ap *ArticulationPoints) FindBridges() []Edge {
	if len(ap.bridges) == 0 {
		ap.FindArticulationPoints() // Will also find bridges
	}
	return ap.bridges
}

// IsArticulationPoint checks if a vertex is an articulation point
func (ap *ArticulationPoints) IsArticulationPoint(v int) bool {
	// For single edge graph, both vertices are articulation points
	if ap.isSingleEdgeGraph() {
		return len(ap.graph.adjList[v]) > 0
	}

	// For other cases, check if vertex is marked as articulation point
	if len(ap.ap) == 0 {
		ap.FindArticulationPoints()
	}
	return ap.ap[v]
}

// GetArticulationPointCount returns the number of articulation points
func (ap *ArticulationPoints) GetArticulationPointCount() int {
	points := ap.FindArticulationPoints()
	return len(points)
}

// GetBridgeCount returns the number of bridges
func (ap *ArticulationPoints) GetBridgeCount() int {
	bridges := ap.FindBridges()
	return len(bridges)
}

// IsBridge checks if an edge is a bridge
func (ap *ArticulationPoints) IsBridge(from, to int) bool {
	bridges := ap.FindBridges()
	for _, bridge := range bridges {
		if (bridge.From == from && bridge.To == to) ||
			(bridge.From == to && bridge.To == from) {
			return true
		}
	}
	return false
}
