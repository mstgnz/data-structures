package Graph

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

// Edge represents a weighted edge in the graph
type Edge struct {
	From   int
	To     int
	Weight int
}

// Graph represents a graph data structure
type Graph struct {
	vertices int
	directed bool
	adjList  map[int][]Edge
}

// NewGraph creates a new graph with n vertices
func NewGraph(vertices int, directed bool) *Graph {
	return &Graph{
		vertices: vertices,
		directed: directed,
		adjList:  make(map[int][]Edge),
	}
}

// AddEdge adds an edge between vertices v1 and v2 with given weight
func (g *Graph) AddEdge(v1, v2, weight int) {
	// Komşuluk listesine kenarı ekle
	g.adjList[v1] = append(g.adjList[v1], Edge{To: v2, Weight: weight})

	// Yönsüz graf ise, ters kenarı da ekle
	if !g.directed {
		g.adjList[v2] = append(g.adjList[v2], Edge{To: v1, Weight: weight})
	}
}

// BFS performs Breadth First Search starting from vertex v
func (g *Graph) BFS(start int) []int {
	visited := make(map[int]bool)
	queue := []int{start}
	result := []int{}

	visited[start] = true

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		result = append(result, vertex)

		for _, edge := range g.adjList[vertex] {
			if !visited[edge.To] {
				visited[edge.To] = true
				queue = append(queue, edge.To)
			}
		}
	}

	return result
}

// DFS performs Depth First Search starting from vertex v
func (g *Graph) DFS(start int) []int {
	visited := make(map[int]bool)
	result := []int{}
	g.dfsUtil(start, visited, &result)
	return result
}

func (g *Graph) dfsUtil(vertex int, visited map[int]bool, result *[]int) {
	visited[vertex] = true
	*result = append(*result, vertex)

	for _, edge := range g.adjList[vertex] {
		if !visited[edge.To] {
			g.dfsUtil(edge.To, visited, result)
		}
	}
}

// Dijkstra finds shortest paths from source vertex to all other vertices
func (g *Graph) Dijkstra(source int) map[int]int {
	distances := make(map[int]int)
	for i := 0; i < g.vertices; i++ {
		distances[i] = math.MaxInt32
	}
	distances[source] = 0

	// Priority queue için yardımcı yapılar
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{vertex: source, priority: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		vertex := current.vertex

		// Eğer daha kısa bir yol bulunmuşsa, bu düğümü atla
		if current.priority > distances[vertex] {
			continue
		}

		// Komşuları kontrol et
		for _, edge := range g.adjList[vertex] {
			distance := distances[vertex] + edge.Weight
			if distance < distances[edge.To] {
				distances[edge.To] = distance
				heap.Push(pq, &Item{vertex: edge.To, priority: distance})
			}
		}
	}

	return distances
}

// Kruskal finds Minimum Spanning Tree using Kruskal's algorithm
func (g *Graph) Kruskal() []Edge {
	if g.directed {
		return nil // Kruskal sadece yönsüz graflar için çalışır
	}

	// Tüm kenarları topla ve ağırlığa göre sırala
	edges := g.getAllEdges()
	result := make([]Edge, 0)

	// Union-Find veri yapısını başlat
	uf := NewUnionFind(g.vertices)

	// Kenarları ağırlığa göre sırala
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	// En küçük ağırlıklı kenarları seç
	edgeCount := 0
	for _, edge := range edges {
		if !uf.Connected(edge.From, edge.To) {
			uf.Union(edge.From, edge.To)
			result = append(result, edge)
			edgeCount++
			if edgeCount == g.vertices-1 {
				break
			}
		}
	}

	return result
}

// Prim finds Minimum Spanning Tree using Prim's algorithm
func (g *Graph) Prim(start int) []Edge {
	if g.directed {
		return nil // Prim sadece yönsüz graflar için çalışır
	}

	visited := make(map[int]bool)
	result := make([]Edge, 0)

	// Priority queue başlat
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Başlangıç düğümünden başla
	visited[start] = true
	for _, edge := range g.adjList[start] {
		heap.Push(pq, &Item{vertex: edge.To, priority: edge.Weight, from: start})
	}

	for pq.Len() > 0 && len(result) < g.vertices-1 {
		item := heap.Pop(pq).(*Item)
		if visited[item.vertex] {
			continue
		}

		// Kenarı MST'ye ekle
		visited[item.vertex] = true
		result = append(result, Edge{From: item.from, To: item.vertex, Weight: item.priority})

		// Yeni düğümün komşularını queue'ya ekle
		for _, edge := range g.adjList[item.vertex] {
			if !visited[edge.To] {
				heap.Push(pq, &Item{vertex: edge.To, priority: edge.Weight, from: item.vertex})
			}
		}
	}

	return result
}

// GetNeighbors returns all neighbors of a vertex
func (g *Graph) GetNeighbors(vertex int) []int {
	neighbors := make([]int, 0)
	for _, edge := range g.adjList[vertex] {
		neighbors = append(neighbors, edge.To)
	}
	return neighbors
}

// GetVertices returns the number of vertices
func (g *Graph) GetVertices() int {
	return g.vertices
}

// IsDirected returns whether the graph is directed
func (g *Graph) IsDirected() bool {
	return g.directed
}

// getAllEdges returns all edges in the graph
func (g *Graph) getAllEdges() []Edge {
	edges := make([]Edge, 0)
	seen := make(map[string]bool)

	for from, adjEdges := range g.adjList {
		for _, edge := range adjEdges {
			// Yönsüz graf için kenarları sadece bir kez ekle
			key := fmt.Sprintf("%d-%d", min(from, edge.To), max(from, edge.To))
			if !seen[key] {
				edges = append(edges, Edge{From: from, To: edge.To, Weight: edge.Weight})
				seen[key] = true
			}
		}
	}

	// Kenarları ağırlığa göre sırala
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	return edges
}

// Priority Queue implementation for Dijkstra and Prim
type Item struct {
	vertex   int
	priority int
	from     int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// Union-Find implementation for Kruskal's algorithm
type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent: parent, rank: rank}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	px, py := uf.Find(x), uf.Find(y)
	if px == py {
		return
	}
	if uf.rank[px] < uf.rank[py] {
		uf.parent[px] = py
	} else if uf.rank[px] > uf.rank[py] {
		uf.parent[py] = px
	} else {
		uf.parent[py] = px
		uf.rank[px]++
	}
}

func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
