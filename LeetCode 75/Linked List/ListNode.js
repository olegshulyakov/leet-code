/**
 * Definition for singly-linked list.
 * @param {number} val
 * @param {ListNode} next
 */
class ListNode {
    constructor(val = 0, next = null) {
        this.val = val;
        this.next = next;
    }

    /**
     * Generate singly-linked list.
     * @param {Array} nums
     * @return {ListNode}
     */
    static from(nums) {
        if (!Array.isArray(nums)) {
            throw new Error('Input must be an array');
        }

        if (nums.length === 0) {
            return undefined;
        }

        let head = undefined;
        for (let i = nums.length - 1; i >= 0; i--) {
            head = new ListNode(nums[i], head);
        }
        return head;
    }

    /**
     * Convert singly-linked list to array.
     * @return {Array}
     */
    static toArray(node) {
        if (node == null) {
            throw new Error('Input node cannot be null');
        }
        if (!(node instanceof ListNode)) {
            throw new Error('Input must be a ListNode');
        }

        const res = [];
        let pointer = node;
        while (pointer != null) {
            res.push(pointer.val);
            pointer = pointer.next;
        }
        return res;
    }
}



module.exports = ListNode;