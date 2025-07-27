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

/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isSubsequence = function (s, t) {
    let currentIndex = 0; // index of s element to find
    for (let i = 0; i < t.length && currentIndex < s.length; i++) {
        if (t[i] === s[currentIndex]) {
            currentIndex++;
        }
    }

    return currentIndex === s.length;
};

describe('392. Is Subsequence', () => {
    test.each([
        { s: "abc", t: "ahbgdc", expected: true },
        { s: "axc", t: "ahbgdc", expected: false },
    ])('$s, $t', ({ s, t, expected }) => expect(isSubsequence(s, t)).toEqual(expected));
})