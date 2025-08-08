package main

import "testing"

/*

Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.

Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.



Example 1:

Input: s = "abciiidef", k = 3
Output: 3
Explanation: The substring "iii" contains 3 vowel letters.
Example 2:

Input: s = "aeiou", k = 2
Output: 2
Explanation: Any substring of length 2 contains 2 vowels.
Example 3:

Input: s = "leetcode", k = 3
Output: 2
Explanation: "lee", "eet" and "ode" contain 2 vowels.


Constraints:

1 <= s.length <= 10^5
s consists of lowercase English letters.
1 <= k <= s.length

*/

func maxVowels(s string, k int) int {
	isVowel := func(letter byte) bool {
		return letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u'
	}

	maxVowelsCount := 0
	temp := 0

	for i := range len(s) {
		if i >= k {
			if isVowel(s[i-k]) {
				temp--
			}
		}
		if isVowel(s[i]) {
			temp++
		}
		if maxVowelsCount < temp {
			maxVowelsCount = temp
		}
	}

	return maxVowelsCount
}

func TestMaxVowels(t *testing.T) {
	const name = "1456. Maximum Number of Vowels in a Substring of Given Length"
	testCases := []struct {
		s    string
		k    int
		want int
	}{
		{"abciiidef", 3, 3},
		{"aeiou", 2, 2},
		{"leetcode", 3, 2},
	}

	for _, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Logf("fun(%v)", tc.s)
			out := maxVowels(tc.s, tc.k)
			if out != tc.want {
				t.Errorf("maxVowels(%s, %d) = %d; expected %d", tc.s, tc.k, out, tc.want)
			}
		})
	}
}
