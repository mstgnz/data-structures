package tree

import (
	"sync"
	"testing"
)

func TestThreadSafeAVL(t *testing.T) {
	avl := NewThreadSafeAVL[int]()

	// Test initial state
	if !avl.IsEmpty() {
		t.Error("New AVL tree should be empty")
	}

	if size := avl.Size(); size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}

	// Test Insert and Contains
	avl.Insert(5)
	avl.Insert(3)
	avl.Insert(7)
	avl.Insert(1)
	avl.Insert(9)

	if !avl.Contains(5) {
		t.Error("AVL tree should contain 5")
	}
	if !avl.Contains(3) {
		t.Error("AVL tree should contain 3")
	}
	if !avl.Contains(7) {
		t.Error("AVL tree should contain 7")
	}
	if avl.Contains(4) {
		t.Error("AVL tree should not contain 4")
	}

	// Test FindMin and FindMax
	if min, ok := avl.FindMin(); !ok || min != 1 {
		t.Errorf("Expected min value 1, got %v", min)
	}
	if max, ok := avl.FindMax(); !ok || max != 9 {
		t.Errorf("Expected max value 9, got %v", max)
	}

	// Test Remove
	if !avl.Remove(3) {
		t.Error("Remove should return true for existing value")
	}
	if avl.Contains(3) {
		t.Error("AVL tree should not contain 3 after removal")
	}
	if avl.Remove(4) {
		t.Error("Remove should return false for non-existing value")
	}

	// Test height after rotations
	if height := avl.GetHeight(); height < 2 || height > 3 {
		t.Errorf("Expected height between 2 and 3, got %d", height)
	}

	// Test Clear
	avl.Clear()
	if !avl.IsEmpty() {
		t.Error("AVL tree should be empty after clear")
	}
}

func TestThreadSafeAVLConcurrent(t *testing.T) {
	avl := NewThreadSafeAVL[int]()
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
				avl.Insert(value)
			}
		}(i)
	}
	wg.Wait()

	expectedSize := numGoroutines * numOperations
	if size := avl.Size(); size != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, size)
	}

	// Test concurrent operations (Contains, Remove)
	wg.Add(numGoroutines * 2)
	for i := 0; i < numGoroutines; i++ {
		go func(n int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				value := n*numOperations + j
				avl.Contains(value)
			}
		}(i)
		go func(n int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				value := n*numOperations + j
				avl.Remove(value)
			}
		}(i)
	}
	wg.Wait()

	if !avl.IsEmpty() {
		t.Error("AVL tree should be empty after all removes")
	}
}

func TestThreadSafeAVLRotations(t *testing.T) {
	avl := NewThreadSafeAVL[int]()

	// Test Left-Left case
	avl.Insert(30)
	avl.Insert(20)
	avl.Insert(10)
	if height := avl.GetHeight(); height != 1 {
		t.Errorf("Expected height 1 after LL rotation, got %d", height)
	}

	avl.Clear()

	// Test Right-Right case
	avl.Insert(10)
	avl.Insert(20)
	avl.Insert(30)
	if height := avl.GetHeight(); height != 1 {
		t.Errorf("Expected height 1 after RR rotation, got %d", height)
	}

	avl.Clear()

	// Test Left-Right case
	avl.Insert(30)
	avl.Insert(10)
	avl.Insert(20)
	if height := avl.GetHeight(); height != 1 {
		t.Errorf("Expected height 1 after LR rotation, got %d", height)
	}

	avl.Clear()

	// Test Right-Left case
	avl.Insert(10)
	avl.Insert(30)
	avl.Insert(20)
	if height := avl.GetHeight(); height != 1 {
		t.Errorf("Expected height 1 after RL rotation, got %d", height)
	}
}

func TestThreadSafeAVLEdgeCases(t *testing.T) {
	avl := NewThreadSafeAVL[int]()

	// Test operations on empty tree
	if min, ok := avl.FindMin(); ok {
		t.Errorf("FindMin should return false on empty tree, got %v", min)
	}
	if max, ok := avl.FindMax(); ok {
		t.Errorf("FindMax should return false on empty tree, got %v", max)
	}

	// Test removing root node
	avl.Insert(1)
	if !avl.Remove(1) {
		t.Error("Failed to remove root node")
	}
	if !avl.IsEmpty() {
		t.Error("AVL tree should be empty after removing root")
	}

	// Test removing node with two children
	avl.Insert(2)
	avl.Insert(1)
	avl.Insert(3)
	if !avl.Remove(2) {
		t.Error("Failed to remove node with two children")
	}
	if size := avl.Size(); size != 2 {
		t.Errorf("Expected size 2 after removal, got %d", size)
	}
}
