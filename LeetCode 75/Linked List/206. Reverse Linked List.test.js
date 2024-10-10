/*

Given the head of a singly linked list, reverse the list, and return the reversed list.


Example 1:

Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]


Example 2:

Input: head = [1,2]
Output: [2,1]


Example 3:

Input: head = []
Output: []


Constraints:

The number of nodes in the list is the range [0, 5000].
-5000 <= Node.val <= 5000


Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?

*/

const ListNode = require('./ListNode');

/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var reverseList = function (head) {
    let current = head,
        previous = null;
    while (current != null) {
        const temp = current.next;
        current.next = previous;
        previous = current;
        if (temp == null) {
            return current;
        } else {
            current = temp;
        }
    }
    return head;
};

describe('206. Reverse Linked List', () => {
    test.each([
        { nums: [1, 2, 3, 4, 5], expected: [5, 4, 3, 2, 1] },
        { nums: [1, 2], expected: [2, 1] },
        { nums: [], expected: [] },
        { nums: [1], expected: [1] },
    ])('$nums', ({ nums, expected }) => expect(ListNode.toArray(reverseList(ListNode.from(nums)))).toEqual(expected));
})