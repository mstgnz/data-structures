package tree

import (
	"reflect"
	"sort"
	"testing"
)

func TestTernarySearchTreeBasicOperations(t *testing.T) {
	tst := NewTernarySearchTree()

	// Test initial state
	if !tst.IsEmpty() {
		t.Error("New tree should be empty")
	}

	// Test Insert and Search
	tst.Insert("test", 1)
	if value, found := tst.Search("test"); !found || value.(int) != 1 {
		t.Error("Failed to find inserted value")
	}

	// Test overwrite
	tst.Insert("test", 2)
	if value, found := tst.Search("test"); !found || value.(int) != 2 {
		t.Error("Failed to overwrite value")
	}

	// Test non-existent key
	if _, found := tst.Search("nonexistent"); found {
		t.Error("Found non-existent key")
	}
}

func TestTernarySearchTreePrefixOperations(t *testing.T) {
	tst := NewTernarySearchTree()

	// Insert strings with common prefixes
	words := []string{"test", "testing", "tester", "team"}
	for i, word := range words {
		tst.Insert(word, i)
	}

	// Test each word
	for i, word := range words {
		if value, found := tst.Search(word); !found || value.(int) != i {
			t.Errorf("Failed to find correct value for %s", word)
		}
	}

	// Test size
	if tst.Size() != len(words) {
		t.Errorf("Expected size %d, got %d", len(words), tst.Size())
	}
}

func TestTernarySearchTreeDelete(t *testing.T) {
	tst := NewTernarySearchTree()

	// Insert some values
	tst.Insert("test", 1)
	tst.Insert("testing", 2)
	tst.Insert("team", 3)

	// Delete existing key
	if !tst.Delete("test") {
		t.Error("Delete returned false for existing key")
	}

	// Verify deletion
	if _, found := tst.Search("test"); found {
		t.Error("Found deleted key")
	}

	// Delete non-existent key
	if tst.Delete("nonexistent") {
		t.Error("Delete returned true for non-existent key")
	}

	// Verify other keys still exist
	if _, found := tst.Search("testing"); !found {
		t.Error("Lost existing key after deletion")
	}
}

func TestTernarySearchTreeClear(t *testing.T) {
	tst := NewTernarySearchTree()

	// Insert some values
	tst.Insert("test", 1)
	tst.Insert("testing", 2)
	tst.Insert("team", 3)

	// Clear the tree
	tst.Clear()

	// Verify tree is empty
	if !tst.IsEmpty() {
		t.Error("Tree should be empty after Clear()")
	}

	// Verify size is 0
	if tst.Size() != 0 {
		t.Errorf("Expected size 0 after Clear(), got %d", tst.Size())
	}

	// Verify no keys exist
	if _, found := tst.Search("test"); found {
		t.Error("Found key after Clear()")
	}
}

func TestTernarySearchTreeStartsWith(t *testing.T) {
	tst := NewTernarySearchTree()

	// Insert some values
	words := []string{"test", "testing", "tester", "team", "toast"}
	for _, word := range words {
		tst.Insert(word, true)
	}

	// Test StartsWith
	testCases := []struct {
		prefix   string
		expected []string
	}{
		{"te", []string{"team", "test", "tester", "testing"}},
		{"test", []string{"test", "tester", "testing"}},
		{"tea", []string{"team"}},
		{"toast", []string{"toast"}},
		{"z", []string{}},
	}

	for _, tc := range testCases {
		result := tst.StartsWith(tc.prefix)
		sort.Strings(result)
		sort.Strings(tc.expected)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("StartsWith(%q) = %v, want %v", tc.prefix, result, tc.expected)
		}
	}
}

func TestTernarySearchTreeKeys(t *testing.T) {
	tst := NewTernarySearchTree()

	// Insert some values
	expected := []string{"team", "test", "testing", "tester", "toast"}
	for _, word := range expected {
		tst.Insert(word, true)
	}

	// Get all keys
	keys := tst.Keys()
	sort.Strings(keys)
	sort.Strings(expected)

	// Compare keys
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("Keys() = %v, want %v", keys, expected)
	}
}

func TestTernarySearchTreeEmptyString(t *testing.T) {
	tst := NewTernarySearchTree()

	// Test empty string
	tst.Insert("", 1)
	if _, found := tst.Search(""); found {
		t.Error("Empty string should not be inserted")
	}

	// Test empty prefix
	result := tst.StartsWith("")
	if result != nil {
		t.Error("StartsWith empty string should return nil")
	}
}

func TestTernarySearchTreeCaseSensitive(t *testing.T) {
	tst := NewTernarySearchTree()

	// Insert mixed case strings
	tst.Insert("Test", 1)
	tst.Insert("test", 2)
	tst.Insert("TEST", 3)

	// Verify case sensitivity
	testCases := []struct {
		key      string
		expected int
	}{
		{"Test", 1},
		{"test", 2},
		{"TEST", 3},
	}

	for _, tc := range testCases {
		if value, found := tst.Search(tc.key); !found || value.(int) != tc.expected {
			t.Errorf("Failed to find correct value for %s", tc.key)
		}
	}
}
