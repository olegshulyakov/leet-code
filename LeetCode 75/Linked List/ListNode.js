/**
 * Definition for singly-linked list.
 * @param {number} val
 * @param {ListNode} next
 */
class ListNode {
    constructor(val, next) {
        this.val = (val === undefined ? 0 : val);
        this.next = (next === undefined ? null : next);
    }

    /**
     * Generate singly-linked list.
     * @param {Array} nums
     * @return {ListNode}
     */
    static from(nums) {
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