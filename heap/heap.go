// Package heap implements various heap data structures
package heap

// IHeap interface defines the basic operations that all heap implementations must provide
type IHeap interface {
	// Insert adds a new element to the heap
	Insert(value int)

	// Extract removes and returns the root element (min for MinHeap, max for MaxHeap)
	Extract() (int, error)

	// Peek returns the root element without removing it
	Peek() (int, error)

	// Size returns the number of elements in the heap
	Size() int

	// IsEmpty returns true if the heap has no elements
	IsEmpty() bool
}

// HeapType represents the type of heap (MinHeap or MaxHeap)
type HeapType int

const (
	// MinHeap represents a min heap where the root is the minimum element
	MinHeap HeapType = iota
	// MaxHeap represents a max heap where the root is the maximum element
	MaxHeap
)

// Helper functions for heap operations

// swap exchanges two elements in a slice
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// parent returns the index of the parent node
func parent(i int) int {
	return (i - 1) / 2
}

// leftChild returns the index of the left child
func leftChild(i int) int {
	return 2*i + 1
}

// rightChild returns the index of the right child
func rightChild(i int) int {
	return 2*i + 2
}
