package tree

import (
	"sync"
	"testing"
)

func TestThreadSafeBST(t *testing.T) {
	bst := NewThreadSafeBST[int]()

	// Test initial state
	if !bst.IsEmpty() {
		t.Error("New BST should be empty")
	}

	if size := bst.Size(); size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}

	// Test Insert and Contains
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)

	if !bst.Contains(5) {
		t.Error("BST should contain 5")
	}
	if !bst.Contains(3) {
		t.Error("BST should contain 3")
	}
	if !bst.Contains(7) {
		t.Error("BST should contain 7")
	}
	if bst.Contains(4) {
		t.Error("BST should not contain 4")
	}

	// Test FindMin and FindMax
	if min, ok := bst.FindMin(); !ok || min != 3 {
		t.Errorf("Expected min value 3, got %v", min)
	}
	if max, ok := bst.FindMax(); !ok || max != 7 {
		t.Errorf("Expected max value 7, got %v", max)
	}

	// Test Remove
	if !bst.Remove(3) {
		t.Error("Remove should return true for existing value")
	}
	if bst.Contains(3) {
		t.Error("BST should not contain 3 after removal")
	}
	if bst.Remove(4) {
		t.Error("Remove should return false for non-existing value")
	}

	// Test Clear
	bst.Clear()
	if !bst.IsEmpty() {
		t.Error("BST should be empty after clear")
	}
}

func TestThreadSafeBSTConcurrent(t *testing.T) {
	bst := NewThreadSafeBST[int]()
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
				bst.Insert(value)
			}
		}(i)
	}
	wg.Wait()

	expectedSize := numGoroutines * numOperations
	if size := bst.Size(); size != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, size)
	}

	// Test concurrent operations (Contains, Remove)
	wg.Add(numGoroutines * 2)
	for i := 0; i < numGoroutines; i++ {
		go func(n int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				value := n*numOperations + j
				bst.Contains(value)
			}
		}(i)
		go func(n int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				value := n*numOperations + j
				bst.Remove(value)
			}
		}(i)
	}
	wg.Wait()

	if !bst.IsEmpty() {
		t.Error("BST should be empty after all removes")
	}
}

func TestThreadSafeBSTEdgeCases(t *testing.T) {
	bst := NewThreadSafeBST[int]()

	// Test operations on empty tree
	if min, ok := bst.FindMin(); ok {
		t.Errorf("FindMin should return false on empty tree, got %v", min)
	}
	if max, ok := bst.FindMax(); ok {
		t.Errorf("FindMax should return false on empty tree, got %v", max)
	}

	// Test removing root node
	bst.Insert(1)
	if !bst.Remove(1) {
		t.Error("Failed to remove root node")
	}
	if !bst.IsEmpty() {
		t.Error("BST should be empty after removing root")
	}

	// Test removing node with two children
	bst.Insert(2)
	bst.Insert(1)
	bst.Insert(3)
	if !bst.Remove(2) {
		t.Error("Failed to remove node with two children")
	}
	if size := bst.Size(); size != 2 {
		t.Errorf("Expected size 2 after removal, got %d", size)
	}
}
