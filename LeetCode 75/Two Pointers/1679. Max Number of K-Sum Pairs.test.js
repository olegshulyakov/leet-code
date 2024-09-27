/*

You are given an integer array nums and an integer k.

In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.

Return the maximum number of operations you can perform on the array.



Example 1:

Input: nums = [1,2,3,4], k = 5
Output: 2
Explanation: Starting with nums = [1,2,3,4]:
- Remove numbers 1 and 4, then nums = [2,3]
- Remove numbers 2 and 3, then nums = []
There are no more pairs that sum up to 5, hence a total of 2 operations.
Example 2:

Input: nums = [3,1,3,4,3], k = 6
Output: 1
Explanation: Starting with nums = [3,1,3,4,3]:
- Remove the first two 3's, then nums = [1,4,3]
There are no more pairs that sum up to 6, hence a total of 1 operation.


Constraints:

1 <= nums.length <= 10^5
1 <= nums[i] <= 10^9
1 <= k <= 10^9

*/

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var maxOperations = function (nums, k) {
    if (nums.length < 2) return 0;

    let ops = 0;
    const map = new Map();
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] > k) {
            continue;
        }

        const el = nums[i];
        const difference = k - el;
        if (map.has(difference)) {
            let counter = map.get(difference);
            if (counter > 1) {
                map.set(difference, --counter);
            } else {
                map.delete(difference);
            }
            ops++;
        } else {
            if (map.has(el)) {
                let counter = map.get(el);
                map.set(el, ++counter);
            } else {
                map.set(el, 1);
            }
        }
    }

    return ops;
};

describe('1679. Max Number of K-Sum Pairs', () => {
    test('[1, 2, 3, 4], 5', () => expect(maxOperations([1, 2, 3, 4], 5)).toEqual(2));
    test('[3, 1, 3, 4, 3], 6', () => expect(maxOperations([3, 1, 3, 4, 3], 6)).toEqual(1));
    test('[[2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2], 3', () => expect(maxOperations([2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2], 3)).toEqual(4));
    test('[5000000 .. 5000000], 10000000', () => expect(maxOperations(new Array(10000).fill(5000000), 10000000)).toEqual(5000));
})