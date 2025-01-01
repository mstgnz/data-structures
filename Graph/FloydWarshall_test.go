package Graph

import (
	"math"
	"reflect"
	"testing"
)

func TestFloydWarshall(t *testing.T) {
	// Test 1: Simple weighted directed graph
	g := NewGraph(4, true)
	g.AddEdge(0, 1, 5)
	g.AddEdge(0, 3, 10)
	g.AddEdge(1, 2, 3)
	g.AddEdge(2, 3, 1)

	fw := NewFloydWarshall(g)
	fw.ComputeShortestPaths()

	// Test distances
	testCases := []struct {
		from     int
		to       int
		expected float64
	}{
		{0, 1, 5},
		{0, 2, 8},
		{0, 3, 9},
		{1, 3, 4},
		{2, 3, 1},
	}

	for _, tc := range testCases {
		dist := fw.GetDistance(tc.from, tc.to)
		if dist != tc.expected {
			t.Errorf("Distance from %d to %d: expected %f, got %f",
				tc.from, tc.to, tc.expected, dist)
		}
	}

	// Test paths
	path := fw.GetPath(0, 3)
	expectedPath := []int{0, 1, 2, 3}
	if !reflect.DeepEqual(path, expectedPath) {
		t.Errorf("Path from 0 to 3: expected %v, got %v", expectedPath, path)
	}

	// Test 2: Negative weighted graph (no cycle)
	g2 := NewGraph(4, true)
	g2.AddEdge(0, 1, 1)
	g2.AddEdge(1, 2, -3)
	g2.AddEdge(2, 3, 2)
	g2.AddEdge(0, 2, 4)
	g2.AddEdge(1, 3, 3)

	fw2 := NewFloydWarshall(g2)
	fw2.ComputeShortestPaths()

	if fw2.HasNegativeCycle() {
		t.Error("Expected no negative cycle")
	}

	// Test distances
	dist02 := fw2.GetDistance(0, 2)
	expectedDist02 := -2.0 // 0->1->2 path
	if dist02 != expectedDist02 {
		t.Errorf("Distance from 0 to 2: expected %f, got %f",
			expectedDist02, dist02)
	}

	// Test 3: Negative cycle graph
	g3 := NewGraph(3, true)
	g3.AddEdge(0, 1, 1)
	g3.AddEdge(1, 2, -1)
	g3.AddEdge(2, 0, -1)

	fw3 := NewFloydWarshall(g3)
	fw3.ComputeShortestPaths()

	if !fw3.HasNegativeCycle() {
		t.Error("Expected to detect negative cycle")
	}

	// Test 4: Disconnected graph
	g4 := NewGraph(4, true)
	g4.AddEdge(0, 1, 1)
	g4.AddEdge(2, 3, 1)

	fw4 := NewFloydWarshall(g4)
	fw4.ComputeShortestPaths()

	// There should be no path from 0 to 3
	dist03 := fw4.GetDistance(0, 3)
	if dist03 != math.Inf(1) {
		t.Errorf("Expected infinite distance from 0 to 3, got %f", dist03)
	}

	path03 := fw4.GetPath(0, 3)
	if path03 != nil {
		t.Errorf("Expected nil path from 0 to 3, got %v", path03)
	}
}
