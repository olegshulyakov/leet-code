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

1 <= s.length <= 3 * 105
s consist of printable ASCII characters.

*/

function reverseVowels(s) {
    const vowels = ['a', 'e', 'i', 'o', 'u'];
    const source = Array.from(s);
    let lastChangedInx = source.length;
    for (let i = 0; i < source.length; i++) {
        if (vowels.includes(source[i].toLowerCase())) {
            for (let j = lastChangedInx - 1; j > i; j--) {
                if (vowels.includes(source[j].toLowerCase())) {
                    const mem = source[i];
                    source[i] = source[j];
                    source[j] = mem;
                    lastChangedInx = j;
                    break;
                }
            }
        }
    }
    return source.join("");
};

describe('345. Reverse Vowels of a String', () => {
    test('"IceCreAm"', () => expect(reverseVowels("IceCreAm")).toEqual("AceCreIm"));
    test('"leetcode"', () => expect(reverseVowels("leetcode")).toEqual("leotcede"));
})