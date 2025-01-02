package tree

import "testing"

func TestBPlusTreeInsert(t *testing.T) {
	bt := NewBPlusTree()

	// Test inserting elements
	values := []int{5, 15, 25, 35, 45}
	for _, v := range values {
		bt.Insert(v)
	}

	// Test if all values can be found
	for _, v := range values {
		if !bt.Search(v) {
			t.Errorf("Value %d not found after insertion", v)
		}
	}

	if bt.Size() != len(values) {
		t.Errorf("Expected size to be %d, got %d", len(values), bt.Size())
	}
}

func TestBPlusTreeSearch(t *testing.T) {
	bt := NewBPlusTree()
	values := []int{10, 20, 30, 40, 50, 60, 70}

	// Insert test values
	for _, v := range values {
		bt.Insert(v)
	}

	// Test searching existing values
	for _, v := range values {
		if !bt.Search(v) {
			t.Errorf("Search failed to find existing value %d", v)
		}
	}

	// Test searching non-existing values
	nonExisting := []int{15, 25, 35, 45, 55, 65}
	for _, v := range nonExisting {
		if bt.Search(v) {
			t.Errorf("Search found non-existing value %d", v)
		}
	}
}

func TestBPlusTreeEmpty(t *testing.T) {
	bt := NewBPlusTree()

	if !bt.IsEmpty() {
		t.Error("New tree should be empty")
	}

	bt.Insert(1)
	if bt.IsEmpty() {
		t.Error("Tree with one element should not be empty")
	}
}

func TestBPlusTreeClear(t *testing.T) {
	bt := NewBPlusTree()
	values := []int{10, 20, 30, 40, 50}

	// Insert test values
	for _, v := range values {
		bt.Insert(v)
	}

	bt.Clear()
	if !bt.IsEmpty() {
		t.Error("Tree should be empty after Clear()")
	}
	if bt.Size() != 0 {
		t.Errorf("Expected size to be 0 after Clear(), got %d", bt.Size())
	}
}

func TestBPlusTreeNodeSplitting(t *testing.T) {
	bt := NewBPlusTree()

	// Insert enough values to cause node splitting
	values := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for _, v := range values {
		bt.Insert(v)
	}

	// Verify all values are still accessible after splits
	for _, v := range values {
		if !bt.Search(v) {
			t.Errorf("Value %d not found after node splitting", v)
		}
	}
}

func TestBPlusTreeDuplicateInsert(t *testing.T) {
	bt := NewBPlusTree()

	// Insert same value multiple times
	bt.Insert(10)
	initialSize := bt.Size()

	bt.Insert(10)
	if bt.Size() != initialSize+1 {
		t.Error("Duplicate insertion should increase size")
	}

	if !bt.Search(10) {
		t.Error("Value 10 should be found after duplicate insertion")
	}
}

func TestBPlusTreeOrdering(t *testing.T) {
	bt := NewBPlusTree()

	// Insert values in random order
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, v := range values {
		bt.Insert(v)
	}

	// All values should be found
	for _, v := range values {
		if !bt.Search(v) {
			t.Errorf("Value %d not found after insertion", v)
		}
	}
}
