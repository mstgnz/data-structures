package Graph

import (
	"container/heap"
	"math"
)

// PrimMST implements Prim's algorithm for finding Minimum Spanning Tree
type PrimMST struct {
	graph    *Graph
	key      []float64 // Anahtar değerleri (minimum ağırlıklar)
	parent   []int     // MST'deki ebeveyn düğümler
	inMST    []bool    // MST'ye dahil olan düğümler
	mstEdges []Edge    // MST'deki kenarlar
	mstCost  float64   // MST'nin toplam maliyeti
	infinity float64
}

// NewPrimMST creates a new Prim's MST instance
func NewPrimMST(g *Graph) *PrimMST {
	if g.IsDirected() {
		return nil // Prim algoritması yönsüz graflar için çalışır
	}
	return &PrimMST{
		graph:    g,
		infinity: math.Inf(1),
		mstCost:  0,
	}
}

// minHeapNode represents a vertex with its key value
type minHeapNode struct {
	vertex int
	key    float64
}

// minHeap is a min-heap of vertices
type minHeap []*minHeapNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].key < h[j].key }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*minHeapNode))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// FindMST finds the Minimum Spanning Tree
func (p *PrimMST) FindMST() bool {
	n := p.graph.GetVertices()
	p.key = make([]float64, n)
	p.parent = make([]int, n)
	p.inMST = make([]bool, n)
	p.mstEdges = make([]Edge, 0)
	p.mstCost = 0

	// Tüm anahtarları sonsuz olarak başlat
	for i := 0; i < n; i++ {
		p.key[i] = p.infinity
		p.parent[i] = -1
	}

	// İlk düğümü başlangıç olarak seç
	p.key[0] = 0

	// Min-heap oluştur
	h := &minHeap{}
	heap.Init(h)
	heap.Push(h, &minHeapNode{0, 0})

	// MST'yi oluştur
	for h.Len() > 0 {
		// En küçük anahtarlı düğümü al
		u := heap.Pop(h).(*minHeapNode).vertex
		if p.inMST[u] {
			continue
		}

		p.inMST[u] = true

		// Eğer bu ilk düğüm değilse, MST'ye kenar ekle
		if u != 0 {
			p.mstEdges = append(p.mstEdges, Edge{
				From:   p.parent[u],
				To:     u,
				Weight: int(p.key[u]),
			})
			p.mstCost += p.key[u]
		}

		// Komşu düğümleri güncelle
		for _, edge := range p.graph.adjList[u] {
			v := edge.To
			weight := float64(edge.Weight)

			if !p.inMST[v] && weight < p.key[v] {
				p.key[v] = weight
				p.parent[v] = u
				heap.Push(h, &minHeapNode{v, weight})
			}
		}
	}

	// Tüm düğümlerin MST'ye dahil olup olmadığını kontrol et
	for i := 0; i < n; i++ {
		if !p.inMST[i] {
			return false // Graf bağlantılı değil
		}
	}

	return true
}

// GetMSTEdges returns the edges in the MST
func (p *PrimMST) GetMSTEdges() []Edge {
	return p.mstEdges
}

// GetMSTCost returns the total cost of the MST
func (p *PrimMST) GetMSTCost() float64 {
	return p.mstCost
}

// GetParent returns the parent of a vertex in the MST
func (p *PrimMST) GetParent(v int) int {
	return p.parent[v]
}

// IsInMST checks if a vertex is in the MST
func (p *PrimMST) IsInMST(v int) bool {
	return p.inMST[v]
}
