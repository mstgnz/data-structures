package tree

import (
	"testing"
)

func TestAVLTree(t *testing.T) {
	avl := NewAVLTree()

	// Test Insert ve Search
	values := []int{10, 20, 30, 40, 50, 25}
	for _, v := range values {
		avl.Insert(v)
	}

	// Search testi
	for _, v := range values {
		if !avl.Search(v) {
			t.Errorf("Value %d should be found in the tree", v)
		}
	}

	// Olmayan değer testi
	if avl.Search(100) {
		t.Error("Value 100 should not be found in the tree")
	}

	// InOrder traversal testi
	var result []int
	avl.InOrderTraversal(avl.Root, &result)
	expected := []int{10, 20, 25, 30, 40, 50}

	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("At index %d, expected %d, got %d", i, expected[i], result[i])
		}
	}

	// Denge kontrolü
	if avl.Root == nil {
		t.Error("Root should not be nil")
		return
	}

	balance := avl.Root.getBalance()
	if balance < -1 || balance > 1 {
		t.Errorf("Tree is not balanced. Balance factor: %d", balance)
	}
}
