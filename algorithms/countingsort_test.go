package algorithms

import (
	"reflect"
	"testing"
)

func TestCountingSort(t *testing.T) {
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
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Multiple elements",
			input:    []int{4, 2, 2, 8, 3, 3, 1},
			expected: []int{1, 2, 2, 3, 3, 4, 8},
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

			result := CountingSort(input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CountingSort() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCountingSortWithRange(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		min      int
		max      int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			min:      0,
			max:      0,
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{5},
			min:      0,
			max:      10,
			expected: []int{5},
		},
		{
			name:     "Multiple elements",
			input:    []int{4, 2, 2, 8, 3, 3, 1},
			min:      1,
			max:      8,
			expected: []int{1, 2, 2, 3, 3, 4, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.input))
			copy(input, tt.input)

			result := CountingSortWithRange(input, tt.min, tt.max)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CountingSortWithRange() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCountingSortString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "Multiple characters",
			input:    "datastructures",
			expected: "aacderrssttu",
		},
		{
			name:     "With spaces",
			input:    "counting sort",
			expected: " cginnorstu",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountingSortString(tt.input)
			if result != tt.expected {
				t.Errorf("CountingSortString() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCountingSortBytes(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "Empty array",
			input:    []byte{},
			expected: []byte{},
		},
		{
			name:     "Single byte",
			input:    []byte{65},
			expected: []byte{65},
		},
		{
			name:     "Multiple bytes",
			input:    []byte{70, 65, 67, 66, 69, 68},
			expected: []byte{65, 66, 67, 68, 69, 70},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]byte, len(tt.input))
			copy(input, tt.input)

			result := CountingSortBytes(input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CountingSortBytes() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCountingSortPanic(t *testing.T) {
	// Test panic with negative numbers
	defer func() {
		if r := recover(); r == nil {
			t.Error("CountingSort did not panic with negative numbers")
		}
	}()

	CountingSort([]int{1, -2, 3})
}

func TestCountingSortWithRangePanic(t *testing.T) {
	// Test panic with out of range numbers
	defer func() {
		if r := recover(); r == nil {
			t.Error("CountingSortWithRange did not panic with out of range numbers")
		}
	}()

	CountingSortWithRange([]int{1, 2, 10}, 1, 5)
}
