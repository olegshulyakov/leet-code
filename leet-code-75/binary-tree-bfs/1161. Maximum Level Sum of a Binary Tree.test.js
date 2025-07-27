/*

Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.

Return the smallest level x such that the sum of all the values of nodes at level x is maximal.


Example 1:

Input: root = [1,7,0,7,-8,null,null]
Output: 2
Explanation:
Level 1 sum = 1.
Level 2 sum = 7 + 0 = 7.
Level 3 sum = 7 + -8 = -1.
So we return the level with the maximum sum which is level 2.


Example 2:

Input: root = [989,null,10250,98693,-89388,null,null,null,-32127]
Output: 2


Constraints:

The number of nodes in the tree is in the range [1, 10^4].
-10^5 <= Node.val <= 10^5

*/

const TreeNode = require('../../src/DataStructures/TreeNode');

/**
 * @param {TreeNode} root
 * @return {number}
 */
var maxLevelSum = function (root) {
    if (!root) return 0;

    let max = root.val,
        res = 1,
        depth = 1;
    let queue = [root.left, root.right];
    while (queue.length > 0) {
        depth++;
        const sum = queue.reduce((acc, current) => acc + (current?.val || 0), 0);

        if (sum > max) {
            max = sum;
            res = depth;
        }

        const newQueue = [];
        queue.forEach(current => {
            if (current == null) return;
            if (current.left != null) newQueue.push(current.left);
            if (current.right != null) newQueue.push(current.right);
        });
        queue = newQueue;
    }

    return res;
};

describe('1161. Maximum Level Sum of a Binary Tree', () => {
    test.each([
        { nums: [1, 7, 0, 7, -8, null, null], expected: 2 },
        { nums: [989, null, 10250, 98693, -89388, null, null, null, -32127], expected: 2 },
        { nums: [], expected: 0 },
        { nums: [1], expected: 1 },
    ])('$nums', ({ nums, expected }) => expect(maxLevelSum(TreeNode.from(nums))).toEqual(expected));
})