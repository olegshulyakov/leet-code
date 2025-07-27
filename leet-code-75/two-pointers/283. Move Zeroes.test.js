/*

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.



Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]


Constraints:

1 <= nums.length <= 10^4
-231 <= nums[i] <= 231 - 1


Follow up: Could you minimize the total number of operations done?

*/

/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var moveZeroes = function (nums) {
    if (nums.length < 2) return nums;

    let i = 0;
    let counter = 0;
    while (counter < nums.length - 1) {
        counter++;
        if (nums[i] === 0) {
            nums.splice(i, 1);
            nums.push(0);
        } else {
            i++;
        }
    }

    return nums;
};

describe('283. Move Zeroes', () => {
    test.each([
        { nums: [0, 1, 0, 3, 12], expected: [1, 3, 12, 0, 0] },
        { nums: [0], expected: [0] },
    ])('$nums', ({ nums, expected }) => expect(moveZeroes(nums)).toEqual(expected));
})