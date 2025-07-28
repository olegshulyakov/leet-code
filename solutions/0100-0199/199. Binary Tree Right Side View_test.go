package main

import (
	"testing"

	"github.com/olegshulyakov/leet-code.go/structures"
)

/*

Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

Example 1:
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]

Example 2:
Input: root = [1,null,3]
Output: [1,3]

Example 3:
Input: root = []
Output: []

Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

*/

// rightSideView returns the right side view of a binary tree.
// The right side view consists of the values of the nodes visible
// when the tree is viewed from the right side, ordered from top to bottom.
//
// The algorithm uses a recursive approach:
// 1. For each node, we get the right side view of left and right subtrees
// 2. The result starts with the current node's value
// 3. We add all elements from the right subtree's view
// 4. If the left subtree's view is longer, we add the remaining elements
//
// Parameters:
//
//	root *TreeNode - Root of the binary tree
//
// Returns:
//
//	[]int - Values of nodes visible from the right side
func rightSideView(root *structures.TreeNode) []int {
	// Base case: empty tree
	if root == nil {
		return []int{}
	}

	// Start result with current node's value
	res := []int{root.Val}

	// Get right side views of left and right subtrees
	right := rightSideView(root.Right)
	left := rightSideView(root.Left)

	// Add all elements from right subtree view
	res = append(res, right...)

	// If left subtree view is longer, add the remaining elements
	if len(left) > len(right) {
		res = append(res, left[len(right):]...)
	}

	return res
}

func TestRightSideView(t *testing.T) {
	testcases := []struct {
		in   []any
		want []int
	}{
		{in: []any{1, 2, 3, nil, 5, nil, 4}, want: []int{1, 3, 4}},
		{in: []any{1, nil, 3}, want: []int{1, 3}},
		{in: []any{}, want: []int{}},
		{in: []any{1, 2}, want: []int{1, 2}},
		{in: []any{1, 2, 3, 4}, want: []int{1, 3, 4}},
	}
	for _, tc := range testcases {
		t.Run("199. Binary Tree Right Side View", func(t *testing.T) {
			t.Logf("rightSideView: in=%v", tc.in)
			out := rightSideView(structures.NewTree(tc.in))
			if len(out) != len(tc.want) {
				t.Errorf("len(fun() = %v, want: %v", len(out), len(tc.want))
			}
			for i := range out {
				if out[i] != tc.want[i] {
					t.Errorf("rightSideView()[%v] = %v, want: %v", i+1, out[i], tc.want[i])
				}
			}
		})
	}
}
