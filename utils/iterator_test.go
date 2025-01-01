package utils

import (
	"sort"
	"testing"
)

func TestSliceIterator(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	it := NewSliceIterator(data)

	// Test iteration
	expected := make([]int, len(data))
	copy(expected, data)

	var result []int
	for it.HasNext() {
		result = append(result, it.Next())
	}

	if !slicesEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test reset
	it.Reset()
	result = nil
	for it.HasNext() {
		result = append(result, it.Next())
	}

	if !slicesEqual(result, expected) {
		t.Errorf("After reset: Expected %v, got %v", expected, result)
	}
}

func TestMapIterator(t *testing.T) {
	data := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	// Test with ordered keys
	keyOrder := func(keys []string) {
		sort.Strings(keys)
	}

	it := NewMapIterator(data, keyOrder)

	expectedKeys := []string{"one", "three", "two"}
	expectedValues := []int{1, 3, 2}

	var keys []string
	var values []int

	for it.HasNext() {
		k, v := it.Next()
		keys = append(keys, k)
		values = append(values, v)
	}

	if !slicesEqual(keys, expectedKeys) {
		t.Errorf("Expected keys %v, got %v", expectedKeys, keys)
	}
	if !slicesEqual(values, expectedValues) {
		t.Errorf("Expected values %v, got %v", expectedValues, values)
	}

	// Test reset
	it.Reset()
	keys = nil
	values = nil

	for it.HasNext() {
		k, v := it.Next()
		keys = append(keys, k)
		values = append(values, v)
	}

	if !slicesEqual(keys, expectedKeys) {
		t.Errorf("After reset: Expected keys %v, got %v", expectedKeys, keys)
	}
	if !slicesEqual(values, expectedValues) {
		t.Errorf("After reset: Expected values %v, got %v", expectedValues, values)
	}
}

func TestLinkedListIterator(t *testing.T) {
	// Create a linked list: 1 -> 2 -> 3
	head := &Node[int]{
		Value: 1,
		Next: &Node[int]{
			Value: 2,
			Next: &Node[int]{
				Value: 3,
			},
		},
	}

	it := NewLinkedListIterator(head)
	expected := []int{1, 2, 3}

	var result []int
	for it.HasNext() {
		result = append(result, it.Next())
	}

	if !slicesEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test reset
	it.Reset()
	result = nil
	for it.HasNext() {
		result = append(result, it.Next())
	}

	if !slicesEqual(result, expected) {
		t.Errorf("After reset: Expected %v, got %v", expected, result)
	}
}

func TestTreeIterator(t *testing.T) {
	// Create a binary tree:
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	root := &TreeNode[int]{
		Value: 1,
		Left: &TreeNode[int]{
			Value: 2,
			Left: &TreeNode[int]{
				Value: 4,
			},
			Right: &TreeNode[int]{
				Value: 5,
			},
		},
		Right: &TreeNode[int]{
			Value: 3,
		},
	}

	testCases := []struct {
		order    TraversalOrder
		expected []int
	}{
		{PreOrder, []int{1, 2, 4, 5, 3}},
		{InOrder, []int{4, 2, 5, 1, 3}},
		{PostOrder, []int{4, 5, 2, 3, 1}},
		{LevelOrder, []int{1, 2, 3, 4, 5}},
	}

	for _, tc := range testCases {
		it := NewTreeIterator(root, tc.order)
		var result []int

		for it.HasNext() {
			result = append(result, it.Next())
		}

		if !slicesEqual(result, tc.expected) {
			t.Errorf("%v traversal: Expected %v, got %v", tc.order, tc.expected, result)
		}

		// Test reset
		it.Reset()
		result = nil
		for it.HasNext() {
			result = append(result, it.Next())
		}

		if !slicesEqual(result, tc.expected) {
			t.Errorf("%v traversal after reset: Expected %v, got %v", tc.order, tc.expected, result)
		}
	}
}

// Helper function to compare slices
func slicesEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
