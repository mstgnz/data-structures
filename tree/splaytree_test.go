package tree

import "testing"

func TestSplayTreeInsert(t *testing.T) {
	st := NewSplayTree()

	// Test inserting elements
	st.Insert(10)
	if st.root.key != 10 {
		t.Errorf("Expected root key to be 10, got %d", st.root.key)
	}

	st.Insert(5)
	if st.root.key != 5 {
		t.Errorf("Expected root key to be 5, got %d", st.root.key)
	}

	st.Insert(15)
	if st.root.key != 15 {
		t.Errorf("Expected root key to be 15, got %d", st.root.key)
	}

	if st.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", st.Size())
	}
}

func TestSplayTreeSearch(t *testing.T) {
	st := NewSplayTree()
	values := []int{10, 5, 15, 3, 7, 12, 17}

	// Insert test values
	for _, v := range values {
		st.Insert(v)
	}

	// Test searching existing values
	for _, v := range values {
		if !st.Search(v) {
			t.Errorf("Search failed to find existing value %d", v)
		}
		if st.root.key != v {
			t.Errorf("After searching %d, expected root key to be %d, got %d", v, v, st.root.key)
		}
	}

	// Test searching non-existing value
	if st.Search(100) {
		t.Error("Search found non-existing value 100")
	}
}

func TestSplayTreeDelete(t *testing.T) {
	st := NewSplayTree()
	values := []int{10, 5, 15, 3, 7, 12, 17}

	// Insert test values
	for _, v := range values {
		st.Insert(v)
	}

	// Test deleting existing values
	if !st.Delete(10) {
		t.Error("Delete failed for existing value 10")
	}
	if st.Search(10) {
		t.Error("Found deleted value 10")
	}

	if !st.Delete(5) {
		t.Error("Delete failed for existing value 5")
	}
	if st.Search(5) {
		t.Error("Found deleted value 5")
	}

	// Test deleting non-existing value
	if st.Delete(100) {
		t.Error("Delete succeeded for non-existing value 100")
	}

	if st.Size() != len(values)-2 {
		t.Errorf("Expected size to be %d, got %d", len(values)-2, st.Size())
	}
}

func TestSplayTreeEmpty(t *testing.T) {
	st := NewSplayTree()

	if !st.IsEmpty() {
		t.Error("New tree should be empty")
	}

	st.Insert(1)
	if st.IsEmpty() {
		t.Error("Tree with one element should not be empty")
	}

	st.Delete(1)
	if !st.IsEmpty() {
		t.Error("Tree after deleting all elements should be empty")
	}
}

func TestSplayTreeClear(t *testing.T) {
	st := NewSplayTree()
	values := []int{10, 5, 15, 3, 7, 12, 17}

	// Insert test values
	for _, v := range values {
		st.Insert(v)
	}

	st.Clear()
	if !st.IsEmpty() {
		t.Error("Tree should be empty after Clear()")
	}
	if st.Size() != 0 {
		t.Errorf("Expected size to be 0 after Clear(), got %d", st.Size())
	}
	if st.root != nil {
		t.Error("Root should be nil after Clear()")
	}
}
