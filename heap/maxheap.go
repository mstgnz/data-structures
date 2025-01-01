package heap

import "errors"

// MaxHeapImpl represents a binary heap data structure that maintains the max-heap property
type MaxHeapImpl struct {
	items []int
}

// NewMaxHeap creates and returns a new instance of MaxHeapImpl
func NewMaxHeap() *MaxHeapImpl {
	return &MaxHeapImpl{
		items: make([]int, 0),
	}
}

// Insert adds a new value to the heap while maintaining the max-heap property
func (h *MaxHeapImpl) Insert(value int) {
	h.items = append(h.items, value)
	h.heapifyUp(len(h.items) - 1)
}

// Extract removes and returns the maximum element from the heap
func (h *MaxHeapImpl) Extract() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("heap is empty")
	}

	max := h.items[0]
	lastIdx := len(h.items) - 1
	h.items[0] = h.items[lastIdx]
	h.items = h.items[:lastIdx]

	if !h.IsEmpty() {
		h.heapifyDown(0)
	}

	return max, nil
}

// Peek returns the maximum element without removing it
func (h *MaxHeapImpl) Peek() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("heap is empty")
	}
	return h.items[0], nil
}

// Size returns the number of elements in the heap
func (h *MaxHeapImpl) Size() int {
	return len(h.items)
}

// IsEmpty returns true if the heap contains no elements
func (h *MaxHeapImpl) IsEmpty() bool {
	return len(h.items) == 0
}

// heapifyUp maintains the max-heap property by moving a node up the tree
func (h *MaxHeapImpl) heapifyUp(index int) {
	for index > 0 {
		parentIdx := parent(index)
		if h.items[parentIdx] < h.items[index] {
			swap(h.items, parentIdx, index)
			index = parentIdx
		} else {
			break
		}
	}
}

// heapifyDown maintains the max-heap property by moving a node down the tree
func (h *MaxHeapImpl) heapifyDown(index int) {
	lastIdx := len(h.items) - 1
	for {
		largest := index
		leftIdx := leftChild(index)
		rightIdx := rightChild(index)

		if leftIdx <= lastIdx && h.items[leftIdx] > h.items[largest] {
			largest = leftIdx
		}
		if rightIdx <= lastIdx && h.items[rightIdx] > h.items[largest] {
			largest = rightIdx
		}

		if largest == index {
			break
		}

		swap(h.items, index, largest)
		index = largest
	}
}
