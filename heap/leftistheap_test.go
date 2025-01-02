package heap

import "testing"

func TestLeftistHeapBasicOperations(t *testing.T) {
	lh := NewLeftistHeap()

	// Test initial state
	if !lh.IsEmpty() {
		t.Error("New heap should be empty")
	}

	// Test Insert and FindMin
	lh.Insert(5)
	if min, ok := lh.FindMin(); !ok || min != 5 {
		t.Errorf("Expected min to be 5, got %d", min)
	}

	// Test multiple inserts
	values := []int{3, 7, 1, 9, 4}
	for _, v := range values {
		lh.Insert(v)
	}

	// Test size
	expectedSize := len(values) + 1 // +1 for the first insert
	if lh.Size() != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, lh.Size())
	}

	// Test DeleteMin order
	expected := []int{1, 3, 4, 5, 7, 9}
	for _, exp := range expected {
		if min, ok := lh.DeleteMin(); !ok || min != exp {
			t.Errorf("Expected DeleteMin to return %d, got %d", exp, min)
		}
	}

	// Test empty after all deletes
	if !lh.IsEmpty() {
		t.Error("Heap should be empty after deleting all elements")
	}
}

func TestLeftistHeapMerge(t *testing.T) {
	h1 := NewLeftistHeap()
	h2 := NewLeftistHeap()

	// Insert elements into first heap
	for _, v := range []int{5, 3, 7} {
		h1.Insert(v)
	}

	// Insert elements into second heap
	for _, v := range []int{4, 6, 2} {
		h2.Insert(v)
	}

	// Merge heaps
	h1.Merge(h2)

	// Test size after merge
	expectedSize := 6
	if h1.Size() != expectedSize {
		t.Errorf("Expected merged size to be %d, got %d", expectedSize, h1.Size())
	}

	// Test if h2 is empty after merge
	if !h2.IsEmpty() {
		t.Error("Second heap should be empty after merge")
	}

	// Test if elements come out in correct order
	expected := []int{2, 3, 4, 5, 6, 7}
	for _, exp := range expected {
		if min, ok := h1.DeleteMin(); !ok || min != exp {
			t.Errorf("Expected DeleteMin to return %d, got %d", exp, min)
		}
	}
}

func TestLeftistHeapClear(t *testing.T) {
	lh := NewLeftistHeap()

	// Insert some elements
	values := []int{5, 3, 7, 1, 9}
	for _, v := range values {
		lh.Insert(v)
	}

	// Clear the heap
	lh.Clear()

	// Test if heap is empty after clear
	if !lh.IsEmpty() {
		t.Error("Heap should be empty after Clear()")
	}

	if lh.Size() != 0 {
		t.Errorf("Expected size 0 after Clear(), got %d", lh.Size())
	}

	// Test if FindMin returns false
	if _, ok := lh.FindMin(); ok {
		t.Error("FindMin should return false on empty heap")
	}
}

func TestLeftistHeapDeleteMinEmpty(t *testing.T) {
	lh := NewLeftistHeap()

	// Try to delete from empty heap
	if _, ok := lh.DeleteMin(); ok {
		t.Error("DeleteMin should return false on empty heap")
	}
}

func TestLeftistHeapNullPathLength(t *testing.T) {
	lh := NewLeftistHeap()

	// Insert elements to create a specific structure
	values := []int{5, 3, 7, 1, 9, 4, 6, 8, 2}
	for _, v := range values {
		lh.Insert(v)
	}

	// Helper function to check NPL property
	var checkNPL func(*LeftistNode) bool
	checkNPL = func(node *LeftistNode) bool {
		if node == nil {
			return true
		}

		// Calculate NPL
		rightNPL := -1
		if node.rightChild != nil {
			rightNPL = node.rightChild.npl
		}

		// Verify NPL value
		if node.npl != rightNPL+1 {
			return false
		}

		// Verify leftist property
		leftNPL := -1
		if node.leftChild != nil {
			leftNPL = node.leftChild.npl
		}
		if leftNPL < rightNPL {
			return false
		}

		// Recursively check children
		return checkNPL(node.leftChild) && checkNPL(node.rightChild)
	}

	if !checkNPL(lh.root) {
		t.Error("Leftist heap property violation detected")
	}
}
