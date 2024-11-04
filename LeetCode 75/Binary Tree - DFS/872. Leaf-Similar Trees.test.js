/*

Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value sequence.


For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).

Two binary trees are considered leaf-similar if their leaf value sequence is the same.

Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.


Example 1:

Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
Output: true


Example 2:

Input: root1 = [1,2,3], root2 = [1,3,2]
Output: false


Constraints:

The number of nodes in each tree will be in the range [1, 200].
Both of the given trees will have values in the range [0, 200].

*/

const TreeNode = require('./TreeNode');

/**
 * @param {TreeNode} root1
 * @param {TreeNode} root2
 * @return {boolean}
 */
var leafSimilar = function (root1, root2) {
    const toLeafArray = (node) => {
        if (!node) return [];

        const queue = [node];
        const res = [];

        while (queue.length > 0) {
            const node = queue.shift();
            if (node == null) continue;

            if (node.left == null && node.right == null) {
                res.push(node.val);
            } else {
                queue.unshift(node.left, node.right);
            }
        }

        while (res.length > 0 && res[res.length - 1] === null) {
            res.pop();
        }

        return res;
    }

    const leafs1 = toLeafArray(root1);
    const leafs2 = toLeafArray(root2);

    return JSON.stringify(leafs1) == JSON.stringify(leafs2);
};

describe('872. Leaf-Similar Trees', () => {
    test.each([
        { root1: [], root2: [], expected: true },
        { root1: [3, 5, 1, 6, 2, 9, 8, null, null, 7, 4], root2: [3, 5, 1, 6, 7, 4, 2, null, null, null, null, null, null, 9, 8], expected: true },
        { root1: [1, 2, 3], root2: [1, 3, 2], expected: false },
        { root1: [1, 2], root2: [2, 2], expected: true },
        { root1: [1], root2: [2], expected: false },
        { root1: [1], root2: [1], expected: true },
    ])('$root1, $root2', ({ root1, root2, expected }) => expect(leafSimilar(TreeNode.from(root1), TreeNode.from(root2))).toEqual(expected));
})