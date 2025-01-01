package heap

import (
	"errors"
	"sync"
)

// MinHeapImpl represents a binary heap data structure that maintains the min-heap property
type MinHeapImpl struct {
	items []int
	mutex sync.RWMutex
}

// NewMinHeap creates and returns a new instance of MinHeapImpl
func NewMinHeap() *MinHeapImpl {
	return &MinHeapImpl{
		items: make([]int, 0),
		mutex: sync.RWMutex{},
	}
}

// Insert adds a new value to the heap while maintaining the min-heap property
func (h *MinHeapImpl) Insert(value int) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	h.items = append(h.items, value)
	h.heapifyUp(len(h.items) - 1)
}

// Extract removes and returns the minimum element from the heap
func (h *MinHeapImpl) Extract() (int, error) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	if h.IsEmpty() {
		return 0, errors.New("heap is empty")
	}

	min := h.items[0]
	lastIdx := len(h.items) - 1
	h.items[0] = h.items[lastIdx]
	h.items = h.items[:lastIdx]

	if !h.IsEmpty() {
		h.heapifyDown(0)
	}

	return min, nil
}

// Peek returns the minimum element without removing it
func (h *MinHeapImpl) Peek() (int, error) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	if h.IsEmpty() {
		return 0, errors.New("heap is empty")
	}
	return h.items[0], nil
}

// Size returns the number of elements in the heap
func (h *MinHeapImpl) Size() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return len(h.items)
}

// IsEmpty returns true if the heap contains no elements
func (h *MinHeapImpl) IsEmpty() bool {
	return len(h.items) == 0
}

// heapifyUp maintains the min-heap property by moving a node up the tree
func (h *MinHeapImpl) heapifyUp(index int) {
	for index > 0 {
		parentIdx := parent(index)
		if h.items[parentIdx] > h.items[index] {
			swap(h.items, parentIdx, index)
			index = parentIdx
		} else {
			break
		}
	}
}

// heapifyDown maintains the min-heap property by moving a node down the tree
func (h *MinHeapImpl) heapifyDown(index int) {
	lastIdx := len(h.items) - 1
	for {
		smallest := index
		leftIdx := leftChild(index)
		rightIdx := rightChild(index)

		if leftIdx <= lastIdx && h.items[leftIdx] < h.items[smallest] {
			smallest = leftIdx
		}
		if rightIdx <= lastIdx && h.items[rightIdx] < h.items[smallest] {
			smallest = rightIdx
		}

		if smallest == index {
			break
		}

		swap(h.items, index, smallest)
		index = smallest
	}
}
