package main

import (
	"fmt"

	"github.com/mstgnz/data-structures/heap"
)

// RunExamples demonstrates various heap implementations
func RunExamples() {
	// Example 1: Min Heap
	fmt.Println("Min Heap Example:")
	minHeap := heap.NewMinHeap()

	fmt.Println("Inserting elements:")
	elements := []int{10, 5, 15, 3, 8, 12, 20, 1, 7}
	for _, element := range elements {
		minHeap.Insert(element)
	}
	fmt.Printf("Min Heap: %v\n", minHeap)

	fmt.Println("\nExtracting minimum elements:")
	for i := 0; i < 3; i++ {
		min, _ := minHeap.Extract()
		fmt.Printf("Extracted min: %d\n", min)
	}
	fmt.Printf("Min Heap after extractions: %v\n\n", minHeap)

	// Example 2: Max Heap
	fmt.Println("Max Heap Example:")
	maxHeap := heap.NewMaxHeap()

	fmt.Println("Inserting elements:")
	for _, element := range elements {
		maxHeap.Insert(element)
	}
	fmt.Printf("Max Heap: %v\n", maxHeap)

	fmt.Println("\nExtracting maximum elements:")
	for i := 0; i < 3; i++ {
		max, _ := maxHeap.Extract()
		fmt.Printf("Extracted max: %d\n", max)
	}
	fmt.Printf("Max Heap after extractions: %v\n\n", maxHeap)

	// Example 3: Binomial Heap
	fmt.Println("Binomial Heap Example:")
	binHeap := heap.NewBinomialHeap()

	fmt.Println("Inserting elements:")
	for _, element := range elements {
		binHeap.Insert(element)
	}

	fmt.Println("\nExtracting minimum elements:")
	for i := 0; i < 3; i++ {
		if min, err := binHeap.Extract(); err == nil {
			fmt.Printf("Extracted min: %d\n", min)
		}
	}
}
