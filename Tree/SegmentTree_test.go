package Tree

import (
	"testing"
)

func TestSegmentTree(t *testing.T) {
	t.Run("Sum Segment Tree", func(t *testing.T) {
		arr := []int{1, 3, 5, 7, 9, 11}
		st := NewSegmentTree(arr, SumCombine)

		// Test range sum queries
		testCases := []struct {
			left     int
			right    int
			expected int
		}{
			{0, 2, 9},   // 1 + 3 + 5
			{1, 4, 24},  // 3 + 5 + 7 + 9
			{0, 5, 36},  // sum of all elements
			{3, 5, 27},  // 7 + 9 + 11
			{2, 2, 5},   // single element
			{-1, 5, 36}, // invalid left bound
			{0, 6, 36},  // invalid right bound
			{4, 2, 0},   // invalid range
		}

		for _, tc := range testCases {
			if got := st.Query(tc.left, tc.right); got != tc.expected {
				t.Errorf("Sum query [%d, %d] = %d; want %d", tc.left, tc.right, got, tc.expected)
			}
		}

		// Test update
		st.Update(2, 6) // Change 5 to 6
		if got := st.Query(0, 2); got != 10 {
			t.Errorf("After update, sum query [0, 2] = %d; want 10", got)
		}
	})

	t.Run("Min Segment Tree", func(t *testing.T) {
		arr := []int{5, 2, 8, 1, 9, 3}
		st := NewSegmentTree(arr, MinCombine)

		// Test range minimum queries
		testCases := []struct {
			left     int
			right    int
			expected int
		}{
			{0, 2, 2}, // min(5, 2, 8)
			{1, 4, 1}, // min(2, 8, 1, 9)
			{0, 5, 1}, // min of all elements
			{3, 5, 1}, // min(1, 9, 3)
			{2, 2, 8}, // single element
		}

		for _, tc := range testCases {
			if got := st.Query(tc.left, tc.right); got != tc.expected {
				t.Errorf("Min query [%d, %d] = %d; want %d", tc.left, tc.right, got, tc.expected)
			}
		}

		// Test update
		st.Update(3, 7) // Change 1 to 7
		if got := st.Query(0, 5); got != 2 {
			t.Errorf("After update, min query [0, 5] = %d; want 2", got)
		}
	})

	t.Run("Max Segment Tree", func(t *testing.T) {
		arr := []int{5, 2, 8, 1, 9, 3}
		st := NewSegmentTree(arr, MaxCombine)

		// Test range maximum queries
		testCases := []struct {
			left     int
			right    int
			expected int
		}{
			{0, 2, 8}, // max(5, 2, 8)
			{1, 4, 9}, // max(2, 8, 1, 9)
			{0, 5, 9}, // max of all elements
			{3, 5, 9}, // max(1, 9, 3)
			{2, 2, 8}, // single element
		}

		for _, tc := range testCases {
			if got := st.Query(tc.left, tc.right); got != tc.expected {
				t.Errorf("Max query [%d, %d] = %d; want %d", tc.left, tc.right, got, tc.expected)
			}
		}

		// Test update
		st.Update(4, 4) // Change 9 to 4
		if got := st.Query(0, 5); got != 8 {
			t.Errorf("After update, max query [0, 5] = %d; want 8", got)
		}
	})

	t.Run("Empty Array", func(t *testing.T) {
		arr := []int{}
		st := NewSegmentTree(arr, SumCombine)

		if got := st.Query(0, 0); got != 0 {
			t.Errorf("Query on empty tree = %d; want 0", got)
		}
	})

	t.Run("Single Element", func(t *testing.T) {
		arr := []int{5}
		st := NewSegmentTree(arr, SumCombine)

		if got := st.Query(0, 0); got != 5 {
			t.Errorf("Query on single element = %d; want 5", got)
		}

		st.Update(0, 10)
		if got := st.Query(0, 0); got != 10 {
			t.Errorf("After update, query = %d; want 10", got)
		}
	})
}
