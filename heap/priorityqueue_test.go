package heap

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	t.Run("New PriorityQueue should be empty", func(t *testing.T) {
		pq := NewPriorityQueue()
		if !pq.IsEmpty() {
			t.Error("New priority queue should be empty")
		}
		if pq.Size() != 0 {
			t.Error("New priority queue should have size 0")
		}
	})

	t.Run("Enqueue and Dequeue should maintain priority order", func(t *testing.T) {
		pq := NewPriorityQueue()

		// Insert items with priorities
		items := []struct {
			value    string
			priority int
		}{
			{"Low", 3},
			{"High", 1},
			{"Medium", 2},
			{"VeryHigh", 0},
		}

		for _, item := range items {
			pq.Enqueue(item.value, item.priority)
		}

		// Expected order based on priority (lowest first)
		expected := []string{"VeryHigh", "High", "Medium", "Low"}

		// Dequeue and verify order
		for _, want := range expected {
			got, err := pq.Dequeue()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if got != want {
				t.Errorf("Expected value %s, got %v", want, got)
			}
		}
	})

	t.Run("Peek should return highest priority item without removing it", func(t *testing.T) {
		pq := NewPriorityQueue()
		pq.Enqueue("Medium", 2)
		pq.Enqueue("High", 1)
		pq.Enqueue("Low", 3)

		// Peek should return "High"
		value, err := pq.Peek()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if value != "High" {
			t.Errorf("Expected peek to return High, got %v", value)
		}

		// Size should still be 3
		if pq.Size() != 3 {
			t.Errorf("Expected size to be 3, got %d", pq.Size())
		}
	})

	t.Run("Dequeue from empty queue should return error", func(t *testing.T) {
		pq := NewPriorityQueue()
		_, err := pq.Dequeue()
		if err == nil {
			t.Error("Expected error when dequeuing from empty queue")
		}
	})

	t.Run("Peek at empty queue should return error", func(t *testing.T) {
		pq := NewPriorityQueue()
		_, err := pq.Peek()
		if err == nil {
			t.Error("Expected error when peeking at empty queue")
		}
	})

	t.Run("Should handle multiple items with same priority", func(t *testing.T) {
		pq := NewPriorityQueue()
		pq.Enqueue("First", 1)
		pq.Enqueue("Second", 1)
		pq.Enqueue("Third", 1)

		size := pq.Size()
		if size != 3 {
			t.Errorf("Expected size 3, got %d", size)
		}

		// Items with same priority should maintain FIFO order
		expected := []string{"First", "Second", "Third"}
		for _, want := range expected {
			got, err := pq.Dequeue()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if got != want {
				t.Errorf("Expected %s, got %v", want, got)
			}
		}
	})
}
