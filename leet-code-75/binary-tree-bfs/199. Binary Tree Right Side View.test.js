/*

Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.


Example 1:

Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]


Example 2:

Input: root = [1,null,3]
Output: [1,3]


Example 3:

Input: root = []
Output: []


Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

*/

const TreeNode = require('../../src/DataStructures/TreeNode');

/**
 * @param {TreeNode} root
 * @return {number[]}
 */
var rightSideView = function (root) {
    if (!root) return [];

    const res = [root.val];
    const right = rightSideView(root.right);
    const left = rightSideView(root.left);

    res.push(...right);
    if (left.length > right.length) {
        res.push(...left.slice(right.length));
    }

    return res;
};

describe('199. Binary Tree Right Side View', () => {
    test.each([
        { nums: [1, 2, 3, null, 5, null, 4], expected: [1, 3, 4] },
        { nums: [1, null, 3], expected: [1, 3] },
        { nums: [], expected: [] },
        { nums: [1, 2], expected: [1, 2] },
        { nums: [1, 2, 3, 4], expected: [1, 3, 4] },
    ])('$nums', ({ nums, expected }) => expect(rightSideView(TreeNode.from(nums))).toEqual(expected));
})