package hash

import (
	"hash"
	"hash/fnv"
	"math"
)

// BloomFilter represents a Bloom filter data structure
type BloomFilter struct {
	bitArray  []bool
	size      uint
	numHash   uint // number of hash functions
	hashFuncs []hash.Hash64
}

// NewBloomFilter creates a new Bloom filter with the given size and desired false positive rate
func NewBloomFilter(expectedElements int, falsePositiveRate float64) *BloomFilter {
	// Calculate optimal size and number of hash functions
	size := calculateOptimalSize(expectedElements, falsePositiveRate)
	numHash := calculateOptimalHashFunctions(expectedElements, size)

	// Initialize hash functions
	hashFuncs := make([]hash.Hash64, numHash)
	for i := range hashFuncs {
		hashFuncs[i] = fnv.New64a()
	}

	return &BloomFilter{
		bitArray:  make([]bool, size),
		size:      uint(size),
		numHash:   uint(numHash),
		hashFuncs: hashFuncs,
	}
}

// calculateOptimalSize calculates the optimal size for the bit array
func calculateOptimalSize(n int, p float64) uint {
	return uint(math.Ceil(-float64(n) * math.Log(p) / math.Pow(math.Log(2), 2)))
}

// calculateOptimalHashFunctions calculates the optimal number of hash functions
func calculateOptimalHashFunctions(n int, m uint) uint {
	return uint(math.Ceil(float64(m) / float64(n) * math.Log(2)))
}

// getHashValues generates hash values for the given data
func (bf *BloomFilter) getHashValues(data []byte) []uint {
	hashValues := make([]uint, bf.numHash)
	for i, h := range bf.hashFuncs {
		h.Reset()
		h.Write(data)
		hashValues[i] = uint(h.Sum64() % uint64(bf.size))
	}
	return hashValues
}

// Add adds an element to the Bloom filter
func (bf *BloomFilter) Add(data []byte) {
	for _, hashValue := range bf.getHashValues(data) {
		bf.bitArray[hashValue] = true
	}
}

// Contains checks if an element might be in the set
func (bf *BloomFilter) Contains(data []byte) bool {
	for _, hashValue := range bf.getHashValues(data) {
		if !bf.bitArray[hashValue] {
			return false
		}
	}
	return true
}

// Clear resets the Bloom filter
func (bf *BloomFilter) Clear() {
	bf.bitArray = make([]bool, bf.size)
}

// EstimateFalsePositiveRate estimates the current false positive rate
func (bf *BloomFilter) EstimateFalsePositiveRate(numElements int) float64 {
	k := float64(bf.numHash)
	m := float64(bf.size)
	n := float64(numElements)
	return math.Pow(1-math.Exp(-k*n/m), k)
}
