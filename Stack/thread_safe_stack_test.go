package stack

import (
	"sync"
	"testing"
)

func TestThreadSafeStack(t *testing.T) {
	stack := NewThreadSafeStack[int]()

	// Test initial state
	if !stack.IsEmpty() {
		t.Error("New stack should be empty")
	}

	if size := stack.Size(); size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}

	// Test Push and Peek
	stack.Push(1)
	if val, ok := stack.Peek(); !ok || val != 1 {
		t.Errorf("Expected peek value 1, got %v", val)
	}

	// Test Pop
	if val, ok := stack.Pop(); !ok || val != 1 {
		t.Errorf("Expected pop value 1, got %v", val)
	}

	// Test empty stack behavior
	if _, ok := stack.Pop(); ok {
		t.Error("Pop on empty stack should return false")
	}

	if _, ok := stack.Peek(); ok {
		t.Error("Peek on empty stack should return false")
	}
}

func TestThreadSafeStackConcurrent(t *testing.T) {
	stack := NewThreadSafeStack[int]()
	var wg sync.WaitGroup
	numGoroutines := 100
	numOperations := 1000

	// Test concurrent pushes
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(n int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				stack.Push(n*numOperations + j)
			}
		}(i)
	}
	wg.Wait()

	expectedSize := numGoroutines * numOperations
	if size := stack.Size(); size != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, size)
	}

	// Test concurrent pops
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				stack.Pop()
			}
		}()
	}
	wg.Wait()

	if !stack.IsEmpty() {
		t.Error("Stack should be empty after all pops")
	}
}

func TestThreadSafeStackClear(t *testing.T) {
	stack := NewThreadSafeStack[string]()

	// Add some items
	items := []string{"a", "b", "c"}
	for _, item := range items {
		stack.Push(item)
	}

	// Clear the stack
	stack.Clear()

	if !stack.IsEmpty() {
		t.Error("Stack should be empty after clear")
	}

	if size := stack.Size(); size != 0 {
		t.Errorf("Expected size 0 after clear, got %d", size)
	}
}
