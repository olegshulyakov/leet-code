package main

import (
	"math"
	"testing"

	"github.com/olegshulyakov/leet-code.go/structures"
)

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

func pairSum(head *structures.ListNode) int {
	// First, we need to reverse the linked list
	var n int
	var tail *structures.ListNode
	current := &structures.ListNode{Val: head.Val, Next: head.Next}

	for current != nil {
		next := current.Next
		current.Next = tail
		tail = current
		if next != nil {
			current = &structures.ListNode{Val: next.Val, Next: next.Next}
		} else {
			current = nil
		}
		n++
	}

	// Now find the maximum twin sum
	maximum := math.MinInt32
	for i := 0; i <= (n/2)-1; i++ {
		sum := head.Val + tail.Val
		if sum > maximum {
			maximum = sum
		}
		head = head.Next
		tail = tail.Next
	}

	return maximum
}

func TestPairSum(t *testing.T) {
	const name = "2130. Maximum Twin Sum of a Linked List"
	testCases := []struct {
		nums []int
		want int
	}{
		{[]int{5, 4, 2, 1}, 6},
		{[]int{4, 2, 2, 3}, 7},
		{[]int{1, 100000}, 100001},
		{
			[]int{47, 22, 81, 46, 94, 95, 90, 22, 55, 91, 6, 83, 49, 65, 10, 32, 41, 26, 83, 99, 14, 85, 42, 99, 89, 69, 30, 92, 32, 74, 9, 81, 5, 9},
			182,
		},
	}
	for _, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Logf("pairSum(%v)", tc.nums)
			out := pairSum(structures.NewLinkedList(tc.nums))
			if out != tc.want {
				t.Errorf("deleteMiddle(%v) = %v, want: %v", tc.nums, out, tc.want)
			}
		})
	}
}
