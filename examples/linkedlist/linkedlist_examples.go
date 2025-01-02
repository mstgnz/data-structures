package main

import (
	"fmt"
	"strings"

	"github.com/mstgnz/data-structures/linkedlist"
)

func main() {
	fmt.Println("=== Linear Linked List Examples ===")
	linearListExamples()

	fmt.Println("\n=== Double Linked List Examples ===")
	doubleListExamples()

	fmt.Println("\n=== Circular Linked List Examples ===")
	circularListExamples()

	fmt.Println("\n=== Custom Type Examples ===")
	customTypeExamples()
}

func linearListExamples() {
	// Create a new Linear List with integers
	intList := linkedlist.NewLinear[int](10)

	// Define comparison functions
	intLess := func(a, b int) bool { return a < b }
	intEquals := func(a, b int) bool { return a == b }

	// Add elements sequentially
	fmt.Println("\nAdding elements sequentially:")
	intList.AddToSequentially(5, intLess)
	intList.AddToSequentially(15, intLess)
	intList.AddToSequentially(12, intLess)
	intList.Print() // Expected: 5 10 12 15

	// Add element after a specific value
	fmt.Println("\nAdding 13 after 12:")
	intList.AddToAfter(13, 12, intEquals)
	intList.Print() // Expected: 5 10 12 13 15

	// Search for elements
	fmt.Printf("\nSearch for 12: %v\n", intList.Search(12, intEquals))
	fmt.Printf("Search for 99: %v\n", intList.Search(99, intEquals))

	// Delete elements
	fmt.Println("\nDeleting 12:")
	intList.Delete(12, intEquals)
	intList.Print()

	// Create a Linear List with strings
	strList := linkedlist.NewLinear[string]("Hello")

	// Define string comparison functions
	strLess := func(a, b string) bool { return strings.Compare(a, b) < 0 }
	strEquals := func(a, b string) bool { return a == b }

	fmt.Println("\nString List Operations:")
	strList.AddToSequentially("World", strLess)
	strList.AddToSequentially("Go", strLess)
	strList.Print() // Expected: Go Hello World

	// Search and delete operations with strings
	fmt.Printf("Search for 'Hello': %v\n", strList.Search("Hello", strEquals))
	strList.Delete("Hello", strEquals)
	strList.Print()
}

func doubleListExamples() {
	// Create a new Double List with integers
	intList := linkedlist.NewDouble[int](10)

	// Define comparison functions
	intLess := func(a, b int) bool { return a < b }
	intEquals := func(a, b int) bool { return a == b }

	// Add elements sequentially
	fmt.Println("\nAdding elements sequentially:")
	intList.AddToSequentially(5, intLess)
	intList.AddToSequentially(15, intLess)
	intList.AddToSequentially(12, intLess)

	// Print forward and backward
	fmt.Println("Forward traversal:")
	intList.Print(false)
	fmt.Println("Backward traversal:")
	intList.Print(true)

	// Add element after a specific value
	fmt.Println("\nAdding 13 after 12:")
	intList.AddToAfter(13, 12, intEquals)
	intList.Print(false)

	// Delete elements
	fmt.Println("\nDeleting 12:")
	intList.Delete(12, intEquals)
	intList.Print(false)
}

func circularListExamples() {
	// Create a new Circular List with integers
	intList := linkedlist.NewCircular[int](10)

	// Define comparison functions
	intLess := func(a, b int) bool { return a < b }
	intEquals := func(a, b int) bool { return a == b }

	// Add elements sequentially
	fmt.Println("\nAdding elements sequentially:")
	intList.AddToSequentially(5, intLess)
	intList.AddToSequentially(15, intLess)
	intList.AddToSequentially(12, intLess)
	intList.Print()

	// Add element after a specific value
	fmt.Println("\nAdding 13 after 12:")
	intList.AddToAfter(13, 12, intEquals)
	intList.Print()

	// Delete elements
	fmt.Println("\nDeleting 12:")
	intList.Delete(12, intEquals)
	intList.Print()
}

func customTypeExamples() {
	// Define a custom type
	type Person struct {
		Name string
		Age  int
	}

	// Create comparison functions for Person type
	personLess := func(a, b Person) bool { return a.Age < b.Age }
	personEquals := func(a, b Person) bool {
		return a.Name == b.Name && a.Age == b.Age
	}

	// Create lists of different types with Person
	linearList := linkedlist.NewLinear[Person](Person{Name: "Alice", Age: 25})
	doubleList := linkedlist.NewDouble[Person](Person{Name: "Bob", Age: 30})
	circularList := linkedlist.NewCircular[Person](Person{Name: "Charlie", Age: 35})

	// Add elements to Linear List
	fmt.Println("\nLinear List with Person type:")
	linearList.AddToSequentially(Person{Name: "David", Age: 20}, personLess)
	linearList.AddToSequentially(Person{Name: "Eve", Age: 28}, personLess)
	linearList.Print()

	// Search and delete operations with Person type
	searchPerson := Person{Name: "David", Age: 20}
	fmt.Printf("Search for %v: %v\n", searchPerson, linearList.Search(searchPerson, personEquals))
	linearList.Delete(searchPerson, personEquals)
	linearList.Print()

	// Add elements to Double List
	fmt.Println("\nDouble List with Person type:")
	doubleList.AddToSequentially(Person{Name: "Frank", Age: 22}, personLess)
	doubleList.AddToSequentially(Person{Name: "Grace", Age: 33}, personLess)
	doubleList.Print(false)
	fmt.Println("Backward traversal:")
	doubleList.Print(true)

	// Add elements to Circular List
	fmt.Println("\nCircular List with Person type:")
	circularList.AddToSequentially(Person{Name: "Henry", Age: 27}, personLess)
	circularList.AddToSequentially(Person{Name: "Ivy", Age: 40}, personLess)
	circularList.Print()
}
