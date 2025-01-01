package graph

import "sync"

// HamiltonianPath implements algorithms for finding Hamiltonian paths and circuits
type HamiltonianPath struct {
	graph   *Graph
	path    []int
	visited []bool
	mutex   sync.RWMutex
}

// NewHamiltonianPath creates a new HamiltonianPath instance
func NewHamiltonianPath(g *Graph) *HamiltonianPath {
	n := g.GetVertices()
	return &HamiltonianPath{
		graph:   g,
		path:    make([]int, 0),
		visited: make([]bool, n),
		mutex:   sync.RWMutex{},
	}
}

// FindHamiltonianPath finds a Hamiltonian path in the graph if it exists
func (hp *HamiltonianPath) FindHamiltonianPath() []int {
	hp.mutex.Lock()
	defer hp.mutex.Unlock()

	n := hp.graph.GetVertices()
	hp.path = make([]int, 0)
	hp.visited = make([]bool, n)

	// Try starting from each node
	for start := 0; start < n; start++ {
		hp.path = []int{start}
		hp.visited = make([]bool, n)
		hp.visited[start] = true

		if hp.hamiltonianPathUtil(1) {
			return hp.path
		}
	}

	return nil
}

// FindHamiltonianCircuit finds a Hamiltonian circuit in the graph if it exists
func (hp *HamiltonianPath) FindHamiltonianCircuit() []int {
	hp.mutex.Lock()
	defer hp.mutex.Unlock()

	n := hp.graph.GetVertices()
	hp.path = make([]int, 0)
	hp.visited = make([]bool, n)

	// Start from 0
	hp.path = []int{0}
	hp.visited[0] = true

	if hp.hamiltonianCircuitUtil(1) {
		return hp.path
	}

	return nil
}

// hamiltonianPathUtil performs backtracking to find Hamiltonian path
func (hp *HamiltonianPath) hamiltonianPathUtil(pos int) bool {
	// All nodes visited?
	if pos == hp.graph.GetVertices() {
		return true
	}

	// Check neighbors of the last added node
	lastVertex := hp.path[len(hp.path)-1]
	for _, edge := range hp.graph.adjList[lastVertex] {
		if !hp.visited[edge.To] {
			hp.visited[edge.To] = true
			hp.path = append(hp.path, edge.To)

			if hp.hamiltonianPathUtil(pos + 1) {
				return true
			}

			// Backtrack
			hp.visited[edge.To] = false
			hp.path = hp.path[:len(hp.path)-1]
		}
	}

	return false
}

// hamiltonianCircuitUtil performs backtracking to find Hamiltonian circuit
func (hp *HamiltonianPath) hamiltonianCircuitUtil(pos int) bool {
	// All nodes visited?
	if pos == hp.graph.GetVertices() {
		// Check if there's an edge from the last node to the start node
		lastVertex := hp.path[len(hp.path)-1]
		hasEdgeToStart := false
		for _, edge := range hp.graph.adjList[lastVertex] {
			if edge.To == hp.path[0] {
				hasEdgeToStart = true
				break
			}
		}
		if hasEdgeToStart {
			hp.path = append(hp.path, hp.path[0])
			return true
		}
		return false
	}

	// Check neighbors of the last added node
	lastVertex := hp.path[len(hp.path)-1]
	for _, edge := range hp.graph.adjList[lastVertex] {
		if !hp.visited[edge.To] {
			hp.visited[edge.To] = true
			hp.path = append(hp.path, edge.To)

			if hp.hamiltonianCircuitUtil(pos + 1) {
				return true
			}

			// Backtrack
			hp.visited[edge.To] = false
			hp.path = hp.path[:len(hp.path)-1]
		}
	}

	return false
}

// HasHamiltonianPath checks if the graph has a Hamiltonian path
func (hp *HamiltonianPath) HasHamiltonianPath() bool {
	hp.mutex.Lock()
	defer hp.mutex.Unlock()

	path := hp.FindHamiltonianPath()
	return path != nil
}

// HasHamiltonianCircuit checks if the graph has a Hamiltonian circuit
func (hp *HamiltonianPath) HasHamiltonianCircuit() bool {
	hp.mutex.Lock()
	defer hp.mutex.Unlock()

	circuit := hp.FindHamiltonianCircuit()
	return circuit != nil
}

// GetPath returns the last found path
func (hp *HamiltonianPath) GetPath() []int {
	hp.mutex.RLock()
	defer hp.mutex.RUnlock()
	return hp.path
}

// IsHamiltonianPath checks if a given path is a valid Hamiltonian path
func (hp *HamiltonianPath) IsHamiltonianPath(path []int) bool {
	hp.mutex.RLock()
	defer hp.mutex.RUnlock()

	if len(path) != hp.graph.GetVertices() {
		return false
	}

	// Check if each node is used once
	visited := make([]bool, hp.graph.GetVertices())
	for _, v := range path {
		if visited[v] {
			return false
		}
		visited[v] = true
	}

	// Check if there's an edge between consecutive nodes
	for i := 0; i < len(path)-1; i++ {
		hasEdge := false
		for _, edge := range hp.graph.adjList[path[i]] {
			if edge.To == path[i+1] {
				hasEdge = true
				break
			}
		}
		if !hasEdge {
			return false
		}
	}

	return true
}

// IsHamiltonianCircuit checks if a given circuit is a valid Hamiltonian circuit
func (hp *HamiltonianPath) IsHamiltonianCircuit(circuit []int) bool {
	hp.mutex.RLock()
	defer hp.mutex.RUnlock()

	if len(circuit) != hp.graph.GetVertices()+1 {
		return false
	}

	if circuit[0] != circuit[len(circuit)-1] {
		return false
	}

	// Check if the path is a valid Hamiltonian path without the last node
	return hp.IsHamiltonianPath(circuit[:len(circuit)-1])
}
