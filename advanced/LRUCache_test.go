package advanced

import "testing"

// TestLRUCacheBasicOperations tests basic operations of LRU cache
func TestLRUCacheBasicOperations(t *testing.T) {
	cache := NewLRUCache(3)

	// Test initial state
	if cache.GetSize() != 0 {
		t.Error("Initial size should be 0")
	}
	if cache.GetCapacity() != 3 {
		t.Error("Capacity should be 3")
	}

	// Test put and get operations
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	if val, ok := cache.Get("key1"); !ok || val.(string) != "value1" {
		t.Error("Failed to get value1")
	}
	if val, ok := cache.Get("key2"); !ok || val.(string) != "value2" {
		t.Error("Failed to get value2")
	}

	// Test non-existent key
	if _, ok := cache.Get("key3"); ok {
		t.Error("Should not find non-existent key")
	}
}

// TestLRUCacheEviction tests the LRU eviction policy
func TestLRUCacheEviction(t *testing.T) {
	cache := NewLRUCache(2)

	// Fill the cache
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	// Access key1 to make it most recently used
	cache.Get("key1")

	// Add new item, should evict key2
	cache.Put("key3", "value3")

	// Check eviction
	if cache.Contains("key2") {
		t.Error("key2 should have been evicted")
	}
	if !cache.Contains("key1") {
		t.Error("key1 should still be in cache")
	}
	if !cache.Contains("key3") {
		t.Error("key3 should be in cache")
	}
}

// TestLRUCacheUpdate tests updating existing keys
func TestLRUCacheUpdate(t *testing.T) {
	cache := NewLRUCache(2)

	// Add initial items
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	// Update existing key
	cache.Put("key1", "newvalue1")

	// Check if value was updated
	if val, ok := cache.Get("key1"); !ok || val.(string) != "newvalue1" {
		t.Error("Failed to update value")
	}

	// Size should remain the same
	if cache.GetSize() != 2 {
		t.Error("Size should not change after update")
	}
}

// TestLRUCacheClear tests clearing the cache
func TestLRUCacheClear(t *testing.T) {
	cache := NewLRUCache(3)

	// Add items
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")
	cache.Put("key3", "value3")

	// Clear cache
	cache.Clear()

	// Check if cache is empty
	if cache.GetSize() != 0 {
		t.Error("Cache should be empty after clear")
	}
	if cache.Contains("key1") || cache.Contains("key2") || cache.Contains("key3") {
		t.Error("Cache should not contain any items after clear")
	}
}

// TestLRUCacheOverflow tests behavior when adding more items than capacity
func TestLRUCacheOverflow(t *testing.T) {
	cache := NewLRUCache(2)

	// Add more items than capacity
	cache.Put("key1", 1)
	cache.Put("key2", 2)
	cache.Put("key3", 3)
	cache.Put("key4", 4)

	// Check size
	if cache.GetSize() > cache.GetCapacity() {
		t.Error("Cache size should not exceed capacity")
	}

	// Check most recently added items are present
	if !cache.Contains("key3") || !cache.Contains("key4") {
		t.Error("Most recently added items should be in cache")
	}

	// Check older items were evicted
	if cache.Contains("key1") || cache.Contains("key2") {
		t.Error("Older items should have been evicted")
	}
}
