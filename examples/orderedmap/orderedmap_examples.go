package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/mstgnz/data-structures/orderedmap"
)

func main() {
	// Basic Operations Example
	fmt.Println("=== Basic Operations ===")
	basicOperationsExample()

	fmt.Println("\n=== Order Preservation Example ===")
	orderPreservationExample()

	fmt.Println("\n=== Concurrent Operations Example ===")
	concurrentOperationsExample()

	fmt.Println("\n=== Advanced Features Example ===")
	advancedFeaturesExample()

	fmt.Println("\n=== JSON Examples ===")
	jsonExamples()
}

func basicOperationsExample() {
	om := orderedmap.NewOrderedMap()

	// Adding elements
	om.Set("name", "John")
	om.Set("age", 30)
	om.Set("city", "New York")

	// Getting values
	if name, exists := om.Get("name"); exists {
		fmt.Printf("Name: %v\n", name)
	}

	// Checking existence
	if om.Has("age") {
		fmt.Println("Age exists in the map")
	}

	// Deleting an element
	om.Delete("city")
	fmt.Printf("After deleting 'city': %v\n", om)

	// Length of the map
	fmt.Printf("Map size: %d\n", om.Len())
}

func orderPreservationExample() {
	om := orderedmap.NewOrderedMap()

	// Adding elements in specific order
	data := []struct {
		key string
		val int
	}{
		{"first", 1},
		{"second", 2},
		{"third", 3},
		{"fourth", 4},
	}

	for _, item := range data {
		om.Set(item.key, item.val)
	}

	// Demonstrating order preservation
	fmt.Println("Original order of elements:", om)

	// Getting ordered keys
	fmt.Println("Keys in order:", om.Keys())

	// Getting ordered values
	fmt.Println("Values in order:", om.Values())

	// Using Range to iterate in order
	fmt.Println("Iterating in order:")
	om.Range(func(key, value any) bool {
		fmt.Printf("  %v: %v\n", key, value)
		return true
	})
}

func concurrentOperationsExample() {
	om := orderedmap.NewOrderedMap()
	var wg sync.WaitGroup

	// Concurrent writes
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", val)
			om.Set(key, val)
			time.Sleep(100 * time.Millisecond) // Simulate work
		}(i)
	}

	// Concurrent reads
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", val)
			if value, exists := om.Get(key); exists {
				fmt.Printf("Read: %s = %v\n", key, value)
			}
			time.Sleep(50 * time.Millisecond) // Simulate work
		}(i)
	}

	wg.Wait()
	fmt.Println("Final map state:", om)
}

func advancedFeaturesExample() {
	om := orderedmap.NewOrderedMap()

	// Adding some initial data
	om.Set("one", 1)
	om.Set("two", 2)
	om.Set("three", 3)

	// Creating a copy
	copyMap := om.Copy()
	fmt.Println("Original map:", om)
	fmt.Println("Copied map:", copyMap)

	// Modifying original doesn't affect copy
	om.Set("four", 4)
	fmt.Println("\nAfter modifying original:")
	fmt.Println("Original map:", om)
	fmt.Println("Copied map:", copyMap)

	// Early exit from Range
	fmt.Println("\nEarly exit from Range (stop after 2 elements):")
	count := 0
	om.Range(func(key, value any) bool {
		fmt.Printf("  %v: %v\n", key, value)
		count++
		return count < 2
	})

	// Clear the map
	om.Clear()
	fmt.Println("\nAfter clearing the map:", om)
	fmt.Printf("Map size after clear: %d\n", om.Len())
}

// Person represents a sample struct
type JsonPerson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func jsonExamples() {
	om := orderedmap.NewOrderedMap()

	// Let's add some data
	om.Set("user1", JsonPerson{Name: "Jhon", Age: 35})
	om.Set("user2", JsonPerson{Name: "Jane", Age: 28})
	om.Set("settings", map[string]string{
		"theme": "dark",
		"lang":  "en",
	})

	// Marshal example
	fmt.Println("=== Marshal Example ===")
	data := make(map[string]interface{})

	// Convert data from OrderedMap to standard map
	om.Range(func(key, value interface{}) bool {
		data[fmt.Sprint(key)] = value
		return true
	})

	// Convert to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Marshal error: %v\n", err)
		return
	}
	fmt.Printf("JSON output:\n%s\n\n", string(jsonData))

	// Unmarshal example
	fmt.Println("=== Unmarshal Example ===")
	jsonStr := `{
		"user1": {"name": "Jhon", "age": 35},
		"user2": {"name": "Jane", "age": 28},
		"settings": {"theme": "light", "lang": "en"}
	}`

	// Parse JSON string
	var parsedData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &parsedData); err != nil {
		fmt.Printf("Unmarshal error: %v\n", err)
		return
	}

	// Create a new OrderedMap and add the parse data
	newOm := orderedmap.NewOrderedMap()
	for k, v := range parsedData {
		newOm.Set(k, v)
	}

	// Show results
	fmt.Println("After Unmarshal, OrderedMap content:")
	newOm.Range(func(key, value interface{}) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})
}
