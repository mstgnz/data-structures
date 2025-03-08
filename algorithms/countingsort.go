package algorithms

import (
	"strings"
)

// CountingSort implements the Counting Sort algorithm for non-negative integers
// Time Complexity: O(n + k) where n is the number of elements and k is the range of input
// Space Complexity: O(k)
func CountingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	// Find the maximum element to determine the count array size
	max := arr[0]
	for _, num := range arr {
		if num < 0 {
			panic("Counting sort only works with non-negative integers")
		}
		if num > max {
			max = num
		}
	}

	// Create count array to store count of each unique object
	count := make([]int, max+1)

	// Store count of each object
	for _, num := range arr {
		count[num]++
	}

	// Modify count array to store actual position of each object
	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}

	// Create output array
	output := make([]int, len(arr))

	// Build the output array
	// To make it stable, we process the input array from end to start
	for i := len(arr) - 1; i >= 0; i-- {
		output[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}

	// Copy the output array to input array
	copy(arr, output)

	return arr
}

// CountingSortWithRange implements Counting Sort for a known range of integers
// This version is more efficient when the range is known and smaller than the array size
func CountingSortWithRange(arr []int, minVal, maxVal int) []int {
	if len(arr) == 0 {
		return arr
	}

	// Validate the range
	for _, num := range arr {
		if num < minVal || num > maxVal {
			panic("Array contains elements outside the specified range")
		}
	}

	// Create count array for the range
	count := make([]int, maxVal-minVal+1)

	// Store count of each object
	for _, num := range arr {
		count[num-minVal]++
	}

	// Modify count array to store actual position of each object
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	// Create output array
	output := make([]int, len(arr))

	// Build the output array
	for i := len(arr) - 1; i >= 0; i-- {
		output[count[arr[i]-minVal]-1] = arr[i]
		count[arr[i]-minVal]--
	}

	// Copy the output array to input array
	copy(arr, output)

	return arr
}

// CountingSortString sorts a string using counting sort algorithm
func CountingSortString(str string) string {
	if len(str) == 0 {
		return ""
	}

	// Create a count array for ASCII characters (256 possible values)
	count := make([]int, 256)

	// Count occurrences of each character in the input string
	for i := 0; i < len(str); i++ {
		count[str[i]]++
	}

	// Build the sorted string
	var result strings.Builder
	result.Grow(len(str)) // Pre-allocate space for efficiency

	// Add characters in sorted order with their original frequency
	for i := 0; i < 256; i++ {
		for j := 0; j < count[i]; j++ {
			result.WriteByte(byte(i))
		}
	}

	return result.String()
}

// CountingSortBytes implements Counting Sort for byte slices
// This is useful for sorting binary data or custom encodings
func CountingSortBytes(arr []byte) []byte {
	if len(arr) == 0 {
		return arr
	}

	// Create count array for bytes (256 possible values)
	count := make([]int, 256)

	// Store count of each byte
	for _, b := range arr {
		count[b]++
	}

	// Modify count array to store actual position of each byte
	for i := 1; i < 256; i++ {
		count[i] += count[i-1]
	}

	// Create output array
	output := make([]byte, len(arr))

	// Build the output array
	for i := len(arr) - 1; i >= 0; i-- {
		output[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}

	// Copy the output array to input array
	copy(arr, output)

	return arr
}
