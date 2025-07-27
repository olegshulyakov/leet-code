/*

You are given the root of a binary search tree (BST) and an integer val.

Find the node in the BST that the node's value equals val and return the subtree rooted with that node. If such a node does not exist, return null.


Example 1:

Input: root = [4,2,7,1,3], val = 2
Output: [2,1,3]


Example 2:

Input: root = [4,2,7,1,3], val = 5
Output: []


Constraints:

The number of nodes in the tree is in the range [1, 5000].
1 <= Node.val <= 10^7
root is a binary search tree.
1 <= val <= 10^7

*/

const TreeNode = require('../../src/DataStructures/TreeNode');

/**
 * @param {TreeNode} root
 * @param {number} val
 * @return {TreeNode}
 */
var searchBST = function (root, val) {
    if (root == null) return null;
    if (root.val == val) return root;
    return searchBST(root.left, val) || searchBST(root.right, val);
};

describe('700. Search in a Binary Search Tree', () => {
    test.each([
        { nums: [4, 2, 7, 1, 3], val: 2, expected: [2, 1, 3] },
        { nums: [4, 2, 7, 1, 3], val: 5, expected: [] },
    ])('$nums, $val', ({ nums, val, expected }) => expect(TreeNode.toArray(searchBST(TreeNode.from(nums), val))).toEqual(expected));
})