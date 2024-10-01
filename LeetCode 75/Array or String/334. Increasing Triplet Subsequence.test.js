/*

Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.



Example 1:

Input: nums = [1,2,3,4,5]
Output: true
Explanation: Any triplet where i < j < k is valid.
Example 2:

Input: nums = [5,4,3,2,1]
Output: false
Explanation: No triplet exists.
Example 3:

Input: nums = [2,1,5,0,4,6]
Output: true
Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.


Constraints:

1 <= nums.length <= 5 * 10^5
-231 <= nums[i] <= 231 - 1


Follow up: Could you implement a solution that runs in O(n) time complexity and O(1) space complexity?

*/

/**
 * @param {number[]} nums
 * @return {boolean}
 */
var increasingTriplet = function (nums) {
    if (nums.length < 3) return false;

    let smallestNumber = nums[0];
    const pairs = new Array();
    pairs.push([nums[0], undefined]);

    for (let i = 1; i < nums.length; i++) {
        const numberToCheck = nums[i];

        // Add new pair
        if (numberToCheck < smallestNumber) {
            pairs.unshift([numberToCheck, undefined]);
            smallestNumber = numberToCheck;
            continue;
        }

        // Check if nums[i] is last number in triple
        if (pairs.some(pair => pair[1] !== undefined && pair[1] < numberToCheck)) {
            return true;
        }

        // Replace second number in pair
        pairs.filter(pair => pair[0] < numberToCheck && (pair[1] === undefined || numberToCheck < pair[1])).forEach(pair => pair[1] = numberToCheck);
    }

    return false;
};

describe('334. Increasing Triplet Subsequence', () => {
    test('[1, 2, 3, 4, 5]', () => expect(increasingTriplet([1, 2, 3, 4, 5])).toEqual(true));
    test('[5, 4, 3, 2, 1]', () => expect(increasingTriplet([5, 4, 3, 2, 1])).toEqual(false));
    test('[2, 1, 5, 0, 4, 6]', () => expect(increasingTriplet([2, 1, 5, 0, 4, 6])).toEqual(true));
    test('[1, 2, 3]', () => expect(increasingTriplet([1, 2, 3])).toEqual(true));
    test('[1, 2]', () => expect(increasingTriplet([1, 2])).toEqual(false));
    test('[1, 3, 2, 4]', () => expect(increasingTriplet([1, 3, 2, 4])).toEqual(true));
    test('[20, 100, 10, 12, 5, 13]', () => expect(increasingTriplet([20, 100, 10, 12, 5, 13])).toEqual(true));
    test('[1, 1, -2, 6]', () => expect(increasingTriplet([1, 1, -2, 6])).toEqual(false));

    test('Array(50).fill(1)', () => expect(increasingTriplet(Array(50).fill(1))).toEqual(false));
})