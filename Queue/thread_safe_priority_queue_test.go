package queue

import (
	"sync"
	"testing"
)

func TestThreadSafePriorityQueue(t *testing.T) {
	t.Run("Basic Operations", func(t *testing.T) {
		pq := NewThreadSafePriorityQueue[string]()

		// Test initial state
		if !pq.IsEmpty() {
			t.Error("New queue should be empty")
		}
		if pq.Size() != 0 {
			t.Error("New queue should have size 0")
		}

		// Test Enqueue
		pq.Enqueue("first", 1)
		pq.Enqueue("second", 2)
		pq.Enqueue("third", 3)

		if pq.Size() != 3 {
			t.Error("Queue size should be 3")
		}

		// Test Peek
		value, ok := pq.Peek()
		if !ok || value != "third" {
			t.Error("Peek should return the highest priority item")
		}

		// Test Dequeue
		value, ok = pq.Dequeue()
		if !ok || value != "third" {
			t.Error("Dequeue should return the highest priority item")
		}

		value, ok = pq.Dequeue()
		if !ok || value != "second" {
			t.Error("Dequeue should return the second highest priority item")
		}

		value, ok = pq.Dequeue()
		if !ok || value != "first" {
			t.Error("Dequeue should return the lowest priority item")
		}

		// Test empty queue
		_, ok = pq.Dequeue()
		if ok {
			t.Error("Dequeue on empty queue should return false")
		}
	})

	t.Run("Priority Updates", func(t *testing.T) {
		pq := NewThreadSafePriorityQueue[string]()
		equals := func(a, b string) bool { return a == b }

		pq.Enqueue("task1", 1)
		pq.Enqueue("task2", 2)

		// Test GetPriority
		priority, exists := pq.GetPriority("task1", equals)
		if !exists || priority != 1 {
			t.Error("GetPriority failed for existing item")
		}

		// Test UpdatePriority
		success := pq.UpdatePriority("task1", 3, equals)
		if !success {
			t.Error("UpdatePriority should succeed for existing item")
		}

		value, _ := pq.Peek()
		if value != "task1" {
			t.Error("After priority update, task1 should be at the top")
		}
	})

	t.Run("Clear Operation", func(t *testing.T) {
		pq := NewThreadSafePriorityQueue[int]()
		pq.Enqueue(1, 1)
		pq.Enqueue(2, 2)

		pq.Clear()
		if !pq.IsEmpty() {
			t.Error("Queue should be empty after Clear")
		}
		if pq.Size() != 0 {
			t.Error("Queue size should be 0 after Clear")
		}
	})

	t.Run("Concurrent Operations", func(t *testing.T) {
		pq := NewThreadSafePriorityQueue[int]()
		const numOperations = 1000
		var wg sync.WaitGroup

		// Concurrent Enqueue
		wg.Add(numOperations)
		for i := 0; i < numOperations; i++ {
			go func(val int) {
				defer wg.Done()
				pq.Enqueue(val, val)
			}(i)
		}
		wg.Wait()

		if pq.Size() != numOperations {
			t.Errorf("Expected size %d, got %d", numOperations, pq.Size())
		}

		// Concurrent Dequeue
		wg.Add(numOperations)
		for i := 0; i < numOperations; i++ {
			go func() {
				defer wg.Done()
				pq.Dequeue()
			}()
		}
		wg.Wait()

		if !pq.IsEmpty() {
			t.Error("Queue should be empty after all dequeues")
		}
	})

	t.Run("Mixed Priority Items", func(t *testing.T) {
		pq := NewThreadSafePriorityQueue[string]()

		// Add items with mixed priorities
		pq.Enqueue("low", 1)
		pq.Enqueue("highest", 10)
		pq.Enqueue("medium", 5)
		pq.Enqueue("high", 8)

		// Verify dequeue order
		expected := []string{"highest", "high", "medium", "low"}
		for _, exp := range expected {
			value, ok := pq.Dequeue()
			if !ok || value != exp {
				t.Errorf("Expected %s, got %s", exp, value)
			}
		}
	})
}
