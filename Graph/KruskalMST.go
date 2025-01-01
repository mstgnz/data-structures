package Graph

import "sort"

// KruskalMST implements Kruskal's algorithm for finding Minimum Spanning Tree
type KruskalMST struct {
	graph    *Graph
	parent   []int   // Union-Find için ebeveyn dizisi
	rank     []int   // Union-Find için rank dizisi
	mstEdges []Edge  // MST'deki kenarlar
	mstCost  float64 // MST'nin toplam maliyeti
}

// NewKruskalMST creates a new Kruskal's MST instance
func NewKruskalMST(g *Graph) *KruskalMST {
	if g.IsDirected() {
		return nil // Kruskal algoritması yönsüz graflar için çalışır
	}
	return &KruskalMST{
		graph:   g,
		mstCost: 0,
	}
}

// initialize prepares the Union-Find data structure
func (k *KruskalMST) initialize() {
	n := k.graph.GetVertices()
	k.parent = make([]int, n)
	k.rank = make([]int, n)
	k.mstEdges = make([]Edge, 0)

	// Her düğümü kendi kümesinde başlat
	for i := 0; i < n; i++ {
		k.parent[i] = i
		k.rank[i] = 0
	}
}

// find returns the representative of the set containing x
func (k *KruskalMST) find(x int) int {
	if k.parent[x] != x {
		k.parent[x] = k.find(k.parent[x]) // Yol sıkıştırma
	}
	return k.parent[x]
}

// union merges sets containing x and y
func (k *KruskalMST) union(x, y int) {
	rootX := k.find(x)
	rootY := k.find(y)

	if rootX != rootY {
		// Rank'a göre birleştirme
		if k.rank[rootX] < k.rank[rootY] {
			k.parent[rootX] = rootY
		} else if k.rank[rootX] > k.rank[rootY] {
			k.parent[rootY] = rootX
		} else {
			k.parent[rootY] = rootX
			k.rank[rootX]++
		}
	}
}

// FindMST finds the Minimum Spanning Tree
func (k *KruskalMST) FindMST() bool {
	k.initialize()

	// Tüm kenarları topla ve ağırlığa göre sırala
	edges := make([]Edge, 0)
	for v := 0; v < k.graph.GetVertices(); v++ {
		for _, edge := range k.graph.adjList[v] {
			// Yönsüz grafta her kenarı bir kez ekle
			if edge.From < edge.To {
				edges = append(edges, edge)
			}
		}
	}

	// Kenarları ağırlığa göre sırala
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	edgeCount := 0
	// Kenarları MST'ye ekle
	for _, edge := range edges {
		if k.find(edge.From) != k.find(edge.To) {
			k.union(edge.From, edge.To)
			k.mstEdges = append(k.mstEdges, edge)
			k.mstCost += float64(edge.Weight)
			edgeCount++
		}
	}

	// MST tam olarak oluştu mu kontrol et
	return edgeCount == k.graph.GetVertices()-1
}

// GetMSTEdges returns the edges in the MST
func (k *KruskalMST) GetMSTEdges() []Edge {
	return k.mstEdges
}

// GetMSTCost returns the total cost of the MST
func (k *KruskalMST) GetMSTCost() float64 {
	return k.mstCost
}

// IsConnected checks if two vertices are connected in the MST
func (k *KruskalMST) IsConnected(x, y int) bool {
	return k.find(x) == k.find(y)
}

// GetNumComponents returns the number of connected components
func (k *KruskalMST) GetNumComponents() int {
	components := make(map[int]bool)
	for v := 0; v < k.graph.GetVertices(); v++ {
		components[k.find(v)] = true
	}
	return len(components)
}
