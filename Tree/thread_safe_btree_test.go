package tree

import (
	"sync"
	"testing"
)

func TestThreadSafeBTree(t *testing.T) {
	btree := NewThreadSafeBTree[int](3) // degree = 3

	// Test initial state
	if !btree.IsEmpty() {
		t.Error("New B-tree should be empty")
	}

	if size := btree.Size(); size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}

	// Test Insert and Contains
	values := []int{5, 3, 7, 1, 9, 2, 4, 6, 8, 10}
	for _, v := range values {
		btree.Insert(v)
	}

	for _, v := range values {
		if !btree.Contains(v) {
			t.Errorf("B-tree should contain %d", v)
		}
	}

	if btree.Contains(11) {
		t.Error("B-tree should not contain 11")
	}

	// Test FindMin and FindMax
	if min, ok := btree.FindMin(); !ok || min != 1 {
		t.Errorf("Expected min value 1, got %v", min)
	}
	if max, ok := btree.FindMax(); !ok || max != 10 {
		t.Errorf("Expected max value 10, got %v", max)
	}

	// Test Clear
	btree.Clear()
	if !btree.IsEmpty() {
		t.Error("B-tree should be empty after clear")
	}
}

func TestThreadSafeBTreeConcurrent(t *testing.T) {
	btree := NewThreadSafeBTree[int](3)
	var wg sync.WaitGroup
	numGoroutines := 10
	numOperations := 100

	// Test concurrent insertions
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(n int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				value := n*numOperations + j
				btree.Insert(value)
			}
		}(i)
	}
	wg.Wait()

	expectedSize := numGoroutines * numOperations
	if size := btree.Size(); size != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, size)
	}

	// Test concurrent operations (Contains)
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(n int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				value := n*numOperations + j
				if !btree.Contains(value) {
					t.Errorf("B-tree should contain value %d", value)
				}
			}
		}(i)
	}
	wg.Wait()
}

func TestThreadSafeBTreeProperties(t *testing.T) {
	btree := NewThreadSafeBTree[int](3)

	// Test B-tree properties with sequential insertions
	values := []int{7, 3, 18, 10, 22, 8, 11, 26, 2, 6, 13}
	for _, v := range values {
		btree.Insert(v)
	}

	// Test that all values are still accessible
	for _, v := range values {
		if !btree.Contains(v) {
			t.Errorf("B-tree should contain value %d", v)
		}
	}
}

func TestThreadSafeBTreeEdgeCases(t *testing.T) {
	btree := NewThreadSafeBTree[int](2) // minimum degree

	// Test operations on empty tree
	if min, ok := btree.FindMin(); ok {
		t.Errorf("FindMin should return false on empty tree, got %v", min)
	}
	if max, ok := btree.FindMax(); ok {
		t.Errorf("FindMax should return false on empty tree, got %v", max)
	}

	// Test duplicate insertions
	btree.Insert(1)
	btree.Insert(1) // Should not affect the tree
	if size := btree.Size(); size != 2 {
		t.Errorf("Expected size 2 after duplicate insertion, got %d", size)
	}

	// Test Clear and reuse
	btree.Clear()
	if !btree.IsEmpty() {
		t.Error("B-tree should be empty after clear")
	}

	btree.Insert(5)
	if !btree.Contains(5) {
		t.Error("B-tree should contain 5 after clear and reinsert")
	}

	// Test with different degrees
	btree = NewThreadSafeBTree[int](4)
	for i := 1; i <= 100; i++ {
		btree.Insert(i)
	}
	for i := 1; i <= 100; i++ {
		if !btree.Contains(i) {
			t.Errorf("B-tree should contain value %d", i)
		}
	}
}
