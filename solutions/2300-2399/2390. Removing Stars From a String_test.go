package main

import "testing"

/*

You are given a string s, which contains stars *.

In one operation, you can:

Choose a star in s.
Remove the closest non-star character to its left, as well as remove the star itself.
Return the string after all stars have been removed.

Note:

The input will be generated such that the operation is always possible.
It can be shown that the resulting string will always be unique.


Example 1:

Input: s = "leet**cod*e"
Output: "lecoe"
Explanation: Performing the removals from left to right:
- The closest character to the 1st star is 't' in "leet**cod*e". s becomes "lee*cod*e".
- The closest character to the 2nd star is 'e' in "lee*cod*e". s becomes "lecod*e".
- The closest character to the 3rd star is 'd' in "lecod*e". s becomes "lecoe".
There are no more stars, so we return "lecoe".


Example 2:

Input: s = "erase*****"
Output: ""
Explanation: The entire string is removed, so we return an empty string.


Constraints:

1 <= s.length <= 10^5
s consists of lowercase English letters and stars *.
The operation above can be performed on s.

*/

// removeStars removes characters based on star operations
// For each star encountered, it removes the closest non-star character to its left.
func removeStars(s string) string {
	stack := []rune{}

	for _, char := range s {
		if char == '*' {
			// Remove the last character if stack is not empty
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, char)
		}
	}

	return string(stack)
}

func TestRemoveStars(t *testing.T) {
	testCases := []struct {
		s    string
		want string
	}{
		{"leet**cod*e", "lecoe"},
		{"erase*****", ""},
		{"", ""},
	}
	for _, tc := range testCases {
		t.Run("2390. Removing Stars From a String", func(t *testing.T) {
			out := removeStars(tc.s)
			if out != tc.want {
				t.Errorf("removeStars(%v) = %v, want: %v", tc.s, out, tc.want)
			}
		})
	}
}
