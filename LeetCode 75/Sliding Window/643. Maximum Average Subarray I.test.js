/*

You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.



Example 1:

Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
Example 2:

Input: nums = [5], k = 1
Output: 5.00000


Constraints:

n == nums.length
1 <= k <= n <= 10^5
-104 <= nums[i] <= 10^4

*/

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var findMaxAverage = function (nums, k) {
    if (nums.length < k) {
        return 0;
    }

    let sum = temp = 0;
    for (let i = 0; i < nums.length; i++) {
        temp += i < k ? nums[i] : nums[i] - nums[i - k];
        sum = i < k ? temp : sum > temp ? sum : temp;
    }

    return sum / k;
};

describe('643. Maximum Average Subarray I', () => {
    test.each([
        { nums: [1, 12, -5, -6, 50, 3], k: 4, expected: 12.75 },
        { nums: [5], k: 1, expected: 5 },
        { nums: [0, 1, 1, 3, 3], k: 4, expected: 2 },
        { nums: [0, 4, 0, 3, 2], k: 1, expected: 4 },
    ])('$nums, $k', ({ nums, k, expected }) => expect(findMaxAverage(nums, k)).toEqual(expected));
})