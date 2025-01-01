package graph

import "sync"

// TarjanSCC implements Tarjan's algorithm for finding Strongly Connected Components
type TarjanSCC struct {
	graph      *Graph
	index      int
	stack      []int
	inStack    map[int]bool
	indices    map[int]int
	lowLink    map[int]int
	components [][]int
	mutex      sync.RWMutex
}

// NewTarjanSCC creates a new Tarjan's SCC instance
func NewTarjanSCC(g *Graph) *TarjanSCC {
	if !g.IsDirected() {
		return nil // Tarjan algorithm works for directed graphs
	}
	return &TarjanSCC{
		graph:      g,
		index:      0,
		stack:      make([]int, 0),
		inStack:    make(map[int]bool),
		indices:    make(map[int]int),
		lowLink:    make(map[int]int),
		components: make([][]int, 0),
		mutex:      sync.RWMutex{},
	}
}

// FindComponents finds all strongly connected components
func (t *TarjanSCC) FindComponents() [][]int {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	// Call DFS for each node
	for v := 0; v < t.graph.GetVertices(); v++ {
		if _, exists := t.indices[v]; !exists {
			t.strongConnect(v)
		}
	}
	return t.components
}

// strongConnect performs the recursive part of Tarjan's algorithm
func (t *TarjanSCC) strongConnect(v int) {
	// Initialize v
	t.indices[v] = t.index
	t.lowLink[v] = t.index
	t.index++
	t.stack = append(t.stack, v)
	t.inStack[v] = true

	// Visit neighbors of v
	for _, edge := range t.graph.adjList[v] {
		w := edge.To
		if _, exists := t.indices[w]; !exists {
			// w has not been visited yet
			t.strongConnect(w)
			// Update v's lowLink value
			if t.lowLink[w] < t.lowLink[v] {
				t.lowLink[v] = t.lowLink[w]
			}
		} else if t.inStack[w] {
			// w is in the stack
			if t.indices[w] < t.lowLink[v] {
				t.lowLink[v] = t.indices[w]
			}
		}
	}

	// Check if v is the root of an SCC
	if t.lowLink[v] == t.indices[v] {
		// Create a new SCC
		component := make([]int, 0)
		for {
			w := t.stack[len(t.stack)-1]
			t.stack = t.stack[:len(t.stack)-1]
			t.inStack[w] = false
			component = append(component, w)
			if w == v {
				break
			}
		}
		t.components = append(t.components, component)
	}
}

// GetComponents returns all found components
func (t *TarjanSCC) GetComponents() [][]int {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if len(t.components) == 0 {
		return t.FindComponents()
	}
	return t.components
}

// GetComponentCount returns the number of strongly connected components
func (t *TarjanSCC) GetComponentCount() int {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if len(t.components) == 0 {
		t.FindComponents()
	}
	return len(t.components)
}

// IsStronglyConnected checks if the graph is strongly connected
func (t *TarjanSCC) IsStronglyConnected() bool {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if len(t.components) == 0 {
		t.FindComponents()
	}
	return len(t.components) == 1
}

// GetLargestComponent returns the largest strongly connected component
func (t *TarjanSCC) GetLargestComponent() []int {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if len(t.components) == 0 {
		t.FindComponents()
	}

	if len(t.components) == 0 {
		return nil
	}

	largest := t.components[0]
	for _, comp := range t.components {
		if len(comp) > len(largest) {
			largest = comp
		}
	}
	return largest
}
