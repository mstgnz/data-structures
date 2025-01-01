package Graph

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
		return nil // SCC sadece yönlü graflarda anlamlıdır
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
	// 1. İlk DFS ile bitiş zamanlarını hesapla
	scc.firstDFS()

	// 2. Grafın transpozunu al
	transpose := scc.getTranspose()

	// 3. İkinci DFS ile bileşenleri bul
	scc.visited = make(map[int]bool) // Ziyaret haritasını sıfırla
	scc.components = make([][]int, 0)

	// Bitiş zamanlarına göre ters sırada DFS çağır
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

	// Komşuları ziyaret et
	for _, edge := range scc.graph.adjList[v] {
		if !scc.visited[edge.To] {
			scc.firstDFSUtil(edge.To)
		}
	}

	// Bitiş zamanını kaydet
	scc.finishTime = append(scc.finishTime, v)
}

// getTranspose returns the transpose of the graph
func (scc *StronglyConnectedComponents) getTranspose() *Graph {
	transpose := NewGraph(scc.graph.GetVertices(), true)

	// Her kenarı ters çevir
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

	// Komşuları ziyaret et
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
