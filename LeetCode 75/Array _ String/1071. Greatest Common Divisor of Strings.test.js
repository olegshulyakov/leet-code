/**

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

function checkDivide(s, t) {
    if (s.length % t.length != 0) {
        return false;
    }

    const multiplier = s.length / t.length;
    if (s !== t.repeat(multiplier)) {
        return false;
    }

    return true;
}

function gcdOfStrings(str1, str2) {
    const s = str1.length > str2.length ? str1 : str2; // long
    const t = str1.length > str2.length ? str2 : str1; // short

    // Find t smallest divider
    let res = '';
    for (let i = 0; i < t.length; i++) {
        const test = t.slice(0, i + 1);
        if (checkDivide(t, test) && checkDivide(s, test)) {
            res = test;
        }
    }

    return res;
};

describe("1071. Greatest Common Divisor of Strings", () => {
    test('"ABCABC", "ABC"', () => expect(gcdOfStrings("ABCABC", "ABC")).toBe("ABC"));
    test('"ABABAB", "ABAB"', () => expect(gcdOfStrings("ABABAB", "ABAB")).toBe("AB"));
    test('"LEET", "CODE"', () => expect(gcdOfStrings("LEET", "CODE")).toBe(""));
})