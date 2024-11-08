/*

Given the root of a binary tree and an integer targetSum, return the number of paths where the sum of the values along the path equals targetSum.

The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from parent nodes to child nodes).


Example 1:

Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
Output: 3
Explanation: The paths that sum to 8 are shown.


Example 2:

Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
Output: 3


Constraints:

The number of nodes in the tree is in the range [0, 1000].
-10^9 <= Node.val <= 10^9
-1000 <= targetSum <= 1000

*/

const TreeNode = require('./TreeNode');

/**
 * @param {TreeNode} root
 * @param {number} targetSum
 * @return {number}
 */
var pathSum = function (root, targetSum) {
    if (root == null) return 0;

    const subSum = (node, target) => {
        if (node == null) return 0;
        let count = node.val === target ? 1 : 0;

        target -= node.val;
        return count + (node.left ? subSum(node.left, target) : 0) + (node.right ? subSum(node.right, target) : 0);
    }

    return subSum(root, targetSum)
        + (root.left ? pathSum(root.left, targetSum) : 0)
        + (root.right ? pathSum(root.right, targetSum) : 0);
};

describe('437. Path Sum III', () => {
    test.each([
        { nums: [10, 5, -3, 3, 2, null, 11, 3, -2, null, 1], targetSum: 8, expected: 3 },
        { nums: [5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1], targetSum: 22, expected: 3 },
        { nums: [], targetSum: 1, expected: 0 },
        { nums: [-2, null, -3], targetSum: -5, expected: 1 },
        { nums: [1, -2, -3, 1, 3, -2, null, -1], targetSum: -1, expected: 4 },
    ])('$nums, $targetSum', ({ nums, targetSum, expected }) => expect(pathSum(TreeNode.from(nums), targetSum)).toEqual(expected));
})