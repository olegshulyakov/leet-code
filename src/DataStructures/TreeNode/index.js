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

    static printTree(node, level = 0, isLast = false) {
        if (!node) return;

        const indent = ' '.repeat(level * 4);
        let output = `${indent}${isLast ? '└── ' : '├── '}${node.val}\n`;

        if (node.left || node.right) {
            output += this.printTree(node.left, level + 1, false);
            output += this.printTree(node.right, level + 1, true);
        }

        return output;
    }

    /**
     * Generate binary tree.
     * @param {Array} nums
     * @return {TreeNode}
     */
    static from(arr) {
        if (arr.length === 0 || arr[0] == null) return null;

        const root = new TreeNode(arr[0]);
        let queue = [root];
        let index = 1;

        while (queue.length > 0 && index < arr.length) {
            let node = queue.shift();

            if (index < arr.length && arr[index] != null) {
                node.left = new TreeNode(arr[index]);
                queue.push(node.left);
            }
            index++;

            if (index < arr.length && arr[index] != null) {
                node.right = new TreeNode(arr[index]);
                queue.push(node.right);
            }
            index++;
        }

        return root;
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
