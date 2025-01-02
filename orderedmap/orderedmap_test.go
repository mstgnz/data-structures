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
