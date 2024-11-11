/*

Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”


Example 1:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.


Example 2:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
Output: 5
Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.


Example 3:

Input: root = [1,2], p = 1, q = 2
Output: 1


Constraints:

The number of nodes in the tree is in the range [2, 10^5].
-10^9 <= Node.val <= 10^9
All Node.val are unique.
p != q
p and q will exist in the tree.

*/

const TreeNode = require('./TreeNode');

/**
 * @param {TreeNode} root
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {TreeNode}
 */
var lowestCommonAncestor = function (root, p, q) {
    if (root == null || root.val === p.val || root.val === q.val) return root;

    const left = lowestCommonAncestor(root.left, p, q);
    const right = lowestCommonAncestor(root.right, p, q);

    return left != null && right != null ? root : left || right;
};

describe('236. Lowest Common Ancestor of a Binary Tree', () => {
    test.each([
        { nums: [3, 5, 1, 6, 2, 0, 8, null, null, 7, 4], p: 5, q: 1, expected: 3 },
        { nums: [3, 5, 1, 6, 2, 0, 8, null, null, 7, 4], p: 5, q: 4, expected: 5 },
        { nums: [1, 2], p: 1, q: 2, expected: 1 },
    ])('$nums, $p, $q', ({ nums, p, q, expected }) => expect(lowestCommonAncestor(TreeNode.from(nums), new TreeNode(p), new TreeNode(q))?.val).toEqual(expected));
})