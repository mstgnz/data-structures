package algorithms

// LinearSearch performs linear search on a slice
func LinearSearch(arr []int, target int) int {
	for i, num := range arr {
		if num == target {
			return i
		}
	}
	return -1
}

// BinarySearch performs binary search on a sorted slice
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	// Keep track of the leftmost occurrence
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			result = mid
			right = mid - 1 // Continue searching on the left side
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

// BinarySearchRecursive performs recursive binary search on a sorted slice
func BinarySearchRecursive(arr []int, target int) int {
	return binarySearchHelper(arr, target, 0, len(arr)-1)
}

func binarySearchHelper(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if arr[mid] == target {
		// Check if there's another occurrence to the left
		if leftResult := binarySearchHelper(arr, target, left, mid-1); leftResult != -1 {
			return leftResult
		}
		return mid
	}

	if arr[mid] < target {
		return binarySearchHelper(arr, target, mid+1, right)
	}

	return binarySearchHelper(arr, target, left, mid-1)
}

// JumpSearch performs jump search on a sorted slice
func JumpSearch(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	// Finding the optimal jump step size
	step := int(float64(n) / 4)
	if step < 1 {
		step = 1
	}

	// Finding the block where element is present (if exists)
	prev := 0
	curr := step
	for curr < n && arr[curr] < target {
		prev = curr
		curr += step
	}

	// Adjust curr if it exceeds array bounds
	if curr >= n {
		curr = n - 1
	}

	// Linear search for target in block beginning with prev
	// Start from prev and search until we find target or exceed curr
	for i := prev; i <= curr && i < n; i++ {
		if arr[i] == target {
			// Found a match, but check if there's an earlier occurrence
			j := i - 1
			for j >= prev && arr[j] == target {
				i = j
				j--
			}
			return i
		}
	}

	return -1
}

// InterpolationSearch performs interpolation search on a sorted slice
func InterpolationSearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	result := -1

	for left <= right && target >= arr[left] && target <= arr[right] {
		if left == right {
			if arr[left] == target {
				result = left
			}
			break
		}

		// Probing the position with keeping uniform distribution in mind
		pos := left + ((right - left) * (target - arr[left]) / (arr[right] - arr[left]))

		if arr[pos] == target {
			result = pos
			right = pos - 1 // Continue searching on the left side
		} else if arr[pos] < target {
			left = pos + 1
		} else {
			right = pos - 1
		}
	}

	return result
}

// ExponentialSearch performs exponential search on a sorted slice
func ExponentialSearch(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	if arr[0] == target {
		return 0
	}

	// Find range for binary search
	i := 1
	for i < n && arr[i] < target {
		i *= 2
	}

	// Adjust i if it exceeds array bounds
	if i >= n {
		i = n - 1
	}

	// Find the leftmost occurrence in the range [i/2, i]
	left := i / 2
	right := i
	result := -1

	// First, check if there are any occurrences in [0, left-1]
	for j := 0; j < left; j++ {
		if arr[j] == target {
			return j
		}
	}

	// Then search in [left, right]
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			result = mid
			right = mid - 1 // Continue searching on the left side
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

// FibonacciSearch performs Fibonacci search on a sorted slice
func FibonacciSearch(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	// Initialize Fibonacci numbers
	fibM2 := 0 // (m-2)'th Fibonacci number
	fibM1 := 1 // (m-1)'th Fibonacci number
	fibM := 1  // m'th Fibonacci number

	// Find the smallest Fibonacci number greater than or equal to n
	for fibM < n {
		fibM2 = fibM1
		fibM1 = fibM
		fibM = fibM2 + fibM1
	}

	// Initialize the eliminated range
	offset := -1
	result := -1

	// While there are elements to be inspected
	for fibM > 1 {
		// Check if fibM2 is a valid location
		i := min(offset+fibM2, n-1)

		if arr[i] < target {
			fibM = fibM1
			fibM1 = fibM2
			fibM2 = fibM - fibM1
			offset = i
		} else if arr[i] > target {
			fibM = fibM2
			fibM1 = fibM1 - fibM2
			fibM2 = fibM - fibM1
		} else {
			// Found target, but continue searching on the left side
			result = i
			fibM = fibM2
			fibM1 = fibM1 - fibM2
			fibM2 = fibM - fibM1
		}
	}

	// Check the final element if we haven't found a better result
	if fibM1 == 1 && arr[offset+1] == target && (result == -1 || offset+1 < result) {
		result = offset + 1
	}

	return result
}

// Helper functions for min and max values
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
