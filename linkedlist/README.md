# LinkedList Package

This package provides three different linked list implementations in Go: Linear (Singly), Double (Doubly), and Circular. All implementations are designed to be thread-safe and support generic types.

## Features

### Core Structures
- Linear (Singly) Linked List
- Double (Doubly) Linked List
- Circular Linked List

### Generic Type Support
- Support for all data types (`[T any]`)
- Custom comparison functions
  - `less` function: For sorting operations
  - `equals` function: For equality checks

### Common Operations
- AddToStart: Add element to the beginning
- AddToSequentially: Add element in sorted order
- AddToAfter: Add element after a specific value
- AddToEnd: Add element to the end
- Delete: Remove element
- Search: Find element
- List: Get all elements
- Print: Display elements

### Thread Safety
- Safe read/write operations with RWMutex
- Concurrent access support for all structures
- Deadlock prevention mechanisms

## Usage Examples

### Using with Integer Type
```go
// Linear List example
intList := linkedlist.NewLinear[int](10)
intList.AddToEnd(20)
intList.AddToEnd(30)

// Comparison functions
intLess := func(a, b int) bool { return a < b }
intEquals := func(a, b int) bool { return a == b }

// Sequential adding and searching
intList.AddToSequentially(15, intLess)
exists := intList.Search(20, intEquals)
```

### Using with String Type
```go
// Double List example
strList := linkedlist.NewDouble[string]("Hello")
strList.AddToEnd("World")

// Comparison functions
strLess := func(a, b string) bool { return strings.Compare(a, b) < 0 }
strEquals := func(a, b string) bool { return a == b }

// Sequential adding
strList.AddToSequentially("Go", strLess)

// Forward and backward listing
strList.Print(false)  // Forward direction
strList.Print(true)   // Backward direction
```

### Using with Custom Struct
```go
// Custom struct definition
type Person struct {
    Name string
    Age  int
}

// Circular List example
personList := linkedlist.NewCircular[Person](Person{Name: "John", Age: 25})

// Comparison functions
personLess := func(a, b Person) bool { return a.Age < b.Age }
personEquals := func(a, b Person) bool { 
    return a.Name == b.Name && a.Age == b.Age 
}

// Adding and removing elements
personList.AddToEnd(Person{Name: "Jane", Age: 30})
personList.AddToSequentially(Person{Name: "Bob", Age: 28}, personLess)
personList.Delete(Person{Name: "John", Age: 25}, personEquals)
```

## Implementation Details

### Data Structures
- Linear: Single next pointer
- Double: Both next and prev pointers
- Circular: Last element connected to first with next pointer

### Time Complexities

#### Linear and Double List
- AddToStart: O(1)
- AddToEnd: O(n)
- AddToSequentially: O(n)
- AddToAfter: O(n)
- Delete: O(n)
- Search: O(n)
- List: O(n)

#### Circular List
- AddToStart: O(1)
- AddToEnd: O(n)
- AddToSequentially: O(n)
- AddToAfter: O(n)
- Delete: O(n)
- List: O(n)

### Thread Safety Details
- RLock for read operations
- Lock for write operations
- Automatic unlock with defer
- Safe design for concurrent access

### Memory Management
- Dynamic node creation and deletion
- Pointer management
- Circular reference prevention
- Memory leak prevention mechanisms

## Testing
The package comes with comprehensive test coverage. To run tests:
```bash
go test ./...
```

## Contributing
Contributions are welcome! Please ensure that any new features or modifications come with:
- Proper documentation
- Thread safety considerations
- Comprehensive test cases
- Example usage
- Performance analysis

## License
This package is distributed under the MIT license. See the LICENSE file for more details. 