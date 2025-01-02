package algorithms

// RecursiveFactorial calculates n! recursively
func RecursiveFactorial(n int) int {
	if n < 0 {
		panic("Factorial is not defined for negative numbers")
	}
	if n <= 1 {
		return 1
	}
	return n * RecursiveFactorial(n-1)
}

// RecursiveReverseString reverses a string using recursion
func RecursiveReverseString(s string) string {
	if len(s) <= 1 {
		return s
	}
	return RecursiveReverseString(s[1:]) + s[:1]
}

// RecursiveNQueens solves the N-Queens problem and returns all solutions
// Each solution is represented as a slice of integers where the index represents the row
// and the value represents the column where a queen is placed
func RecursiveNQueens(n int) [][]int {
	if n < 1 {
		return nil
	}

	solutions := make([][]int, 0)
	board := make([]int, n)
	recursiveSolveNQueens(n, 0, board, &solutions)
	return solutions
}

// recursiveSolveNQueens is a helper function that recursively solves the N-Queens problem
func recursiveSolveNQueens(n, row int, board []int, solutions *[][]int) {
	if row == n {
		// Found a solution, make a copy and add it to solutions
		solution := make([]int, n)
		copy(solution, board)
		*solutions = append(*solutions, solution)
		return
	}

	// Try placing a queen in each column of the current row
	for col := 0; col < n; col++ {
		if recursiveIsSafe(board, row, col) {
			board[row] = col
			recursiveSolveNQueens(n, row+1, board, solutions)
		}
	}
}

// recursiveIsSafe checks if it's safe to place a queen at the given position
func recursiveIsSafe(board []int, row, col int) bool {
	// Check previous rows
	for i := 0; i < row; i++ {
		// Check vertical and diagonal attacks
		if board[i] == col ||
			board[i]-i == col-row ||
			board[i]+i == col+row {
			return false
		}
	}
	return true
}

// RecursiveGetNQueensBoard converts a solution to a 2D board representation
// Returns a slice of strings where 'Q' represents a queen and '.' represents an empty cell
func RecursiveGetNQueensBoard(solution []int) []string {
	n := len(solution)
	board := make([]string, n)

	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			if solution[i] == j {
				row[j] = 'Q'
			} else {
				row[j] = '.'
			}
		}
		board[i] = string(row)
	}

	return board
}

// RecursiveCountNQueensSolutions returns the number of solutions for the N-Queens problem
func RecursiveCountNQueensSolutions(n int) int {
	if n < 1 {
		return 0
	}

	count := 0
	board := make([]int, n)
	recursiveCountNQueensSolutions(n, 0, board, &count)
	return count
}

// recursiveCountNQueensSolutions is a helper function that counts solutions recursively
func recursiveCountNQueensSolutions(n, row int, board []int, count *int) {
	if row == n {
		*count++
		return
	}

	for col := 0; col < n; col++ {
		if recursiveIsSafe(board, row, col) {
			board[row] = col
			recursiveCountNQueensSolutions(n, row+1, board, count)
		}
	}
}

// RecursiveIsValidNQueensSolution verifies if a given solution is valid
func RecursiveIsValidNQueensSolution(solution []int) bool {
	n := len(solution)

	// Check if each value is within bounds
	for i := 0; i < n; i++ {
		if solution[i] < 0 || solution[i] >= n {
			return false
		}
	}

	// Check for conflicts
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// Check vertical and diagonal attacks
			if solution[i] == solution[j] ||
				solution[i]-i == solution[j]-j ||
				solution[i]+i == solution[j]+j {
				return false
			}
		}
	}

	return true
}
