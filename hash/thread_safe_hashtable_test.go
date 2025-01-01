package hash

import (
	"sync"
	"testing"
)

func TestThreadSafeHashTable(t *testing.T) {
	table := NewThreadSafeHashTable[string, int](16)

	// Test initial state
	if !table.IsEmpty() {
		t.Error("New table should be empty")
	}

	if size := table.Size(); size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}

	// Test Put and Get
	table.Put("one", 1)
	if val, ok := table.Get("one"); !ok || val != 1 {
		t.Errorf("Expected value 1, got %v", val)
	}

	// Test Contains
	if !table.Contains("one") {
		t.Error("Table should contain key 'one'")
	}
	if table.Contains("two") {
		t.Error("Table should not contain key 'two'")
	}

	// Test Update
	table.Put("one", 100)
	if val, ok := table.Get("one"); !ok || val != 100 {
		t.Errorf("Expected updated value 100, got %v", val)
	}

	// Test Remove
	if !table.Remove("one") {
		t.Error("Remove should return true for existing key")
	}
	if table.Remove("two") {
		t.Error("Remove should return false for non-existing key")
	}

	// Test empty table behavior
	if _, ok := table.Get("one"); ok {
		t.Error("Get should return false for removed key")
	}
}

func TestThreadSafeHashTableConcurrent(t *testing.T) {
	table := NewThreadSafeHashTable[string, int](16)
	var wg sync.WaitGroup
	numGoroutines := 100
	numOperations := 100

	// Test concurrent Put operations
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(n int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				key := string(rune('A' + n%26))
				table.Put(key, n*numOperations+j)
			}
		}(i)
	}
	wg.Wait()

	// Test concurrent Get and Remove operations
	wg.Add(numGoroutines * 2)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				key := string(rune('A' + j%26))
				table.Get(key)
			}
		}()
		go func() {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				key := string(rune('A' + j%26))
				table.Remove(key)
			}
		}()
	}
	wg.Wait()
}

func TestThreadSafeHashTableClear(t *testing.T) {
	table := NewThreadSafeHashTable[string, string](16)

	// Add some items
	items := map[string]string{
		"a": "apple",
		"b": "banana",
		"c": "cherry",
	}
	for k, v := range items {
		table.Put(k, v)
	}

	// Clear the table
	table.Clear()

	if !table.IsEmpty() {
		t.Error("Table should be empty after clear")
	}

	if size := table.Size(); size != 0 {
		t.Errorf("Expected size 0 after clear, got %d", size)
	}

	// Verify all items are removed
	for k := range items {
		if table.Contains(k) {
			t.Errorf("Table should not contain key %s after clear", k)
		}
	}
}
