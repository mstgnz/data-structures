package hash

import (
	"testing"
)

// TestBloomFilterBasicOperations tests basic Bloom filter operations
func TestBloomFilterBasicOperations(t *testing.T) {
	// Create a new Bloom filter with 1000 expected elements and 0.01 false positive rate
	bf := NewBloomFilter(1000, 0.01)

	// Test adding and checking elements
	testData := []string{
		"test1",
		"test2",
		"test3",
		"test4",
		"test5",
	}

	// Add elements
	for _, data := range testData {
		bf.Add([]byte(data))
	}

	// Test for presence of added elements
	for _, data := range testData {
		if !bf.Contains([]byte(data)) {
			t.Errorf("Bloom filter should contain %s", data)
		}
	}

	// Test for absence of non-added elements
	nonExistentData := []string{
		"nonexistent1",
		"nonexistent2",
		"nonexistent3",
	}

	falsePositives := 0
	for _, data := range nonExistentData {
		if bf.Contains([]byte(data)) {
			falsePositives++
		}
	}

	// Calculate actual false positive rate
	actualFPR := float64(falsePositives) / float64(len(nonExistentData))
	expectedFPR := 0.01 // Our target false positive rate

	// The actual FPR might be higher than expected due to the small sample size
	if actualFPR > expectedFPR*2 {
		t.Errorf("False positive rate too high: got %v, expected around %v", actualFPR, expectedFPR)
	}
}

// TestBloomFilterClear tests the clear operation
func TestBloomFilterClear(t *testing.T) {
	bf := NewBloomFilter(100, 0.01)

	// Add some elements
	testData := "test data"
	bf.Add([]byte(testData))

	// Verify element is present
	if !bf.Contains([]byte(testData)) {
		t.Error("Element should be present before clear")
	}

	// Clear the filter
	bf.Clear()

	// Verify element is no longer present
	if bf.Contains([]byte(testData)) {
		t.Error("Element should not be present after clear")
	}
}

// TestBloomFilterFalsePositiveRate tests the estimated false positive rate
func TestBloomFilterFalsePositiveRate(t *testing.T) {
	expectedElements := 1000
	targetFPR := 0.01
	bf := NewBloomFilter(expectedElements, targetFPR)

	// Add expectedElements number of elements
	for i := 0; i < expectedElements; i++ {
		bf.Add([]byte(string(rune(i))))
	}

	// Get estimated FPR
	estimatedFPR := bf.EstimateFalsePositiveRate(expectedElements)

	// The estimated FPR should be close to our target
	if estimatedFPR < targetFPR*0.5 || estimatedFPR > targetFPR*1.5 {
		t.Errorf("Estimated false positive rate %v is not close to target %v", estimatedFPR, targetFPR)
	}
}
