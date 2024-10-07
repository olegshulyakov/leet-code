/*

Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.

The first node is considered odd, and the second node is even, and so on.

Note that the relative order inside both the even and odd groups should remain as it was in the input.

You must solve the problem in O(1) extra space complexity and O(n) time complexity.


Example 1:

Input: head = [1,2,3,4,5]
Output: [1,3,5,2,4]


Example 2:

Input: head = [2,1,3,5,6,4,7]
Output: [2,3,6,7,1,5,4]


Constraints:

The number of nodes in the linked list is in the range [0, 10^4].
-10^6 <= Node.val <= 10^6

*/

const ListNode = require('./ListNode');

/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var oddEvenList = function (head) {
    if (head == null || head.next == null || head.next.next == null) {
        return head;
    }

    let current = head.next,
        oddNode = head,
        evenNode = head.next,
        memoryNode = head.next;
    while (current != null) {
        current = current.next;
        oddNode.next = current;
        oddNode = current != null ? current : oddNode;

        current = current != null ? current.next : null;
        evenNode.next = current;
        evenNode = current != null ? current : evenNode;
    }

    oddNode.next = memoryNode;

    return head;
};

describe('328. Odd Even Linked List', () => {
    test.each([
        { nums: [1, 2, 3, 4, 5], expected: [1, 3, 5, 2, 4] },
        { nums: [2, 1, 3, 5, 6, 4, 7], expected: [2, 3, 6, 7, 1, 5, 4] },
        { nums: [1, 2, 3, 4, 5, 6, 7, 8], expected: [1, 3, 5, 7, 2, 4, 6, 8] },
        { nums: [], expected: [] },
        { nums: [1], expected: [1] },
        { nums: [1, 2], expected: [1, 2] },
    ])('$nums', ({ nums, expected }) => expect(ListNode.toArray(oddEvenList(ListNode.from(nums)))).toEqual(expected));
})