package algorithms

import (
	"reflect"
	"testing"
)

// Test cases for pattern searching algorithms
var patternSearchTestCases = []struct {
	name     string
	text     string
	pattern  string
	expected []int
}{
	{
		name:     "Empty pattern",
		text:     "hello",
		pattern:  "",
		expected: []int{},
	},
	{
		name:     "Empty text",
		text:     "",
		pattern:  "pattern",
		expected: []int{},
	},
	{
		name:     "Pattern not found",
		text:     "hello world",
		pattern:  "xyz",
		expected: []int{},
	},
	{
		name:     "Single match",
		text:     "hello world",
		pattern:  "world",
		expected: []int{6},
	},
	{
		name:     "Multiple matches",
		text:     "abababa",
		pattern:  "aba",
		expected: []int{0, 2, 4},
	},
	{
		name:     "Overlapping matches",
		text:     "aaa",
		pattern:  "aa",
		expected: []int{0, 1},
	},
}

func TestKMPSearch(t *testing.T) {
	for _, tc := range patternSearchTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := KMPSearch(tc.text, tc.pattern)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("KMPSearch(%q, %q) = %v; want %v", tc.text, tc.pattern, result, tc.expected)
			}
		})
	}
}

func TestRabinKarpSearch(t *testing.T) {
	for _, tc := range patternSearchTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := RabinKarpSearch(tc.text, tc.pattern)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("RabinKarpSearch(%q, %q) = %v; want %v", tc.text, tc.pattern, result, tc.expected)
			}
		})
	}
}

func TestBoyerMooreSearch(t *testing.T) {
	for _, tc := range patternSearchTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := BoyerMooreSearch(tc.text, tc.pattern)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("BoyerMooreSearch(%q, %q) = %v; want %v", tc.text, tc.pattern, result, tc.expected)
			}
		})
	}
}

// Test cases for LCS
var lcsTestCases = []struct {
	name     string
	text1    string
	text2    string
	expected string
}{
	{
		name:     "Empty strings",
		text1:    "",
		text2:    "",
		expected: "",
	},
	{
		name:     "One empty string",
		text1:    "abc",
		text2:    "",
		expected: "",
	},
	{
		name:     "No common subsequence",
		text1:    "abc",
		text2:    "def",
		expected: "",
	},
	{
		name:     "Simple case",
		text1:    "abcde",
		text2:    "ace",
		expected: "ace",
	},
	{
		name:     "Complex case",
		text1:    "AGGTAB",
		text2:    "GXTXAYB",
		expected: "GTAB",
	},
}

func TestLongestCommonSubsequence(t *testing.T) {
	for _, tc := range lcsTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := LongestCommonSubsequence(tc.text1, tc.text2)
			if result != tc.expected {
				t.Errorf("LongestCommonSubsequence(%q, %q) = %q; want %q", tc.text1, tc.text2, result, tc.expected)
			}
		})
	}
}

// Test cases for Levenshtein distance
var levenshteinTestCases = []struct {
	name     string
	str1     string
	str2     string
	expected int
}{
	{
		name:     "Empty strings",
		str1:     "",
		str2:     "",
		expected: 0,
	},
	{
		name:     "One empty string",
		str1:     "abc",
		str2:     "",
		expected: 3,
	},
	{
		name:     "Same strings",
		str1:     "abc",
		str2:     "abc",
		expected: 0,
	},
	{
		name:     "Single operation",
		str1:     "cat",
		str2:     "cut",
		expected: 1,
	},
	{
		name:     "Multiple operations",
		str1:     "sunday",
		str2:     "saturday",
		expected: 3,
	},
}

func TestLevenshteinDistance(t *testing.T) {
	for _, tc := range levenshteinTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := LevenshteinDistance(tc.str1, tc.str2)
			if result != tc.expected {
				t.Errorf("LevenshteinDistance(%q, %q) = %d; want %d", tc.str1, tc.str2, result, tc.expected)
			}
		})
	}
}
