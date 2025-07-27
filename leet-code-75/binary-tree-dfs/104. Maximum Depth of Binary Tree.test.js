/*

Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.


Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: 3


Example 2:

Input: root = [1,null,2]
Output: 2


Constraints:

The number of nodes in the tree is in the range [0, 10^4].
-100 <= Node.val <= 100

*/

const TreeNode = require('../../src/DataStructures/TreeNode');

/**
 * @param {TreeNode} root
 * @return {number}
 */
function maxDepth(root) {
    if (root == null) return 0;

    const l = maxDepth(root.left);
    const r = maxDepth(root.right);

    return (l > r ? l : r) + 1;
}

describe('104. Maximum Depth of Binary Tree', () => {
    test.each([
        { nums: [3, 9, 20, null, null, 15, 7], expected: 3 },
        { nums: [1, null, 2], expected: 2 },
        { nums: [], expected: 0 },
    ])('$nums', ({ nums, expected }) => expect(maxDepth(TreeNode.from(nums))).toEqual(expected));
})