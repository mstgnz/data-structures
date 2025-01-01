package graph

import (
	"testing"
)

func TestKruskalMST(t *testing.T) {
	// Test 1: Simple connected undirected graph
	g := NewGraph(5, false)
	g.AddEdge(0, 1, 2)
	g.AddEdge(0, 3, 6)
	g.AddEdge(1, 2, 3)
	g.AddEdge(1, 3, 8)
	g.AddEdge(1, 4, 5)
	g.AddEdge(2, 4, 7)
	g.AddEdge(3, 4, 9)

	kruskal := NewKruskalMST(g)

	if !kruskal.FindMST() {
		t.Error("Expected to find MST in connected graph")
	}

	expectedCost := 16.0 // 2 + 3 + 6 + 5
	actualCost := kruskal.GetMSTCost()

	if actualCost != expectedCost {
		t.Errorf("Expected MST cost %f, got %f", expectedCost, actualCost)
	}

	// Check number of edges in MST
	edges := kruskal.GetMSTEdges()
	expectedEdges := 4 // n-1 edges
	if len(edges) != expectedEdges {
		t.Errorf("Expected %d edges in MST, got %d", expectedEdges, len(edges))
	}

	// Test 2: Directed graph (should not work)
	g2 := NewGraph(3, true)
	kruskal2 := NewKruskalMST(g2)

	if kruskal2 != nil {
		t.Error("Expected nil KruskalMST for directed graph")
	}

	// Test 3: Disconnected graph
	g3 := NewGraph(4, false)
	g3.AddEdge(0, 1, 1)
	g3.AddEdge(2, 3, 1)

	kruskal3 := NewKruskalMST(g3)

	if kruskal3.FindMST() {
		t.Error("Expected to fail finding MST in disconnected graph")
	}

	expectedComponents := 2
	actualComponents := kruskal3.GetNumComponents()
	if actualComponents != expectedComponents {
		t.Errorf("Expected %d components, got %d", expectedComponents, actualComponents)
	}

	// Test 4: Single vertex graph
	g4 := NewGraph(1, false)
	kruskal4 := NewKruskalMST(g4)

	if !kruskal4.FindMST() {
		t.Error("Expected to find MST in single-vertex graph")
	}

	if kruskal4.GetMSTCost() != 0 {
		t.Error("Expected zero cost MST for single-vertex graph")
	}

	// Test 5: Multiple edges and cycles
	g5 := NewGraph(3, false)
	g5.AddEdge(0, 1, 2)
	g5.AddEdge(0, 1, 3) // Same nodes with different weights
	g5.AddEdge(1, 2, 4)
	g5.AddEdge(0, 2, 7)

	kruskal5 := NewKruskalMST(g5)
	kruskal5.FindMST()

	expectedCost5 := 6.0 // 2 + 4
	actualCost5 := kruskal5.GetMSTCost()

	if actualCost5 != expectedCost5 {
		t.Errorf("Expected MST cost %f, got %f", expectedCost5, actualCost5)
	}

	// Connection check
	if !kruskal5.IsConnected(0, 2) {
		t.Error("Vertices 0 and 2 should be connected in MST")
	}

	// Test 6: Equal weight edges
	g6 := NewGraph(4, false)
	g6.AddEdge(0, 1, 1)
	g6.AddEdge(1, 2, 1)
	g6.AddEdge(2, 3, 1)
	g6.AddEdge(3, 0, 1)

	kruskal6 := NewKruskalMST(g6)
	kruskal6.FindMST()

	expectedCost6 := 3.0 // Any three edges
	actualCost6 := kruskal6.GetMSTCost()

	if actualCost6 != expectedCost6 {
		t.Errorf("Expected MST cost %f, got %f", expectedCost6, actualCost6)
	}
}
