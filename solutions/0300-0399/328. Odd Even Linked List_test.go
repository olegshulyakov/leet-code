package main

import (
	"testing"

	"github.com/olegshulyakov/leet-code.go/structures"
)

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

func oddEvenList(head *structures.ListNode) *structures.ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	current := head.Next
	oddNode := head
	evenNode := head.Next
	memoryNode := head.Next

	for current != nil {
		current = current.Next
		oddNode.Next = current
		if current != nil {
			oddNode = current
		}

		if current != nil {
			current = current.Next
		} else {
			current = nil
		}
		evenNode.Next = current
		if current != nil {
			evenNode = current
		}
	}

	oddNode.Next = memoryNode
	return head
}

func Test(t *testing.T) {
	const name = "328. Odd Even Linked List"
	testCases := []struct {
		nums []int
		want []int
	}{
		{nums: []int{1, 2, 3, 4, 5}, want: []int{1, 3, 5, 2, 4}},
		{nums: []int{2, 1, 3, 5, 6, 4, 7}, want: []int{2, 3, 6, 7, 1, 5, 4}},
		{nums: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: []int{1, 3, 5, 7, 2, 4, 6, 8}},
		{nums: []int{}, want: []int{}},
		{nums: []int{1}, want: []int{1}},
		{nums: []int{1, 2}, want: []int{1, 2}},
	}
	for _, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Logf("oddEvenList(%v)", tc.nums)
			out := oddEvenList(structures.NewLinkedList(tc.nums)).ToArray()
			if len(out) != len(tc.want) {
				t.Errorf("len(oddEvenList() = %v, want: %v", len(out), len(tc.want))
			}
			for i := range out {
				if out[i] != tc.want[i] {
					t.Errorf("oddEvenList()[%v] = %v, want: %v", i+1, out[i], tc.want[i])
				}
			}
		})
	}
}
