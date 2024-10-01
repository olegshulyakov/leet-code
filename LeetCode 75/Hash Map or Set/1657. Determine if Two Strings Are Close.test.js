/*

Two strings are considered close if you can attain one from the other using the following operations:

Operation 1: Swap any two existing characters.
For example, abcde -> aecdb
Operation 2: Transform every occurrence of one existing character into another existing character, and do the same with the other character.
For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)
You can use the operations on either string as many times as necessary.

Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise.



Example 1:

Input: word1 = "abc", word2 = "bca"
Output: true
Explanation: You can attain word2 from word1 in 2 operations.
Apply Operation 1: "abc" -> "acb"
Apply Operation 1: "acb" -> "bca"

Example 2:

Input: word1 = "a", word2 = "aa"
Output: false
Explanation: It is impossible to attain word2 from word1, or vice versa, in any number of operations.
Example 3:

Input: word1 = "cabbba", word2 = "abbccc"
Output: true
Explanation: You can attain word2 from word1 in 3 operations.
Apply Operation 1: "cabbba" -> "caabbb"
Apply Operation 2: "caabbb" -> "baaccc"
Apply Operation 2: "baaccc" -> "abbccc"


Constraints:

1 <= word1.length, word2.length <= 10^5
word1 and word2 contain only lowercase English letters.

*/

/**
 * @param {string} word1
 * @param {string} word2
 * @return {boolean}
 */
var closeStrings = function (word1, word2) {
    // Step 1. Compare length
    // if (word1.length !== word2.length) return false;

    const getMapCharToCount = (word) => {
        const map = new Map();
        for (const c of word) {
            let count = map.get(c);
            if (!count) {
                map.set(c, 1);
            } else {
                map.set(c, count + 1);
            }
        }
        return map;
    }

    const map1 = getMapCharToCount(word1);
    const map2 = getMapCharToCount(word2);

    // Step 2. Check that both string contains of same characters
    for (const [c] of map1) {
        if (!map2.has(c)) {
            return false;
        }
    }
    for (const [c] of map2) {
        if (!map1.has(c)) {
            return false;
        }
    }

    // Step 3. Find two symbols that have same count in both words and filter them
    for (const [char1, count1] of map1) {
        let isFound = false;
        for (const [char2, count2] of map2) {
            if (count1 === count2) {
                map1.delete(char1);
                map2.delete(char2);
                isFound = true;
                break;
            }
        }
        if (!isFound) {
            return false;
        }
    }

    return map1.size === 0 && map2.size === 0;
};

describe('1657. Determine if Two Strings Are Close', () => {
    test('"abc", "bca"', () => expect(closeStrings("abc", "bca")).toEqual(true));
    test('"a", "aa"', () => expect(closeStrings("a", "aa")).toEqual(false));
    test('"cabbba", "abbccc"', () => expect(closeStrings("cabbba", "abbccc")).toEqual(true));
    test('"abc", "dca"', () => expect(closeStrings("abc", "dca")).toEqual(false));
    test('"ccbbaa", "abbccc"', () => expect(closeStrings("ccbbaa", "abbccc")).toEqual(false));
})