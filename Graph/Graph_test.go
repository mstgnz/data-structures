package Graph

import (
	"reflect"
	"testing"
)

func TestGraph(t *testing.T) {
	// Undirected graph test
	t.Run("Undirected Graph Basic Operations", func(t *testing.T) {
		g := NewGraph(5, false)

		// Edge addition test
		g.AddEdge(0, 1, 1)
		g.AddEdge(0, 2, 2)
		g.AddEdge(1, 2, 3)
		g.AddEdge(2, 3, 4)
		g.AddEdge(3, 4, 5)

		// Neighbors check
		neighbors := g.GetNeighbors(0)
		expectedNeighbors := []int{1, 2}
		if !reflect.DeepEqual(neighbors, expectedNeighbors) {
			t.Errorf("Expected neighbors %v, got %v", expectedNeighbors, neighbors)
		}

		// BFS test
		bfsResult := g.BFS(0)
		expectedBFS := []int{0, 1, 2, 3, 4}
		if !reflect.DeepEqual(bfsResult, expectedBFS) {
			t.Errorf("Expected BFS %v, got %v", expectedBFS, bfsResult)
		}

		// DFS test
		dfsResult := g.DFS(0)
		expectedDFS := []int{0, 1, 2, 3, 4}
		if !reflect.DeepEqual(dfsResult, expectedDFS) {
			t.Errorf("Expected DFS %v, got %v", expectedDFS, dfsResult)
		}
	})

	// Dijkstra test
	t.Run("Dijkstra's Algorithm", func(t *testing.T) {
		g := NewGraph(5, true)
		g.AddEdge(0, 1, 4)
		g.AddEdge(0, 2, 1)
		g.AddEdge(2, 1, 2)
		g.AddEdge(1, 3, 1)
		g.AddEdge(2, 3, 5)
		g.AddEdge(3, 4, 3)

		distances := g.Dijkstra(0)
		expected := map[int]int{
			0: 0, // Başlangıç noktası
			1: 3, // 0->2->1
			2: 1, // 0->2
			3: 4, // 0->2->1->3
			4: 7, // 0->2->1->3->4
		}

		if !reflect.DeepEqual(distances, expected) {
			t.Errorf("Expected distances %v, got %v", expected, distances)
		}
	})

	// Kruskal test
	t.Run("Kruskal's Algorithm", func(t *testing.T) {
		g := NewGraph(4, false)
		g.AddEdge(0, 1, 10)
		g.AddEdge(0, 2, 6)
		g.AddEdge(0, 3, 5)
		g.AddEdge(1, 3, 15)
		g.AddEdge(2, 3, 4)

		mst := g.Kruskal()
		totalWeight := 0
		for _, edge := range mst {
			totalWeight += edge.Weight
		}

		expectedWeight := 19 // 4 + 5 + 10
		if totalWeight != expectedWeight {
			t.Errorf("Expected MST weight %d, got %d", expectedWeight, totalWeight)
		}
	})

	// Prim test
	t.Run("Prim's Algorithm", func(t *testing.T) {
		g := NewGraph(4, false)
		g.AddEdge(0, 1, 10)
		g.AddEdge(0, 2, 6)
		g.AddEdge(0, 3, 5)
		g.AddEdge(1, 3, 15)
		g.AddEdge(2, 3, 4)

		mst := g.Prim(0)
		totalWeight := 0
		for _, edge := range mst {
			totalWeight += edge.Weight
		}

		expectedWeight := 19 // 4 + 5 + 10
		if totalWeight != expectedWeight {
			t.Errorf("Expected MST weight %d, got %d", expectedWeight, totalWeight)
		}
	})

	// Directed graph test
	t.Run("Directed Graph", func(t *testing.T) {
		g := NewGraph(4, true)

		g.AddEdge(0, 1, 1)
		g.AddEdge(1, 2, 1)
		g.AddEdge(2, 3, 1)
		g.AddEdge(3, 0, 1)

		if !g.IsDirected() {
			t.Error("Expected directed graph")
		}

		neighbors := g.GetNeighbors(0)
		expectedNeighbors := []int{1}
		if !reflect.DeepEqual(neighbors, expectedNeighbors) {
			t.Errorf("Expected neighbors %v, got %v", expectedNeighbors, neighbors)
		}

		// MST algorithms should not work on directed graphs
		if mst := g.Kruskal(); mst != nil {
			t.Error("Kruskal should return nil for directed graphs")
		}

		if mst := g.Prim(0); mst != nil {
			t.Error("Prim should return nil for directed graphs")
		}
	})
}
