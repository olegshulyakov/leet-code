package main

import (
	"testing"
)

/*
Given a string s, reverse only all the vowels in the string and return it.
The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.


Example 1:
Input: s = "IceCreAm"
Output: "AceCreIm"
Explanation:
The vowels in s are ['I', 'e', 'e', 'A']. On reversing the vowels, s becomes "AceCreIm".

Example 2:
Input: s = "leetcode"
Output: "leotcede"


Constraints:

1 <= s.length <= 3 * 10^5
s consist of printable ASCII characters.

*/

// reverseVowels reverses only the vowels in a given string.
// Vowels are 'a', 'e', 'i', 'o', 'u' (case-insensitive).
// All other characters remain in their original positions.
//
// Parameters:
//
//	s string - Input string to process
//
// Returns:
//
//	string - String with vowels reversed
func reverseVowels(s string) string {
	// Define vowels set for efficient lookup
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	// Convert string to rune slice for proper handling of Unicode characters
	source := []rune(s)

	// Track the last position where a vowel was changed
	lastChangedIdx := len(source)

	// Iterate through the string from left to right
	for i := 0; i < len(source); i++ {
		// Check if current character is a vowel
		if vowels[source[i]] {
			// Search for next vowel from the right side
			for j := lastChangedIdx - 1; j > i; j-- {
				// Check if character at position j is a vowel
				if vowels[source[j]] {
					// Swap the vowels
					source[i], source[j] = source[j], source[i]

					// Update last changed position
					lastChangedIdx = j

					// Break inner loop after swap
					break
				}
			}
		}
	}

	// Convert back to string and return
	return string(source)
}

func TestReverseVowels(t *testing.T) {
	name := "345. Reverse Vowels of a String"
	testCases := []struct{ s, want string }{
		{s: "IceCreAm", want: "AceCreIm"},
		{s: "leetcode", want: "leotcede"},
	}
	for _, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			out := reverseVowels(tc.s)
			if out != tc.want {
				t.Errorf("reverseVowels(%v) = %v, want: %v", tc.s, out, tc.want)
			}
		})
	}
}
