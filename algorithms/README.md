# Algorithms Package

This package provides implementations of various fundamental algorithms in Go, including sorting, searching, and string processing algorithms.

## Features

### Sorting Algorithms
- QuickSort: Divide-and-conquer sorting with pivot selection
- MergeSort: Divide-and-conquer sorting with merging
- HeapSort: Binary heap based sorting
- InsertionSort: Simple sorting by building sorted array
- BubbleSort: Simple comparison-based sorting
- CountingSort: Integer sorting using counting
- RadixSort: Non-comparative integer sorting
- ShellSort: Improved insertion sort with gap sequence

### Searching Algorithms
- Linear Search: Simple sequential search
- Binary Search (Iterative and Recursive):
  - Efficient search in sorted arrays
  - Returns leftmost occurrence
- Jump Search: Block-based search for sorted arrays
- Interpolation Search: Improved binary search with interpolation
- Exponential Search: Search in unbounded arrays
- Fibonacci Search: Divide-and-conquer with Fibonacci numbers

### String Algorithms
- Knuth-Morris-Pratt (KMP): Pattern matching with prefix function
- Rabin-Karp: Pattern matching with rolling hash
- Boyer-Moore: Pattern matching with bad character rule
- Longest Common Subsequence (LCS)
- Levenshtein Distance (Edit Distance)

## Usage Examples

### Sorting Examples
```go
// QuickSort
arr := []int{64, 34, 25, 12, 22, 11, 90}
sorted := QuickSort(arr)

// MergeSort
sorted = MergeSort(arr)

// HeapSort
sorted = HeapSort(arr)

// InsertionSort
sorted = InsertionSort(arr)

// ShellSort
sorted = ShellSort(arr)
```

### Searching Examples
```go
arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// Binary Search
index := BinarySearch(arr, 5) // Returns index of 5

// Jump Search
index = JumpSearch(arr, 7) // Returns index of 7

// Interpolation Search
index = InterpolationSearch(arr, 3) // Returns index of 3

// Exponential Search
index = ExponentialSearch(arr, 8) // Returns index of 8
```

### String Algorithm Examples
```go
// KMP Pattern Search
text := "AABAACAADAABAAABAA"
pattern := "AABA"
matches := KMPSearch(text, pattern) // Returns all occurrences

// Rabin-Karp Search
matches = RabinKarpSearch(text, pattern)

// Boyer-Moore Search
matches = BoyerMooreSearch(text, pattern)

// Longest Common Subsequence
lcs := LongestCommonSubsequence("ABCDGH", "AEDFHR")

// Levenshtein Distance
distance := LevenshteinDistance("kitten", "sitting")
```

## Implementation Details

### Time Complexities

#### Sorting Algorithms
- QuickSort: O(n log n) average, O(n²) worst case
- MergeSort: O(n log n)
- HeapSort: O(n log n)
- InsertionSort: O(n²)
- BubbleSort: O(n²)
- CountingSort: O(n + k) where k is the range
- RadixSort: O(d * (n + k)) where d is number of digits
- ShellSort: O(n log n) to O(n²) depending on gap sequence

#### Searching Algorithms
- Linear Search: O(n)
- Binary Search: O(log n)
- Jump Search: O(√n)
- Interpolation Search: O(log log n) average, O(n) worst case
- Exponential Search: O(log n)
- Fibonacci Search: O(log n)

#### String Algorithms
- KMP Search: O(n + m) where n is text length, m is pattern length
- Rabin-Karp: O(n + m) average, O(nm) worst case
- Boyer-Moore: O(n/m) best case, O(nm) worst case
- LCS: O(mn) where m, n are string lengths
- Levenshtein Distance: O(mn)

### Space Complexities
- QuickSort: O(log n)
- MergeSort: O(n)
- HeapSort: O(1)
- String Pattern Matching: O(m)
- LCS and Edit Distance: O(mn)

## Testing
Each algorithm comes with comprehensive test coverage. Run tests using:
```bash
go test ./...
```

## Contributing
Contributions are welcome! Please ensure that any new features or modifications come with:
- Proper documentation
- Comprehensive test cases
- Example usage
- Time and space complexity analysis

## License
This package is distributed under the MIT license. See the LICENSE file for more details. 