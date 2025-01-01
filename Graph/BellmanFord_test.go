package Graph

import (
	"math"
	"reflect"
	"testing"
)

func TestBellmanFord(t *testing.T) {
	// Test 1: Basit yönlü ağırlıklı graf
	g := NewGraph(5, true)
	g.AddEdge(0, 1, 6)
	g.AddEdge(0, 3, 7)
	g.AddEdge(1, 2, 5)
	g.AddEdge(1, 3, 8)
	g.AddEdge(1, 4, -4)
	g.AddEdge(2, 1, -2)
	g.AddEdge(3, 2, -3)
	g.AddEdge(3, 4, 9)
	g.AddEdge(4, 0, 2)
	g.AddEdge(4, 2, 7)

	source := 0
	bf := NewBellmanFord(g, source)

	if !bf.ComputeShortestPaths() {
		t.Error("Expected no negative cycle")
	}

	// Test mesafeleri
	expectedDist := []float64{0, 2, 4, 7, -2}
	actualDist := bf.GetAllDistances()

	for i := 0; i < len(expectedDist); i++ {
		if actualDist[i] != expectedDist[i] {
			t.Errorf("Distance to vertex %d: expected %f, got %f",
				i, expectedDist[i], actualDist[i])
		}
	}

	// Test yolları
	testPaths := []struct {
		to       int
		expected []int
	}{
		{1, []int{0, 1}},
		{2, []int{0, 1, 2}},
		{3, []int{0, 3}},
		{4, []int{0, 1, 4}},
	}

	for _, tp := range testPaths {
		path := bf.GetPath(tp.to)
		if !reflect.DeepEqual(path, tp.expected) {
			t.Errorf("Path to vertex %d: expected %v, got %v",
				tp.to, tp.expected, path)
		}
	}

	// Test 2: Negatif çevrimli graf
	g2 := NewGraph(4, true)
	g2.AddEdge(0, 1, 1)
	g2.AddEdge(1, 2, -1)
	g2.AddEdge(2, 3, -1)
	g2.AddEdge(3, 1, -1)

	bf2 := NewBellmanFord(g2, 0)

	if bf2.ComputeShortestPaths() {
		t.Error("Expected to detect negative cycle")
	}

	// Test 3: Bağlantısız graf
	g3 := NewGraph(4, true)
	g3.AddEdge(0, 1, 1)
	g3.AddEdge(2, 3, 1)

	bf3 := NewBellmanFord(g3, 0)
	bf3.ComputeShortestPaths()

	// 0'dan 3'e yol olmamalı
	if bf3.IsReachable(3) {
		t.Error("Vertex 3 should not be reachable from source 0")
	}

	dist03 := bf3.GetDistance(3)
	if dist03 != math.Inf(1) {
		t.Errorf("Expected infinite distance to vertex 3, got %f", dist03)
	}

	path03 := bf3.GetPath(3)
	if path03 != nil {
		t.Errorf("Expected nil path to vertex 3, got %v", path03)
	}

	// Test 4: Sıfır ağırlıklı kenarlar
	g4 := NewGraph(3, true)
	g4.AddEdge(0, 1, 0)
	g4.AddEdge(1, 2, 0)

	bf4 := NewBellmanFord(g4, 0)
	bf4.ComputeShortestPaths()

	expectedDist4 := []float64{0, 0, 0}
	actualDist4 := bf4.GetAllDistances()

	for i := 0; i < len(expectedDist4); i++ {
		if actualDist4[i] != expectedDist4[i] {
			t.Errorf("Distance to vertex %d: expected %f, got %f",
				i, expectedDist4[i], actualDist4[i])
		}
	}
}
