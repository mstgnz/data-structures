package graph

import (
	"reflect"
	"sort"
	"testing"
)

func TestTarjanSCC(t *testing.T) {
	// Test 1: Simple directed graph
	g := NewGraph(5, true)
	g.AddEdge(1, 0, 1)
	g.AddEdge(0, 2, 1)
	g.AddEdge(2, 1, 1)
	g.AddEdge(0, 3, 1)
	g.AddEdge(3, 4, 1)

	tarjan := NewTarjanSCC(g)
	components := tarjan.FindComponents()

	// Expected components: [[0,1,2], [3], [4]]
	expectedComponents := [][]int{
		{0, 1, 2},
		{3},
		{4},
	}

	if len(components) != len(expectedComponents) {
		t.Errorf("Expected %d components, got %d", len(expectedComponents), len(components))
	}

	// Sort each component and compare
	for i := range components {
		sort.Ints(components[i])
	}
	for i := range expectedComponents {
		sort.Ints(expectedComponents[i])
	}

	// Sort components by size
	sort.Slice(components, func(i, j int) bool {
		return len(components[i]) > len(components[j])
	})
	sort.Slice(expectedComponents, func(i, j int) bool {
		return len(expectedComponents[i]) > len(expectedComponents[j])
	})

	if !reflect.DeepEqual(components, expectedComponents) {
		t.Errorf("Expected components %v, got %v", expectedComponents, components)
	}

	// Test 2: Single component cyclic graph
	g2 := NewGraph(3, true)
	g2.AddEdge(0, 1, 1)
	g2.AddEdge(1, 2, 1)
	g2.AddEdge(2, 0, 1)

	tarjan2 := NewTarjanSCC(g2)
	components2 := tarjan2.FindComponents()

	if len(components2) != 1 {
		t.Errorf("Expected 1 component, got %d", len(components2))
	}

	if !tarjan2.IsStronglyConnected() {
		t.Error("Graph should be strongly connected")
	}

	// Test 3: Undirected graph
	g3 := NewGraph(3, false)
	tarjan3 := NewTarjanSCC(g3)

	if tarjan3 != nil {
		t.Error("Expected nil TarjanSCC for undirected graph")
	}

	// Test 4: Disconnected graph
	g4 := NewGraph(6, true)
	g4.AddEdge(0, 1, 1)
	g4.AddEdge(1, 0, 1)
	g4.AddEdge(2, 3, 1)
	g4.AddEdge(3, 2, 1)
	g4.AddEdge(4, 5, 1)
	g4.AddEdge(5, 4, 1)

	tarjan4 := NewTarjanSCC(g4)
	components4 := tarjan4.FindComponents()

	if len(components4) != 3 {
		t.Errorf("Expected 3 components, got %d", len(components4))
	}

	// Each component should have size 2
	for i, comp := range components4 {
		if len(comp) != 2 {
			t.Errorf("Component %d: expected size 2, got %d", i, len(comp))
		}
	}

	// Test 5: Check largest component
	g5 := NewGraph(7, true)
	g5.AddEdge(0, 1, 1)
	g5.AddEdge(1, 2, 1)
	g5.AddEdge(2, 0, 1)
	g5.AddEdge(3, 4, 1)
	g5.AddEdge(4, 3, 1)
	g5.AddEdge(5, 6, 1)

	tarjan5 := NewTarjanSCC(g5)
	largest := tarjan5.GetLargestComponent()
	sort.Ints(largest)

	expectedLargest := []int{0, 1, 2}
	if !reflect.DeepEqual(largest, expectedLargest) {
		t.Errorf("Expected largest component %v, got %v", expectedLargest, largest)
	}
}
