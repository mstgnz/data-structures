package Graph

import (
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	// Test 1: Basit yönlü asiklik graf (DAG)
	g := NewGraph(6, true)
	g.AddEdge(5, 2, 1)
	g.AddEdge(5, 0, 1)
	g.AddEdge(4, 0, 1)
	g.AddEdge(4, 1, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 1, 1)

	ts := NewTopologicalSort(g)
	order := ts.Sort()

	// Beklenen sıralama: [5 4 2 3 1 0] veya [4 5 2 3 1 0]
	if len(order) != 6 {
		t.Errorf("Expected order length 6, got %d", len(order))
	}

	// Bağımlılıkları kontrol et
	dependencies := map[int][]int{
		5: {2, 0},
		4: {0, 1},
		2: {3},
		3: {1},
	}

	for vertex, deps := range dependencies {
		vertexIndex := -1
		for i, v := range order {
			if v == vertex {
				vertexIndex = i
				break
			}
		}

		for _, dep := range deps {
			depIndex := -1
			for i, v := range order {
				if v == dep {
					depIndex = i
					break
				}
			}
			if depIndex < vertexIndex {
				t.Errorf("Dependency violation: %d should come after %d", dep, vertex)
			}
		}
	}

	// Test 2: Çevrim içeren graf
	g2 := NewGraph(3, true)
	g2.AddEdge(0, 1, 1)
	g2.AddEdge(1, 2, 1)
	g2.AddEdge(2, 0, 1)

	ts2 := NewTopologicalSort(g2)
	order2 := ts2.Sort()

	if order2 != nil {
		t.Error("Expected nil result for cyclic graph")
	}

	if !ts2.HasCycle() {
		t.Error("Expected HasCycle() to return true for cyclic graph")
	}

	// Test 3: Yönsüz graf
	g3 := NewGraph(3, false)
	ts3 := NewTopologicalSort(g3)

	if ts3 != nil {
		t.Error("Expected nil TopologicalSort for undirected graph")
	}
}
