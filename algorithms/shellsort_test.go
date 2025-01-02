package algorithms

import (
	"reflect"
	"testing"
)

func TestShellSort(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of input to avoid modifying the test case
			input := make([]int, len(tt.input))
			copy(input, tt.input)

			result := ShellSort(input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ShellSort() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestShellSortWithGaps(t *testing.T) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Test with different gap sequences
	gapSequences := []struct {
		name string
		gaps []int
	}{
		{"Hibbard", HibbardGaps(len(arr))},
		{"Sedgewick", SedgewickGaps(len(arr))},
		{"Pratt", PrattGaps(len(arr))},
	}

	for _, seq := range gapSequences {
		t.Run(seq.name, func(t *testing.T) {
			// Make a copy of input
			input := make([]int, len(arr))
			copy(input, arr)

			result := ShellSortWithGaps(input, seq.gaps)

			if !reflect.DeepEqual(result, expected) {
				t.Errorf("ShellSortWithGaps() with %s sequence = %v, want %v",
					seq.name, result, expected)
			}
		})
	}
}

func TestGapSequences(t *testing.T) {
	n := 100

	t.Run("Hibbard", func(t *testing.T) {
		gaps := HibbardGaps(n)
		if len(gaps) == 0 {
			t.Error("HibbardGaps returned empty sequence")
		}
		// Check if sequence is in descending order
		for i := 1; i < len(gaps); i++ {
			if gaps[i-1] <= gaps[i] {
				t.Errorf("HibbardGaps not in descending order at index %d", i)
			}
		}
	})

	t.Run("Sedgewick", func(t *testing.T) {
		gaps := SedgewickGaps(n)
		if len(gaps) == 0 {
			t.Error("SedgewickGaps returned empty sequence")
		}
		// Check if sequence is in descending order
		for i := 1; i < len(gaps); i++ {
			if gaps[i-1] <= gaps[i] {
				t.Errorf("SedgewickGaps not in descending order at index %d", i)
			}
		}
	})

	t.Run("Pratt", func(t *testing.T) {
		gaps := PrattGaps(n)
		if len(gaps) == 0 {
			t.Error("PrattGaps returned empty sequence")
		}
		// Check if sequence is in descending order
		for i := 1; i < len(gaps); i++ {
			if gaps[i-1] <= gaps[i] {
				t.Errorf("PrattGaps not in descending order at index %d", i)
			}
		}
	})
}

func TestPow3(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 3},
		{2, 9},
		{3, 27},
		{4, 81},
	}

	for _, tt := range tests {
		result := pow3(tt.input)
		if result != tt.expected {
			t.Errorf("pow3(%d) = %d, want %d", tt.input, result, tt.expected)
		}
	}
}
