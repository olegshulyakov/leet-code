/*

Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.

Return the number of good nodes in the binary tree.


Example 1:

Input: root = [3,1,4,3,null,1,5]
Output: 4
Explanation: Nodes in blue are good.
Root Node (3) is always a good node.
Node 4 -> (3,4) is the maximum value in the path starting from the root.
Node 5 -> (3,4,5) is the maximum value in the path
Node 3 -> (3,1,3) is the maximum value in the path.


Example 2:

Input: root = [3,3,null,4,2]
Output: 3
Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.


Example 3:

Input: root = [1]
Output: 1
Explanation: Root is considered as good.


Constraints:

The number of nodes in the binary tree is in the range [1, 10^5].
Each node's value is between [-10^4, 10^4].

*/

const TreeNode = require('./TreeNode');

/**
 * @param {TreeNode} root
 * @return {number}
 */
var goodNodes = function (root) {
    if (!root) return 0;

    const checkSubTree = function (root, maxValue) {
        const maxVal = root.val > maxValue ? root.val : maxValue;
        return (root.val >= maxValue ? 1 : 0)
            + (root.left ? checkSubTree(root.left, maxVal) : 0)
            + (root.right ? checkSubTree(root.right, maxVal) : 0);
    }

    return checkSubTree(root, Number.NEGATIVE_INFINITY);
};

describe('1448. Count Good Nodes in Binary Tree', () => {
    test.each([
        { nums: [3, 1, 4, 3, null, 1, 5], expected: 4 },
        { nums: [3, 3, null, 4, 2], expected: 3 },
        { nums: [1], expected: 1 },
        { nums: [9, null, 3, 6], expected: 1 },
    ])('$nums', ({ nums, expected }) => expect(goodNodes(TreeNode.from(nums))).toEqual(expected));
})