# Utility Package

This package provides utility functions and tools for data structure implementations in Go. It includes iterators, serialization helpers, and data structure converters.

## Features

### Iterators
- Generic iterator interface with `HasNext()`, `Next()`, and `Reset()` methods
- Implementations for different data structures:
  - `SliceIterator`: For array/slice iteration
  - `MapIterator`: For map iteration with optional key ordering
  - `LinkedListIterator`: For linked list traversal
  - `TreeIterator`: For tree traversal with multiple orders (PreOrder, InOrder, PostOrder, LevelOrder)

### Serialization
- Support for multiple formats:
  - JSON serialization with pretty print option
  - XML serialization with pretty print option
  - GOB serialization for Go-specific binary format
- Helper functions for string serialization/deserialization
- Deep copy functionality using serialization

### Type Converters
- Array/Slice conversions:
  - `ArrayToSet`: Convert array to set
  - `SetToArray`: Convert set to array
  - `ArrayToSortedArray`: Sort array of orderable types
  - `ArrayToFrequencyMap`: Create frequency map from array
  - `ArrayToMatrix`: Convert 1D array to 2D matrix
- Data structure conversions:
  - `ArrayToLinkedList`: Convert array to linked list
  - `LinkedListToArray`: Convert linked list to array
  - `ArrayToBinaryTree`: Convert array to binary tree (level-order)
  - `BinaryTreeToArray`: Convert binary tree to array (level-order)
- Map conversions:
  - `MapToArray`: Convert map to array of key-value pairs
  - `ArrayToMap`: Convert array of key-value pairs to map

## Usage Examples

### Iterator Usage
```go
// Slice Iterator
numbers := []int{1, 2, 3, 4, 5}
iter := NewSliceIterator(numbers)
for iter.HasNext() {
    value := iter.Next()
    // Process value
}

// Map Iterator
myMap := map[string]int{"a": 1, "b": 2}
mapIter := NewMapIterator(myMap, nil)
for mapIter.HasNext() {
    key, value := mapIter.Next()
    // Process key-value pair
}

// Tree Iterator with different traversal orders
treeIter := NewTreeIterator(root, InOrder)
for treeIter.HasNext() {
    value := treeIter.Next()
    // Process value
}
```

### Serialization Usage
```go
// JSON Serialization
jsonSerializer := JSONSerializer{PrettyPrint: true}
data := MyStruct{...}
bytes, err := jsonSerializer.Serialize(data)

// Using SerializationHelper
helper := NewSerializationHelper(FormatJSON, true)
str, err := helper.Serialize(data)

// Deep Copy
src := MyStruct{...}
var dst MyStruct
err := DeepCopy(&src, &dst)
```

### Converter Usage
```go
// Array to Set conversion
arr := []int{1, 2, 2, 3, 3, 4}
set := ArrayToSet(arr)

// Array to Sorted Array
unsorted := []int{3, 1, 4, 1, 5, 9}
sorted := ArrayToSortedArray(unsorted)

// Array to Binary Tree
arr := []int{1, 2, 3, 4, 5}
tree := ArrayToBinaryTree(arr)

// Matrix conversions
matrix := ArrayToMatrix([]int{1, 2, 3, 4, 5, 6}, 3)
array := MatrixToArray(matrix)
```

## Implementation Details

### Generic Support
- All implementations use Go generics for type safety
- Type constraints:
  - `comparable` for map keys and set elements
  - `orderable` for sortable types
  - `any` for general purpose use

### Performance Considerations
- Efficient implementations focusing on minimal allocations
- Appropriate data structure choices for different operations
- Optimized algorithms for common operations

## Testing
Each component comes with comprehensive test coverage. Run tests using:
```bash
go test ./...
```

## Contributing
Contributions are welcome! Please ensure that any new features or modifications come with:
- Proper documentation
- Comprehensive test cases
- Example usage
- Generic type support where applicable

## License
This package is distributed under the MIT license. See the LICENSE file for more details. 