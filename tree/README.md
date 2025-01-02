# Tree Data Structures Package

This package provides implementations of various tree data structures in Go. All implementations are thread-safe and support concurrent operations.

## Features

### Binary Search Tree
- Basic binary search tree implementation
- Thread-safe operations with RWMutex
- Supports:
  - Insert
  - Search/Exists
  - Delete
  - Min/Max value finding
  - Multiple traversal orders (Infix, Prefix, Postfix)
  - Custom traversal directions (LNR, RNL, NLR, NRL, LRN, RLN)

### AVL Tree
- Self-balancing binary search tree
- Maintains height balance property
- Operations:
  - Insert with automatic rebalancing
  - Search
  - InOrder traversal
- Balance operations:
  - Left rotation
  - Right rotation
  - Height updates
  - Balance factor calculation

### Red-Black Tree
- Self-balancing binary search tree with color properties
- Maintains Red-Black properties:
  - Root is black
  - Red nodes can't have red children
  - All paths have same number of black nodes
- Operations:
  - Insert with color fixing
  - Search
  - InOrder traversal
- Balance operations:
  - Left rotation
  - Right rotation
  - Color adjustments

### B+ Tree
- Optimized for storage systems
- Multiple keys per node
- All data stored in leaf nodes
- Leaf nodes linked for range queries

### Segment Tree
- Efficient for range queries
- Supports range updates
- Used for computational geometry

### Radix Tree
- Compressed prefix tree
- Memory efficient for strings
- Fast prefix matching

### Ternary Search Tree
- Hybrid between binary tree and trie
- Efficient for string operations
- Space-efficient compared to standard tries

## Usage Examples

### Binary Search Tree
```go
// Create a new binary tree
tree := BinaryTree(10)

// Insert elements
tree.Insert(5)
tree.Insert(15)
tree.Insert(3)

// Search for elements
tree.Search(5)  // prints: 5: available in the tree
exists := tree.Exists(15) // returns: true

// Find min/max
min := tree.Min() // returns: 3
max := tree.Max() // returns: 15

// Print in different orders
tree.Print("LNR") // Inorder: 3 5 10 15
tree.Print("NLR") // Preorder: 10 5 3 15
```

### AVL Tree
```go
// Create a new AVL tree
avl := NewAVLTree()

// Insert elements (tree automatically balances)
avl.Insert(10)
avl.Insert(20)
avl.Insert(30)

// Search for elements
found := avl.Search(20) // returns: true

// Get sorted elements
var result []int
avl.InOrderTraversal(&result)
```

### Red-Black Tree
```go
// Create a new Red-Black tree
rb := NewRedBlackTree()

// Insert elements (tree maintains Red-Black properties)
rb.Insert(10)
rb.Insert(20)
rb.Insert(30)

// Search for elements
found := rb.Search(20) // returns: true

// Get sorted elements
var result []int
rb.InOrderTraversal(&result)
```

## Implementation Details

### Thread Safety
- All tree implementations use sync.RWMutex
- Read operations use RLock
- Write operations use Lock
- Proper lock/unlock handling with defer

### Performance Characteristics

#### Binary Search Tree
- Average case: O(log n) for all operations
- Worst case: O(n) when unbalanced

#### AVL Tree
- All operations: O(log n) guaranteed
- Extra space for height information
- More rotations than Red-Black tree

#### Red-Black Tree
- All operations: O(log n) guaranteed
- Less rotations than AVL tree
- Slightly more space for color information

## Testing
Each tree implementation comes with comprehensive test coverage. Run tests using:
```bash
go test ./...
```

## Contributing
Contributions are welcome! Please ensure that any new features or modifications come with:
- Proper documentation
- Thread safety considerations
- Comprehensive test cases
- Example usage

## License
This package is distributed under the MIT license. See the LICENSE file for more details. 