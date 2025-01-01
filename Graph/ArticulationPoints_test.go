package Graph

import (
	"reflect"
	"sort"
	"testing"
)

func TestArticulationPoints(t *testing.T) {
	// Test 1: Simple undirected graph
	g := NewGraph(5, false)
	g.AddEdge(1, 0, 1)
	g.AddEdge(0, 2, 1)
	g.AddEdge(2, 1, 1)
	g.AddEdge(0, 3, 1)
	g.AddEdge(3, 4, 1)

	ap := NewArticulationPoints(g)
	points := ap.FindArticulationPoints()
	sort.Ints(points)

	// Expected articulation points: [0, 3]
	expectedPoints := []int{0, 3}
	if !reflect.DeepEqual(points, expectedPoints) {
		t.Errorf("Expected articulation points %v, got %v", expectedPoints, points)
	}

	// Check bridges
	bridges := ap.FindBridges()
	if len(bridges) != 2 {
		t.Errorf("Expected 2 bridges, got %d", len(bridges))
	}

	// Test 2: Directed graph (should not work)
	g2 := NewGraph(3, true)
	ap2 := NewArticulationPoints(g2)

	if ap2 != nil {
		t.Error("Expected nil ArticulationPoints for directed graph")
	}

	// Test 3: Cyclic graph
	g3 := NewGraph(3, false)
	g3.AddEdge(0, 1, 1)
	g3.AddEdge(1, 2, 1)
	g3.AddEdge(2, 0, 1)

	ap3 := NewArticulationPoints(g3)
	points3 := ap3.FindArticulationPoints()

	if len(points3) != 0 {
		t.Errorf("Expected no articulation points in cycle, got %d", len(points3))
	}

	bridges3 := ap3.FindBridges()
	if len(bridges3) != 0 {
		t.Errorf("Expected no bridges in cycle, got %d", len(bridges3))
	}

	// Test 4: Single edge
	g4 := NewGraph(2, false)
	g4.AddEdge(0, 1, 1)

	ap4 := NewArticulationPoints(g4)

	if !ap4.IsArticulationPoint(0) || !ap4.IsArticulationPoint(1) {
		t.Error("Both vertices should be articulation points in single edge graph")
	}

	if !ap4.IsBridge(0, 1) {
		t.Error("Edge should be a bridge in single edge graph")
	}

	// Test 5: Complex graph
	g5 := NewGraph(7, false)
	g5.AddEdge(0, 1, 1)
	g5.AddEdge(1, 2, 1)
	g5.AddEdge(2, 0, 1)
	g5.AddEdge(1, 3, 1)
	g5.AddEdge(1, 4, 1)
	g5.AddEdge(1, 6, 1)
	g5.AddEdge(3, 5, 1)
	g5.AddEdge(4, 5, 1)

	ap5 := NewArticulationPoints(g5)
	points5 := ap5.FindArticulationPoints()
	sort.Ints(points5)

	expectedPoints5 := []int{1, 5}
	if !reflect.DeepEqual(points5, expectedPoints5) {
		t.Errorf("Expected articulation points %v, got %v", expectedPoints5, points5)
	}

	expectedBridgeCount := 1
	if ap5.GetBridgeCount() != expectedBridgeCount {
		t.Errorf("Expected %d bridges, got %d", expectedBridgeCount, ap5.GetBridgeCount())
	}

	// Test 6: Disconnected graph
	g6 := NewGraph(4, false)
	g6.AddEdge(0, 1, 1)
	g6.AddEdge(2, 3, 1)

	ap6 := NewArticulationPoints(g6)
	points6 := ap6.FindArticulationPoints()
	sort.Ints(points6)

	expectedPoints6 := []int{0, 1, 2, 3}
	if !reflect.DeepEqual(points6, expectedPoints6) {
		t.Errorf("Expected articulation points %v, got %v", expectedPoints6, points6)
	}

	bridges6 := ap6.FindBridges()
	if len(bridges6) != 2 {
		t.Errorf("Expected 2 bridges in disconnected graph, got %d", len(bridges6))
	}
}
