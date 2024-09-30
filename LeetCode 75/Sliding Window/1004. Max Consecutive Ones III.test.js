/*

Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.


Example 1:

Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [0,1,1,1,1,0]


Example 2:

Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Explanation: [1,1,0,0,1,1,1,0,1,1]


Constraints:

1 <= nums.length <= 10^5
nums[i] is either 0 or 1.
0 <= k <= nums.length

*/

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var longestOnes = function (nums, k) {
    let left = max = zeros = 0;
    for (let i = 0; i < nums.length; i++) {
        if (0 === nums[i]) {
            if (zeros === k) {
                while (nums[left++] !== 0) { }
            } else {
                zeros++;
            }
        }

        const diff = i - left + 1;
        if (max < diff) {
            max = diff;
        }
    }

    return max;
};

describe('1004. Max Consecutive Ones III', () => {
    test('[1, 1, 1], 2', () => expect(longestOnes([1, 1, 1], 2)).toEqual(3));
    test('[1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0], 2', () => expect(longestOnes([1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0], 2)).toEqual(6));
    test('[0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1], 3', () => expect(longestOnes([0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1], 3)).toEqual(10));
    test('[0 .. 0], 2', () => expect(longestOnes(new Array(10).fill(0), 2)).toEqual(2));
    test('[0 .. 0], 10', () => expect(longestOnes(new Array(10).fill(0), 10)).toEqual(10));
})