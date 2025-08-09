package main

import (
	"testing"

	"github.com/olegshulyakov/leet-code.go/structures"
)

/*

You are given the head of a linked list. Delete the middle node, and return the head of the modified linked list.

The middle node of a linked list of size n is the ⌊n / 2⌋th node from the start using 0-based indexing, where ⌊x⌋ denotes the largest integer less than or equal to x.

For n = 1, 2, 3, 4, and 5, the middle nodes are 0, 1, 1, 2, and 2, respectively.


Example 1:

Input: head = [1,3,4,7,1,2,6]
Output: [1,3,4,1,2,6]
Explanation:
The above figure represents the given linked list. The indices of the nodes are written below.
Since n = 7, node 3 with value 7 is the middle node, which is marked in red.
We return the new list after removing this node.


Example 2:

Input: head = [1,2,3,4]
Output: [1,2,4]
Explanation:
The above figure represents the given linked list.
For n = 4, node 2 with value 3 is the middle node, which is marked in red.


Example 3:

Input: head = [2,1]
Output: [2]
Explanation:
The above figure represents the given linked list.
For n = 2, node 1 with value 1 is the middle node, which is marked in red.
Node 0 with value 2 is the only node remaining after removing node 1.


Constraints:

The number of nodes in the list is in the range [1, 10^5].
1 <= Node.val <= 10^5

*/

func deleteMiddle(head *structures.ListNode) *structures.ListNode {
	// Handle edge cases
	if head == nil || head.Next == nil {
		return nil
	}

	// Count total size
	temp := head
	index := 0.5
	for temp != nil {
		temp = temp.Next
		index += 0.5
	}

	// Adjust index for middle calculation
	if index-float64(int64(index)) > 0 {
		index += 0.5
	}

	// Remove the middle node
	temp = head
	i := 1
	for i < int(index-1) && temp.Next.Next != nil {
		temp = temp.Next
		i++
	}

	temp.Next = temp.Next.Next
	return head
}

func Test(t *testing.T) {
	const name = "2095. Delete the Middle Node of a Linked List"
	testCases := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 3, 4, 7, 1, 2, 6}, []int{1, 3, 4, 1, 2, 6}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 4}},
		{[]int{2, 1}, []int{2}},
		{[]int{1}, []int{}},
	}
	for _, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Logf("deleteMiddle(%v)", tc.nums)
			out := deleteMiddle(structures.NewLinkedList(tc.nums)).ToArray()
			if len(out) != len(tc.want) {
				t.Errorf("len(deleteMiddle() = %v, want: %v", len(out), len(tc.want))
			}
			for i := range out {
				if out[i] != tc.want[i] {
					t.Errorf("deleteMiddle()[%v] = %v, want: %v", i+1, out[i], tc.want[i])
				}
			}
		})
	}
}
