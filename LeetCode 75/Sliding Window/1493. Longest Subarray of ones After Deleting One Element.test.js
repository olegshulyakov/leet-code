/*

Given a binary array nums, you should delete one element from it.

Return the size of the longest non-empty subarray containing only 1's in the resulting array. Return 0 if there is no such subarray.



Example 1:

Input: nums = [1,1,0,1]
Output: 3
Explanation: After deleting the number in position 2, [1,1,1] contains 3 numbers with value of 1's.
Example 2:

Input: nums = [0,1,1,1,0,1,1,0,1]
Output: 5
Explanation: After deleting the number in position 4, [0,1,1,1,1,1,0,1] longest subarray with value of 1's is [1,1,1,1,1].
Example 3:

Input: nums = [1,1,1]
Output: 2
Explanation: You must delete one element.


Constraints:

1 <= nums.length <= 10^5
nums[i] is either 0 or 1.

*/

/**
 * @param {number[]} nums
 * @return {number}
 */
var longestSubarray = function (nums) {
    let left = max = 0;
    let hasZero = false;
    for (let i = 0; i < nums.length; i++) {
        if (0 === nums[i]) {
            if (hasZero) {
                while (nums[left++] !== 0) { }
            } else {
                hasZero = true;
            }
        }

        const diff = i - left + 1;
        if (max < diff) {
            max = diff;
        }
    }

    return max - 1;
};

describe('1493. Longest Subarray of ones After Deleting One Element', () => {
    test('[1, 1, 0, 1]', () => expect(longestSubarray([1, 1, 0, 1])).toEqual(3));
    test('[0, 1, 1, 1, 0, 1, 1, 0, 1]', () => expect(longestSubarray([0, 1, 1, 1, 0, 1, 1, 0, 1])).toEqual(5));
    test('[1, 1, 1]', () => expect(longestSubarray([1, 1, 1])).toEqual(2));
})