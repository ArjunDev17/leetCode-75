package main

import (
	"fmt"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	// Frequency arrays for 'a' to 'z'
	count1, count2 := make([]int, 26), make([]int, 26)

	for _, ch := range word1 {
		count1[ch-'a']++
	}
	for _, ch := range word2 {
		count2[ch-'a']++
	}

	// Ensure both words have the same set of characters
	for i := 0; i < 26; i++ {
		if (count1[i] == 0) != (count2[i] == 0) {
			return false
		}
	}

	// Sort the frequency counts and compare
	sort.Ints(count1)
	sort.Ints(count2)

	for i := 0; i < 26; i++ {
		if count1[i] != count2[i] {
			return false
		}
	}
	return true
}

// Main function to test the implementation
func main() {
	// Test cases
	tests := []struct {
		word1, word2 string
		expected     bool
	}{
		{"abc", "bca", true},
		{"a", "aa", false},
		{"cabbba", "abbccc", true},
		{"ab", "ba", true},
		{"abbzccca", "babzzczc", true},
		{"aabbcc", "ccbbaa", true},
		{"abcd", "abdc", true},
		{"abcd", "abcde", false},
		{"xyz", "zyx", true},
		{"abb", "bba", true},
	}

	// Run tests
	for _, test := range tests {
		result := closeStrings(test.word1, test.word2)
		fmt.Printf("closeStrings(%q, %q) = %v | Expected: %v\n",
			test.word1, test.word2, result, test.expected)
	}
}
