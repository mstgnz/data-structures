package heap

import (
	"testing"
)

func TestFibonacciHeap(t *testing.T) {
	t.Run("New FibonacciHeap should be empty", func(t *testing.T) {
		fh := NewFibonacciHeap()
		if !fh.IsEmpty() {
			t.Error("New Fibonacci heap should be empty")
		}
		if fh.Size() != 0 {
			t.Error("New Fibonacci heap should have size 0")
		}
	})

	t.Run("Insert and Extract should maintain heap property", func(t *testing.T) {
		fh := NewFibonacciHeap()
		values := []int{5, 3, 7, 1, 4, 6, 2}
		expected := []int{1, 2, 3, 4, 5, 6, 7}

		// Insert values
		for _, v := range values {
			fh.Insert(v)
		}

		// Size should be correct
		if fh.Size() != len(values) {
			t.Errorf("Expected size %d, got %d", len(values), fh.Size())
		}

		// Extract values and verify they come out in ascending order
		for _, want := range expected {
			got, err := fh.Extract()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if got != want {
				t.Errorf("Expected %d, got %d", want, got)
			}
		}
	})

	t.Run("Peek should return minimum element without removing it", func(t *testing.T) {
		fh := NewFibonacciHeap()
		fh.Insert(3)
		fh.Insert(1)
		fh.Insert(2)

		// Peek should return 1
		min, err := fh.Peek()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if min != 1 {
			t.Errorf("Expected peek to return 1, got %d", min)
		}

		// Size should still be 3
		if fh.Size() != 3 {
			t.Errorf("Expected size to be 3, got %d", fh.Size())
		}
	})

	t.Run("Extract from empty heap should return error", func(t *testing.T) {
		fh := NewFibonacciHeap()
		_, err := fh.Extract()
		if err == nil {
			t.Error("Expected error when extracting from empty heap")
		}
	})

	t.Run("Peek at empty heap should return error", func(t *testing.T) {
		fh := NewFibonacciHeap()
		_, err := fh.Peek()
		if err == nil {
			t.Error("Expected error when peeking at empty heap")
		}
	})

	t.Run("DecreaseKey should maintain heap property", func(t *testing.T) {
		fh := NewFibonacciHeap()

		// Insert some values and keep track of nodes
		fh.Insert(5)
		node1 := fh.min
		fh.Insert(3)
		_ = fh.min // Skip storing node2
		fh.Insert(7)
		node3 := fh.min.right

		// Decrease key of node3 from 7 to 1
		err := fh.DecreaseKey(node3, 1)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		// Minimum should now be 1
		min, err := fh.Peek()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if min != 1 {
			t.Errorf("Expected minimum to be 1 after DecreaseKey, got %d", min)
		}

		// Try to increase key (should fail)
		err = fh.DecreaseKey(node1, 10)
		if err == nil {
			t.Error("Expected error when trying to increase key")
		}
	})

	t.Run("Should handle large number of operations", func(t *testing.T) {
		fh := NewFibonacciHeap()
		n := 100 // Test with 100 elements

		// Insert values in reverse order
		for i := n; i > 0; i-- {
			fh.Insert(i)
		}

		// Extract all values and verify they come out sorted
		for i := 1; i <= n; i++ {
			got, err := fh.Extract()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if got != i {
				t.Errorf("Expected %d, got %d", i, got)
			}
		}

		// Heap should be empty now
		if !fh.IsEmpty() {
			t.Error("Heap should be empty after extracting all elements")
		}
	})

	t.Run("Should maintain heap property after multiple operations", func(t *testing.T) {
		fh := NewFibonacciHeap()

		// Insert values
		fh.Insert(5)
		node1 := fh.min
		fh.Insert(3)
		fh.Insert(7)

		// Decrease key of 5 to 2
		err := fh.DecreaseKey(node1, 2)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		// Extract all values and verify order
		expected := []int{2, 3, 7}
		for _, want := range expected {
			got, err := fh.Extract()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if got != want {
				t.Errorf("Expected %d, got %d", want, got)
			}
		}
	})
}
