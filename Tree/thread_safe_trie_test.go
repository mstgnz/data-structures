package tree

import (
	"sort"
	"sync"
	"testing"
)

func TestThreadSafeTrie(t *testing.T) {
	trie := NewThreadSafeTrie()

	// Test initial state
	if !trie.IsEmpty() {
		t.Error("New Trie should be empty")
	}

	if size := trie.Size(); size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}

	// Test Insert and Contains
	words := []string{"apple", "app", "apricot", "banana", "bat"}
	for _, word := range words {
		trie.Insert(word)
	}

	for _, word := range words {
		if !trie.Contains(word) {
			t.Errorf("Trie should contain word %s", word)
		}
	}

	if trie.Contains("orange") {
		t.Error("Trie should not contain word 'orange'")
	}

	// Test StartsWith
	if !trie.StartsWith("app") {
		t.Error("Trie should have words starting with 'app'")
	}
	if !trie.StartsWith("ba") {
		t.Error("Trie should have words starting with 'ba'")
	}
	if trie.StartsWith("cat") {
		t.Error("Trie should not have words starting with 'cat'")
	}

	// Test GetWordsWithPrefix
	appWords := trie.GetWordsWithPrefix("app")
	expectedAppWords := []string{"app", "apple"}
	sort.Strings(appWords)
	sort.Strings(expectedAppWords)
	if len(appWords) != len(expectedAppWords) {
		t.Errorf("Expected %v words with prefix 'app', got %v", expectedAppWords, appWords)
	}
	for i := range appWords {
		if appWords[i] != expectedAppWords[i] {
			t.Errorf("Expected word %s, got %s", expectedAppWords[i], appWords[i])
		}
	}

	// Test Remove
	if !trie.Remove("app") {
		t.Error("Remove should return true for existing word")
	}
	if trie.Contains("app") {
		t.Error("Trie should not contain 'app' after removal")
	}
	if trie.Remove("orange") {
		t.Error("Remove should return false for non-existing word")
	}

	// Test Clear
	trie.Clear()
	if !trie.IsEmpty() {
		t.Error("Trie should be empty after clear")
	}
}

func TestThreadSafeTrieConcurrent(t *testing.T) {
	trie := NewThreadSafeTrie()
	var wg sync.WaitGroup
	numGoroutines := 10
	wordsPerGoroutine := 100

	// Generate test words
	generateWord := func(n int) string {
		return string(rune('a' + n%26))
	}

	// Test concurrent insertions
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(n int) {
			defer wg.Done()
			for j := 0; j < wordsPerGoroutine; j++ {
				word := generateWord(n*wordsPerGoroutine + j)
				trie.Insert(word)
			}
		}(i)
	}
	wg.Wait()

	// Test concurrent operations (Contains, StartsWith)
	wg.Add(numGoroutines * 2)
	for i := 0; i < numGoroutines; i++ {
		go func(n int) {
			defer wg.Done()
			for j := 0; j < wordsPerGoroutine; j++ {
				word := generateWord(n*wordsPerGoroutine + j)
				if !trie.Contains(word) {
					t.Errorf("Trie should contain word %s", word)
				}
			}
		}(i)
		go func(n int) {
			defer wg.Done()
			for j := 0; j < wordsPerGoroutine; j++ {
				prefix := generateWord(n*wordsPerGoroutine + j)
				if !trie.StartsWith(prefix) {
					t.Errorf("Trie should have words starting with %s", prefix)
				}
			}
		}(i)
	}
	wg.Wait()
}

func TestThreadSafeTrieEdgeCases(t *testing.T) {
	trie := NewThreadSafeTrie()

	// Test empty string
	trie.Insert("")
	if !trie.Contains("") {
		t.Error("Trie should contain empty string after insertion")
	}

	// Test single character
	trie.Insert("a")
	if !trie.Contains("a") {
		t.Error("Trie should contain single character word")
	}

	// Test duplicate insertions
	size := trie.Size()
	trie.Insert("a")
	if trie.Size() != size {
		t.Error("Size should not change after duplicate insertion")
	}

	// Test prefix of existing word
	trie.Insert("apple")
	trie.Insert("app")
	if !trie.Contains("app") || !trie.Contains("apple") {
		t.Error("Trie should contain both word and its prefix")
	}

	// Test GetAllWords
	trie.Clear()
	words := []string{"a", "ab", "abc", "b", "bc"}
	for _, word := range words {
		trie.Insert(word)
	}
	allWords := trie.GetAllWords()
	sort.Strings(allWords)
	sort.Strings(words)
	if len(allWords) != len(words) {
		t.Errorf("Expected %v words, got %v", words, allWords)
	}
	for i := range allWords {
		if allWords[i] != words[i] {
			t.Errorf("Expected word %s, got %s", words[i], allWords[i])
		}
	}
}
