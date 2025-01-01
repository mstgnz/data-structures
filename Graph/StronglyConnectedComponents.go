package graph

// StronglyConnectedComponents implements Kosaraju's algorithm for finding SCCs
type StronglyConnectedComponents struct {
	graph      *Graph
	visited    map[int]bool
	finishTime []int
	components [][]int
}

// NewSCC creates a new SCC instance
func NewSCC(g *Graph) *StronglyConnectedComponents {
	if !g.IsDirected() {
		return nil // SCC is only meaningful for directed graphs
	}
	return &StronglyConnectedComponents{
		graph:      g,
		visited:    make(map[int]bool),
		finishTime: make([]int, 0),
		components: make([][]int, 0),
	}
}

// FindComponents finds all strongly connected components
func (scc *StronglyConnectedComponents) FindComponents() [][]int {
	// 1. First DFS to calculate finish times
	scc.firstDFS()

	// 2. Get the transpose of the graph
	transpose := scc.getTranspose()

	// 3. Second DFS to find components
	scc.visited = make(map[int]bool) // Reset visited map
	scc.components = make([][]int, 0)

	// Call DFS in reverse order of finish times
	for i := len(scc.finishTime) - 1; i >= 0; i-- {
		v := scc.finishTime[i]
		if !scc.visited[v] {
			component := make([]int, 0)
			scc.secondDFS(transpose, v, &component)
			scc.components = append(scc.components, component)
		}
	}

	return scc.components
}

// firstDFS performs first DFS pass to compute finish times
func (scc *StronglyConnectedComponents) firstDFS() {
	for v := 0; v < scc.graph.GetVertices(); v++ {
		if !scc.visited[v] {
			scc.firstDFSUtil(v)
		}
	}
}

func (scc *StronglyConnectedComponents) firstDFSUtil(v int) {
	scc.visited[v] = true

	// Visit neighbors
	for _, edge := range scc.graph.adjList[v] {
		if !scc.visited[edge.To] {
			scc.firstDFSUtil(edge.To)
		}
	}

	// Save finish time
	scc.finishTime = append(scc.finishTime, v)
}

// getTranspose returns the transpose of the graph
func (scc *StronglyConnectedComponents) getTranspose() *Graph {
	transpose := NewGraph(scc.graph.GetVertices(), true)

	// Reverse each edge
	for v := 0; v < scc.graph.GetVertices(); v++ {
		for _, edge := range scc.graph.adjList[v] {
			transpose.AddEdge(edge.To, v, edge.Weight)
		}
	}

	return transpose
}

// secondDFS performs second DFS pass to find components
func (scc *StronglyConnectedComponents) secondDFS(g *Graph, v int, component *[]int) {
	scc.visited[v] = true
	*component = append(*component, v)

	// Visit neighbors
	for _, edge := range g.adjList[v] {
		if !scc.visited[edge.To] {
			scc.secondDFS(g, edge.To, component)
		}
	}
}

// GetComponents returns all found components
func (scc *StronglyConnectedComponents) GetComponents() [][]int {
	if len(scc.components) == 0 {
		return scc.FindComponents()
	}
	return scc.components
}

// GetComponentCount returns the number of strongly connected components
func (scc *StronglyConnectedComponents) GetComponentCount() int {
	if len(scc.components) == 0 {
		scc.FindComponents()
	}
	return len(scc.components)
}
