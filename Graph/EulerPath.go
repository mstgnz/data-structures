package Graph

// EulerPath implements algorithms for finding Euler paths and circuits
type EulerPath struct {
	graph     *Graph
	path      []int
	edgeCount map[string]int // Kullanılan kenar sayısı
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
			// Yönsüz grafta her kenar bir kez kullanılmalı
			if count == 0 {
				return &edge
			}
		} else {
			// Yönlü grafta her kenar yönüne göre bir kez kullanılmalı
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

	// Başlangıç düğümünü bul
	start := ep.findStartVertex()
	ep.dfs(start)

	// Yolu ters çevir (DFS sonucu ters sırada)
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
		// Yönlü graf için giriş-çıkış derecesi kontrolü
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
		// Yönlü graf için giriş-çıkış derecesi kontrolü
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

	// Yönsüz graf için tüm düğümlerin çift dereceli olması gerekir
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

	// İlk düğümden DFS başlat
	start := 0
	for v := 0; v < ep.graph.GetVertices(); v++ {
		if len(ep.graph.adjList[v]) > 0 {
			start = v
			break
		}
	}

	ep.dfsUtil(start, visited)

	// Tüm düğümlerin ziyaret edilip edilmediğini kontrol et
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

		// Çıkış derecesi giriş derecesinden 1 fazla olan düğümü bul
		for v := 0; v < ep.graph.GetVertices(); v++ {
			if len(ep.graph.adjList[v])-inDegree[v] == 1 {
				return v
			}
		}
		// Bulunamazsa herhangi bir düğümden başla
		return 0
	}

	// Yönsüz graf için tek dereceli düğümü bul
	for v := 0; v < ep.graph.GetVertices(); v++ {
		if len(ep.graph.adjList[v])%2 != 0 {
			return v
		}
	}
	// Bulunamazsa herhangi bir düğümden başla
	return 0
}
