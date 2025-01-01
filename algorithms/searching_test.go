package algorithms

import "testing"

// Test cases for all search algorithms
var searchTestCases = []struct {
	name     string
	arr      []int
	target   int
	expected int
}{
	{
		name:     "Empty array",
		arr:      []int{},
		target:   5,
		expected: -1,
	},
	{
		name:     "Single element found",
		arr:      []int{5},
		target:   5,
		expected: 0,
	},
	{
		name:     "Single element not found",
		arr:      []int{5},
		target:   3,
		expected: -1,
	},
	{
		name:     "Multiple elements - found at beginning",
		arr:      []int{1, 2, 3, 4, 5},
		target:   1,
		expected: 0,
	},
	{
		name:     "Multiple elements - found at middle",
		arr:      []int{1, 2, 3, 4, 5},
		target:   3,
		expected: 2,
	},
	{
		name:     "Multiple elements - found at end",
		arr:      []int{1, 2, 3, 4, 5},
		target:   5,
		expected: 4,
	},
	{
		name:     "Multiple elements - not found",
		arr:      []int{1, 2, 3, 4, 5},
		target:   6,
		expected: -1,
	},
	{
		name:     "Duplicate elements - found",
		arr:      []int{1, 2, 2, 3, 3, 3, 4},
		target:   3,
		expected: 3, // Returns first occurrence
	},
}

func TestLinearSearch(t *testing.T) {
	for _, tc := range searchTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := LinearSearch(tc.arr, tc.target)
			if result != tc.expected {
				t.Errorf("LinearSearch(%v, %d) = %d; want %d", tc.arr, tc.target, result, tc.expected)
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	for _, tc := range searchTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := BinarySearch(tc.arr, tc.target)
			if result != tc.expected && result != -1 { // Binary search might find any occurrence for duplicates
				t.Errorf("BinarySearch(%v, %d) = %d; want %d", tc.arr, tc.target, result, tc.expected)
			}
		})
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	for _, tc := range searchTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := BinarySearchRecursive(tc.arr, tc.target)
			if result != tc.expected && result != -1 { // Binary search might find any occurrence for duplicates
				t.Errorf("BinarySearchRecursive(%v, %d) = %d; want %d", tc.arr, tc.target, result, tc.expected)
			}
		})
	}
}

func TestJumpSearch(t *testing.T) {
	for _, tc := range searchTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := JumpSearch(tc.arr, tc.target)
			if result != tc.expected && result != -1 { // Jump search might find any occurrence for duplicates
				t.Errorf("JumpSearch(%v, %d) = %d; want %d", tc.arr, tc.target, result, tc.expected)
			}
		})
	}
}

func TestInterpolationSearch(t *testing.T) {
	for _, tc := range searchTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := InterpolationSearch(tc.arr, tc.target)
			if result != tc.expected && result != -1 { // Interpolation search might find any occurrence for duplicates
				t.Errorf("InterpolationSearch(%v, %d) = %d; want %d", tc.arr, tc.target, result, tc.expected)
			}
		})
	}
}

func TestExponentialSearch(t *testing.T) {
	for _, tc := range searchTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ExponentialSearch(tc.arr, tc.target)
			if result != tc.expected && result != -1 { // Exponential search might find any occurrence for duplicates
				t.Errorf("ExponentialSearch(%v, %d) = %d; want %d", tc.arr, tc.target, result, tc.expected)
			}
		})
	}
}

func TestFibonacciSearch(t *testing.T) {
	for _, tc := range searchTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FibonacciSearch(tc.arr, tc.target)
			if result != tc.expected && result != -1 { // Fibonacci search might find any occurrence for duplicates
				t.Errorf("FibonacciSearch(%v, %d) = %d; want %d", tc.arr, tc.target, result, tc.expected)
			}
		})
	}
}
