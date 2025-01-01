package Graph

import (
	"testing"
)

func TestPrimMST(t *testing.T) {
	// Test 1: Basit bağlantılı yönsüz graf
	g := NewGraph(5, false)
	g.AddEdge(0, 1, 2)
	g.AddEdge(0, 3, 6)
	g.AddEdge(1, 2, 3)
	g.AddEdge(1, 3, 8)
	g.AddEdge(1, 4, 5)
	g.AddEdge(2, 4, 7)
	g.AddEdge(3, 4, 9)

	prim := NewPrimMST(g)

	if !prim.FindMST() {
		t.Error("Expected to find MST in connected graph")
	}

	expectedCost := 16.0 // 2 + 3 + 6 + 5
	actualCost := prim.GetMSTCost()

	if actualCost != expectedCost {
		t.Errorf("Expected MST cost %f, got %f", expectedCost, actualCost)
	}

	// MST'deki kenar sayısını kontrol et
	edges := prim.GetMSTEdges()
	expectedEdges := 4 // n-1 kenar olmalı
	if len(edges) != expectedEdges {
		t.Errorf("Expected %d edges in MST, got %d", expectedEdges, len(edges))
	}

	// Test 2: Yönlü graf (çalışmamalı)
	g2 := NewGraph(3, true)
	prim2 := NewPrimMST(g2)

	if prim2 != nil {
		t.Error("Expected nil PrimMST for directed graph")
	}

	// Test 3: Bağlantısız graf
	g3 := NewGraph(4, false)
	g3.AddEdge(0, 1, 1)
	g3.AddEdge(2, 3, 1)

	prim3 := NewPrimMST(g3)

	if prim3.FindMST() {
		t.Error("Expected to fail finding MST in disconnected graph")
	}

	// Test 4: Tek düğümlü graf
	g4 := NewGraph(1, false)
	prim4 := NewPrimMST(g4)

	if !prim4.FindMST() {
		t.Error("Expected to find MST in single-vertex graph")
	}

	if prim4.GetMSTCost() != 0 {
		t.Error("Expected zero cost MST for single-vertex graph")
	}

	// Test 5: Çoklu kenarlar ve döngüler
	g5 := NewGraph(3, false)
	g5.AddEdge(0, 1, 2)
	g5.AddEdge(0, 1, 3) // Aynı düğümler arasında farklı ağırlık
	g5.AddEdge(1, 2, 4)
	g5.AddEdge(0, 2, 7)

	prim5 := NewPrimMST(g5)
	prim5.FindMST()

	expectedCost5 := 6.0 // 2 + 4
	actualCost5 := prim5.GetMSTCost()

	if actualCost5 != expectedCost5 {
		t.Errorf("Expected MST cost %f, got %f", expectedCost5, actualCost5)
	}

	// Parent ilişkilerini kontrol et
	if !prim5.IsInMST(1) {
		t.Error("Vertex 1 should be in MST")
	}

	parent1 := prim5.GetParent(1)
	if parent1 != 0 {
		t.Errorf("Expected parent of vertex 1 to be 0, got %d", parent1)
	}
}
