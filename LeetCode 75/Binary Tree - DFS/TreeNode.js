/**
* Definition for a binary tree node.
* @param {number} val
* @param {TreeNode} left
* @param {TreeNode} right
*/
class TreeNode {
    constructor(val = 0, left = null, right = null) {
        this.val = val;
        this.left = left;
        this.right = right;
    }

    /**
     * Generate binary tree.
     * @param {Array} nums
     * @return {TreeNode}
     */
    static from(nums) {
        if (!nums.length) return null;

        const treeNodes = nums.map(val => val === null ? null : new TreeNode(val));

        let i = 0;
        while (2 * i + 1 < treeNodes.length) {
            if (treeNodes[2 * i + 1]) {
                treeNodes[i].left = treeNodes[2 * i + 1];
            }
            if (treeNodes[2 * i + 2]) {
                treeNodes[i].right = treeNodes[2 * i + 2];
            }
            i++;
        }

        return treeNodes[0];
    }

    /**
     * Convert binary tree to array.
     * @return {Array}
     */
    static toArray(node) {
        if (!node) return [];

        const queue = [node];
        const res = [];

        while (queue.length > 0) {
            const node = queue.shift();
            if (node) {
                res.push(node.val);
                queue.push(node.left, node.right);
            } else {
                res.push(null);
            }
        }

        while (res.length > 0 && res[res.length - 1] === null) {
            res.pop();
        }

        return res;
    }
}



module.exports = TreeNode;
