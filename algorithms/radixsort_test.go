package algorithms

import (
	"reflect"
	"testing"
)

func TestRadixSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{170},
			expected: []int{170},
		},
		{
			name:     "Multiple elements",
			input:    []int{170, 45, 75, 90, 802, 24, 2, 66},
			expected: []int{2, 24, 45, 66, 75, 90, 170, 802},
		},
		{
			name:     "Already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.input))
			copy(input, tt.input)

			result := RadixSort(input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("RadixSort() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestRadixSortString(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Empty array",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "Single string",
			input:    []string{"abc"},
			expected: []string{"abc"},
		},
		{
			name:     "Multiple strings",
			input:    []string{"cat", "dog", "bat", "ant"},
			expected: []string{"ant", "bat", "cat", "dog"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]string, len(tt.input))
			copy(input, tt.input)

			result := RadixSortString(input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("RadixSortString() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestRadixSortBytes(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]byte
		expected [][]byte
	}{
		{
			name:     "Empty array",
			input:    [][]byte{},
			expected: [][]byte{},
		},
		{
			name:     "Single byte array",
			input:    [][]byte{{1, 2, 3}},
			expected: [][]byte{{1, 2, 3}},
		},
		{
			name: "Multiple byte arrays",
			input: [][]byte{
				{3, 2, 1},
				{1, 2, 3},
				{2, 1, 3},
			},
			expected: [][]byte{
				{1, 2, 3},
				{2, 1, 3},
				{3, 2, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Deep copy input
			input := make([][]byte, len(tt.input))
			for i, arr := range tt.input {
				input[i] = make([]byte, len(arr))
				copy(input[i], arr)
			}

			result := RadixSortBytes(input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("RadixSortBytes() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestRadixSortStringPanic(t *testing.T) {
	// Test panic with strings of different lengths
	defer func() {
		if r := recover(); r == nil {
			t.Error("RadixSortString did not panic with strings of different lengths")
		}
	}()

	RadixSortString([]string{"abc", "de"})
}

func TestRadixSortBytesPanic(t *testing.T) {
	// Test panic with byte slices of different lengths
	defer func() {
		if r := recover(); r == nil {
			t.Error("RadixSortBytes did not panic with byte slices of different lengths")
		}
	}()

	RadixSortBytes([][]byte{{1, 2, 3}, {1, 2}})
}

func TestGetMax(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Single element",
			input:    []int{5},
			expected: 5,
		},
		{
			name:     "Multiple elements",
			input:    []int{1, 5, 3, 9, 2},
			expected: 9,
		},
		{
			name:     "All same elements",
			input:    []int{4, 4, 4, 4},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getMax(tt.input)
			if result != tt.expected {
				t.Errorf("getMax() = %v, want %v", result, tt.expected)
			}
		})
	}
}
