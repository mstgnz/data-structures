package main

import (
	"fmt"
)

func main() {
	solution := []int{0, 2, 1, 3}
	fmt.Printf("Checking solution %v\n", solution)

	// Manuel kontrol
	n := len(solution)
	valid := true

	// Sütun çakışmalarını kontrol et
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if solution[i] == solution[j] {
				fmt.Printf("Column conflict: queens at (%d,%d) and (%d,%d)\n", i, solution[i], j, solution[j])
				valid = false
			}
		}
	}

	// Çapraz çakışmaları kontrol et
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			rowDiff := j - i
			colDiff := abs(solution[j] - solution[i])

			if rowDiff == colDiff {
				fmt.Printf("Diagonal conflict: queens at (%d,%d) and (%d,%d)\n", i, solution[i], j, solution[j])
				fmt.Printf("rowDiff = %d, colDiff = %d\n", rowDiff, colDiff)
				valid = false
			}
		}
	}

	// Çapraz çakışmaları farklı şekilde kontrol et
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			rowDiff := abs(j - i)
			colDiff := abs(solution[j] - solution[i])

			if rowDiff == colDiff {
				fmt.Printf("Diagonal conflict (abs): queens at (%d,%d) and (%d,%d)\n", i, solution[i], j, solution[j])
				fmt.Printf("abs rowDiff = %d, abs colDiff = %d\n", rowDiff, colDiff)
				valid = false
			}
		}
	}

	fmt.Printf("Solution is valid: %v\n", valid)

	// Görsel olarak tahtayı göster
	printBoard(solution)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func printBoard(solution []int) {
	n := len(solution)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if solution[i] == j {
				fmt.Print("Q ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}
