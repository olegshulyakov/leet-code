/*

In a linked list of size n, where n is even, the ith node (0-indexed) of the linked list is known as the twin of the (n-1-i)th node, if 0 <= i <= (n / 2) - 1.

For example, if n = 4, then node 0 is the twin of node 3, and node 1 is the twin of node 2. These are the only nodes with twins for n = 4.
The twin sum is defined as the sum of a node and its twin.

Given the head of a linked list with even length, return the maximum twin sum of the linked list.


Example 1:

Input: head = [5,4,2,1]
Output: 6
Explanation:
Nodes 0 and 1 are the twins of nodes 3 and 2, respectively. All have twin sum = 6.
There are no other nodes with twins in the linked list.
Thus, the maximum twin sum of the linked list is 6.


Example 2:

Input: head = [4,2,2,3]
Output: 7
Explanation:
The nodes with twins present in this linked list are:
- Node 0 is the twin of node 3 having a twin sum of 4 + 3 = 7.
- Node 1 is the twin of node 2 having a twin sum of 2 + 2 = 4.
Thus, the maximum twin sum of the linked list is max(7, 4) = 7.


Example 3:

Input: head = [1,100000]
Output: 100001
Explanation:
There is only one node with a twin in the linked list having twin sum of 1 + 100000 = 100001.


Constraints:

The number of nodes in the list is an even integer in the range [2, 10^5].
1 <= Node.val <= 10^5

*/

const ListNode = require('./ListNode');

/**
 * Reverse linked list
 * @param {ListNode} head
 * @return {ListNode}
 */
function reverse(head) {
    let prev = null;
    let current = new ListNode(head.val, head.next);
    while (current) {
        let next = current.next;
        current.next = prev;
        prev = current;
        current = next != null ? new ListNode(next.val, next.next) : null;
    }
    return prev;
}
/**
 * @param {ListNode} head
 * @return {number}
 */
var pairSum = function (head) {
    let max = Number.NEGATIVE_INFINITY;
    let tail = reverse(head);
    while (head != null && tail != null) {
        const sum = head.val + tail.val;
        max = sum > max ? sum : max;
        head = head.next;
        tail = tail.next;
    }
    return max;
};

describe('2130. Maximum Twin Sum of a Linked List', () => {
    test.each([
        { nums: [5, 4, 2, 1], expected: 6 },
        { nums: [4, 2, 2, 3], expected: 7 },
        { nums: [1, 100000], expected: 100001 },
        { nums: [47, 22, 81, 46, 94, 95, 90, 22, 55, 91, 6, 83, 49, 65, 10, 32, 41, 26, 83, 99, 14, 85, 42, 99, 89, 69, 30, 92, 32, 74, 9, 81, 5, 9], expected: 182 },
    ])('$nums', ({ nums, expected }) => expect(pairSum(ListNode.from(nums))).toEqual(expected));
})