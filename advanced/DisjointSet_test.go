package advanced

import "testing"

// TestDisjointSetBasicOperations tests basic operations of disjoint set
func TestDisjointSetBasicOperations(t *testing.T) {
	ds := NewDisjointSet(5)

	// Test initial state
	if ds.GetSize() != 5 {
		t.Error("Wrong initial size")
	}
	if ds.GetSetCount() != 5 {
		t.Error("Wrong initial set count")
	}

	// Test union operation
	if !ds.Union(0, 1) {
		t.Error("Union of 0 and 1 should succeed")
	}
	if !ds.Union(2, 3) {
		t.Error("Union of 2 and 3 should succeed")
	}
	if !ds.Union(0, 2) {
		t.Error("Union of 0 and 2 should succeed")
	}

	// Test connected operation
	if !ds.Connected(1, 3) {
		t.Error("Elements 1 and 3 should be connected")
	}
	if ds.Connected(1, 4) {
		t.Error("Elements 1 and 4 should not be connected")
	}

	// Test set count after unions
	if ds.GetSetCount() != 2 {
		t.Error("Should have 2 sets after unions")
	}
}

// TestDisjointSetPathCompression tests path compression optimization
func TestDisjointSetPathCompression(t *testing.T) {
	ds := NewDisjointSet(5)

	// Create a chain: 0 -> 1 -> 2 -> 3 -> 4
	ds.parent[1] = 0
	ds.parent[2] = 1
	ds.parent[3] = 2
	ds.parent[4] = 3

	// Find operation should compress the path
	root := ds.Find(4)
	if root != 0 {
		t.Error("Root should be 0")
	}

	// After path compression, all elements should point directly to root
	for i := 1; i < 5; i++ {
		if ds.parent[i] != 0 {
			t.Errorf("Element %d should point directly to root after path compression", i)
		}
	}
}

// TestDisjointSetUnionByRank tests union by rank optimization
func TestDisjointSetUnionByRank(t *testing.T) {
	ds := NewDisjointSet(6)

	// Create two trees
	ds.Union(0, 1) // Tree 1: 0-1
	ds.Union(0, 2) // Tree 1: 0-1-2

	ds.Union(3, 4) // Tree 2: 3-4
	ds.Union(3, 5) // Tree 2: 3-4-5

	// Union the trees
	ds.Union(0, 3)

	// Check if the union was done by rank
	root := ds.Find(0)
	for i := 1; i < 6; i++ {
		if ds.Find(i) != root {
			t.Errorf("Element %d should be in the same set with root %d", i, root)
		}
	}
}

// TestDisjointSetInvalidOperations tests handling of invalid operations
func TestDisjointSetInvalidOperations(t *testing.T) {
	ds := NewDisjointSet(3)

	// Test invalid Find operations
	if ds.Find(-1) != -1 {
		t.Error("Find with negative index should return -1")
	}
	if ds.Find(3) != -1 {
		t.Error("Find with out of bounds index should return -1")
	}

	// Test invalid Union operations
	if ds.Union(-1, 1) {
		t.Error("Union with negative index should return false")
	}
	if ds.Union(1, 3) {
		t.Error("Union with out of bounds index should return false")
	}

	// Test invalid Connected operations
	if ds.Connected(-1, 1) {
		t.Error("Connected with negative index should return false")
	}
	if ds.Connected(1, 3) {
		t.Error("Connected with out of bounds index should return false")
	}
}
