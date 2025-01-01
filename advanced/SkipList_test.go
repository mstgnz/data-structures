package advanced

import (
	"testing"
)

// TestSkipListBasicOperations tests basic operations of skip list
func TestSkipListBasicOperations(t *testing.T) {
	sl := NewSkipList()

	// Test insertion
	sl.Insert(3, "value3")
	sl.Insert(1, "value1")
	sl.Insert(2, "value2")

	// Test search
	if val, found := sl.Search(2); !found || val.(string) != "value2" {
		t.Error("Failed to find value2")
	}

	// Test update
	sl.Insert(2, "newvalue2")
	if val, found := sl.Search(2); !found || val.(string) != "newvalue2" {
		t.Error("Failed to update value2")
	}

	// Test non-existent key
	if _, found := sl.Search(4); found {
		t.Error("Found non-existent key")
	}

	// Test deletion
	if !sl.Delete(2) {
		t.Error("Failed to delete existing key")
	}
	if sl.Delete(4) {
		t.Error("Deleted non-existent key")
	}
	if _, found := sl.Search(2); found {
		t.Error("Found deleted key")
	}
}

// TestSkipListLargeDataset tests skip list with larger dataset
func TestSkipListLargeDataset(t *testing.T) {
	sl := NewSkipList()
	n := 1000

	// Insert n elements
	for i := 0; i < n; i++ {
		sl.Insert(i, i*10)
	}

	// Verify all elements
	for i := 0; i < n; i++ {
		if val, found := sl.Search(i); !found || val.(int) != i*10 {
			t.Errorf("Failed to find correct value for key %d", i)
		}
	}

	// Delete even numbers
	for i := 0; i < n; i += 2 {
		if !sl.Delete(i) {
			t.Errorf("Failed to delete key %d", i)
		}
	}

	// Verify deletion
	for i := 0; i < n; i++ {
		_, found := sl.Search(i)
		if i%2 == 0 && found {
			t.Errorf("Found deleted key %d", i)
		} else if i%2 != 0 && !found {
			t.Errorf("Could not find existing key %d", i)
		}
	}
}

// TestSkipListLevel tests if the skip list maintains proper levels
func TestSkipListLevel(t *testing.T) {
	sl := NewSkipList()

	// Insert elements and check if level is within bounds
	for i := 0; i < 100; i++ {
		sl.Insert(i, i)
		if sl.level >= maxLevel {
			t.Error("Skip list level exceeded maximum level")
		}
	}

	// Delete all elements and check if level is properly adjusted
	for i := 0; i < 100; i++ {
		sl.Delete(i)
	}

	if sl.level != 0 {
		t.Error("Skip list level should be 0 after deleting all elements")
	}
}
