package main

import (
	"testing"
)

func closeStrings(word1 string, word2 string) bool {
	// Step 1. Compare length
	if len(word1) != len(word2) {
		return false
	}

	// Helper function to count character frequencies
	getCharCount := func(word string) map[rune]int {
		counts := make(map[rune]int)
		for _, c := range word {
			counts[c]++
		}
		return counts
	}

	map1 := getCharCount(word1)
	map2 := getCharCount(word2)

	// Step 2. Check that both string contains of same characters
	for char := range map1 {
		if map2[char] == 0 {
			return false
		}
	}
	for char := range map2 {
		if map1[char] == 0 {
			return false
		}
	}

	// Step 3. Find two symbols that have same count in both words and filter them
	for char1, count1 := range map1 {
		isFound := false
		for char2, count2 := range map2 {
			if count1 == count2 {
				delete(map1, char1)
				delete(map2, char2)
				isFound = true
				break
			}
		}
		if !isFound {
			return false
		}
	}

	return len(map1) == 0 && len(map2) == 0
}

func TestCloseStrings(t *testing.T) {
	testCases := []struct {
		word1 string
		word2 string
		want  bool
	}{
		{"abc", "bca", true},
		{"a", "aa", false},
		{"cabbba", "abbccc", true},
		{"abc", "dca", false},
		{"ccbbaa", "abbccc", false},
	}
	for _, tc := range testCases {
		t.Run("1657. Determine if Two Strings Are Close", func(t *testing.T) {
			out := closeStrings(tc.word1, tc.word2)
			if out != tc.want {
				t.Errorf("closeStrings(%v, %v) = %v, want: %v", tc.word1, tc.word2, out, tc.want)
			}
		})
	}
}
