package heap

import (
	"testing"
)

func TestBinomialHeap(t *testing.T) {
	t.Run("New BinomialHeap should be empty", func(t *testing.T) {
		bh := NewBinomialHeap()
		if !bh.IsEmpty() {
			t.Error("New binomial heap should be empty")
		}
		if bh.Size() != 0 {
			t.Error("New binomial heap should have size 0")
		}
	})

	t.Run("Insert and Extract should maintain heap property", func(t *testing.T) {
		bh := NewBinomialHeap()
		values := []int{5, 3, 7, 1, 4, 6, 2}
		expected := []int{1, 2, 3, 4, 5, 6, 7}

		// Insert values
		for _, v := range values {
			bh.Insert(v)
		}

		// Size should be correct
		if bh.Size() != len(values) {
			t.Errorf("Expected size %d, got %d", len(values), bh.Size())
		}

		// Extract values and verify they come out in ascending order
		for _, want := range expected {
			got, err := bh.Extract()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if got != want {
				t.Errorf("Expected %d, got %d", want, got)
			}
		}
	})

	t.Run("Peek should return minimum element without removing it", func(t *testing.T) {
		bh := NewBinomialHeap()
		bh.Insert(3)
		bh.Insert(1)
		bh.Insert(2)

		// Peek should return 1
		min, err := bh.Peek()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if min != 1 {
			t.Errorf("Expected peek to return 1, got %d", min)
		}

		// Size should still be 3
		if bh.Size() != 3 {
			t.Errorf("Expected size to be 3, got %d", bh.Size())
		}
	})

	t.Run("Extract from empty heap should return error", func(t *testing.T) {
		bh := NewBinomialHeap()
		_, err := bh.Extract()
		if err == nil {
			t.Error("Expected error when extracting from empty heap")
		}
	})

	t.Run("Peek at empty heap should return error", func(t *testing.T) {
		bh := NewBinomialHeap()
		_, err := bh.Peek()
		if err == nil {
			t.Error("Expected error when peeking at empty heap")
		}
	})

	t.Run("Should handle large number of insertions and extractions", func(t *testing.T) {
		bh := NewBinomialHeap()
		n := 100 // Test with 100 elements

		// Insert values in reverse order
		for i := n; i > 0; i-- {
			bh.Insert(i)
		}

		// Extract all values and verify they come out sorted
		for i := 1; i <= n; i++ {
			got, err := bh.Extract()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if got != i {
				t.Errorf("Expected %d, got %d", i, got)
			}
		}

		// Heap should be empty now
		if !bh.IsEmpty() {
			t.Error("Heap should be empty after extracting all elements")
		}
	})

	t.Run("Should maintain heap property after multiple operations", func(t *testing.T) {
		bh := NewBinomialHeap()

		// Insert some values
		bh.Insert(5)
		bh.Insert(3)
		bh.Insert(7)

		// Extract minimum (should be 3)
		min, err := bh.Extract()
		if err != nil || min != 3 {
			t.Errorf("Expected 3, got %d with error: %v", min, err)
		}

		// Insert more values
		bh.Insert(1)
		bh.Insert(4)

		// Extract all remaining values and verify order
		expected := []int{1, 4, 5, 7}
		for _, want := range expected {
			got, err := bh.Extract()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if got != want {
				t.Errorf("Expected %d, got %d", want, got)
			}
		}
	})
}
