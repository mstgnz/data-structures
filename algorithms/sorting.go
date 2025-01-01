package algorithms

// QuickSort performs quick sort on an integer slice
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)-1]
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, pivot), right...)
}

// MergeSort performs merge sort on an integer slice
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

// merge merges two sorted arrays
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// HeapSort performs heap sort on an integer slice
func HeapSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)

	// Build heap
	for i := len(result)/2 - 1; i >= 0; i-- {
		heapify(result, len(result), i)
	}

	// Extract elements from heap
	for i := len(result) - 1; i > 0; i-- {
		result[0], result[i] = result[i], result[0]
		heapify(result, i, 0)
	}

	return result
}

// heapify maintains heap property
func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// InsertionSort performs insertion sort on an integer slice
func InsertionSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)

	for i := 1; i < len(result); i++ {
		key := result[i]
		j := i - 1

		for j >= 0 && result[j] > key {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}

	return result
}

// BubbleSort performs bubble sort on an integer slice
func BubbleSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)

	for i := 0; i < len(result)-1; i++ {
		for j := 0; j < len(result)-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}

	return result
}

// CountingSort performs counting sort on an integer slice
func CountingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	// Find the range of input array
	max := arr[0]
	min := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	// Create counting array
	count := make([]int, max-min+1)
	for _, num := range arr {
		count[num-min]++
	}

	// Modify counting array to contain actual positions
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	// Create output array
	result := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		result[count[arr[i]-min]-1] = arr[i]
		count[arr[i]-min]--
	}

	return result
}
