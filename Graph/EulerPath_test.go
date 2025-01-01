package Graph

import (
	"testing"
)

func TestEulerPath(t *testing.T) {
	// Test 1: Euler circuit in undirected graph
	g := NewGraph(4, false)
	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 0, 1)

	ep := NewEulerPath(g)

	if !ep.HasEulerCircuit() {
		t.Error("Expected to have Euler circuit")
	}

	circuit := ep.FindEulerCircuit()
	if circuit == nil {
		t.Error("Expected to find Euler circuit")
	}

	if len(circuit) != 5 { // n+1 nodes (starting node repeated)
		t.Errorf("Expected circuit length 5, got %d", len(circuit))
	}

	if circuit[0] != circuit[len(circuit)-1] {
		t.Error("Circuit should start and end at the same vertex")
	}

	// Test 2: Euler path but no circuit in undirected graph
	g2 := NewGraph(5, false)
	g2.AddEdge(0, 1, 1)
	g2.AddEdge(1, 2, 1)
	g2.AddEdge(2, 3, 1)
	g2.AddEdge(3, 4, 1)

	ep2 := NewEulerPath(g2)

	if !ep2.HasEulerPath() {
		t.Error("Expected to have Euler path")
	}

	if ep2.HasEulerCircuit() {
		t.Error("Expected not to have Euler circuit")
	}

	path := ep2.FindEulerPath()
	if path == nil {
		t.Error("Expected to find Euler path")
	}

	if len(path) != 5 { // n nodes
		t.Errorf("Expected path length 5, got %d", len(path))
	}

	// Test 3: No Euler path graph
	g3 := NewGraph(5, false)
	g3.AddEdge(0, 1, 1)
	g3.AddEdge(0, 2, 1)
	g3.AddEdge(0, 3, 1)
	g3.AddEdge(1, 2, 1)

	ep3 := NewEulerPath(g3)

	if ep3.HasEulerPath() {
		t.Error("Expected not to have Euler path")
	}

	path3 := ep3.FindEulerPath()
	if path3 != nil {
		t.Error("Expected nil path for graph with no Euler path")
	}

	// Test 4: Euler circuit in directed graph
	g4 := NewGraph(3, true)
	g4.AddEdge(0, 1, 1)
	g4.AddEdge(1, 2, 1)
	g4.AddEdge(2, 0, 1)

	ep4 := NewEulerPath(g4)

	if !ep4.HasEulerCircuit() {
		t.Error("Expected to have Euler circuit in directed graph")
	}

	circuit4 := ep4.FindEulerCircuit()
	if circuit4 == nil {
		t.Error("Expected to find Euler circuit in directed graph")
	}

	if len(circuit4) != 4 { // n+1 nodes
		t.Errorf("Expected circuit length 4, got %d", len(circuit4))
	}

	// Test 5: Disconnected graph
	g5 := NewGraph(4, false)
	g5.AddEdge(0, 1, 1)
	g5.AddEdge(2, 3, 1)

	ep5 := NewEulerPath(g5)

	if ep5.HasEulerPath() {
		t.Error("Expected no Euler path in disconnected graph")
	}

	// Test 6: Single vertex graph
	g6 := NewGraph(1, false)
	ep6 := NewEulerPath(g6)

	if !ep6.HasEulerCircuit() {
		t.Error("Expected to have Euler circuit in single vertex graph")
	}

	circuit6 := ep6.FindEulerCircuit()
	if len(circuit6) != 1 {
		t.Errorf("Expected circuit length 1 for single vertex, got %d", len(circuit6))
	}

	// Test 7: Euler path in directed graph
	g7 := NewGraph(4, true)
	g7.AddEdge(0, 1, 1)
	g7.AddEdge(1, 2, 1)
	g7.AddEdge(2, 3, 1)

	ep7 := NewEulerPath(g7)

	if !ep7.HasEulerPath() {
		t.Error("Expected to have Euler path in directed graph")
	}

	path7 := ep7.FindEulerPath()
	if path7 == nil {
		t.Error("Expected to find Euler path in directed graph")
	}

	if len(path7) != 4 {
		t.Errorf("Expected path length 4, got %d", len(path7))
	}
}
