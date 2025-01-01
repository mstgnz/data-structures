package tree

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

// tree.List(pType string) -> pType -> Infix: LNR-RNL, Prefix: NLR-NRL, Postfix: LRN, RLN

func TestBinaryTree(t *testing.T) {
	tree := BinaryTree(1)
	expect := []int{1}
	if got := tree.List("NRL"); !reflect.DeepEqual(got, expect) {
		t.Errorf("BinaryTree() = %v, want %v", got, expect)
	}
}

func Test_binaryTree_Delete(t *testing.T) {
	tree := BinaryTree(1)
	expect := []int{1, 3} // for NRL
	tree.Insert(2)
	tree.Insert(3)
	tree.Delete(2)
	if got := tree.List("NRL"); !reflect.DeepEqual(got, expect) {
		t.Errorf("Delete() = %v, want %v", got, expect)
	}
}

func Test_binaryTree_Exists(t *testing.T) {
	tree := BinaryTree(1)
	expect := false
	got := tree.Exists(13)
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Exists() = %v, want %v", got, expect)
	}
}

func Test_binaryTree_Insert(t *testing.T) {
	tree := BinaryTree(1)
	expect := []int{1, 2, 3} // for NRL
	tree.Insert(2)
	tree.Insert(3)
	if got := tree.List("NRL"); !reflect.DeepEqual(got, expect) {
		t.Errorf("Insert() = %v, want %v", got, expect)
	}
}

func Test_binaryTree_Max(t *testing.T) {
	tree := BinaryTree(1)
	expect := 3
	tree.Insert(2)
	tree.Insert(3)
	if got := tree.Max(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Max() = %v, want %v", got, expect)
	}
}

func Test_binaryTree_Min(t *testing.T) {
	tree := BinaryTree(1)
	expect := 1
	tree.Insert(2)
	tree.Insert(3)
	if got := tree.Min(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Min() = %v, want %v", got, expect)
	}
}

func Test_binaryTree_Search(t *testing.T) {
	tree := BinaryTree(1)
	expect := "1: available in the tree\n"
	got := captureOutput(func() {
		tree.Search(1)
	})
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Search() = %v, want %v", got, expect)
	}
}

func TestBinaryTree_Print(t *testing.T) {
	tests := []struct {
		name string
		data []int
	}{
		{
			name: "print empty tree",
			data: []int{},
		},
		{
			name: "print single node",
			data: []int{1},
		},
		{
			name: "print multiple nodes",
			data: []int{5, 3, 7, 2, 4, 6, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tree IBinaryTree
			if len(tt.data) > 0 {
				tree = BinaryTree(tt.data[0])
				for i := 1; i < len(tt.data); i++ {
					tree.Insert(tt.data[i])
				}
			} else {
				tree = BinaryTree(0)
			}
			tree.Print("LNR") // Infix traversal
		})
	}
}

func TestBinaryTree_PrefixPrint(t *testing.T) {
	tests := []struct {
		name string
		data []int
	}{
		{
			name: "prefix print empty tree",
			data: []int{},
		},
		{
			name: "prefix print single node",
			data: []int{1},
		},
		{
			name: "prefix print balanced tree",
			data: []int{5, 3, 7, 2, 4, 6, 8},
		},
		{
			name: "prefix print left-heavy tree",
			data: []int{5, 4, 3, 2, 1},
		},
		{
			name: "prefix print right-heavy tree",
			data: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tree IBinaryTree
			if len(tt.data) > 0 {
				tree = BinaryTree(tt.data[0])
				for i := 1; i < len(tt.data); i++ {
					tree.Insert(tt.data[i])
				}
			} else {
				tree = BinaryTree(0)
			}
			tree.Print("NLR") // Prefix traversal
		})
	}
}

func TestBinaryTree_PostfixPrint(t *testing.T) {
	tests := []struct {
		name string
		data []int
	}{
		{
			name: "postfix print empty tree",
			data: []int{},
		},
		{
			name: "postfix print single node",
			data: []int{1},
		},
		{
			name: "postfix print balanced tree",
			data: []int{5, 3, 7, 2, 4, 6, 8},
		},
		{
			name: "postfix print left-heavy tree",
			data: []int{5, 4, 3, 2, 1},
		},
		{
			name: "postfix print right-heavy tree",
			data: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tree IBinaryTree
			if len(tt.data) > 0 {
				tree = BinaryTree(tt.data[0])
				for i := 1; i < len(tt.data); i++ {
					tree.Insert(tt.data[i])
				}
			} else {
				tree = BinaryTree(0)
			}
			tree.Print("LRN") // Postfix traversal
		})
	}
}

func TestBinaryTree_Search(t *testing.T) {
	tests := []struct {
		name     string
		init     int
		add      []int
		search   int
		expected bool
	}{
		{
			name:     "search in empty tree",
			init:     0,
			add:      []int{},
			search:   1,
			expected: false,
		},
		{
			name:     "search existing value",
			init:     5,
			add:      []int{3, 7, 2, 4, 6, 8},
			search:   4,
			expected: true,
		},
		{
			name:     "search non-existing value",
			init:     5,
			add:      []int{3, 7, 2, 4, 6, 8},
			search:   9,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := BinaryTree(tt.init)
			for _, v := range tt.add {
				tree.Insert(v)
			}
			tree.Search(tt.search)
		})
	}
}

func TestBinaryTree_PostfixTraversal(t *testing.T) {
	tests := []struct {
		name string
		data []int
	}{
		{
			name: "postfix print empty tree",
			data: []int{},
		},
		{
			name: "postfix print single node",
			data: []int{1},
		},
		{
			name: "postfix print balanced tree",
			data: []int{5, 3, 7, 2, 4, 6, 8},
		},
		{
			name: "postfix print left-heavy tree",
			data: []int{5, 4, 3, 2, 1},
		},
		{
			name: "postfix print right-heavy tree",
			data: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tree IBinaryTree
			if len(tt.data) > 0 {
				tree = BinaryTree(tt.data[0])
				for i := 1; i < len(tt.data); i++ {
					tree.Insert(tt.data[i])
				}
			} else {
				tree = BinaryTree(0)
			}
			tree.Print("LRN") // Postfix traversal
		})
	}
}

func TestBinaryTree_Max(t *testing.T) {
	tests := []struct {
		name     string
		init     int
		add      []int
		expected int
	}{
		{
			name:     "max in single node tree",
			init:     5,
			add:      []int{},
			expected: 5,
		},
		{
			name:     "max in balanced tree",
			init:     5,
			add:      []int{3, 7, 2, 4, 6, 8},
			expected: 8,
		},
		{
			name:     "max in left-heavy tree",
			init:     5,
			add:      []int{4, 3, 2, 1},
			expected: 5,
		},
		{
			name:     "max in right-heavy tree",
			init:     1,
			add:      []int{2, 3, 4, 5},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := BinaryTree(tt.init)
			for _, v := range tt.add {
				tree.Insert(v)
			}
			if got := tree.Max(); got != tt.expected {
				t.Errorf("Max() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestBinaryTree_Min(t *testing.T) {
	tests := []struct {
		name     string
		init     int
		add      []int
		expected int
	}{
		{
			name:     "min in single node tree",
			init:     5,
			add:      []int{},
			expected: 5,
		},
		{
			name:     "min in balanced tree",
			init:     5,
			add:      []int{3, 7, 2, 4, 6, 8},
			expected: 2,
		},
		{
			name:     "min in left-heavy tree",
			init:     5,
			add:      []int{4, 3, 2, 1},
			expected: 1,
		},
		{
			name:     "min in right-heavy tree",
			init:     1,
			add:      []int{2, 3, 4, 5},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := BinaryTree(tt.init)
			for _, v := range tt.add {
				tree.Insert(v)
			}
			if got := tree.Min(); got != tt.expected {
				t.Errorf("Min() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// captureOutput print
func captureOutput(f func()) string {
	// Create a pipe
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	f()

	// Restore stdout and read the output
	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String()
}
