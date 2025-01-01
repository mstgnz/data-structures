package Graph

// TarjanSCC implements Tarjan's algorithm for finding Strongly Connected Components
type TarjanSCC struct {
	graph      *Graph
	index      int
	stack      []int
	inStack    map[int]bool
	indices    map[int]int
	lowLink    map[int]int
	components [][]int
}

// NewTarjanSCC creates a new Tarjan's SCC instance
func NewTarjanSCC(g *Graph) *TarjanSCC {
	if !g.IsDirected() {
		return nil // Tarjan algoritması yönlü graflar için çalışır
	}
	return &TarjanSCC{
		graph:      g,
		index:      0,
		stack:      make([]int, 0),
		inStack:    make(map[int]bool),
		indices:    make(map[int]int),
		lowLink:    make(map[int]int),
		components: make([][]int, 0),
	}
}

// FindComponents finds all strongly connected components
func (t *TarjanSCC) FindComponents() [][]int {
	// Her düğüm için DFS çağır
	for v := 0; v < t.graph.GetVertices(); v++ {
		if _, exists := t.indices[v]; !exists {
			t.strongConnect(v)
		}
	}
	return t.components
}

// strongConnect performs the recursive part of Tarjan's algorithm
func (t *TarjanSCC) strongConnect(v int) {
	// v'yi başlat
	t.indices[v] = t.index
	t.lowLink[v] = t.index
	t.index++
	t.stack = append(t.stack, v)
	t.inStack[v] = true

	// v'nin komşularını ziyaret et
	for _, edge := range t.graph.adjList[v] {
		w := edge.To
		if _, exists := t.indices[w]; !exists {
			// w henüz ziyaret edilmemiş
			t.strongConnect(w)
			// v'nin lowLink değerini güncelle
			if t.lowLink[w] < t.lowLink[v] {
				t.lowLink[v] = t.lowLink[w]
			}
		} else if t.inStack[w] {
			// w stack'te
			if t.indices[w] < t.lowLink[v] {
				t.lowLink[v] = t.indices[w]
			}
		}
	}

	// v bir SCC'nin kökü mü kontrol et
	if t.lowLink[v] == t.indices[v] {
		// Yeni bir SCC oluştur
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
	if len(t.components) == 0 {
		return t.FindComponents()
	}
	return t.components
}

// GetComponentCount returns the number of strongly connected components
func (t *TarjanSCC) GetComponentCount() int {
	if len(t.components) == 0 {
		t.FindComponents()
	}
	return len(t.components)
}

// IsStronglyConnected checks if the graph is strongly connected
func (t *TarjanSCC) IsStronglyConnected() bool {
	if len(t.components) == 0 {
		t.FindComponents()
	}
	return len(t.components) == 1
}

// GetLargestComponent returns the largest strongly connected component
func (t *TarjanSCC) GetLargestComponent() []int {
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
