package algorithms

// KMPSearch performs Knuth-Morris-Pratt pattern searching
func KMPSearch(text, pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}

	// Compute LPS array
	lps := computeLPSArray(pattern)
	matches := make([]int, 0)

	i, j := 0, 0 // i for text, j for pattern
	for i < len(text) {
		if pattern[j] == text[i] {
			i++
			j++
		}

		if j == len(pattern) {
			matches = append(matches, i-j)
			j = lps[j-1]
		} else if i < len(text) && pattern[j] != text[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return matches
}

// computeLPSArray computes Longest Proper Prefix which is also Suffix array
func computeLPSArray(pattern string) []int {
	lps := make([]int, len(pattern))
	length := 0 // Length of previous longest prefix suffix
	i := 1

	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}

// RabinKarpSearch performs Rabin-Karp pattern searching
func RabinKarpSearch(text, pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}

	// Numbers used for hashing
	d := 256 // Number of characters in alphabet
	q := 101 // A prime number

	matches := make([]int, 0)
	n := len(text)
	m := len(pattern)
	if n < m {
		return matches
	}

	// Calculate hash value for pattern and first window of text
	patternHash := 0
	textHash := 0
	h := 1

	// Calculate h = pow(d, m-1) % q
	for i := 0; i < m-1; i++ {
		h = (h * d) % q
	}

	// Calculate initial hash values
	for i := 0; i < m; i++ {
		patternHash = (d*patternHash + int(pattern[i])) % q
		textHash = (d*textHash + int(text[i])) % q
	}

	// Slide pattern over text one by one
	for i := 0; i <= n-m; i++ {
		if patternHash == textHash {
			// Check character by character
			match := true
			for j := 0; j < m; j++ {
				if text[i+j] != pattern[j] {
					match = false
					break
				}
			}
			if match {
				matches = append(matches, i)
			}
		}

		// Calculate hash value for next window
		if i < n-m {
			textHash = (d*(textHash-int(text[i])*h) + int(text[i+m])) % q
			if textHash < 0 {
				textHash += q
			}
		}
	}

	return matches
}

// BoyerMooreSearch performs Boyer-Moore pattern searching
func BoyerMooreSearch(text, pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}

	matches := make([]int, 0)
	n := len(text)
	m := len(pattern)
	if n < m {
		return matches
	}

	// Create bad character array
	badChar := make([]int, 256)
	for i := 0; i < 256; i++ {
		badChar[i] = -1
	}
	for i := 0; i < m; i++ {
		badChar[pattern[i]] = i
	}

	// Process pattern
	s := 0 // s is shift of the pattern with respect to text
	for s <= n-m {
		j := m - 1

		// Reduce j while characters of pattern and text are matching
		for j >= 0 && pattern[j] == text[s+j] {
			j--
		}

		if j < 0 {
			matches = append(matches, s)
			// Shift pattern so that the next character in text aligns with the last occurrence of it in pattern
			if s+m < n {
				s += m - badChar[text[s+m]]
			} else {
				s++
			}
		} else {
			// Shift pattern so that the bad character in text aligns with the last occurrence of it in pattern
			shift := j - badChar[text[s+j]]
			if shift > 1 {
				s += shift
			} else {
				s++
			}
		}
	}

	return matches
}

// LongestCommonSubsequence finds the longest common subsequence of two strings
func LongestCommonSubsequence(text1, text2 string) string {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Fill dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// Reconstruct LCS
	lcs := make([]byte, dp[m][n])
	i, j := m, n
	index := len(lcs) - 1

	for i > 0 && j > 0 {
		if text1[i-1] == text2[j-1] {
			lcs[index] = text1[i-1]
			i--
			j--
			index--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return string(lcs)
}

// LevenshteinDistance calculates the minimum number of single-character edits required to change one string into another
func LevenshteinDistance(str1, str2 string) int {
	m, n := len(str1), len(str2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize first row and column
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// Fill dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min3(
					dp[i-1][j],   // deletion
					dp[i][j-1],   // insertion
					dp[i-1][j-1], // substitution
				)
			}
		}
	}

	return dp[m][n]
}

// Helper function for min of three values
func min3(a, b, c int) int {
	return min(min(a, b), c)
}
