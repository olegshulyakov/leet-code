package arrayorstring_test

import (
	"testing"
)

/*

You are given two strings word1 and word2.Merge the strings by adding letters in alternating order, starting with word1.If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.

Example 1:
Input: word1 = "abc", word2 = "pqr"
Output: "apbqcr"
Explanation: The merged string will be merged as so:
word1:  a   b   c
word2:    p   q   r
merged: a p b q c r

Example 2:
Input: word1 = "ab", word2 = "pqrs"
Output: "apbqrs"
Explanation: Notice that as word2 is longer, "rs" is appended to the end.
word1:  a   b
word2:    p   q   r   s
merged: a p b q   r   s

Example 3:
Input: word1 = "abcd", word2 = "pq"
Output: "apbqcd"
Explanation: Notice that as word1 is longer, "cd" is appended to the end.
word1:  a   b   c   d
word2:    p   q
merged: a p b q c   d


Constraints:

1 <= word1.length, word2.length <= 100
word1 and word2 consist of lowercase English letters.

*/

// mergeAlternately merges two strings by alternating characters from each string.
// If one string is longer than the other, the remaining characters are appended to the result.
//
// Example:
//
//	word1 = "abc", word2 = "pqr" → result = "apbqcr"
//	word1 = "ab", word2 = "pqrs" → result = "apbqrs"
//
// Parameters:
//
//	word1 string - First input string
//	word2 string - Second input string
//
// Returns:
//
//	string - Merged string with alternating characters
func mergeAlternately(word1, word2 string) string {
	// Initialize result string builder for efficient string concatenation
	var res string
	i := 0

	// Continue until all characters from both strings are processed
	for i < len(word1) || i < len(word2) {
		// Add character from word1 if available
		if i < len(word1) {
			res += string(word1[i])
		}

		// Add character from word2 if available
		if i < len(word2) {
			res += string(word2[i])
		}

		// Move to next character position
		i++
	}

	return res
}

func TestMergeAlternately(t *testing.T) {
	testCases := []struct {
		word1 string
		word2 string
		want  string
	}{
		{word1: "abc", word2: "pqr", want: "apbqcr"},
		{word1: "ab", word2: "pqrs", want: "apbqrs"},
		{word1: "abcd", word2: "pq", want: "apbqcd"},
	}
	for _, tc := range testCases {
		t.Run("1768. Merge Strings Alternately", func(t *testing.T) {
			result := mergeAlternately(tc.word1, tc.word2)
			if result != tc.want {
				t.Errorf("mergeAlternately(%q, %q) = %q, want: %q", tc.word1, tc.word2, result, tc.want)
			}
		})
	}
}
