/*

Given an array of integers arr, return true if the number of occurrences of each value in the array is unique or false otherwise.



Example 1:

Input: arr = [1,2,2,1,1,3]
Output: true
Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.
Example 2:

Input: arr = [1,2]
Output: false
Example 3:

Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
Output: true


Constraints:

1 <= arr.length <= 1000
-1000 <= arr[i] <= 1000

*/

/**
 * @param {number[]} arr
 * @return {boolean}
 */
var uniqueOccurrences = function (arr) {
    const mapNumberToCount = new Map();
    for (const n of arr) {
        if (mapNumberToCount.has(n)) {
            mapNumberToCount.set(n, mapNumberToCount.get(n) + 1);
        } else {
            mapNumberToCount.set(n, 1);
        }
    }

    const setCounts = new Set();
    for (const count of mapNumberToCount.values()) {
        if (!setCounts.has(count)) {
            setCounts.add(count);
        } else {
            return false;
        }
    }

    return true;
};

describe('1207. Unique Number of Occurrences', () => {
    test.each([
        { arr: [1, 2, 2, 1, 1, 3], expected: true },
        { arr: [1, 2], expected: false },
        { arr: [-3, 0, 1, -3, 1, 1, 1, -3, 10, 0], expected: true },
    ])('$arr', ({ arr, expected }) => expect(uniqueOccurrences(arr)).toEqual(expected));
})