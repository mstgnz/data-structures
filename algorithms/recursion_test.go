package algorithms

import (
	"reflect"
	"testing"
)

func TestRecursiveFactorial(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{7, 5040},
	}

	for _, tt := range tests {
		result := RecursiveFactorial(tt.n)
		if result != tt.expected {
			t.Errorf("RecursiveFactorial(%d) = %d, want %d", tt.n, result, tt.expected)
		}
	}
}

func TestRecursiveFactorialPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("RecursiveFactorial did not panic with negative input")
		}
	}()

	RecursiveFactorial(-1)
}

func TestRecursiveReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"hello", "olleh"},
		{"12345", "54321"},
		{"Go Programming", "gnimmargorP oG"},
	}

	for _, tt := range tests {
		result := RecursiveReverseString(tt.input)
		if result != tt.expected {
			t.Errorf("RecursiveReverseString(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestRecursiveNQueens(t *testing.T) {
	tests := []struct {
		n            int
		solutionSize int // Number of solutions
	}{
		{1, 1},
		{2, 0},
		{3, 0},
		{4, 2},
		{5, 10},
		{6, 4},
		{7, 40},
		{8, 92},
	}

	for _, tt := range tests {
		solutions := RecursiveNQueens(tt.n)
		if len(solutions) != tt.solutionSize {
			t.Errorf("RecursiveNQueens(%d) returned %d solutions, want %d",
				tt.n, len(solutions), tt.solutionSize)
		}

		// Verify each solution
		for _, solution := range solutions {
			if !RecursiveIsValidNQueensSolution(solution) {
				t.Errorf("Invalid solution found for n=%d: %v", tt.n, solution)
			}
		}
	}
}

func TestRecursiveGetNQueensBoard(t *testing.T) {
	solution := []int{1, 3, 0, 2} // A valid solution for 4-queens
	expected := []string{
		".Q..",
		"...Q",
		"Q...",
		"..Q.",
	}

	board := RecursiveGetNQueensBoard(solution)
	if !reflect.DeepEqual(board, expected) {
		t.Errorf("RecursiveGetNQueensBoard() = %v, want %v", board, expected)
	}
}

func TestRecursiveCountNQueensSolutions(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 0},
		{3, 0},
		{4, 2},
		{5, 10},
		{6, 4},
		{7, 40},
		{8, 92},
	}

	for _, tt := range tests {
		count := RecursiveCountNQueensSolutions(tt.n)
		if count != tt.expected {
			t.Errorf("RecursiveCountNQueensSolutions(%d) = %d, want %d",
				tt.n, count, tt.expected)
		}
	}
}

func TestRecursiveIsValidNQueensSolution(t *testing.T) {
	tests := []struct {
		solution []int
		valid    bool
	}{
		{[]int{1, 3, 0, 2}, true},  // Valid 4-queens solution
		{[]int{0, 0, 0, 0}, false}, // Invalid: queens in same column
		{[]int{0, 1, 2, 3}, false}, // Invalid: queens in diagonal
		{[]int{3, 2, 1, 0}, false}, // Invalid: queens in diagonal
		{[]int{0, 2, 1, 3}, false}, // Invalid: queens in diagonal
	}

	for _, tt := range tests {
		result := RecursiveIsValidNQueensSolution(tt.solution)
		if result != tt.valid {
			t.Errorf("RecursiveIsValidNQueensSolution(%v) = %v, want %v",
				tt.solution, result, tt.valid)
		}
	}
}

func TestRecursiveNQueensInvalidInput(t *testing.T) {
	tests := []struct {
		n int
	}{
		{-1},
		{0},
	}

	for _, tt := range tests {
		solutions := RecursiveNQueens(tt.n)
		if solutions != nil {
			t.Errorf("RecursiveNQueens(%d) = %v, want nil", tt.n, solutions)
		}
	}
}
