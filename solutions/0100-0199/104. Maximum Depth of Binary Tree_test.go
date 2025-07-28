package main

import (
	"testing"

	"github.com/olegshulyakov/leet-code.go/structures"
)

/*

Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.


Example 1:

Input: root = [3,9,20,null,null,15,7]
Output: 3


Example 2:

Input: root = [1,null,2]
Output: 2


Constraints:

The number of nodes in the tree is in the range [0, 10^4].
-100 <= Node.val <= 100

*/

// maxDepth calculates the maximum depth of a binary tree.
// The maximum depth is the number of nodes along the longest path
// from the root node down to the farthest leaf node.
//
// The algorithm uses a recursive depth-first search approach:
// 1. If the node is nil, return 0 (base case)
// 2. Recursively calculate the depth of left and right subtrees
// 3. Return the maximum of the two depths plus 1 (for current node)
//
// Parameters:
//
//	root *TreeNode - Root of the binary tree
//
// Returns:
//
//	int - Maximum depth of the tree
func maxDepth(root *structures.TreeNode) int {
	// Base case: empty tree has depth 0
	if root == nil {
		return 0
	}

	// Recursively calculate depths of left and right subtrees
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	// Return the maximum depth plus 1 for current node
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func TestMaxDepth(t *testing.T) {
	testCases := []struct {
		root []any
		want int
	}{
		{root: []any{3, 9, 20, nil, nil, 15, 7}, want: 3},
		{root: []any{1, nil, 2}, want: 2},
		{root: []any{}, want: 0},
	}
	for _, tc := range testCases {
		t.Run("104. Maximum Depth of Binary Tree", func(t *testing.T) {
			out := maxDepth(structures.NewTree(tc.root))
			if out != tc.want {
				t.Errorf("maxDepth(%v) = %v, want: %v", tc.root, out, tc.want)
			}
		})
	}
}
