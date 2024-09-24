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

1 <= nums.length <= 5 * 105
-231 <= nums[i] <= 231 - 1


Follow up: Could you implement a solution that runs in O(n) time complexity and O(1) space complexity?

*/

/**
 * @param {number[]} nums
 * @return {boolean}
 */
var increasingTriplet = function (nums) {
    if (nums.length < 3) return false;

    for (let j = 1; j < nums.length - 1; j++) {
        let isFound = false;
        for (let i = j - 1; i >= 0; i--) {
            if (nums[i] < nums[j]) {
                isFound = true;
                break;
            }
        }
        if (!isFound) {
            continue;
        }

        for (let k = j + 1; k < nums.length; k++) {
            if (nums[j] < nums[k]) {
                return true;
            }
        }
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

    test('Array(50).fill(1)', () => expect(increasingTriplet(Array(50).fill(1))).toEqual(false));
})