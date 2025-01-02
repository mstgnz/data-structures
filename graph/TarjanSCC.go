package graph

import (
	"sort"
	"sync"
)

// TarjanSCC implements Tarjan's algorithm for finding Strongly Connected Components
type TarjanSCC struct {
	graph      *Graph
	index      int
	stack      []int
	inStack    []bool
	indices    []int
	lowLink    []int
	components [][]int
	mutex      sync.RWMutex
}

// NewTarjanSCC creates a new Tarjan's SCC instance
func NewTarjanSCC(g *Graph) *TarjanSCC {
	if !g.IsDirected() {
		return nil // Tarjan algorithm works for directed graphs
	}
	n := g.GetVertices()
	return &TarjanSCC{
		graph:      g,
		index:      0,
		stack:      make([]int, 0),
		inStack:    make([]bool, n),
		indices:    make([]int, n),
		lowLink:    make([]int, n),
		components: make([][]int, 0),
		mutex:      sync.RWMutex{},
	}
}

// initialize resets the state for a new computation
func (t *TarjanSCC) initialize() {
	t.index = 0
	t.stack = make([]int, 0)
	n := t.graph.GetVertices()
	t.inStack = make([]bool, n)
	t.indices = make([]int, n)
	t.lowLink = make([]int, n)
	t.components = make([][]int, 0)

	// Initialize arrays
	for i := 0; i < n; i++ {
		t.indices[i] = -1
		t.lowLink[i] = -1
		t.inStack[i] = false
	}
}

// FindComponents finds all strongly connected components
func (t *TarjanSCC) FindComponents() [][]int {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.initialize()

	// Call DFS for each node
	for v := 0; v < t.graph.GetVertices(); v++ {
		if t.indices[v] == -1 {
			t.strongConnect(v)
		}
	}

	// Sort each component
	for i := range t.components {
		sort.Ints(t.components[i])
	}

	// Sort components by size and content
	sort.Slice(t.components, func(i, j int) bool {
		// First compare by size (larger components first)
		if len(t.components[i]) != len(t.components[j]) {
			return len(t.components[i]) > len(t.components[j])
		}
		// If sizes are equal, compare by content
		for k := 0; k < len(t.components[i]); k++ {
			if t.components[i][k] != t.components[j][k] {
				return t.components[i][k] < t.components[j][k]
			}
		}
		return false
	})

	// Create a copy of the components
	result := make([][]int, len(t.components))
	for i, comp := range t.components {
		result[i] = make([]int, len(comp))
		copy(result[i], comp)
	}

	return result
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
		if t.indices[w] == -1 {
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

	// Return a copy of the components
	result := make([][]int, len(t.components))
	for i, comp := range t.components {
		result[i] = make([]int, len(comp))
		copy(result[i], comp)
	}

	return result
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

	// Find the largest component
	largest := t.components[0]
	for _, comp := range t.components {
		if len(comp) > len(largest) {
			largest = comp
		}
	}

	// Return a copy of the largest component
	result := make([]int, len(largest))
	copy(result, largest)
	return result
}
