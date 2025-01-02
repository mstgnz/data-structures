package algorithms

// getMax returns the maximum element from an array
func getMax(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

// countingSortByDigit performs counting sort on arr based on the digit at position exp
func countingSortByDigit(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10) // 0-9 digits

	// Store count of occurrences in count[]
	for i := 0; i < n; i++ {
		digit := (arr[i] / exp) % 10
		count[digit]++
	}

	// Change count[i] so that count[i] now contains actual
	// position of this digit in output[]
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// Build the output array
	for i := n - 1; i >= 0; i-- {
		digit := (arr[i] / exp) % 10
		output[count[digit]-1] = arr[i]
		count[digit]--
	}

	// Copy the output array to arr[]
	copy(arr, output)
}

// RadixSort implements the Radix Sort algorithm for non-negative integers
// Time Complexity: O(d * (n + k)) where d is the number of digits, n is the number of elements
// and k is the range of values for each digit (10 for decimal)
func RadixSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	// Find the maximum number to know number of digits
	max := getMax(arr)

	// Do counting sort for every digit
	for exp := 1; max/exp > 0; exp *= 10 {
		countingSortByDigit(arr, exp)
	}

	return arr
}

// RadixSortString implements Radix Sort for strings
// This implementation sorts strings of equal length
func RadixSortString(arr []string) []string {
	if len(arr) == 0 {
		return arr
	}

	// Find the length of strings (assuming all strings have same length)
	maxLen := len(arr[0])
	for _, str := range arr {
		if len(str) != maxLen {
			panic("All strings must have the same length")
		}
	}

	// Create output array and count array
	output := make([]string, len(arr))
	count := make([]int, 256) // ASCII characters

	// Process all characters from right to left
	for pos := maxLen - 1; pos >= 0; pos-- {
		// Reset count array
		for i := range count {
			count[i] = 0
		}

		// Count frequencies
		for _, str := range arr {
			count[str[pos]]++
		}

		// Change count[i] so that count[i] now contains actual
		// position of this character in output[]
		for i := 1; i < 256; i++ {
			count[i] += count[i-1]
		}

		// Build the output array
		for i := len(arr) - 1; i >= 0; i-- {
			char := arr[i][pos]
			output[count[char]-1] = arr[i]
			count[char]--
		}

		// Copy the output array to arr[]
		copy(arr, output)
	}

	return arr
}

// RadixSortBytes implements Radix Sort for byte slices of equal length
func RadixSortBytes(arr [][]byte) [][]byte {
	if len(arr) == 0 {
		return arr
	}

	// Find the length of byte slices (assuming all have same length)
	maxLen := len(arr[0])
	for _, bytes := range arr {
		if len(bytes) != maxLen {
			panic("All byte slices must have the same length")
		}
	}

	// Create output array and count array
	output := make([][]byte, len(arr))
	count := make([]int, 256)

	// Process all bytes from right to left
	for pos := maxLen - 1; pos >= 0; pos-- {
		// Reset count array
		for i := range count {
			count[i] = 0
		}

		// Count frequencies
		for _, bytes := range arr {
			count[bytes[pos]]++
		}

		// Change count[i] so that count[i] now contains actual
		// position of this byte in output[]
		for i := 1; i < 256; i++ {
			count[i] += count[i-1]
		}

		// Build the output array
		for i := len(arr) - 1; i >= 0; i-- {
			b := arr[i][pos]
			output[count[b]-1] = arr[i]
			count[b]--
		}

		// Copy the output array to arr[]
		copy(arr, output)
	}

	return arr
}
