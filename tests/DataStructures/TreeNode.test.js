const TreeNode = require('../../src/DataStructures/TreeNode')

describe('TreeNode', () => {
    it('should create a binary tree from an array of numbers', () => {
        const nums = [1, 2, null, 3, 4];
        const root = TreeNode.from(nums);
        expect(TreeNode.toArray(root)).toEqual([1, 2, null, 3, 4]);
    });

    it('should create a binary tree with no nodes', () => {
        const nums = [];
        const root = TreeNode.from(nums);
        expect(TreeNode.toArray(root)).toEqual([]);
    });

    it.skip('should not create a binary tree with duplicate values', () => {
        const nums = [1, 2, null, 3, 4];
        const modifiedNums = [1, 2, 2, 3, 4];
        const root = TreeNode.from(modifiedNums);
        expect(TreeNode.toArray(root)).toEqual(nums);
    });

    it('should  create a binary tree with duplicate values', () => {
        const nums = [1, 2, 2, 3, 4];
        const root = TreeNode.from(nums);
        expect(TreeNode.toArray(root)).toEqual(nums);
    });

    it('should create a binary tree with only one node', () => {
        const nums = [1];
        const root = TreeNode.from(nums);
        expect(TreeNode.toArray(root)).toEqual([1]);
    });

    it('should not throw an error when converting to array from null node', () => {
        expect(() => TreeNode.toArray(null)).not.toThrow();
    });
});