package hash

import "testing"

// TestHashTableLinearProbing tests the linear probing implementation
func TestHashTableLinearProbing(t *testing.T) {
	ht := NewHashTable(10, "linear")

	// Test insertion
	if !ht.Put("test1", 100) {
		t.Error("Failed to insert test1")
	}
	if !ht.Put("test2", 200) {
		t.Error("Failed to insert test2")
	}

	// Test retrieval
	if val, ok := ht.Get("test1"); !ok || val.(int) != 100 {
		t.Error("Failed to get correct value for test1")
	}

	// Test update
	if !ht.Put("test1", 150) {
		t.Error("Failed to update test1")
	}
	if val, ok := ht.Get("test1"); !ok || val.(int) != 150 {
		t.Error("Failed to get updated value for test1")
	}

	// Test removal
	if !ht.Remove("test1") {
		t.Error("Failed to remove test1")
	}
	if _, ok := ht.Get("test1"); ok {
		t.Error("test1 should not exist after removal")
	}
}

// TestHashTableChaining tests the chaining implementation
func TestHashTableChaining(t *testing.T) {
	ht := NewHashTable(10, "chain")

	// Test insertion
	if !ht.Put("test1", 100) {
		t.Error("Failed to insert test1")
	}
	if !ht.Put("test2", 200) {
		t.Error("Failed to insert test2")
	}

	// Force collision by using same hash index
	index := ht.hash("test1")
	collisionKey := ""
	for i := 0; i < 100; i++ {
		testKey := "collision" + string(rune(i))
		if ht.hash(testKey) == index {
			collisionKey = testKey
			break
		}
	}
	if collisionKey != "" {
		if !ht.Put(collisionKey, 300) {
			t.Error("Failed to insert collision key")
		}
		if val, ok := ht.Get(collisionKey); !ok || val.(int) != 300 {
			t.Error("Failed to get collision key value")
		}
	}

	// Test retrieval
	if val, ok := ht.Get("test1"); !ok || val.(int) != 100 {
		t.Error("Failed to get correct value for test1")
	}

	// Test update in chain
	if !ht.Put("test1", 150) {
		t.Error("Failed to update test1")
	}
	if val, ok := ht.Get("test1"); !ok || val.(int) != 150 {
		t.Error("Failed to get updated value for test1")
	}

	// Test removal from chain
	if !ht.Remove("test1") {
		t.Error("Failed to remove test1")
	}
	if _, ok := ht.Get("test1"); ok {
		t.Error("test1 should not exist after removal")
	}
}
