package utils

import (
	"sort"
)

// ArrayToSet converts an array to a set (map with empty struct values)
func ArrayToSet[T comparable](arr []T) map[T]struct{} {
	set := make(map[T]struct{})
	for _, v := range arr {
		set[v] = struct{}{}
	}
	return set
}

// SetToArray converts a set to an array
func SetToArray[T comparable](set map[T]struct{}) []T {
	arr := make([]T, 0, len(set))
	for k := range set {
		arr = append(arr, k)
	}
	return arr
}

// MapToArray converts a map to an array of key-value pairs
func MapToArray[K comparable, V any](m map[K]V) [][2]any {
	arr := make([][2]any, 0, len(m))
	for k, v := range m {
		arr = append(arr, [2]any{k, v})
	}
	return arr
}

// ArrayToMap converts an array of key-value pairs to a map
func ArrayToMap[K comparable, V any](arr [][2]any) map[K]V {
	m := make(map[K]V)
	for _, pair := range arr {
		if k, ok := pair[0].(K); ok {
			if v, ok := pair[1].(V); ok {
				m[k] = v
			}
		}
	}
	return m
}

// ArrayToSortedArray converts an array to a sorted array
func ArrayToSortedArray[T orderable](arr []T) []T {
	result := make([]T, len(arr))
	copy(result, arr)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

// Type constraint for orderable types
type orderable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

// ArrayToFrequencyMap converts an array to a frequency map
func ArrayToFrequencyMap[T comparable](arr []T) map[T]int {
	freq := make(map[T]int)
	for _, v := range arr {
		freq[v]++
	}
	return freq
}

// FrequencyMapToArray converts a frequency map to an array with repeated elements
func FrequencyMapToArray[T comparable](freq map[T]int) []T {
	var result []T
	for v, count := range freq {
		for i := 0; i < count; i++ {
			result = append(result, v)
		}
	}
	return result
}

// ArrayToMatrix converts a 1D array to a 2D matrix with specified number of columns
func ArrayToMatrix[T any](arr []T, cols int) [][]T {
	if cols <= 0 {
		return nil
	}

	rows := (len(arr) + cols - 1) / cols
	matrix := make([][]T, rows)

	for i := range matrix {
		start := i * cols
		end := start + cols
		if end > len(arr) {
			end = len(arr)
		}
		matrix[i] = make([]T, end-start)
		copy(matrix[i], arr[start:end])
	}

	return matrix
}

// MatrixToArray converts a 2D matrix to a 1D array
func MatrixToArray[T any](matrix [][]T) []T {
	if len(matrix) == 0 {
		return nil
	}

	totalLen := 0
	for _, row := range matrix {
		totalLen += len(row)
	}

	result := make([]T, 0, totalLen)
	for _, row := range matrix {
		result = append(result, row...)
	}

	return result
}

// ArrayToLinkedList converts an array to a singly linked list
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func ArrayToLinkedList[T any](arr []T) *Node[T] {
	if len(arr) == 0 {
		return nil
	}

	head := &Node[T]{Value: arr[0]}
	current := head

	for i := 1; i < len(arr); i++ {
		current.Next = &Node[T]{Value: arr[i]}
		current = current.Next
	}

	return head
}

// LinkedListToArray converts a singly linked list to an array
func LinkedListToArray[T any](head *Node[T]) []T {
	var result []T
	current := head

	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}

	return result
}

// ArrayToBinaryTree converts a level-order array representation to a binary tree
type TreeNode[T any] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func ArrayToBinaryTree[T any](arr []T) *TreeNode[T] {
	if len(arr) == 0 {
		return nil
	}

	root := &TreeNode[T]{Value: arr[0]}
	queue := []*TreeNode[T]{root}
	i := 1

	for len(queue) > 0 && i < len(arr) {
		node := queue[0]
		queue = queue[1:]

		if i < len(arr) {
			node.Left = &TreeNode[T]{Value: arr[i]}
			queue = append(queue, node.Left)
			i++
		}

		if i < len(arr) {
			node.Right = &TreeNode[T]{Value: arr[i]}
			queue = append(queue, node.Right)
			i++
		}
	}

	return root
}

// BinaryTreeToArray converts a binary tree to its level-order array representation
func BinaryTreeToArray[T any](root *TreeNode[T]) []T {
	if root == nil {
		return nil
	}

	var result []T
	queue := []*TreeNode[T]{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Value)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}
