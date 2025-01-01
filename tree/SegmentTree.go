package tree

import "math"

// SegmentTree represents a segment tree data structure
type SegmentTree struct {
	tree         []int
	arr          []int
	n            int
	combine      func(int, int) int
	defaultValue int
	combineType  string
}

// NewSegmentTree creates a new segment tree from an array
func NewSegmentTree(arr []int, combine func(int, int) int) *SegmentTree {
	n := len(arr)
	if n == 0 {
		return &SegmentTree{
			tree:         []int{},
			arr:          []int{},
			n:            0,
			combine:      combine,
			defaultValue: 0,
		}
	}

	// Determine combine type and default value
	var defaultValue int
	var combineType string

	// Test with sample values to determine combine type
	if combine(1, 2) == 3 {
		defaultValue = 0
		combineType = "sum"
	} else if combine(1, 2) == 1 {
		defaultValue = math.MaxInt32
		combineType = "min"
	} else {
		defaultValue = math.MinInt32
		combineType = "max"
	}

	// Size of segment tree array
	treeSize := 2*int(math.Pow(2, math.Ceil(math.Log2(float64(n))))) - 1

	st := &SegmentTree{
		tree:         make([]int, treeSize),
		arr:          make([]int, n),
		n:            n,
		combine:      combine,
		defaultValue: defaultValue,
		combineType:  combineType,
	}

	// Initialize tree with default values
	for i := range st.tree {
		st.tree[i] = defaultValue
	}

	// Copy input array and build tree
	copy(st.arr, arr)
	st.buildTree(0, 0, n-1)
	return st
}

// buildTree builds the segment tree recursively
func (st *SegmentTree) buildTree(node int, start int, end int) int {
	if start == end {
		st.tree[node] = st.arr[start]
		return st.arr[start]
	}

	mid := (start + end) / 2
	leftVal := st.buildTree(2*node+1, start, mid)
	rightVal := st.buildTree(2*node+2, mid+1, end)
	st.tree[node] = st.combine(leftVal, rightVal)
	return st.tree[node]
}

// Update updates the value at index i to val
func (st *SegmentTree) Update(i int, val int) {
	if st.n == 0 || i < 0 || i >= st.n {
		return
	}
	st.arr[i] = val
	st.updateTree(0, 0, st.n-1, i, val)
}

func (st *SegmentTree) updateTree(node int, start int, end int, idx int, val int) {
	if start == end {
		st.tree[node] = val
		return
	}

	mid := (start + end) / 2
	if idx <= mid {
		st.updateTree(2*node+1, start, mid, idx, val)
	} else {
		st.updateTree(2*node+2, mid+1, end, idx, val)
	}
	st.tree[node] = st.combine(st.tree[2*node+1], st.tree[2*node+2])
}

// Query returns the result of the combine function over the range [left, right]
func (st *SegmentTree) Query(left int, right int) int {
	// Handle empty array case
	if st.n == 0 {
		return st.defaultValue
	}

	// Handle out of bounds cases
	if left < 0 || right >= st.n || left > right {
		switch st.combineType {
		case "sum":
			// For sum queries, include valid range only
			newLeft := left
			if left < 0 {
				newLeft = 0
			}
			newRight := right
			if right >= st.n {
				newRight = st.n - 1
			}
			if newLeft > newRight {
				return 0
			}
			return st.queryRange(0, 0, st.n-1, newLeft, newRight)
		case "min":
			return math.MaxInt32
		case "max":
			return math.MinInt32
		default:
			return st.defaultValue
		}
	}

	return st.queryRange(0, 0, st.n-1, left, right)
}

func (st *SegmentTree) queryRange(node int, start int, end int, left int, right int) int {
	if left > end || right < start {
		return st.defaultValue
	}
	if left <= start && right >= end {
		return st.tree[node]
	}

	mid := (start + end) / 2
	leftVal := st.queryRange(2*node+1, start, mid, left, right)
	rightVal := st.queryRange(2*node+2, mid+1, end, left, right)
	return st.combine(leftVal, rightVal)
}

// GetArray returns the current array
func (st *SegmentTree) GetArray() []int {
	result := make([]int, st.n)
	copy(result, st.arr)
	return result
}

// Common combine functions
func SumCombine(a, b int) int {
	return a + b
}

func MinCombine(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxCombine(a, b int) int {
	if a > b {
		return a
	}
	return b
}
