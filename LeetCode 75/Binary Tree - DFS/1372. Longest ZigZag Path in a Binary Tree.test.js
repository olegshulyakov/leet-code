/*

You are given the root of a binary tree.

A ZigZag path for a binary tree is defined as follow:

Choose any node in the binary tree and a direction (right or left).
If the current direction is right, move to the right child of the current node; otherwise, move to the left child.
Change the direction from right to left or from left to right.
Repeat the second and third steps until you can't move in the tree.
Zigzag length is defined as the number of nodes visited - 1. (A single node has a length of 0).

Return the longest ZigZag path contained in that tree.


Example 1:

Input: root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1]
Output: 3
Explanation: Longest ZigZag path in blue nodes (right -> left -> right).


Example 2:

Input: root = [1,1,1,null,1,null,null,1,1,null,1]
Output: 4
Explanation: Longest ZigZag path in blue nodes (left -> right -> left -> right).


Example 3:

Input: root = [1]
Output: 0


Constraints:

The number of nodes in the tree is in the range [1, 5 * 10^4].
1 <= Node.val <= 100

*/

const TreeNode = require('./TreeNode');

/**
 * @param {TreeNode} root
 * @return {number}
 */
var longestZigZag = function (root) {
    let max = 0;
    const queue = [root];

    while (queue.length > 0) {
        const cursor = queue.pop();

        let count = 0;
        let node = cursor;
        while (node != null) {
            node = count % 2 ? node.left : node.right;
            count++;
        }
        count--;
        if (max < count) max = count;

        count = 0;
        node = cursor;
        while (node != null) {
            node = count % 2 ? node.right : node.left;
            count++;
        }
        count--;
        if (max < count) max = count;

        if (cursor.left != null) queue.push(cursor.left);
        if (cursor.right != null) queue.push(cursor.right);
    }

    return max;
};

describe('1372. Longest ZigZag Path in a Binary Tree', () => {
    test.each([
        { nums: [1, null, 1, 1, 1, null, null, 1, 1, null, 1, null, null, null, 1], expected: 3 },
        { nums: [1, 1, 1, null, 1, null, null, 1, 1, null, 1], expected: 4 },
        { nums: [1], expected: 0 },
    ])('$nums', ({ nums, expected }) => expect(longestZigZag(TreeNode.from(nums))).toEqual(expected));
})