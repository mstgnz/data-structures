package Graph

import (
	"reflect"
	"testing"
)

func TestGraph(t *testing.T) {
	// Yönsüz graf testi
	t.Run("Undirected Graph Basic Operations", func(t *testing.T) {
		g := NewGraph(5, false)

		// Kenar ekleme testi
		g.AddEdge(0, 1, 1)
		g.AddEdge(0, 2, 2)
		g.AddEdge(1, 2, 3)
		g.AddEdge(2, 3, 4)
		g.AddEdge(3, 4, 5)

		// Komşuluk kontrolü
		neighbors := g.GetNeighbors(0)
		expectedNeighbors := []int{1, 2}
		if !reflect.DeepEqual(neighbors, expectedNeighbors) {
			t.Errorf("Expected neighbors %v, got %v", expectedNeighbors, neighbors)
		}

		// BFS testi
		bfsResult := g.BFS(0)
		expectedBFS := []int{0, 1, 2, 3, 4}
		if !reflect.DeepEqual(bfsResult, expectedBFS) {
			t.Errorf("Expected BFS %v, got %v", expectedBFS, bfsResult)
		}

		// DFS testi
		dfsResult := g.DFS(0)
		expectedDFS := []int{0, 1, 2, 3, 4}
		if !reflect.DeepEqual(dfsResult, expectedDFS) {
			t.Errorf("Expected DFS %v, got %v", expectedDFS, dfsResult)
		}
	})

	// Dijkstra testi
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

	// Kruskal testi
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

	// Prim testi
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

	// Yönlü graf testi
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

		// MST algoritmaları yönlü grafta çalışmamalı
		if mst := g.Kruskal(); mst != nil {
			t.Error("Kruskal should return nil for directed graphs")
		}

		if mst := g.Prim(0); mst != nil {
			t.Error("Prim should return nil for directed graphs")
		}
	})
}
