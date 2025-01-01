package graph

import "fmt"

// EulerPath implements algorithms for finding Euler paths and circuits
type EulerPath struct {
	graph      *Graph
	path       []int
	edgeCount  map[string]int
	edgesUsed  int
	totalEdges int
}

// NewEulerPath creates a new EulerPath instance
func NewEulerPath(g *Graph) *EulerPath {
	totalEdges := 0
	for v := 0; v < g.GetVertices(); v++ {
		totalEdges += len(g.adjList[v])
	}
	if !g.IsDirected() {
		totalEdges /= 2 // Each edge counted twice in undirected graph
	}

	return &EulerPath{
		graph:      g,
		path:       make([]int, 0),
		edgeCount:  make(map[string]int),
		edgesUsed:  0,
		totalEdges: totalEdges,
	}
}

// makeEdgeKey creates a unique key for an edge
func (ep *EulerPath) makeEdgeKey(from, to int) string {
	if ep.graph.IsDirected() {
		return fmt.Sprintf("%d->%d", from, to)
	}
	if from < to {
		return fmt.Sprintf("%d-%d", from, to)
	}
	return fmt.Sprintf("%d-%d", to, from)
}

// getUnusedEdge returns an unused edge from vertex v
func (ep *EulerPath) getUnusedEdge(v int) *Edge {
	for _, edge := range ep.graph.adjList[v] {
		key := ep.makeEdgeKey(v, edge.To)
		count := ep.edgeCount[key]

		if ep.graph.IsDirected() {
			if count == 0 {
				return &edge
			}
		} else {
			if count == 0 {
				return &edge
			}
		}
	}
	return nil
}

// dfs performs depth first search to find Euler path
func (ep *EulerPath) dfs(v int) {
	for {
		edge := ep.getUnusedEdge(v)
		if edge == nil {
			break
		}

		key := ep.makeEdgeKey(v, edge.To)
		ep.edgeCount[key]++
		ep.edgesUsed++
		ep.dfs(edge.To)
	}
	ep.path = append(ep.path, v)
}

// FindEulerPath finds an Euler path in the graph if it exists
func (ep *EulerPath) FindEulerPath() []int {
	if !ep.HasEulerPath() {
		return nil
	}

	// Reset state
	ep.path = make([]int, 0)
	ep.edgeCount = make(map[string]int)
	ep.edgesUsed = 0

	// Find starting vertex
	start := ep.findStartVertex()
	ep.dfs(start)

	// Check if all edges were used
	if ep.edgesUsed != ep.totalEdges {
		return nil
	}

	// Reverse the path
	for i, j := 0, len(ep.path)-1; i < j; i, j = i+1, j-1 {
		ep.path[i], ep.path[j] = ep.path[j], ep.path[i]
	}

	// Verify circuit property if needed
	if ep.HasEulerCircuit() && ep.path[0] != ep.path[len(ep.path)-1] {
		return nil
	}

	return ep.path
}

// FindEulerCircuit finds an Euler circuit in the graph if it exists
func (ep *EulerPath) FindEulerCircuit() []int {
	if !ep.HasEulerCircuit() {
		return nil
	}
	path := ep.FindEulerPath()
	if path == nil || path[0] != path[len(path)-1] {
		return nil
	}
	return path
}

// HasEulerPath checks if the graph has an Euler path
func (ep *EulerPath) HasEulerPath() bool {
	if !ep.isConnected() {
		return false
	}

	if ep.graph.IsDirected() {
		inDegree := make([]int, ep.graph.GetVertices())
		outDegree := make([]int, ep.graph.GetVertices())

		for v := 0; v < ep.graph.GetVertices(); v++ {
			outDegree[v] = len(ep.graph.adjList[v])
			for _, edge := range ep.graph.adjList[v] {
				inDegree[edge.To]++
			}
		}

		startCount := 0
		endCount := 0
		for v := 0; v < ep.graph.GetVertices(); v++ {
			diff := outDegree[v] - inDegree[v]
			if diff > 1 || diff < -1 {
				return false
			}
			if diff == 1 {
				startCount++
			}
			if diff == -1 {
				endCount++
			}
		}
		return (startCount == 0 && endCount == 0) || (startCount == 1 && endCount == 1)
	}

	// For undirected graph
	oddCount := 0
	for v := 0; v < ep.graph.GetVertices(); v++ {
		if len(ep.graph.adjList[v])%2 != 0 {
			oddCount++
		}
	}
	return oddCount == 0 || oddCount == 2
}

// HasEulerCircuit checks if the graph has an Euler circuit
func (ep *EulerPath) HasEulerCircuit() bool {
	if !ep.isConnected() {
		return false
	}

	if ep.graph.IsDirected() {
		// Check in-out degree for directed graph
		inDegree := make([]int, ep.graph.GetVertices())
		for v := 0; v < ep.graph.GetVertices(); v++ {
			for _, edge := range ep.graph.adjList[v] {
				inDegree[edge.To]++
			}
		}

		for v := 0; v < ep.graph.GetVertices(); v++ {
			if len(ep.graph.adjList[v]) != inDegree[v] {
				return false
			}
		}
		return true
	}

	// In undirected graph all nodes should have even degree
	for v := 0; v < ep.graph.GetVertices(); v++ {
		if len(ep.graph.adjList[v])%2 != 0 {
			return false
		}
	}
	return true
}

// isConnected checks if the graph is connected
func (ep *EulerPath) isConnected() bool {
	visited := make([]bool, ep.graph.GetVertices())

	// Start DFS from first node
	start := 0
	for v := 0; v < ep.graph.GetVertices(); v++ {
		if len(ep.graph.adjList[v]) > 0 {
			start = v
			break
		}
	}

	ep.dfsUtil(start, visited)

	// Check if all nodes are visited
	for v := 0; v < ep.graph.GetVertices(); v++ {
		if len(ep.graph.adjList[v]) > 0 && !visited[v] {
			return false
		}
	}
	return true
}

// dfsUtil performs DFS for connectivity check
func (ep *EulerPath) dfsUtil(v int, visited []bool) {
	visited[v] = true
	for _, edge := range ep.graph.adjList[v] {
		if !visited[edge.To] {
			ep.dfsUtil(edge.To, visited)
		}
	}
}

// findStartVertex finds a suitable starting vertex for Euler path
func (ep *EulerPath) findStartVertex() int {
	if ep.graph.IsDirected() {
		inDegree := make([]int, ep.graph.GetVertices())
		for v := 0; v < ep.graph.GetVertices(); v++ {
			for _, edge := range ep.graph.adjList[v] {
				inDegree[edge.To]++
			}
		}

		// First try to find a vertex with out degree = in degree + 1
		for v := 0; v < ep.graph.GetVertices(); v++ {
			if len(ep.graph.adjList[v]) > 0 && len(ep.graph.adjList[v]) == inDegree[v]+1 {
				return v
			}
		}

		// If no such vertex exists, find first vertex with non-zero degree
		for v := 0; v < ep.graph.GetVertices(); v++ {
			if len(ep.graph.adjList[v]) > 0 {
				return v
			}
		}
	} else {
		// For undirected graph, first try to find vertex with odd degree
		for v := 0; v < ep.graph.GetVertices(); v++ {
			if len(ep.graph.adjList[v])%2 != 0 {
				return v
			}
		}

		// If no odd degree vertex exists, find first vertex with non-zero degree
		for v := 0; v < ep.graph.GetVertices(); v++ {
			if len(ep.graph.adjList[v]) > 0 {
				return v
			}
		}
	}
	return 0
}
