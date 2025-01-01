package Graph

import (
	"reflect"
	"sort"
	"testing"
)

func TestStronglyConnectedComponents(t *testing.T) {
	// Test 1: Simple directed graph
	g := NewGraph(5, true)
	g.AddEdge(1, 0, 1)
	g.AddEdge(0, 2, 1)
	g.AddEdge(2, 1, 1)
	g.AddEdge(0, 3, 1)
	g.AddEdge(3, 4, 1)

	scc := NewSCC(g)
	components := scc.FindComponents()

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

	scc2 := NewSCC(g2)
	components2 := scc2.FindComponents()

	if len(components2) != 1 {
		t.Errorf("Expected 1 component, got %d", len(components2))
	}

	expectedSize := 3
	if len(components2[0]) != expectedSize {
		t.Errorf("Expected component size %d, got %d", expectedSize, len(components2[0]))
	}

	// Test 3: Undirected graph
	g3 := NewGraph(3, false)
	scc3 := NewSCC(g3)

	if scc3 != nil {
		t.Error("Expected nil SCC for undirected graph")
	}

	// Test 4: Disconnected graph
	g4 := NewGraph(6, true)
	g4.AddEdge(0, 1, 1)
	g4.AddEdge(1, 0, 1)
	g4.AddEdge(2, 3, 1)
	g4.AddEdge(3, 2, 1)
	g4.AddEdge(4, 5, 1)
	g4.AddEdge(5, 4, 1)

	scc4 := NewSCC(g4)
	components4 := scc4.FindComponents()

	if len(components4) != 3 {
		t.Errorf("Expected 3 components, got %d", len(components4))
	}

	// Each component should have size 2
	for i, comp := range components4 {
		if len(comp) != 2 {
			t.Errorf("Component %d: expected size 2, got %d", i, len(comp))
		}
	}
}
