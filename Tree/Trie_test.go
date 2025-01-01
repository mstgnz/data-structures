package Tree

import (
	"reflect"
	"sort"
	"testing"
)

func TestTrie(t *testing.T) {
	t.Run("Basic Operations", func(t *testing.T) {
		trie := NewTrie()

		// Test insertion and search
		words := []string{"hello", "world", "hi", "hey", "hell", "help"}
		for _, word := range words {
			trie.Insert(word)
		}

		// Test search
		for _, word := range words {
			if !trie.Search(word) {
				t.Errorf("Word %s should be found in the trie", word)
			}
		}

		// Test non-existent word
		if trie.Search("help123") {
			t.Error("Word 'help123' should not be found in the trie")
		}

		// Test prefix search
		prefixTests := []struct {
			prefix   string
			expected bool
		}{
			{"he", true},
			{"hel", true},
			{"world", true},
			{"wor", true},
			{"xyz", false},
		}

		for _, test := range prefixTests {
			if got := trie.StartsWith(test.prefix); got != test.expected {
				t.Errorf("StartsWith(%s) = %v; want %v", test.prefix, got, test.expected)
			}
		}

		// Test deletion
		trie.Delete("hello")
		if trie.Search("hello") {
			t.Error("Word 'hello' should be deleted from the trie")
		}
		if !trie.Search("hell") {
			t.Error("Word 'hell' should still be in the trie")
		}

		// Test GetAllWords
		allWords := trie.GetAllWords()
		sort.Strings(allWords)
		expectedWords := []string{"hell", "help", "hey", "hi", "world"}
		if !reflect.DeepEqual(allWords, expectedWords) {
			t.Errorf("GetAllWords() = %v; want %v", allWords, expectedWords)
		}

		// Test GetWordsWithPrefix
		wordsWithHe := trie.GetWordsWithPrefix("he")
		sort.Strings(wordsWithHe)
		expectedHeWords := []string{"hell", "help", "hey"}
		if !reflect.DeepEqual(wordsWithHe, expectedHeWords) {
			t.Errorf("GetWordsWithPrefix('he') = %v; want %v", wordsWithHe, expectedHeWords)
		}
	})

	t.Run("Empty Trie", func(t *testing.T) {
		trie := NewTrie()

		if trie.Search("test") {
			t.Error("Empty trie should not find any word")
		}

		if trie.StartsWith("t") {
			t.Error("Empty trie should not find any prefix")
		}

		if len(trie.GetAllWords()) != 0 {
			t.Error("Empty trie should return empty list of words")
		}
	})

	t.Run("Unicode Support", func(t *testing.T) {
		trie := NewTrie()
		words := []string{"こんにちは", "世界", "你好", "안녕하세요"}

		for _, word := range words {
			trie.Insert(word)
		}

		for _, word := range words {
			if !trie.Search(word) {
				t.Errorf("Word %s should be found in the trie", word)
			}
		}

		allWords := trie.GetAllWords()
		sort.Strings(allWords)
		sort.Strings(words)
		if !reflect.DeepEqual(allWords, words) {
			t.Errorf("GetAllWords() = %v; want %v", allWords, words)
		}
	})
}
