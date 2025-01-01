package graph

// TopologicalSort performs topological sorting on a directed graph
type TopologicalSort struct {
	graph    *Graph
	visited  map[int]bool
	tempMark map[int]bool // Temporary marking for cycle detection
	order    []int        // Topological sort result
	hasCycle bool         // Has cycle?
}

// NewTopologicalSort creates a new topological sort instance
func NewTopologicalSort(g *Graph) *TopologicalSort {
	if !g.IsDirected() {
		return nil // Topological sort works only for directed graphs
	}
	return &TopologicalSort{
		graph:    g,
		visited:  make(map[int]bool),
		tempMark: make(map[int]bool),
		order:    make([]int, 0),
		hasCycle: false,
	}
}

// Sort performs topological sorting and returns the sorted vertices
// Returns nil if the graph has a cycle
func (ts *TopologicalSort) Sort() []int {
	// Call DFS for each node
	for v := 0; v < ts.graph.GetVertices(); v++ {
		if !ts.visited[v] {
			ts.visit(v)
		}
	}

	if ts.hasCycle {
		return nil
	}

	// Reverse the result (we need to reverse the DFS postorder)
	for i, j := 0, len(ts.order)-1; i < j; i, j = i+1, j-1 {
		ts.order[i], ts.order[j] = ts.order[j], ts.order[i]
	}

	return ts.order
}

// visit performs DFS visit with cycle detection
func (ts *TopologicalSort) visit(v int) {
	if ts.tempMark[v] {
		ts.hasCycle = true
		return
	}
	if ts.visited[v] {
		return
	}

	ts.tempMark[v] = true

	// Visit neighbors
	for _, edge := range ts.graph.adjList[v] {
		ts.visit(edge.To)
	}

	ts.visited[v] = true
	delete(ts.tempMark, v)
	ts.order = append(ts.order, v)
}

// HasCycle returns true if the graph contains a cycle
func (ts *TopologicalSort) HasCycle() bool {
	if ts.order == nil {
		ts.Sort()
	}
	return ts.hasCycle
}

// GetDependencyOrder returns the dependency order of vertices
// For example, if v depends on u, then u will appear before v in the result
func (ts *TopologicalSort) GetDependencyOrder() []int {
	if ts.order == nil {
		return ts.Sort()
	}
	return ts.order
}
