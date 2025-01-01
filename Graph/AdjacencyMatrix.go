package graph

import "math"

// AdjMatrix represents a graph using adjacency matrix
type AdjMatrix struct {
	vertices int
	directed bool
	matrix   [][]int
}

// NewAdjMatrix creates a new graph with adjacency matrix representation
func NewAdjMatrix(vertices int, directed bool) *AdjMatrix {
	// Create matrix and initialize values
	matrix := make([][]int, vertices)
	for i := range matrix {
		matrix[i] = make([]int, vertices)
		for j := range matrix[i] {
			if i != j {
				matrix[i][j] = math.MaxInt32 // Infinite value
			}
		}
	}

	return &AdjMatrix{
		vertices: vertices,
		directed: directed,
		matrix:   matrix,
	}
}

// AddEdge adds an edge between vertices v1 and v2 with given weight
func (g *AdjMatrix) AddEdge(v1, v2, weight int) {
	g.matrix[v1][v2] = weight
	if !g.directed {
		g.matrix[v2][v1] = weight
	}
}

// GetWeight returns the weight of the edge between v1 and v2
func (g *AdjMatrix) GetWeight(v1, v2 int) int {
	return g.matrix[v1][v2]
}

// FloydWarshall finds shortest paths between all pairs of vertices
func (g *AdjMatrix) FloydWarshall() [][]int {
	dist := make([][]int, g.vertices)
	for i := range dist {
		dist[i] = make([]int, g.vertices)
		copy(dist[i], g.matrix[i])
	}

	// Floyd-Warshall algorithm
	for k := 0; k < g.vertices; k++ {
		for i := 0; i < g.vertices; i++ {
			for j := 0; j < g.vertices; j++ {
				if dist[i][k] != math.MaxInt32 && dist[k][j] != math.MaxInt32 {
					newDist := dist[i][k] + dist[k][j]
					if newDist < dist[i][j] {
						dist[i][j] = newDist
					}
				}
			}
		}
	}

	return dist
}

// GetNeighbors returns all neighbors of a vertex
func (g *AdjMatrix) GetNeighbors(vertex int) []int {
	neighbors := make([]int, 0)
	for i := 0; i < g.vertices; i++ {
		if g.matrix[vertex][i] != math.MaxInt32 && vertex != i {
			neighbors = append(neighbors, i)
		}
	}
	return neighbors
}

// GetVertices returns the number of vertices
func (g *AdjMatrix) GetVertices() int {
	return g.vertices
}

// IsDirected returns whether the graph is directed
func (g *AdjMatrix) IsDirected() bool {
	return g.directed
}
