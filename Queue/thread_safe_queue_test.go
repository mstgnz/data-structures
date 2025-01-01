package queue

import (
	"sync"
	"testing"
)

func TestThreadSafeQueue(t *testing.T) {
	queue := NewThreadSafeQueue[int]()

	// Test initial state
	if !queue.IsEmpty() {
		t.Error("New queue should be empty")
	}

	if size := queue.Size(); size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}

	// Test Enqueue and Front
	queue.Enqueue(1)
	if val, ok := queue.Front(); !ok || val != 1 {
		t.Errorf("Expected front value 1, got %v", val)
	}

	// Test Dequeue
	if val, ok := queue.Dequeue(); !ok || val != 1 {
		t.Errorf("Expected dequeue value 1, got %v", val)
	}

	// Test empty queue behavior
	if _, ok := queue.Dequeue(); ok {
		t.Error("Dequeue on empty queue should return false")
	}

	if _, ok := queue.Front(); ok {
		t.Error("Front on empty queue should return false")
	}
}

func TestThreadSafeQueueConcurrent(t *testing.T) {
	queue := NewThreadSafeQueue[int]()
	var wg sync.WaitGroup
	numGoroutines := 100
	numOperations := 1000

	// Test concurrent enqueues
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(n int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				queue.Enqueue(n*numOperations + j)
			}
		}(i)
	}
	wg.Wait()

	expectedSize := numGoroutines * numOperations
	if size := queue.Size(); size != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, size)
	}

	// Test concurrent dequeues
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				queue.Dequeue()
			}
		}()
	}
	wg.Wait()

	if !queue.IsEmpty() {
		t.Error("Queue should be empty after all dequeues")
	}
}

func TestThreadSafeQueueClear(t *testing.T) {
	queue := NewThreadSafeQueue[string]()

	// Add some items
	items := []string{"a", "b", "c"}
	for _, item := range items {
		queue.Enqueue(item)
	}

	// Clear the queue
	queue.Clear()

	if !queue.IsEmpty() {
		t.Error("Queue should be empty after clear")
	}

	if size := queue.Size(); size != 0 {
		t.Errorf("Expected size 0 after clear, got %d", size)
	}
}
