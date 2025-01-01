package heap

import (
	"testing"
)

func TestMaxHeap(t *testing.T) {
	t.Run("New MaxHeap should be empty", func(t *testing.T) {
		heap := NewMaxHeap()
		if !heap.IsEmpty() {
			t.Error("New heap should be empty")
		}
		if heap.Size() != 0 {
			t.Error("New heap should have size 0")
		}
	})

	t.Run("Insert and Extract should maintain max-heap property", func(t *testing.T) {
		heap := NewMaxHeap()
		values := []int{5, 3, 7, 1, 4, 6, 2}
		expected := []int{7, 6, 5, 4, 3, 2, 1}

		// Insert values
		for _, v := range values {
			heap.Insert(v)
		}

		// Extract values and verify they come out in descending order
		for _, want := range expected {
			got, err := heap.Extract()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if got != want {
				t.Errorf("Expected %d, got %d", want, got)
			}
		}
	})

	t.Run("Peek should return maximum element without removing it", func(t *testing.T) {
		heap := NewMaxHeap()
		heap.Insert(3)
		heap.Insert(1)
		heap.Insert(2)

		// Peek should return 3
		max, err := heap.Peek()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if max != 3 {
			t.Errorf("Expected peek to return 3, got %d", max)
		}

		// Size should still be 3
		if heap.Size() != 3 {
			t.Errorf("Expected size to be 3, got %d", heap.Size())
		}
	})

	t.Run("Extract from empty heap should return error", func(t *testing.T) {
		heap := NewMaxHeap()
		_, err := heap.Extract()
		if err == nil {
			t.Error("Expected error when extracting from empty heap")
		}
	})

	t.Run("Peek at empty heap should return error", func(t *testing.T) {
		heap := NewMaxHeap()
		_, err := heap.Peek()
		if err == nil {
			t.Error("Expected error when peeking at empty heap")
		}
	})
}
