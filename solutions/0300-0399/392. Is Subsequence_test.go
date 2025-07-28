package main

import "testing"

/*

Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).



Example 1:

Input: s = "abc", t = "ahbgdc"
Output: true
Example 2:

Input: s = "axc", t = "ahbgdc"
Output: false


Constraints:

0 <= s.length <= 100
0 <= t.length <= 10^4
s and t consist only of lowercase English letters.


Follow up: Suppose there are lots of incoming s, say s1, s2, ..., sk where k >= 10^9, and you want to check one by one to see if t has its subsequence. In this scenario, how would you change your code?

*/

// isSubsequence checks if string s is a subsequence of string t.
// A subsequence is a sequence that can be derived from another sequence
// by deleting some or no elements without changing the order of remaining elements.
//
// The algorithm uses a two-pointer approach:
// 1. One pointer tracks the current position in string s
// 2. Another pointer iterates through string t
// 3. When characters match, advance the pointer in s
// 4. If we've matched all characters in s, it's a valid subsequence
//
// Time Complexity: O(n) where n is the length of t
// Space Complexity: O(1)
//
// Parameters:
//
//	s string - String to check if it's a subsequence
//	t string - String to check against
//
// Returns:
//
//	bool - True if s is a subsequence of t, false otherwise
func isSubsequence(s string, t string) bool {
	// Handle edge case: empty string is always a subsequence
	if len(s) == 0 {
		return true
	}

	left := 0
	for right := range t {
		if t[right] == s[left] {
			left++
		}
		if left == len(s) {
			break
		}
	}

	// Check if all characters in s were found in order
	return left == len(s)
}

func TestIsSubsequence(t *testing.T) {
	testCases := []struct {
		s    string
		t    string
		want bool
	}{
		{s: "abc", t: "ahbgdc", want: true},
		{s: "axc", t: "ahbgdc", want: false},
		{s: "", t: "ahbgdc", want: true},
		{s: "abc", t: "", want: false},
		{s: "a", t: "b", want: false},
		{s: "a", t: "a", want: true},
	}
	for _, tc := range testCases {
		t.Run("392. Is Subsequence", func(t *testing.T) {
			out := isSubsequence(tc.s, tc.t)

			if out != tc.want {
				t.Errorf("isSubsequence(%v, %v) = %v, want: %v", tc.s, tc.t, out, tc.want)
			}
		})
	}
}
