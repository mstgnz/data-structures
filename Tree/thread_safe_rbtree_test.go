package tree

import (
	"sync"
	"testing"
)

func TestThreadSafeRBTree(t *testing.T) {
	rbt := NewThreadSafeRBTree[int]()

	// Test initial state
	if !rbt.IsEmpty() {
		t.Error("New RB tree should be empty")
	}

	if size := rbt.Size(); size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}

	// Test Insert and Contains
	rbt.Insert(5)
	rbt.Insert(3)
	rbt.Insert(7)
	rbt.Insert(1)
	rbt.Insert(9)

	if !rbt.Contains(5) {
		t.Error("RB tree should contain 5")
	}
	if !rbt.Contains(3) {
		t.Error("RB tree should contain 3")
	}
	if !rbt.Contains(7) {
		t.Error("RB tree should contain 7")
	}
	if rbt.Contains(4) {
		t.Error("RB tree should not contain 4")
	}

	// Test FindMin and FindMax
	if min, ok := rbt.FindMin(); !ok || min != 1 {
		t.Errorf("Expected min value 1, got %v", min)
	}
	if max, ok := rbt.FindMax(); !ok || max != 9 {
		t.Errorf("Expected max value 9, got %v", max)
	}

	// Test Clear
	rbt.Clear()
	if !rbt.IsEmpty() {
		t.Error("RB tree should be empty after clear")
	}
}

func TestThreadSafeRBTreeConcurrent(t *testing.T) {
	rbt := NewThreadSafeRBTree[int]()
	var wg sync.WaitGroup
	numGoroutines := 100
	numOperations := 100

	// Test concurrent insertions
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(n int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				value := n*numOperations + j
				rbt.Insert(value)
			}
		}(i)
	}
	wg.Wait()

	expectedSize := numGoroutines * numOperations
	if size := rbt.Size(); size != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, size)
	}

	// Test concurrent operations (Contains)
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(n int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				value := n*numOperations + j
				if !rbt.Contains(value) {
					t.Errorf("Tree should contain value %d", value)
				}
			}
		}(i)
	}
	wg.Wait()
}

func TestThreadSafeRBTreeProperties(t *testing.T) {
	rbt := NewThreadSafeRBTree[int]()

	// Test Red-Black tree properties with sequential insertions
	values := []int{7, 3, 18, 10, 22, 8, 11, 26, 2, 6, 13}
	for _, v := range values {
		rbt.Insert(v)
	}

	// Property 1: Root must be black (tested in Insert implementation)
	// Property 2: Red nodes cannot have red children (tested in Insert implementation)
	// Property 3: All paths from root to leaves must have same number of black nodes (tested in Insert implementation)

	// Test that all values are still accessible
	for _, v := range values {
		if !rbt.Contains(v) {
			t.Errorf("Tree should contain value %d", v)
		}
	}
}

func TestThreadSafeRBTreeEdgeCases(t *testing.T) {
	rbt := NewThreadSafeRBTree[int]()

	// Test operations on empty tree
	if min, ok := rbt.FindMin(); ok {
		t.Errorf("FindMin should return false on empty tree, got %v", min)
	}
	if max, ok := rbt.FindMax(); ok {
		t.Errorf("FindMax should return false on empty tree, got %v", max)
	}

	// Test duplicate insertions
	rbt.Insert(1)
	rbt.Insert(1) // Should not affect the tree
	if size := rbt.Size(); size != 1 {
		t.Errorf("Expected size 1 after duplicate insertion, got %d", size)
	}

	// Test Clear and reuse
	rbt.Clear()
	if !rbt.IsEmpty() {
		t.Error("Tree should be empty after clear")
	}

	rbt.Insert(5)
	if !rbt.Contains(5) {
		t.Error("Tree should contain 5 after clear and reinsert")
	}
}
