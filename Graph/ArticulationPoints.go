package Graph

// ArticulationPoints implements algorithms for finding articulation points and bridges
type ArticulationPoints struct {
	graph   *Graph
	time    int
	disc    []int  // Discovery times
	low     []int  // Lowest discovery times
	parent  []int  // Parent nodes in DFS tree
	ap      []bool // Articulation points
	bridges []Edge // Bridges
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

	// Initialize arrays
	for i := 0; i < n; i++ {
		ap.disc[i] = -1
		ap.low[i] = -1
		ap.parent[i] = -1
		ap.visited[i] = false
		ap.ap[i] = false
	}

	// Call DFS for each connected component
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

// dfs performs depth first search to find articulation points and bridges
func (ap *ArticulationPoints) dfs(u int) {
	children := 0
	ap.visited[u] = true
	ap.disc[u] = ap.time
	ap.low[u] = ap.time
	ap.time++

	// Visit neighbors
	for _, edge := range ap.graph.adjList[u] {
		v := edge.To

		if !ap.visited[v] {
			children++
			ap.parent[v] = u
			ap.dfs(v)

			// Update u's low value
			if ap.low[v] < ap.low[u] {
				ap.low[u] = ap.low[v]
			}

			// Special case for root node
			if ap.parent[u] == -1 && children > 1 {
				ap.ap[u] = true
			}

			// Special case for non-root node
			if ap.parent[u] != -1 && ap.low[v] >= ap.disc[u] {
				ap.ap[u] = true
			}

			// Bridge check
			if ap.low[v] > ap.disc[u] {
				ap.bridges = append(ap.bridges, Edge{
					From: u,
					To:   v,
				})
			}
		} else if v != ap.parent[u] {
			// Back edge case
			if ap.disc[v] < ap.low[u] {
				ap.low[u] = ap.disc[v]
			}
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
