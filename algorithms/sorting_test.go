package algorithms

import (
	"reflect"
	"testing"
)

// Test cases for all sorting algorithms
var testCases = []struct {
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
		name:     "Already sorted",
		input:    []int{1, 2, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "Reverse sorted",
		input:    []int{5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "Random order",
		input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
		expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
	},
	{
		name:     "Duplicate elements",
		input:    []int{3, 3, 3, 1, 1, 2, 2},
		expected: []int{1, 1, 2, 2, 3, 3, 3},
	},
}

func TestQuickSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := make([]int, len(tc.input))
			copy(input, tc.input)
			result := QuickSort(input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("QuickSort(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := make([]int, len(tc.input))
			copy(input, tc.input)
			result := MergeSort(input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("MergeSort(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := make([]int, len(tc.input))
			copy(input, tc.input)
			result := HeapSort(input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("HeapSort(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := make([]int, len(tc.input))
			copy(input, tc.input)
			result := InsertionSort(input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("InsertionSort(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestBubbleSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := make([]int, len(tc.input))
			copy(input, tc.input)
			result := BubbleSort(input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("BubbleSort(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestCountingSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := make([]int, len(tc.input))
			copy(input, tc.input)
			result := CountingSort(input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("CountingSort(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
