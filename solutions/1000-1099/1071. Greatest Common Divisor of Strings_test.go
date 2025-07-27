package main

import (
	"strings"
	"testing"
)

/*

For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).

Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

Example 1:
Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"

Example 2:
Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"

Example 3:
Input: str1 = "LEET", str2 = "CODE"
Output: ""


Constraints:
1 <= str1.length, str2.length <= 1000
str1 and str2 consist of English uppercase letters.

*/

// checkDivide checks if string 's' can be formed by repeating string 't' multiple times.
// A string 't' divides string 's' if s = t + t + ... + t (multiple concatenations).
//
// Parameters:
//
//	s string - The string to be divided
//	t string - The potential divisor string
//
// Returns:
//
//	bool - True if 't' can divide 's', false otherwise
func checkDivide(s, t string) bool {
	// If length of 's' is not divisible by length of 't', 't' cannot divide 's'
	if len(s)%len(t) != 0 {
		return false
	}

	// Calculate how many times 't' needs to be repeated to match length of 's'
	multiplier := len(s) / len(t)

	// Create string by repeating 't' multiplier times
	repeated := strings.Repeat(t, multiplier)

	// Check if the repeated string equals 's'
	if s != repeated {
		return false
	}

	return true
}

// gcdOfStrings finds the greatest common divisor of two strings.
// The GCD string 'x' is the largest string that can divide both input strings.
// A string 't' divides string 's' if s can be formed by concatenating 't' multiple times.
//
// Parameters:
//
//	str1 string - First input string
//	str2 string - Second input string
//
// Returns:
//
//	string - The greatest common divisor string, or empty string if none exists
func gcdOfStrings(str1, str2 string) string {
	// Assign longer string to 's' and shorter to 't'
	var s, t string
	if len(str1) > len(str2) {
		s = str1
		t = str2
	} else {
		s = str2
		t = str1
	}

	// Find the largest common divisor string
	var res string

	// Iterate through all possible prefixes of the shorter string
	for i := 0; i < len(t); i++ {
		// Create test string as prefix of 't' from index 0 to i+1
		test := t[:i+1]

		// Check if this prefix can divide both strings
		if checkDivide(t, test) && checkDivide(s, test) {
			res = test
		}
	}

	return res
}

func TestGcdOfStrings(t *testing.T) {
	testCases := []struct{ str1, str2, want string }{
		{str1: "ABCABC", str2: "ABC", want: "ABC"},
		{str1: "ABABAB", str2: "ABAB", want: "AB"},
		{str1: "LEET", str2: "CODE", want: ""},
	}
	for _, tc := range testCases {
		t.Run("1071. Greatest Common Divisor of Strings", func(t *testing.T) {
			out := gcdOfStrings(tc.str1, tc.str2)
			if out != tc.want {
				t.Errorf("gcdOfStrings(%q, %q) = %q, want: %q", tc.str1, tc.str2, out, tc.want)
			}
		})
	}
}
