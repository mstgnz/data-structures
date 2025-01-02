package tree

import (
	"reflect"
	"sort"
	"testing"
)

func TestRadixTreeBasicOperations(t *testing.T) {
	rt := NewRadixTree()

	// Test initial state
	if !rt.IsEmpty() {
		t.Error("New tree should be empty")
	}

	// Test Insert and Search
	rt.Insert("test", 1)
	if value, found := rt.Search("test"); !found || value.(int) != 1 {
		t.Error("Failed to find inserted value")
	}

	// Test overwrite
	rt.Insert("test", 2)
	if value, found := rt.Search("test"); !found || value.(int) != 2 {
		t.Error("Failed to overwrite value")
	}

	// Test non-existent key
	if _, found := rt.Search("nonexistent"); found {
		t.Error("Found non-existent key")
	}
}

func TestRadixTreePrefixOperations(t *testing.T) {
	rt := NewRadixTree()

	// Insert strings with common prefixes
	words := []string{"test", "testing", "tester", "team"}
	for i, word := range words {
		rt.Insert(word, i)
	}

	// Test each word
	for i, word := range words {
		if value, found := rt.Search(word); !found || value.(int) != i {
			t.Errorf("Failed to find correct value for %s", word)
		}
	}

	// Test size
	if rt.Size() != len(words) {
		t.Errorf("Expected size %d, got %d", len(words), rt.Size())
	}
}

func TestRadixTreeDelete(t *testing.T) {
	rt := NewRadixTree()

	// Insert some values
	rt.Insert("test", 1)
	rt.Insert("testing", 2)
	rt.Insert("team", 3)

	// Delete existing key
	if !rt.Delete("test") {
		t.Error("Delete returned false for existing key")
	}

	// Verify deletion
	if _, found := rt.Search("test"); found {
		t.Error("Found deleted key")
	}

	// Delete non-existent key
	if rt.Delete("nonexistent") {
		t.Error("Delete returned true for non-existent key")
	}

	// Verify other keys still exist
	if _, found := rt.Search("testing"); !found {
		t.Error("Lost existing key after deletion")
	}
}

func TestRadixTreeClear(t *testing.T) {
	rt := NewRadixTree()

	// Insert some values
	rt.Insert("test", 1)
	rt.Insert("testing", 2)
	rt.Insert("team", 3)

	// Clear the tree
	rt.Clear()

	// Verify tree is empty
	if !rt.IsEmpty() {
		t.Error("Tree should be empty after Clear()")
	}

	// Verify size is 0
	if rt.Size() != 0 {
		t.Errorf("Expected size 0 after Clear(), got %d", rt.Size())
	}

	// Verify no keys exist
	if _, found := rt.Search("test"); found {
		t.Error("Found key after Clear()")
	}
}

func TestRadixTreeKeys(t *testing.T) {
	rt := NewRadixTree()

	// Insert some values
	expected := []string{"team", "test", "testing", "tester"}
	for _, word := range expected {
		rt.Insert(word, true)
	}

	// Get all keys
	keys := rt.Keys()
	sort.Strings(keys)
	sort.Strings(expected)

	// Compare keys
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("Keys() = %v, want %v", keys, expected)
	}
}

func TestRadixTreeLongPrefix(t *testing.T) {
	rt := NewRadixTree()

	// Insert strings with long common prefixes
	rt.Insert("aaaaaa", 1)
	rt.Insert("aaaaab", 2)
	rt.Insert("aaaabb", 3)

	// Test each string
	testCases := []struct {
		key   string
		value int
	}{
		{"aaaaaa", 1},
		{"aaaaab", 2},
		{"aaaabb", 3},
	}

	for _, tc := range testCases {
		if value, found := rt.Search(tc.key); !found || value.(int) != tc.value {
			t.Errorf("Failed to find correct value for %s", tc.key)
		}
	}
}

func TestRadixTreeEmptyString(t *testing.T) {
	rt := NewRadixTree()

	// Test empty string
	rt.Insert("", 1)
	if value, found := rt.Search(""); !found || value.(int) != 1 {
		t.Error("Failed to handle empty string")
	}

	// Delete empty string
	if !rt.Delete("") {
		t.Error("Failed to delete empty string")
	}

	// Verify empty string is gone
	if _, found := rt.Search(""); found {
		t.Error("Found empty string after deletion")
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		a, b     string
		expected string
	}{
		{"", "", ""},
		{"a", "", ""},
		{"", "a", ""},
		{"a", "a", "a"},
		{"abc", "abd", "ab"},
		{"hello", "help", "hel"},
		{"test", "testing", "test"},
		{"xyz", "abc", ""},
	}

	for _, tt := range tests {
		result := longestCommonPrefix(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("longestCommonPrefix(%q, %q) = %q, want %q",
				tt.a, tt.b, result, tt.expected)
		}
	}
}
