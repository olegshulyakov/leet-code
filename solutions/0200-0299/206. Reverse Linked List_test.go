package main

import (
	"testing"

	"github.com/olegshulyakov/leet-code.go/structures"
)

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

// reverseList reverses a singly-linked list.
func reverseList(head *structures.ListNode) *structures.ListNode {
	var current = head
	var previous *structures.ListNode

	for current != nil {
		temp := current.Next
		current.Next = previous
		previous = current
		if temp == nil {
			return current
		}
		current = temp
	}
	return head
}

func TestReverseList(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
	}
	for _, tc := range testCases {
		t.Run("206. Reverse Linked List", func(t *testing.T) {
			out := reverseList(structures.NewLinkedList(tc.nums)).ToArray()
			if len(out) != len(tc.want) {
				t.Errorf("len(fun() = %v, want: %v", len(out), len(tc.want))
			}
			for i := range out {
				if out[i] != tc.want[i] {
					t.Errorf("fun()[%v] = %v, want: %v", i+1, out[i], tc.want[i])
				}
			}
		})
	}
}
