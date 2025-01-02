package main

import (
	"fmt"

	"github.com/mstgnz/data-structures/hash"
)

// RunExamples demonstrates various hash table implementations
func RunExamples() {
	// Example 1: Hash Table with djb2 hash function
	fmt.Println("Hash Table Example (djb2):")
	hashTable1 := hash.NewHashTable(10, "djb2")

	fmt.Println("Adding key-value pairs:")
	hashTable1.Put("name", "John")
	hashTable1.Put("age", "30")
	hashTable1.Put("city", "New York")

	fmt.Println("\nRetrieving values:")
	if value, exists := hashTable1.Get("name"); exists {
		fmt.Printf("name: %v\n", value)
	}
	if value, exists := hashTable1.Get("age"); exists {
		fmt.Printf("age: %v\n", value)
	}

	fmt.Println("\nRemoving a key:")
	hashTable1.Remove("age")
	if _, exists := hashTable1.Get("age"); !exists {
		fmt.Println("age was successfully removed")
	}

	// Example 2: Hash Table with sdbm hash function
	fmt.Println("\nHash Table Example (sdbm):")
	hashTable2 := hash.NewHashTable(10, "sdbm")

	fmt.Println("Adding key-value pairs:")
	hashTable2.Put("email", "john@example.com")
	hashTable2.Put("phone", "123-456-7890")
	hashTable2.Put("country", "USA")

	fmt.Println("\nRetrieving values:")
	if value, exists := hashTable2.Get("email"); exists {
		fmt.Printf("email: %v\n", value)
	}
	if value, exists := hashTable2.Get("phone"); exists {
		fmt.Printf("phone: %v\n", value)
	}
}
