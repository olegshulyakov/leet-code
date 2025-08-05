package main

import (
	"strings"
	"testing"
)

/*

Given an input string s, reverse the order of the words.
A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
Return a string of the words in reverse order concatenated by a single space.
Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.


Example 1:
Input: s = "the sky is blue"
Output: "blue is sky the"

Example 2:
Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.

Example 3:
Input: s = "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.


Constraints:

1 <= s.length <= 10^4
s contains English letters (upper-case and lower-case), digits, and spaces ' '.
There is at least one word in s.


Follow-up: If the string data type is mutable in your language, can you solve it in-place with O(1) extra space?

*/

// reverseWords takes a string and returns a new string with the words in reverse order.
// It handles multiple spaces between words and trims leading/trailing spaces.
//
// Parameters:
//
//	s string - Input string that may contain multiple spaces between words
//
// Returns:
//
//	string - String with words in reverse order, separated by single spaces
func reverseWords(s string) string {
	// Split the string by spaces
	source := strings.Split(s, " ")

	// Initialize result slice to store non-empty words
	res := []string{}

	// Iterate through the words in reverse order
	for i := len(source) - 1; i >= 0; i-- {
		// Only add non-empty words to result
		if source[i] != "" {
			res = append(res, source[i])
		}
	}

	// Join the reversed words with single spaces
	return strings.Join(res, " ")
}

func TestReverseWords(t *testing.T) {
	const name = "151. Reverse Words in a String"
	testCases := []struct{ s, want string }{
		{s: "the sky is blue", want: "blue is sky the"},
		{s: "  hello world  ", want: "world hello"},
		{s: "a good   example", want: "example good a"},
	}
	for _, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			out := reverseWords(tc.s)
			if out != tc.want {
				t.Errorf("reverseWords(%v) = %v, want %v", tc.s, out, tc.want)
			}
		})
	}
}
