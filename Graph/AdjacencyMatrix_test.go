package Graph

import (
	"math"
	"reflect"
	"testing"
)

func TestAdjMatrix(t *testing.T) {
	t.Run("Basic Operations", func(t *testing.T) {
		g := NewAdjMatrix(4, false)

		// Kenar ekleme testi
		g.AddEdge(0, 1, 5)
		g.AddEdge(1, 2, 3)
		g.AddEdge(2, 3, 1)

		// Ağırlık kontrolü
		if weight := g.GetWeight(0, 1); weight != 5 {
			t.Errorf("Expected weight 5, got %d", weight)
		}

		// Yönsüz graf kontrolü
		if weight := g.GetWeight(1, 0); weight != 5 {
			t.Errorf("Expected weight 5 for reverse edge, got %d", weight)
		}

		// Komşuluk kontrolü
		neighbors := g.GetNeighbors(1)
		expectedNeighbors := []int{0, 2}
		if !reflect.DeepEqual(neighbors, expectedNeighbors) {
			t.Errorf("Expected neighbors %v, got %v", expectedNeighbors, neighbors)
		}
	})

	t.Run("Floyd-Warshall Algorithm", func(t *testing.T) {
		g := NewAdjMatrix(4, true)

		g.AddEdge(0, 1, 5)
		g.AddEdge(1, 2, 3)
		g.AddEdge(2, 3, 1)
		g.AddEdge(0, 3, 10)

		dist := g.FloydWarshall()

		// 0'dan 3'e en kısa yol kontrolü
		// 0->1->2->3 yolu (toplam: 9) 0->3 yolundan (10) daha kısa olmalı
		if dist[0][3] != 9 {
			t.Errorf("Expected shortest path length 9, got %d", dist[0][3])
		}
	})

	t.Run("Directed Graph", func(t *testing.T) {
		g := NewAdjMatrix(3, true)

		g.AddEdge(0, 1, 2)
		g.AddEdge(1, 2, 3)

		// Yönlü graf kontrolü
		if !g.IsDirected() {
			t.Error("Expected directed graph")
		}

		// Tek yönlü kenar kontrolü
		if weight := g.GetWeight(1, 0); weight != math.MaxInt32 {
			t.Errorf("Expected no reverse edge (MaxInt32), got %d", weight)
		}
	})
}
