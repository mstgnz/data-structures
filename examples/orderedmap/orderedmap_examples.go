package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/mstgnz/data-structures/orderedmap"
)

func main() {
	// Create a new OrderedMap
	om := orderedmap.New()

	// Create a WaitGroup for synchronization
	var wg sync.WaitGroup

	// Concurrent writes
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", val)
			om.Set(key, val)
			time.Sleep(100 * time.Millisecond) // Simulate some work
		}(i)
	}

	// Concurrent reads
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			keys := om.Keys()
			fmt.Printf("Current keys: %v\n", keys)
			time.Sleep(50 * time.Millisecond) // Simulate some work
		}()
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Print final state
	fmt.Println("\nFinal map content:", om)
	fmt.Println("Map size:", om.Len())
	fmt.Println("All keys:", om.Keys())
	fmt.Println("All values:", om.Values())

	// Demonstrate concurrent delete and read
	wg.Add(2)
	go func() {
		defer wg.Done()
		om.Delete("key-1")
		fmt.Println("Deleted key-1")
	}()

	go func() {
		defer wg.Done()
		if val, exists := om.Get("key-2"); exists {
			fmt.Printf("Read while deleting: key-2 = %v\n", val)
		}
	}()

	wg.Wait()
	fmt.Println("\nFinal state after concurrent operations:", om)
}
