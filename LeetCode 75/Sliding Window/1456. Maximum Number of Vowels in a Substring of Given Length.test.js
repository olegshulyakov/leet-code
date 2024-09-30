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

/**
 * @param {string} s
 * @param {number} k
 * @return {number}
 */
var maxVowels = function (s, k) {
    const isVowel = (letter) => letter === 'a' || letter === 'e' || letter === 'i' || letter === 'o' || letter === 'u';

    let maxVowels = temp = 0;
    for (let i = 0; i < s.length; i++) {
        if (i >= k) {
            temp -= isVowel(s[i - k]) ? 1 : 0;
        }
        temp += isVowel(s[i]) ? 1 : 0;
        maxVowels = maxVowels < temp ? temp : maxVowels;
    }

    return maxVowels;
};

describe('1456. Maximum Number of Vowels in a Substring of Given Length', () => {
    test('"abciiidef", 3', () => expect(maxVowels("abciiidef", 3)).toEqual(3));
    test('"aeiou", 2', () => expect(maxVowels("aeiou", 2)).toEqual(2));
    test('"leetcode", 3', () => expect(maxVowels("leetcode", 3)).toEqual(2));
})