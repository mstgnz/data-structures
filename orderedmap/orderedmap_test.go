package orderedmap

import (
	"sync"
	"testing"
)

func TestOrderedMap_BasicOperations(t *testing.T) {
	om := New()

	// Test Set and Get
	t.Run("Set and Get", func(t *testing.T) {
		om.Set("key1", "value1")
		if val, exists := om.Get("key1"); !exists || val != "value1" {
			t.Errorf("Expected value1, got %v", val)
		}
	})

	// Test non-existent key
	t.Run("Get Non-existent Key", func(t *testing.T) {
		if val, exists := om.Get("nonexistent"); exists || val != nil {
			t.Errorf("Expected nil and false for non-existent key")
		}
	})

	// Test update existing key
	t.Run("Update Existing Key", func(t *testing.T) {
		om.Set("key1", "updated_value")
		if val, exists := om.Get("key1"); !exists || val != "updated_value" {
			t.Errorf("Expected updated_value, got %v", val)
		}
	})

	// Test Delete
	t.Run("Delete", func(t *testing.T) {
		om.Delete("key1")
		if val, exists := om.Get("key1"); exists || val != nil {
			t.Errorf("Expected key to be deleted")
		}
	})
}

func TestOrderedMap_Order(t *testing.T) {
	om := New()

	// Add elements in specific order
	elements := []struct {
		key   string
		value int
	}{
		{"first", 1},
		{"second", 2},
		{"third", 3},
	}

	for _, elem := range elements {
		om.Set(elem.key, elem.value)
	}

	// Test Keys order
	t.Run("Keys Order", func(t *testing.T) {
		keys := om.Keys()
		if len(keys) != len(elements) {
			t.Errorf("Expected %d keys, got %d", len(elements), len(keys))
		}
		for i, elem := range elements {
			if keys[i] != elem.key {
				t.Errorf("Expected key %s at position %d, got %v", elem.key, i, keys[i])
			}
		}
	})

	// Test Values order
	t.Run("Values Order", func(t *testing.T) {
		values := om.Values()
		if len(values) != len(elements) {
			t.Errorf("Expected %d values, got %d", len(elements), len(values))
		}
		for i, elem := range elements {
			if values[i] != elem.value {
				t.Errorf("Expected value %d at position %d, got %v", elem.value, i, values[i])
			}
		}
	})
}

func TestOrderedMap_ConcurrentOperations(t *testing.T) {
	om := New()
	var wg sync.WaitGroup
	numGoroutines := 100

	// Test concurrent writes
	t.Run("Concurrent Writes", func(t *testing.T) {
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(val int) {
				defer wg.Done()
				key := val
				om.Set(key, val)
			}(i)
		}
		wg.Wait()

		if om.Len() != numGoroutines {
			t.Errorf("Expected length %d, got %d", numGoroutines, om.Len())
		}
	})

	// Test concurrent reads
	t.Run("Concurrent Reads", func(t *testing.T) {
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(val int) {
				defer wg.Done()
				key := val
				if _, exists := om.Get(key); !exists {
					t.Errorf("Key %v should exist", key)
				}
			}(i)
		}
		wg.Wait()
	})

	// Test concurrent reads and writes
	t.Run("Concurrent Reads and Writes", func(t *testing.T) {
		for i := 0; i < numGoroutines; i++ {
			wg.Add(2)
			// Reader
			go func(val int) {
				defer wg.Done()
				om.Get(val)
			}(i)
			// Writer
			go func(val int) {
				defer wg.Done()
				om.Set(val, val*2)
			}(i)
		}
		wg.Wait()
	})

	// Test concurrent deletes
	t.Run("Concurrent Deletes", func(t *testing.T) {
		initialLen := om.Len()
		for i := 0; i < numGoroutines/2; i++ {
			wg.Add(1)
			go func(val int) {
				defer wg.Done()
				om.Delete(val)
			}(i)
		}
		wg.Wait()

		expectedLen := initialLen - numGoroutines/2
		if om.Len() != expectedLen {
			t.Errorf("Expected length %d after deletes, got %d", expectedLen, om.Len())
		}
	})
}

func TestOrderedMap_String(t *testing.T) {
	om := New()
	om.Set("key1", 1)
	om.Set("key2", 2)

	str := om.String()
	expected := "{key1: 1, key2: 2}"
	if str != expected {
		t.Errorf("Expected string representation %s, got %s", expected, str)
	}
}

func TestOrderedMap_EmptyOperations(t *testing.T) {
	om := New()

	t.Run("Empty Map Operations", func(t *testing.T) {
		if om.Len() != 0 {
			t.Errorf("Expected empty map length 0, got %d", om.Len())
		}

		if len(om.Keys()) != 0 {
			t.Errorf("Expected empty keys slice")
		}

		if len(om.Values()) != 0 {
			t.Errorf("Expected empty values slice")
		}

		if str := om.String(); str != "{}" {
			t.Errorf("Expected empty map string {}, got %s", str)
		}
	})
}

func TestOrderedMap_Range(t *testing.T) {
	om := New()
	elements := []struct {
		key   string
		value int
	}{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}

	// Add elements
	for _, elem := range elements {
		om.Set(elem.key, elem.value)
	}

	// Test Range
	t.Run("Range All Elements", func(t *testing.T) {
		index := 0
		om.Range(func(key, value any) bool {
			if key != elements[index].key || value != elements[index].value {
				t.Errorf("Expected (%v, %v) at index %d, got (%v, %v)",
					elements[index].key, elements[index].value,
					index, key, value)
			}
			index++
			return true
		})
		if index != len(elements) {
			t.Errorf("Expected to iterate over %d elements, got %d", len(elements), index)
		}
	})

	// Test Range Early Stop
	t.Run("Range Early Stop", func(t *testing.T) {
		count := 0
		om.Range(func(key, value any) bool {
			count++
			return count < 2 // Stop after first element
		})
		if count != 2 {
			t.Errorf("Expected to stop after 2 elements, got %d", count)
		}
	})
}

func TestOrderedMap_Clear(t *testing.T) {
	om := New()
	om.Set("key1", 1)
	om.Set("key2", 2)

	t.Run("Clear Map", func(t *testing.T) {
		om.Clear()
		if om.Len() != 0 {
			t.Errorf("Expected empty map after clear, got length %d", om.Len())
		}
		if len(om.Keys()) != 0 {
			t.Errorf("Expected no keys after clear")
		}
		if val, exists := om.Get("key1"); exists {
			t.Errorf("Expected no values after clear, got %v", val)
		}
	})
}

func TestOrderedMap_Copy(t *testing.T) {
	om := New()
	om.Set("key1", 1)
	om.Set("key2", 2)

	t.Run("Copy Map", func(t *testing.T) {
		copy := om.Copy()

		// Check length
		if copy.Len() != om.Len() {
			t.Errorf("Expected copy to have same length")
		}

		// Check all elements
		om.Range(func(key, value any) bool {
			copyVal, exists := copy.Get(key)
			if !exists {
				t.Errorf("Key %v not found in copy", key)
				return false
			}
			if copyVal != value {
				t.Errorf("Value mismatch for key %v: expected %v, got %v", key, value, copyVal)
			}
			return true
		})

		// Verify independence
		om.Set("key3", 3)
		if _, exists := copy.Get("key3"); exists {
			t.Error("Copy should not be affected by changes to original")
		}
	})
}

func TestOrderedMap_Has(t *testing.T) {
	om := New()
	om.Set("key1", 1)

	t.Run("Has Existing Key", func(t *testing.T) {
		if !om.Has("key1") {
			t.Error("Expected Has to return true for existing key")
		}
	})

	t.Run("Has Non-existing Key", func(t *testing.T) {
		if om.Has("nonexistent") {
			t.Error("Expected Has to return false for non-existing key")
		}
	})
}

func TestOrderedMap_ConcurrentRangeAndModify(t *testing.T) {
	om := New()
	for i := 0; i < 100; i++ {
		om.Set(i, i)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	// Concurrent Range
	go func() {
		defer wg.Done()
		om.Range(func(key, value any) bool {
			return true
		})
	}()

	// Concurrent Modification
	go func() {
		defer wg.Done()
		om.Set("new", 1000)
		om.Delete(50)
	}()

	wg.Wait()
}
