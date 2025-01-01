package linkedlist

import (
	"sync"
	"testing"
)

func TestThreadSafeLinkedList(t *testing.T) {
	list := NewThreadSafeLinkedList[int]()

	// Test initial state
	if !list.IsEmpty() {
		t.Error("New list should be empty")
	}

	if size := list.Size(); size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}

	// Test AddFirst and GetFirst
	list.AddFirst(1)
	if val, ok := list.GetFirst(); !ok || val != 1 {
		t.Errorf("Expected first value 1, got %v", val)
	}

	// Test AddLast and GetLast
	list.AddLast(2)
	if val, ok := list.GetLast(); !ok || val != 2 {
		t.Errorf("Expected last value 2, got %v", val)
	}

	// Test Contains
	if !Contains(list, 1) {
		t.Error("List should contain 1")
	}
	if !Contains(list, 2) {
		t.Error("List should contain 2")
	}
	if Contains(list, 3) {
		t.Error("List should not contain 3")
	}

	// Test RemoveFirst
	if val, ok := list.RemoveFirst(); !ok || val != 1 {
		t.Errorf("Expected removed first value 1, got %v", val)
	}

	// Test RemoveLast
	if val, ok := list.RemoveLast(); !ok || val != 2 {
		t.Errorf("Expected removed last value 2, got %v", val)
	}

	// Test empty list behavior
	if _, ok := list.RemoveFirst(); ok {
		t.Error("RemoveFirst on empty list should return false")
	}
	if _, ok := list.RemoveLast(); ok {
		t.Error("RemoveLast on empty list should return false")
	}
}

func TestThreadSafeLinkedListConcurrent(t *testing.T) {
	list := NewThreadSafeLinkedList[int]()
	var wg sync.WaitGroup
	numGoroutines := 100
	numOperations := 100

	// Test concurrent AddFirst and AddLast
	wg.Add(numGoroutines * 2)
	for i := 0; i < numGoroutines; i++ {
		go func(n int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				list.AddFirst(n*numOperations + j)
			}
		}(i)
		go func(n int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				list.AddLast(n*numOperations + j)
			}
		}(i)
	}
	wg.Wait()

	expectedSize := numGoroutines * numOperations * 2
	if size := list.Size(); size != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, size)
	}

	// Test concurrent RemoveFirst and RemoveLast
	wg.Add(numGoroutines * 2)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				list.RemoveFirst()
			}
		}()
		go func() {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				list.RemoveLast()
			}
		}()
	}
	wg.Wait()

	if !list.IsEmpty() {
		t.Error("List should be empty after all removes")
	}
}

func TestThreadSafeLinkedListClear(t *testing.T) {
	list := NewThreadSafeLinkedList[string]()

	// Add some items
	items := []string{"a", "b", "c"}
	for _, item := range items {
		list.AddLast(item)
	}

	// Clear the list
	list.Clear()

	if !list.IsEmpty() {
		t.Error("List should be empty after clear")
	}

	if size := list.Size(); size != 0 {
		t.Errorf("Expected size 0 after clear, got %d", size)
	}
}
