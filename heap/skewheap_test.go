package heap

import "testing"

func TestSkewHeapBasicOperations(t *testing.T) {
	sh := NewSkewHeap()

	// Test initial state
	if !sh.IsEmpty() {
		t.Error("New heap should be empty")
	}

	// Test Insert and FindMin
	sh.Insert(5)
	if min, ok := sh.FindMin(); !ok || min != 5 {
		t.Errorf("Expected min to be 5, got %d", min)
	}

	// Test multiple inserts
	values := []int{3, 7, 1, 9, 4}
	for _, v := range values {
		sh.Insert(v)
	}

	// Test size
	expectedSize := len(values) + 1 // +1 for the first insert
	if sh.Size() != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, sh.Size())
	}

	// Test DeleteMin order
	expected := []int{1, 3, 4, 5, 7, 9}
	for _, exp := range expected {
		if min, ok := sh.DeleteMin(); !ok || min != exp {
			t.Errorf("Expected DeleteMin to return %d, got %d", exp, min)
		}
	}

	// Test empty after all deletes
	if !sh.IsEmpty() {
		t.Error("Heap should be empty after deleting all elements")
	}
}

func TestSkewHeapMerge(t *testing.T) {
	h1 := NewSkewHeap()
	h2 := NewSkewHeap()

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

func TestSkewHeapClear(t *testing.T) {
	sh := NewSkewHeap()

	// Insert some elements
	values := []int{5, 3, 7, 1, 9}
	for _, v := range values {
		sh.Insert(v)
	}

	// Clear the heap
	sh.Clear()

	// Test if heap is empty after clear
	if !sh.IsEmpty() {
		t.Error("Heap should be empty after Clear()")
	}

	if sh.Size() != 0 {
		t.Errorf("Expected size 0 after Clear(), got %d", sh.Size())
	}

	// Test if FindMin returns false
	if _, ok := sh.FindMin(); ok {
		t.Error("FindMin should return false on empty heap")
	}
}

func TestSkewHeapDeleteMinEmpty(t *testing.T) {
	sh := NewSkewHeap()

	// Try to delete from empty heap
	if _, ok := sh.DeleteMin(); ok {
		t.Error("DeleteMin should return false on empty heap")
	}
}

func TestSkewHeapHeapProperty(t *testing.T) {
	sh := NewSkewHeap()

	// Insert elements to create a specific structure
	values := []int{5, 3, 7, 1, 9, 4, 6, 8, 2}
	for _, v := range values {
		sh.Insert(v)
	}

	// Check heap property
	if !sh.checkHeapProperty(sh.root) {
		t.Error("Heap property violation detected")
	}

	// Delete some elements and check heap property again
	for i := 0; i < 4; i++ {
		sh.DeleteMin()
		if !sh.checkHeapProperty(sh.root) {
			t.Error("Heap property violation detected after DeleteMin")
		}
	}
}

func TestSkewHeapStress(t *testing.T) {
	sh := NewSkewHeap()

	// Insert a large number of elements
	for i := 1000; i >= 1; i-- {
		sh.Insert(i)
	}

	// Verify they come out in sorted order
	prev := 0
	for i := 0; i < 1000; i++ {
		if val, ok := sh.DeleteMin(); !ok || val <= prev {
			t.Errorf("Incorrect order: got %d after %d", val, prev)
		} else {
			prev = val
		}
	}
}
