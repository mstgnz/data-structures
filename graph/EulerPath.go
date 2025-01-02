package graph

import (
	"sync"
)

// EulerPath implements algorithms for finding Euler paths and circuits
type EulerPath struct {
	graph   *Graph
	visited map[string]bool
	path    []int
	mutex   sync.RWMutex
}

// NewEulerPath creates a new EulerPath instance
func NewEulerPath(g *Graph) *EulerPath {
	return &EulerPath{
		graph:   g,
		visited: make(map[string]bool),
		path:    make([]int, 0),
		mutex:   sync.RWMutex{},
	}
}

// hierholzer implements Hierholzer's algorithm for finding Euler paths/circuits
func (ep *EulerPath) hierholzer(start int) []int {
	// Create a copy of adjacency list to track remaining edges
	remainingEdges := make([][]Edge, ep.graph.GetVertices())
	for i := 0; i < ep.graph.GetVertices(); i++ {
		remainingEdges[i] = make([]Edge, len(ep.graph.adjList[i]))
		copy(remainingEdges[i], ep.graph.adjList[i])
	}

	// Stack for vertices and final path
	stack := []int{start}
	var path []int

	for len(stack) > 0 {
		current := stack[len(stack)-1]

		// If current vertex has no remaining edges
		if len(remainingEdges[current]) == 0 {
			path = append(path, current)
			stack = stack[:len(stack)-1]
			continue
		}

		// Take the next available edge
		next := remainingEdges[current][0]
		remainingEdges[current] = remainingEdges[current][1:]

		// For undirected graph, remove the reverse edge
		if !ep.graph.IsDirected() {
			for i, edge := range remainingEdges[next.To] {
				if edge.To == current {
					remainingEdges[next.To] = append(remainingEdges[next.To][:i], remainingEdges[next.To][i+1:]...)
					break
				}
			}
		}

		stack = append(stack, next.To)
	}

	// Reverse the path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

// FindEulerPath finds an Euler path in the graph if it exists
func (ep *EulerPath) FindEulerPath() []int {
	ep.mutex.Lock()
	defer ep.mutex.Unlock()

	if !ep.HasEulerPath() {
		return nil
	}

	start := ep.findStartVertex()
	path := ep.hierholzer(start)

	// For Euler circuit, add starting vertex at the end if needed
	if ep.HasEulerCircuit() && len(path) > 0 {
		path = append(path, path[0])
	}

	return path
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
	ep.mutex.RLock()
	defer ep.mutex.RUnlock()

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
	ep.mutex.RLock()
	defer ep.mutex.RUnlock()

	if !ep.isConnected() {
		return false
	}

	if ep.graph.IsDirected() {
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

	// For undirected graph
	for v := 0; v < ep.graph.GetVertices(); v++ {
		if len(ep.graph.adjList[v])%2 != 0 {
			return false
		}
	}
	return true
}

// isConnected checks if the graph is connected
func (ep *EulerPath) isConnected() bool {
	n := ep.graph.GetVertices()
	visited := make([]bool, n)

	// Find first non-zero degree vertex
	start := -1
	for v := 0; v < n; v++ {
		if len(ep.graph.adjList[v]) > 0 {
			start = v
			break
		}
	}

	if start == -1 {
		return true // Empty graph is considered connected
	}

	// Run DFS from start vertex
	ep.dfsUtil(start, visited)

	// Check if all non-zero degree vertices are visited
	for v := 0; v < n; v++ {
		if len(ep.graph.adjList[v]) > 0 && !visited[v] {
			return false
		}
	}
	return true
}

// dfsUtil is a utility function for DFS traversal
func (ep *EulerPath) dfsUtil(v int, visited []bool) {
	visited[v] = true
	for _, edge := range ep.graph.adjList[v] {
		if !visited[edge.To] {
			ep.dfsUtil(edge.To, visited)
		}
	}
}

// findStartVertex finds a valid starting vertex for Euler path
func (ep *EulerPath) findStartVertex() int {
	if !ep.graph.IsDirected() {
		// For undirected graph, start from a vertex with odd degree if exists
		for v := 0; v < ep.graph.GetVertices(); v++ {
			if len(ep.graph.adjList[v])%2 != 0 {
				return v
			}
		}
		// If no odd degree vertex, start from any vertex with non-zero degree
		for v := 0; v < ep.graph.GetVertices(); v++ {
			if len(ep.graph.adjList[v]) > 0 {
				return v
			}
		}
	} else {
		// For directed graph, find vertex with out-degree > in-degree
		inDegree := make([]int, ep.graph.GetVertices())
		for v := 0; v < ep.graph.GetVertices(); v++ {
			for _, edge := range ep.graph.adjList[v] {
				inDegree[edge.To]++
			}
		}

		for v := 0; v < ep.graph.GetVertices(); v++ {
			outDegree := len(ep.graph.adjList[v])
			if outDegree > inDegree[v] {
				return v
			}
		}

		// If no such vertex exists, start from any vertex with non-zero out-degree
		for v := 0; v < ep.graph.GetVertices(); v++ {
			if len(ep.graph.adjList[v]) > 0 {
				return v
			}
		}
	}
	return 0
}
