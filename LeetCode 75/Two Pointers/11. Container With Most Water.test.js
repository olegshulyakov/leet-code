/*

ou are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.



Example 1:


Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
Example 2:

Input: height = [1,1]
Output: 1


Constraints:

n == height.length
2 <= n <= 105
0 <= height[i] <= 104

*/

/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function (height) {
    let num1 = Number.NEGATIVE_INFINITY,
        num2 = Number.NEGATIVE_INFINITY,
        num1Index = -1,
        num2Index = -1,
        areaSize = 0;

    for (let i = 0; i < height.length; i++) {
        if (height[i] > num1) {
            num1 = height[i];
            num1Index = i;
        } else if (height[i] > num2 || (height[i] * (i - num1Index) > areaSize)) {
            num2 = height[i];
            num2Index = i;
            areaSize = num2 * (num2Index - num1Index)
        }
    }

    if (num1 > 0 && num2 > 0) {
        return areaSize;
    }

    return 0;
};

describe('11. Container With Most Water', () => {
    test('[1, 8, 6, 2, 5, 4, 8, 3, 7]', () => expect(maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7])).toEqual(49));
    test('[1, 1]', () => expect(maxArea([1, 1])).toEqual(1));
    test('[]', () => expect(maxArea([])).toEqual(0));
    test('[1, 2, 1]', () => expect(maxArea([1, 2, 1])).toEqual(2));
})