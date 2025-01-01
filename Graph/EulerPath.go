package Graph

// EulerPath implements algorithms for finding Euler paths and circuits
type EulerPath struct {
	graph     *Graph
	path      []int
	edgeCount map[string]int // Count of used edges
}

// NewEulerPath creates a new EulerPath instance
func NewEulerPath(g *Graph) *EulerPath {
	return &EulerPath{
		graph:     g,
		path:      make([]int, 0),
		edgeCount: make(map[string]int),
	}
}

// makeEdgeKey creates a unique key for an edge
func (ep *EulerPath) makeEdgeKey(from, to int) string {
	if from < to {
		return string(rune(from)) + "-" + string(rune(to))
	}
	return string(rune(to)) + "-" + string(rune(from))
}

// getUnusedEdge returns an unused edge from vertex v
func (ep *EulerPath) getUnusedEdge(v int) *Edge {
	for _, edge := range ep.graph.adjList[v] {
		key := ep.makeEdgeKey(edge.From, edge.To)
		count := ep.edgeCount[key]

		if !ep.graph.IsDirected() {
			// In undirected graph each edge should be used once
			if count == 0 {
				return &edge
			}
		} else {
			// In directed graph each edge should be used once in each direction
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

		key := ep.makeEdgeKey(edge.From, edge.To)
		ep.edgeCount[key]++
		ep.dfs(edge.To)
	}
	ep.path = append(ep.path, v)
}

// FindEulerPath finds an Euler path in the graph if it exists
func (ep *EulerPath) FindEulerPath() []int {
	if !ep.HasEulerPath() {
		return nil
	}

	// Find starting vertex
	start := ep.findStartVertex()
	ep.dfs(start)

	// Reverse the path (DFS result in reverse order)
	for i, j := 0, len(ep.path)-1; i < j; i, j = i+1, j-1 {
		ep.path[i], ep.path[j] = ep.path[j], ep.path[i]
	}

	return ep.path
}

// FindEulerCircuit finds an Euler circuit in the graph if it exists
func (ep *EulerPath) FindEulerCircuit() []int {
	if !ep.HasEulerCircuit() {
		return nil
	}
	return ep.FindEulerPath()
}

// HasEulerPath checks if the graph has an Euler path
func (ep *EulerPath) HasEulerPath() bool {
	if !ep.isConnected() {
		return false
	}

	oddCount := 0
	for v := 0; v < ep.graph.GetVertices(); v++ {
		degree := len(ep.graph.adjList[v])
		if degree%2 != 0 {
			oddCount++
		}
	}

	if ep.graph.IsDirected() {
		// Check in-out degree for directed graph
		inDegree := make([]int, ep.graph.GetVertices())
		outDegree := make([]int, ep.graph.GetVertices())

		for v := 0; v < ep.graph.GetVertices(); v++ {
			outDegree[v] = len(ep.graph.adjList[v])
			for _, edge := range ep.graph.adjList[v] {
				inDegree[edge.To]++
			}
		}

		diffCount := 0
		for v := 0; v < ep.graph.GetVertices(); v++ {
			diff := outDegree[v] - inDegree[v]
			if diff > 1 || diff < -1 {
				return false
			}
			if diff != 0 {
				diffCount++
			}
		}
		return diffCount == 0 || diffCount == 2
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

		// Find node with out degree 1 more than in degree
		for v := 0; v < ep.graph.GetVertices(); v++ {
			if len(ep.graph.adjList[v])-inDegree[v] == 1 {
				return v
			}
		}
		// If not found, start from any node
		return 0
	}

	// Find node with odd degree in undirected graph
	for v := 0; v < ep.graph.GetVertices(); v++ {
		if len(ep.graph.adjList[v])%2 != 0 {
			return v
		}
	}
	// If not found, start from any node
	return 0
}
