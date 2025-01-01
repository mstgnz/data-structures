package Tree

import (
	"reflect"
	"testing"
)

func TestBTree(t *testing.T) {
	t.Run("Basic Operations", func(t *testing.T) {
		btree := NewBTree(3) // Minimum degree = 3

		// Test insertion
		keys := []int{10, 20, 5, 6, 12, 30, 7, 17, 25, 22, 16, 27}
		for _, key := range keys {
			btree.Insert(key)
		}

		// Test search
		for _, key := range keys {
			if !btree.Search(key) {
				t.Errorf("Key %d should be found in the B-tree", key)
			}
		}

		// Test non-existent key
		if btree.Search(100) {
			t.Error("Key 100 should not be found in the B-tree")
		}

		// Test in-order traversal
		result := btree.GetInOrder()
		expected := []int{5, 6, 7, 10, 12, 16, 17, 20, 22, 25, 27, 30}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected in-order traversal %v, got %v", expected, result)
		}

		// Test deletion
		deleteKeys := []int{20, 5, 30}
		for _, key := range deleteKeys {
			btree.Delete(key)
			if btree.Search(key) {
				t.Errorf("Key %d should be deleted from the B-tree", key)
			}
		}

		// Test remaining keys
		remainingKeys := []int{6, 7, 10, 12, 16, 17, 22, 25, 27}
		for _, key := range remainingKeys {
			if !btree.Search(key) {
				t.Errorf("Key %d should still be in the B-tree", key)
			}
		}

		// Test in-order traversal after deletion
		result = btree.GetInOrder()
		expected = remainingKeys
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected in-order traversal after deletion %v, got %v", expected, result)
		}
	})

	t.Run("Empty Tree", func(t *testing.T) {
		btree := NewBTree(3)

		if btree.Search(1) {
			t.Error("Empty B-tree should not find any key")
		}

		result := btree.GetInOrder()
		if len(result) != 0 {
			t.Error("Empty B-tree should return empty list of keys")
		}
	})

	t.Run("Single Node Operations", func(t *testing.T) {
		btree := NewBTree(3)

		// Insert single key
		btree.Insert(10)
		if !btree.Search(10) {
			t.Error("Key 10 should be found in the B-tree")
		}

		// Delete single key
		btree.Delete(10)
		if btree.Search(10) {
			t.Error("Key 10 should be deleted from the B-tree")
		}

		result := btree.GetInOrder()
		if len(result) != 0 {
			t.Error("B-tree should be empty after deleting the only key")
		}
	})

	t.Run("Node Splitting and Merging", func(t *testing.T) {
		btree := NewBTree(2) // Minimum degree = 2 (2-3-4 tree)

		// Insert enough keys to force splits
		keys := []int{1, 2, 3, 4, 5, 6, 7}
		for _, key := range keys {
			btree.Insert(key)
		}

		// Verify all keys are present
		for _, key := range keys {
			if !btree.Search(key) {
				t.Errorf("Key %d should be found in the B-tree", key)
			}
		}

		// Delete keys to force merges
		deleteKeys := []int{4, 5, 6}
		for _, key := range deleteKeys {
			btree.Delete(key)
			if btree.Search(key) {
				t.Errorf("Key %d should be deleted from the B-tree", key)
			}
		}

		// Verify remaining keys
		remainingKeys := []int{1, 2, 3, 7}
		result := btree.GetInOrder()
		if !reflect.DeepEqual(result, remainingKeys) {
			t.Errorf("Expected remaining keys %v, got %v", remainingKeys, result)
		}
	})
}
