package algorithms

// ShellSort implements the Shell Sort algorithm
// Shell sort is an optimization of insertion sort that allows the exchange of items that are far apart
func ShellSort(arr []int) []int {
	n := len(arr)
	// Start with a big gap, then reduce the gap
	for gap := n / 2; gap > 0; gap /= 2 {
		// Do a gapped insertion sort for this gap size
		// The first gap elements arr[0..gap-1] are already in gapped order
		// Keep adding one more element until the entire array is gap sorted
		for i := gap; i < n; i++ {
			// Add arr[i] to the elements that have been gap sorted
			// Save arr[i] in temp and make a hole at position i
			temp := arr[i]

			// Shift earlier gap-sorted elements up until the correct location for arr[i] is found
			var j int
			for j = i; j >= gap && arr[j-gap] > temp; j -= gap {
				arr[j] = arr[j-gap]
			}

			// Put temp (the original arr[i]) in its correct location
			arr[j] = temp
		}
	}
	return arr
}

// ShellSortWithGaps implements Shell Sort with custom gap sequence
// This version allows you to specify the gap sequence to use
func ShellSortWithGaps(arr []int, gaps []int) []int {
	n := len(arr)
	// Use the provided gap sequence
	for _, gap := range gaps {
		// Skip gaps that are too large
		if gap >= n {
			continue
		}

		// Perform insertion sort with current gap
		for i := gap; i < n; i++ {
			temp := arr[i]
			var j int
			for j = i; j >= gap && arr[j-gap] > temp; j -= gap {
				arr[j] = arr[j-gap]
			}
			arr[j] = temp
		}
	}
	return arr
}

// Hibbard sequence: 2^k - 1
func HibbardGaps(n int) []int {
	gaps := make([]int, 0)
	k := 1
	for {
		gap := (1 << k) - 1 // 2^k - 1
		if gap >= n {
			break
		}
		gaps = append([]int{gap}, gaps...) // Prepend to get descending order
		k++
	}
	return gaps
}

// Sedgewick sequence: 4^k + 3 * 2^(k-1) + 1
func SedgewickGaps(n int) []int {
	gaps := make([]int, 0)
	k := 0
	for {
		var gap int
		if k%2 == 0 {
			gap = 9*(1<<k) - 9*(1<<(k/2)) + 1
		} else {
			gap = 8*(1<<k) - 6*(1<<((k+1)/2)) + 1
		}
		if gap >= n {
			break
		}
		gaps = append([]int{gap}, gaps...) // Prepend to get descending order
		k++
	}
	return gaps
}

// Pratt sequence: 2^i * 3^j
func PrattGaps(n int) []int {
	// Use map to avoid duplicates
	gapMap := make(map[int]bool)

	// Generate sequence
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			gap := (1 << i) * pow3(j)
			if gap >= n {
				break
			}
			gapMap[gap] = true
		}
		if (1 << i) >= n {
			break
		}
	}

	// Convert map to sorted slice
	gaps := make([]int, 0, len(gapMap))
	for gap := range gapMap {
		gaps = append(gaps, gap)
	}

	// Sort in descending order
	for i := 0; i < len(gaps)-1; i++ {
		for j := i + 1; j < len(gaps); j++ {
			if gaps[i] < gaps[j] {
				gaps[i], gaps[j] = gaps[j], gaps[i]
			}
		}
	}

	return gaps
}

// Helper function to calculate 3^n
func pow3(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 3
	}
	return result
}
