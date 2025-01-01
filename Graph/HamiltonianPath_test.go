package Graph

import (
	"testing"
)

func TestHamiltonianPath(t *testing.T) {
	// Test 1: Hamilton çevrimi olan tam graf
	g := NewGraph(4, false)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 1)
	g.AddEdge(0, 3, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 1)

	hp := NewHamiltonianPath(g)

	if !hp.HasHamiltonianCircuit() {
		t.Error("Expected to have Hamiltonian circuit in complete graph")
	}

	circuit := hp.FindHamiltonianCircuit()
	if circuit == nil {
		t.Error("Expected to find Hamiltonian circuit")
	}

	if len(circuit) != 5 { // n+1 düğüm (başlangıç düğümü tekrar)
		t.Errorf("Expected circuit length 5, got %d", len(circuit))
	}

	if !hp.IsHamiltonianCircuit(circuit) {
		t.Error("Found circuit is not valid")
	}

	// Test 2: Hamilton yolu olan ama çevrimi olmayan graf
	g2 := NewGraph(4, false)
	g2.AddEdge(0, 1, 1)
	g2.AddEdge(1, 2, 1)
	g2.AddEdge(2, 3, 1)

	hp2 := NewHamiltonianPath(g2)

	if !hp2.HasHamiltonianPath() {
		t.Error("Expected to have Hamiltonian path")
	}

	if hp2.HasHamiltonianCircuit() {
		t.Error("Expected not to have Hamiltonian circuit")
	}

	path := hp2.FindHamiltonianPath()
	if path == nil {
		t.Error("Expected to find Hamiltonian path")
	}

	if len(path) != 4 { // n düğüm
		t.Errorf("Expected path length 4, got %d", len(path))
	}

	if !hp2.IsHamiltonianPath(path) {
		t.Error("Found path is not valid")
	}

	// Test 3: Hamilton yolu olmayan graf
	g3 := NewGraph(4, false)
	g3.AddEdge(0, 1, 1)
	g3.AddEdge(0, 2, 1)
	g3.AddEdge(1, 2, 1)

	hp3 := NewHamiltonianPath(g3)

	if hp3.HasHamiltonianPath() {
		t.Error("Expected not to have Hamiltonian path")
	}

	path3 := hp3.FindHamiltonianPath()
	if path3 != nil {
		t.Error("Expected nil path for graph with no Hamiltonian path")
	}

	// Test 4: Yönlü Hamilton çevrimi
	g4 := NewGraph(3, true)
	g4.AddEdge(0, 1, 1)
	g4.AddEdge(1, 2, 1)
	g4.AddEdge(2, 0, 1)

	hp4 := NewHamiltonianPath(g4)

	if !hp4.HasHamiltonianCircuit() {
		t.Error("Expected to have Hamiltonian circuit in directed graph")
	}

	circuit4 := hp4.FindHamiltonianCircuit()
	if circuit4 == nil {
		t.Error("Expected to find Hamiltonian circuit in directed graph")
	}

	if len(circuit4) != 4 { // n+1 düğüm
		t.Errorf("Expected circuit length 4, got %d", len(circuit4))
	}

	// Test 5: Bağlantısız graf
	g5 := NewGraph(4, false)
	g5.AddEdge(0, 1, 1)
	g5.AddEdge(2, 3, 1)

	hp5 := NewHamiltonianPath(g5)

	if hp5.HasHamiltonianPath() {
		t.Error("Expected no Hamiltonian path in disconnected graph")
	}

	// Test 6: Tek düğümlü graf
	g6 := NewGraph(1, false)
	hp6 := NewHamiltonianPath(g6)

	if !hp6.HasHamiltonianPath() {
		t.Error("Expected to have Hamiltonian path in single vertex graph")
	}

	path6 := hp6.FindHamiltonianPath()
	if len(path6) != 1 {
		t.Errorf("Expected path length 1 for single vertex, got %d", len(path6))
	}

	// Test 7: Geçersiz yol kontrolü
	g7 := NewGraph(3, false)
	g7.AddEdge(0, 1, 1)
	g7.AddEdge(1, 2, 1)

	hp7 := NewHamiltonianPath(g7)

	invalidPath := []int{0, 2, 1} // 0'dan 2'ye kenar yok
	if hp7.IsHamiltonianPath(invalidPath) {
		t.Error("Expected invalid path to be rejected")
	}

	invalidCircuit := []int{0, 1, 2, 0} // 2'den 0'a kenar yok
	if hp7.IsHamiltonianCircuit(invalidCircuit) {
		t.Error("Expected invalid circuit to be rejected")
	}
}
