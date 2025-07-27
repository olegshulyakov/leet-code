/*

Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:

Search for a node to remove.
If the node is found, delete the node.


Example 1:

Input: root = [5,3,6,2,4,null,7], key = 3
Output: [5,4,6,2,null,null,7]
Explanation: Given key to delete is 3. So we find the node with value 3 and delete it.
One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.


Example 2:

Input: root = [5,3,6,2,4,null,7], key = 0
Output: [5,3,6,2,4,null,7]
Explanation: The tree does not contain a node with value = 0.


Example 3:

Input: root = [], key = 0
Output: []


Constraints:

The number of nodes in the tree is in the range [0, 10^4].
-10^5 <= Node.val <= 10^5
Each node has a unique value.
root is a valid binary search tree.
-10^5 <= key <= 10^5


Follow up: Could you solve it with time complexity O(height of tree)?

*/

const TreeNode = require("../../src/DataStructures/TreeNode");

/**
 * @param {TreeNode} root
 * @param {number} key
 * @return {TreeNode}
 */
var deleteNode = function (root, key) {
    if (!root) return root;

    if (root.val > key) {
        root.left = deleteNode(root.left, key);
        return root;
    }
    if (root.val < key) {
        root.right = deleteNode(root.right, key);
        return root;
    }

    if (!root.left) {
        return root.right;
    }
    if (!root.right) {
        return root.left;
    }

    let minimal = root.right;
    while (minimal.left) minimal = minimal.left;

    root.val = minimal.val;
    root.right = deleteNode(root.right, minimal.val);

    return root;
}

describe('450. Delete Node in a BST', () => {
    test.each([
        { nums: [], key: 0, expected: [] },
        { nums: [5, 3, 6, 2, 4, null, 7], key: 0, expected: [5, 3, 6, 2, 4, null, 7] },
        { nums: [5, 3, 6, 2, 4, null, 7], key: 3, expected: [5, 4, 6, 2, null, null, 7] },
        { nums: [0], key: 0, expected: [] },
        { nums: [50, 30, 70, null, 40, 60, 80], key: 50, expected: [60, 30, 70, null, 40, null, 80] },
    ])('$nums, $key', ({ nums, key, expected }) => expect(TreeNode.toArray(deleteNode(TreeNode.from(nums), key))).toEqual(expected));
})