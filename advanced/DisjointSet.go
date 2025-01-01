package advanced

// DisjointSet represents a disjoint-set data structure
type DisjointSet struct {
	parent []int
	rank   []int
	size   int
}

// NewDisjointSet creates a new disjoint-set data structure
func NewDisjointSet(size int) *DisjointSet {
	parent := make([]int, size)
	rank := make([]int, size)

	// Initialize each element as a separate set
	for i := 0; i < size; i++ {
		parent[i] = i
		rank[i] = 0
	}

	return &DisjointSet{
		parent: parent,
		rank:   rank,
		size:   size,
	}
}

// Find finds the representative (root) of the set containing element x
// Uses path compression for optimization
func (ds *DisjointSet) Find(x int) int {
	if x < 0 || x >= ds.size {
		return -1
	}

	// Path compression: Make each node point directly to the root
	if ds.parent[x] != x {
		ds.parent[x] = ds.Find(ds.parent[x])
	}
	return ds.parent[x]
}

// Union merges the sets containing elements x and y
// Uses union by rank for optimization
func (ds *DisjointSet) Union(x, y int) bool {
	if x < 0 || x >= ds.size || y < 0 || y >= ds.size {
		return false
	}

	// Find the roots of both sets
	rootX := ds.Find(x)
	rootY := ds.Find(y)

	// If elements are already in the same set
	if rootX == rootY {
		return false
	}

	// Union by rank: Attach smaller rank tree under root of higher rank tree
	if ds.rank[rootX] < ds.rank[rootY] {
		ds.parent[rootX] = rootY
	} else if ds.rank[rootX] > ds.rank[rootY] {
		ds.parent[rootY] = rootX
	} else {
		// If ranks are same, make one as root and increment its rank
		ds.parent[rootY] = rootX
		ds.rank[rootX]++
	}

	return true
}

// Connected checks if two elements are in the same set
func (ds *DisjointSet) Connected(x, y int) bool {
	if x < 0 || x >= ds.size || y < 0 || y >= ds.size {
		return false
	}
	return ds.Find(x) == ds.Find(y)
}

// GetSize returns the total number of elements in the disjoint set
func (ds *DisjointSet) GetSize() int {
	return ds.size
}

// GetSetCount returns the number of disjoint sets
func (ds *DisjointSet) GetSetCount() int {
	count := 0
	for i := 0; i < ds.size; i++ {
		if ds.parent[i] == i {
			count++
		}
	}
	return count
}
